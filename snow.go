package snowflake

import (
    "errors"
    "sync"
    "time"
)

const (
    //系统时间回退错误
    ErrClockMoveBackwards = "refuse_generate_id_by_system_clock_moved_backwards"
    //无效的机器ID
    ErrInvalidWorkerId = "invalid_worker_id"
    //无效的数据中心ID
    ErrInvalidDataCenterId = "invalid_data_center_id"
)

const (
    //开始时间戳
    startTimestamp = 1558930914000
    //机器ID位数
    workerIdBits = 3
    //支持最大机器ID
    maxWorkerId = -1 ^ (-1 << workerIdBits)
    //数据中心ID位数
    dataCenterIdBits = 2
    //支持最大数据中心ID
    maxDataCenterId = -1 ^ (-1 << dataCenterIdBits)
    //序列位数
    sequenceBits = 12
    //机器ID左移位数
    workerIdShift = sequenceBits
    //数据中心ID左移位数
    dataCenterIdShift = workerIdBits + sequenceBits
    //时间戳左移位数
    timestampLeftShift = dataCenterIdBits + workerIdBits + sequenceBits
    //序列最大掩码
    sequenceMask = -1 ^ (-1 << sequenceBits)
)

type Snowflake struct {
    workerId      int64 //机器ID
    dataCenterId  int64 //数据中心ID
    sequence      int64 //序列
    lastTimestamp int64 //上次时间戳
    mutex         *sync.Mutex
}

func NewSnowflake(workerId, dataCenterId int64) (*Snowflake, error) {
    if workerId < 0 || workerId > maxWorkerId {
        return nil, errors.New(ErrInvalidWorkerId)
    }
    if dataCenterId < 0 || dataCenterId > maxDataCenterId {
        return nil, errors.New(ErrInvalidDataCenterId)
    }
    s := &Snowflake{workerId, dataCenterId, 0, -1, new(sync.Mutex)}
    return s, nil
}

func (s *Snowflake) NextId() (int64, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    return s.nextId()
}

func (s *Snowflake) nextId() (id int64, err error) {
    t := timeGen()
    //当前时间戳小于上次，则系统时间发生过回退
    if t < s.lastTimestamp {
        return 0, errors.New(ErrClockMoveBackwards)
    }
    //当前时间戳等于上次，则生成该毫秒内序列
    if t == s.lastTimestamp {
        s.sequence = (s.sequence + 1) & sequenceMask
        //序列溢出
        if s.sequence == 0 {
            //等待下一毫秒
            t = nextMillisecond(s.lastTimestamp)
        }
    } else {
        //时间戳改变，序列重置
        s.sequence = 0
    }
    s.lastTimestamp = t
    return ((t - startTimestamp) << timestampLeftShift) | (s.dataCenterId << dataCenterIdShift) | (s.workerId << workerIdShift) | s.sequence,
        nil
}

func timeGen() int64 {
    return time.Now().UnixNano() / 1000000
}

func nextMillisecond(last int64) int64 {
    t := timeGen()
    for t <= last {
        t = timeGen()
    }
    return t
}

// Package snowflake 分布式ID雪花算法Golang实现
// Powered by CyrivlClth
package snowflake

import (
	"errors"
	"sync"
	"time"

	"github.com/CyrivlClth/snowflake/idgen"
)

var (
	// ErrClockMoveBackwards 系统时间回退错误
	ErrClockMoveBackwards = errors.New("refuse_generate_id_by_system_clock_moved_backwards")
	// ErrInvalidWorkerID 无效的机器ID
	ErrInvalidWorkerID = errors.New("invalid_worker_id")
	// ErrInvalidDataCenterID 无效的数据中心ID
	ErrInvalidDataCenterID = errors.New("invalid_data_center_id")
)

const (
	// 开始时间戳
	startTimestamp = 1558930914000
	// 机器ID位数
	workerIDBits = 5
	// 支持最大机器ID
	maxWorkerID = -1 ^ (-1 << workerIDBits)
	// 数据中心ID位数
	dataCenterIDBits = 5
	// 支持最大数据中心ID
	maxDataCenterID = -1 ^ (-1 << dataCenterIDBits)
	// 序列位数
	sequenceBits = 12
	// 机器ID左移位数
	workerIDShift = sequenceBits
	// 数据中心ID左移位数
	dataCenterIDShift = workerIDBits + sequenceBits
	// 时间戳左移位数
	timestampLeftShift = dataCenterIDBits + workerIDBits + sequenceBits
	// 序列最大掩码
	sequenceMask = -1 ^ (-1 << sequenceBits)
)

// snowflake 雪花算法结构体
type snowflake struct {
	workerID      int64 // 机器ID
	dataCenterID  int64 // 数据中心ID
	sequence      int64 // 序列
	lastTimestamp int64 // 上次时间戳
	mutex         *sync.Mutex
}

func (s *snowflake) GetID() int64 {
	i, _ := s.NextID()
	return i
}

// New 生成新的计算体
// workerID：机器ID标识
// dataCenterID 数据中心ID标识
func New(workerID, dataCenterID int64) (idgen.IDGenerator, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, ErrInvalidWorkerID
	}
	if dataCenterID < 0 || dataCenterID > maxDataCenterID {
		return nil, ErrInvalidDataCenterID
	}
	return &snowflake{workerID, dataCenterID, 0, -1, new(sync.Mutex)}, nil
}

// NextID 获取ID
func (s *snowflake) NextID() (int64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.nextID()
}

func (s *snowflake) nextID() (id int64, err error) {
	t := timeGen()
	// 当前时间戳小于上次，则系统时间发生过回退
	if t < s.lastTimestamp {
		return 0, ErrClockMoveBackwards
	}
	// 当前时间戳等于上次，则生成该毫秒内序列
	if t == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		// 序列溢出
		if s.sequence == 0 {
			// 等待下一毫秒
			t = nextMillisecond(s.lastTimestamp)
		}
	} else {
		// 时间戳改变，序列重置
		s.sequence = 0
	}
	s.lastTimestamp = t
	return ((t - startTimestamp) << timestampLeftShift) | (s.dataCenterID << dataCenterIDShift) | (s.workerID << workerIDShift) | s.sequence, nil
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

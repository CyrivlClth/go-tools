# Snowflake

雪花算法golang实现

## 说明

Twitter开源的雪花算法生成ID，每毫秒每个机器可生成4096个连续不重复ID。

## Benchmark

| Benchmark name          | COUNT   | ns/op     |
| ----------------------- | ------- | --------- |
| BenchmarkNewSnowflake-4 | 5000000 | 244 ns/op |

## 工具

从 [github.com/CyrivlClth/snowserver](https://github.com/CyrivlClth/snowserver.git) 可获取id生成器工作节点的服务。

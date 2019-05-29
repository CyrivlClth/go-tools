# snowflake
雪花算法golang实现
## Introduction
Twitter开源的雪花算法生成ID，每毫秒每个机器可生成4096个连续不重复ID。
## Benchmark
Name|Count|ops
:-|-:|:-:
BenchmarkNewSnowflake-8 |5000000	|243 ns/op
BenchmarkNewSnowflake_Lock-8 |5000000	| 244 ns/op

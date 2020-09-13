# Go Tools

[![Build Status](https://travis-ci.org/CyrivlClth/go-tools.svg?branch=master)](https://travis-ci.org/CyrivlClth/go-tools)

[English Doc](./README.md)

The awesome tool kit for Golang.

等待golang泛型

## 目录

- [Go Tools](#go-tools)

    - [目录](#目录)
    
    - [概览](#概览)
    
    - [开始](#开始)
    
    - [实例](#实例)
    
        - [ID生成器](#ID生成器)
        
            - [雪花算法](#雪花算法)
            
            - [序列号](#序列号)
            
        - [切片工具](#切片工具)
        
            - [Integer Slice](#integer-slice)
    
    - [贡献](#贡献)
    
    - [证书](#证书)

## 概览

- [x] ID生成器

    - [x] 雪花算法
    
    - [x] 序列号

- [ ] 数据结构

    - [ ] 切片

        - [x] 方法 `[]int`: `去重-Distinct`, `合并-Merge`, `排序-Sort`, `筛选-Filter`
        
        - [ ] 通用 `interface`

    - [ ] Map
    
    - [x] JSON对象

- [ ] Bean

    - [ ] 结构体复制



## 开始

```bash
go get -u github.com/CyrivlClth/go-tools
```

如果你只想使用其中部分工具，也可以采用以下方式:

```bash
go get -u github.com/CyrivlClth/go-tools/jsonobj
```

## 实例

### ID生成器

ID生成器包括:

- 雪花算法

- 序列号-达到最大值时归零

#### 雪花算法

```go
...
func main() {
    generator, _ := snowflake.New(0, 0)
    id, _ := generator.NextID()
    fmt.Println(id)
}
...
```

更多可参见 `github.com/CyrivlClth/go-tools/idgen/snowflake/cmd`

#### 序列号

```go
serialnum.New(9).GetID()
```

### 切片工具

#### Integer Slice

```go
package main

import (
 "github.com/CyrivlClth/go-tools/container/set"
 "github.com/CyrivlClth/go-tools/container/slice"
)

func main(){
    slice.Integer([]int{3, 2, 5, 2}).Distinct()
    // [3,2,5]
    slice.Integer([]int{3, 2, 5, 2}).Contains(2)
    // true
    slice.Integer([]int{3, 2, 5, 2}).Sort()
    // [2,2,3,5]
    slice.Integer([]int{3, 2, 5, 2}).Filter(func(x int) bool {return x%2==0})
    // [2,2]
    int1 := slice.Integer([]int{3, 2, 5, 2})
    int2 := slice.Integer([]int{4, 2, 5, 2})
    set.NewIntegerSet(int1).Merge(int2)
    // [3,2,5,4]
}

```

## 贡献

欢迎commit

## 证书

©CyrivlClth, 2019-time.Now()

Released under the [MIT License](./LICENSE)

# Go Tools

[![Build Status](https://api.travis-ci.org/CyrivlClth/go-tools.svg)](https://api.travis-ci.org/CyrivlClth/go-tools)

The awesome tool kit for Golang.

Waiting for generic type of golang.

## Contents

- [Go Tools](#go-tools)

    - [Contents](#contents)
    
    - [Overview](#overview)
    
    - [Getting Started](#getting-started)
    
    - [Examples](#examples)
    
        - [ID Generator](#id-generator)
        
            - [Snowflake](#snowflake-id-generator)
            
            - [Serial Number](#serial-number-generator)
            
        - [Slice Tools](#slice-tools)
        
            - [Integer Slice](#integer-slice)
    
    - [Contributing](#contributing)
    
    - [License](#license)

## Overview

- [x] ID generator

    - [x] snowflake
    
    - [x] serial no

- [ ] Slice tools

    - [x] type `[]int` method: `Distinct`, `Merge`, `Sort`, `Filter`

- [ ] Map tools



## Getting Started

```bash
go get -u github.com/CyrivlClth/go-tools
```

## Examples

### ID generator

ID generator contains:

- snowflake

- simple serial no

#### Snowflake ID Generator

```go
...
func main() {
    generator, _ := snowflake.New(0, 0)
    id, _ := generator.NextID()
    fmt.Println(id)
}
...
```

more detail for package `github.com/CyrivlClth/go-tools/idgen/snowflake/cmd`

#### Serial Number Generator

```go
serialnum.New(9).GetID()
```

### Slice Tools

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

## Contributing

You can just create a new issues and commit your code

## License

©CyrivlClth, 2019-time.Now()

Released under the [MIT License](./LICENSE)

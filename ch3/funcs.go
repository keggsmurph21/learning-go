package main

import "fmt"

type mytype int

func (p mytype) do_something(q int) (int, int) {
    return q, int(p)
}

func incr(n int) int {
    return n + 1
}

func invert(n int) (m int) {
    m = (-1) * n
    return
}

func rec(i int) {
    if i == 10 {
        return
    }
    rec(i+1)
    fmt.Println(i)
}

func main() {
    var a mytype = 1
    fmt.Println(a.do_something(2))
    fmt.Println(incr(5))
    fmt.Println(invert(10))
    rec(100)
}

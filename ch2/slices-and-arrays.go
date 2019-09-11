package main

import "fmt"

func main() {

    const n = 6
    const m = 9

    var arr [m]int
    slice := arr[:n]

    fmt.Println("arr len:", len(arr), " cap:", cap(arr))
    fmt.Println("slice len:", len(slice), " cap:", cap(slice))

    a := [...]int{1, 2, 3, 4, 5}
    s1 := a[2:4]
    fmt.Println(s1)
    s2 := a[1:5]
    fmt.Println(s2)
    s3 := a[:]
    fmt.Println(s3)
    s4 := a[:4]
    fmt.Println(s4)
    s5 := s2[:]
    fmt.Println(s5)
    s6 := a[2:4:5]
    fmt.Println(s6)

    // overrunning boundaries
    var b [100]int
    s := b[0:99]
    s[98] = 1
    // s[99] = 2 // panic: runtime error: index out of range
    fmt.Println(b)
    fmt.Println(s)

    // extending slices
    s7 := []int{0, 0}
    s8 := append(s7, 2)
    fmt.Println(s8)
    s9 := append(s8, 3, 5, 7)
    fmt.Println(s9)
    s10 := append(s9, s7...)
    fmt.Println(s10)

    var d = [...]int{0,1,2,3,4,5,6,7}
    var s11 = make([]int, 6)
    n1 := copy(s11, d[0:])
    fmt.Println(n1, s11)
    n2 := copy(s11, s11[2:])
    fmt.Println(n2, s11)

}

package main

import "fmt"

func main() {

    // 1.1
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    // 1.2
    j := 0
For:
    fmt.Println(j)
    j++
    if j < 10 {
        goto For
    }

    // 1.3
    var arr [10]int
    for k := 0; k < 10; k++ {
        arr[k] = k
    }
    fmt.Println(arr)

    // 1.4
    fmt.Println([...]int{0,1,2,3,4,5,6,7,8,9})

    // 2.1
    var farr [20]float64
    fslice := farr[:]
    var sum float64 = 0
    for i := 0; i < 20; i++ {
        sum += fslice[i]
    }
    avg := sum / 20
    fmt.Println(avg)

    // 2.1 (improved)
    var xs = make([]float64, 10)
    var _avg float64
    _sum := 0.0
    switch len(xs) {
    case 0:
        _avg = 0
    default:
        for _, v := range xs {
            _sum += v
        }
        _avg = _sum / float64(len(xs))
    }
    fmt.Println(_avg)

    // 3.1
    for i := 0; i < 100; i++ {
        var caught bool
        if i % 3 == 0 {
            fmt.Print("Fizz")
            caught = true
        }
        if i % 5 == 0 {
            fmt.Print("Buzz")
            caught = true
        }
        if ! caught {
            fmt.Print(i)
        }
        fmt.Print("\n")
    }

}

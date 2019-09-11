package main

import "fmt"

func main() {

    monthdays := map[string]int {
        "jan": 31, "feb": 28, "mar": 31,
        "apr": 30, "may": 31, "jun": 30,
        "jul": 31, "aug": 31, "sep": 30,
        "oct": 31, "nov": 30, "dec": 31,
    }
    fmt.Println(monthdays)
    fmt.Printf("%d\n", monthdays["dec"])

    // when just declaring, use make()
    // othermonthsdays := make(map[string]int)
    // fmt.Println(othermonthdays) // undefined

    year := 0
    for _, days := range monthdays {
        year += days
    }
    fmt.Printf("year has %d days\n", year)

    // add values
    monthdays["undecim"] = 30

    // overwrite values
    monthdays["feb"] = 29

    // test existence ("comma ok")
    val1, ok1 := monthdays["feb"]
    fmt.Println(val1, ok1)
    val2, ok2 := monthdays["neb"]
    fmt.Println(val2, ok2)

    // remove entries
    delete(monthdays, "jan")
    fmt.Println(monthdays)

}

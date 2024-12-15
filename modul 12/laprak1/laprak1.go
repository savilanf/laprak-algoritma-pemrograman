package main

import "fmt"

func main() {
    var usr, pass string
    fmt.Scan(&usr, &pass)

    i := 0
    for usr != "Admin" || pass != "Admin" {
        fmt.Scan(&usr, &pass)
        i++
    }
    fmt.Printf("%v percobaan gagal login", i)
}
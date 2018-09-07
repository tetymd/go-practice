package main

import (
    "fmt"
)

type Car struct {
    name string
    maker string
}

type Animal struct {
    name string
}

func main() {
    car := Car{}
    var cat Animal

    car.name = "86"
    car.maker = "TOYOTA"

    cat.name = "cat"

    fmt.Println(car.name)
    fmt.Println(car.maker)

    fmt.Println(cat.name)
}

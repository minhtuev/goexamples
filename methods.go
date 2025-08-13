package main

import "fmt"

type rect struct {
    width, height int
}

// Method with pointer receiver
func (r *rect) area() int {
    return r.width * r.height
}

// Method with value receiver
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

// Method with pointer receiver that modifies the struct
func (r *rect) scale(factor int) {
    r.width *= factor
    r.height *= factor
}

// Method with value receiver that returns a copy
func (r rect) double() rect {
    return rect{
        width:  r.width * 2,
        height: r.height * 2,
    }
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())

    // Using the scale method (pointer receiver modifies the original)
    r.scale(2)
    fmt.Println("scaled area:", r.area(), "scaled perim:", r.perim())

    // Using the double method (value receiver returns a new rect)
    r2 := r.double()
    fmt.Println("original after double():", r.area(), r.perim())
    fmt.Println("new doubled rect:", r2.area(), r2.perim())
}

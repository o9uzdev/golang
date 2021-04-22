package main

import "testing"

func TestArea(t *testing.T) {

    areaTests := []struct {
        name    string
        shape   shape
        hasArea float64
    }{
        {name: "Rectangle", shape: rect{width: 12, height: 6}, hasArea: 72.0},
        {name: "Circle", shape: circle{radius: 10}, hasArea: 314.1592653589793},
    }

    for _, tt := range areaTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
            }
        })

    }

}

func TestPerim(t *testing.T) {

    perimTests := []struct {
        name    string
        shape   shape
        hasArea float64
    }{
        {name: "Rectangle", shape: rect{width: 3, height: 4}, hasArea: 14},
        {name: "Circle", shape: circle{radius: 5}, hasArea: 31.41592653589793},
    }

    for _, tt := range perimTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.perim()
            if got != tt.hasArea {
                t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
            }
        })

    }

}
package main

import (
	"fmt"

	geom "github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy"
)

func difftest() {
	coords := []float64{1, 1, 1, 3}
	distance := xy.Distance(geom.Coord(coords[0:2]), geom.Coord(coords[2:4]))
	fmt.Println("first distance", distance)

	coords = []float64{1, 1, 2, 2}
	distance = xy.Distance(geom.Coord(coords[0:2]), geom.Coord(coords[2:4]))
	fmt.Println("second distance", distance)

	coords = []float64{1, 1, 4, 4}
	distance = xy.Distance(geom.Coord(coords[0:2]), geom.Coord(coords[2:4]))
	fmt.Println("third distance", distance)

	coords = []float64{8, 8, 6, 6}
	distance = xy.Distance(geom.Coord(coords[0:2]), geom.Coord(coords[2:4]))
	fmt.Println("fourth distance", distance)
}

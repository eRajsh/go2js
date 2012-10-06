// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"math"
)

var PASS = true

type Rectangle struct {
	width, height float64
}

func noMethod() {
	pass := true

	area := func(r Rectangle) float64 {
		return r.width * r.height
	}

	r1 := Rectangle{12, 2}

	if area(r1) != 24 {
		fmt.Printf("\tFAIL: area r1 => got %v, want 24)\n", area(r1))
		pass, PASS = false, false
	}
	if area(Rectangle{9, 4}) != 36 {
		fmt.Printf("\tFAIL: area Rectangle{9,4} => got %v, want 36)\n",
			area(Rectangle{9, 4}))
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

// * * *

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func method() {
	pass := true

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	tests := []struct {
		msg string
		in  float64
		out float64
	}{
		{"Rectangle{12,2}", r1.area(), 24},
		{"Rectangle{9,4}", r2.area(), 36},
		{"Circle{10}", c1.area(), 314.1592653589793},
		{"Circle{25}", c2.area(), 1963.4954084936207},
	}

	for _, t := range tests {
		if t.in != t.out {
			fmt.Printf("\tFAIL: %s => got %v, want %v\n", t.msg, t.in, t.out)
			pass, PASS = false, false
		}
	}
	if pass {
		fmt.Println("\tpass")
	}
}

// * * *

type sliceOfints []int
type agesByNames map[string]int

func (s sliceOfints) sum() int {
	sum := 0
	for _, value := range s {
		sum += value
	}
	return sum
}

func (people agesByNames) older() string {
	a := 0
	n := ""
	for key, value := range people {
		if value > a {
			a = value
			n = key
		}
	}
	return n
}

func withNamedType() {
	pass := true

	s := sliceOfints{1, 2, 3, 4, 5}
	folks := agesByNames{
		"Bob":   36,
		"Mike":  44,
		"Jane":  30,
		"Popey": 100,
	}

	if s.sum() != 15 {
		fmt.Printf("\tFAIL: s.sum => got %v, want 15)\n", s.sum())
		pass, PASS = false, false
	}
	if folks.older() != "Popey" {
		fmt.Printf("\tFAIL: folks.older => got %s, want Popey)\n",
			folks.older())
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func main() {
	fmt.Print("\n\n== Methods\n\n")

	fmt.Println("=== RUN noMethod")
	noMethod()
	fmt.Println("=== RUN method")
	method()
	fmt.Println("=== RUN withNamedType")
	withNamedType()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Methods")
	}
}

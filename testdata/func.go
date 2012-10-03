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

var x = 10

func init() {
	x = 13
}

func testInit() {
	if x == 13 {
		fmt.Println("\tpass")
	} else {
		fmt.Printf("\tFAIL: got %v, want 13\n", x)
		PASS = false
	}
}

func singleLine() { fmt.Println("\tpasssingleLine") }

func simpleFunc() {
	pass := true

	// Returns the maximum between two int a, and b.
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) // calling max(x, y)
	if max_xy != 4 {
		fmt.Printf("\tFAIL: max(x,y) => got %v, want 4)\n", max_xy)
		pass, PASS = false, false
	}

	max_xz := max(x, z) // calling max(x, z)
	if max_xz != 5 {
		fmt.Printf("\tFAIL: max(x,z) => got %v, want 5)\n", max_xz)
		pass, PASS = false, false
	}

	if max(y, z) != 5 { // just call it here
		fmt.Printf("\tFAIL: max(y,z) => got %v, want 5)\n", max(y, z))
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func twoOuputValues() {
	pass := true

	// Returns A+B and A*B in a single shot.
	SumAndProduct := func(A, B int) (int, int) {
		return A + B, A * B
	}

	x := 3
	y := 4
	xPLUSy, xTIMESy := SumAndProduct(x, y)

	if xPLUSy != 7 {
		fmt.Printf("\tFAIL: sum => got %v, want 7)\n", xPLUSy)
		pass, PASS = false, false
	}
	if xTIMESy != 12 {
		fmt.Printf("\tFAIL: product => got %v, want 12)\n", xTIMESy)
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func resultVariable() {
	pass := true

	// Returns a bool that is set to true when Sqrt is possible and false when not,
	// and the actual square root of a float64.
	MySqrt := func(f float64) (s float64, ok bool) {
		if f > 0 {
			s, ok = math.Sqrt(f), true
		}
		return s, ok
	}

	tests := map[float64]float64{
		1:  1,
		2:  1.4142135623730951,
		3:  1.7320508075688772,
		4:  2,
		5:  2.23606797749979,
		6:  2.449489742783178,
		7:  2.6457513110645907,
		8:  2.8284271247461903,
		9:  3,
		10: 3.1622776601683795,
	}

	for i := -2.0; i <= 10; i++ {
		sqroot, ok := MySqrt(i)
		if ok {
			if sqroot != tests[i] {
				fmt.Printf("\tFAIL: square(%v) => got %v, want %v\n",
					i, sqroot, tests[i])
				pass, PASS = false, false
			}
		} else {
			if i != -2.0 && i != -1.0 && i != 0 {
				fmt.Printf("\tFAIL: square(%v) => should no be run\n", i)
				pass, PASS = false, false
			}
		}
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func testReturn() {
	pass := true

	MySqrt := func(f float64) (squareroot float64, ok bool) {
		if f > 0 {
			squareroot, ok = math.Sqrt(f), true
		}
		return // Omitting the output named variables, but keeping the "return".
	}

	_, ok := MySqrt(5)
	if !ok {
		fmt.Printf("\tFAIL: MySqrt(5) => got %v, want %v\n", ok, !ok)
		pass, PASS = false, false
	}

	if _, ok := MySqrt(0); ok {
		fmt.Printf("\tFAIL: MySqrt(0) => got %v, want %v\n", ok, !ok)
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func testPanic() {
	panic("unreachable")
	panic(fmt.Sprintf("not implemented: %s", "foo"))
}

func main() {
	fmt.Print("\n\n== Functions\n")

	fmt.Print("\n=== RUN testInit\n")
	testInit()
	fmt.Print("\n=== RUN singleLine\n")
	singleLine()
	fmt.Print("\n=== RUN simpleFunc\n")
	simpleFunc()
	fmt.Print("\n=== RUN twoOuputValues\n")
	twoOuputValues()
	fmt.Print("\n=== RUN resultVariable\n")
	resultVariable()
	fmt.Print("\n=== RUN testReturn\n")
	testReturn()
	fmt.Print("\n=== RUN testPanic\n")
	testPanic()

	if PASS {
		fmt.Print("\nPASS\n")
	} else {
		fmt.Print("\nFAIL\n")
		print("Fail: Functions")
	}
}

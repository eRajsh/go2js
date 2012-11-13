// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "fmt"

var PASS = true

var (
	u1, u2 uint = 1, 2

	u   uint   = 1
	u_         = uint(1)
	u8  uint8  = 8
	u16 uint16 = 16
	u32 uint32 = 32

	i   int   = 1
	i_        = int(1)
	i8  int8  = 8
	i16 int16 = 16
	i32 int32 = 32

	f32  float32 = 3.2
	f32_         = float32(3.2)

	b  byte = 8
	b_      = byte(8)

	r  rune = 32
	r_      = rune(32)
)

//b3      = '1'
//b4      = 'a'
//r3      = '9'
//r4      = '€'

func value() {
	pass := true

	if u != 1 || u_ != 1 || u8 != 8 || u16 != 16 || u32 != 32 {
		fmt.Print("\tFAIL: uint\n")
		pass, PASS = false, false
	}
	if i != 1 || i_ != 1 || i8 != 8 || i16 != 16 || i32 != 32 {
		fmt.Print("\tFAIL: int\n")
		pass, PASS = false, false
	}
	if f32 != 3.2 || f32_ != 3.2 {
		fmt.Print("\tFAIL: float\n")
		pass, PASS = false, false
	}
	if b != 8 || b_ != 8 {
		fmt.Print("\tFAIL: byte\n")
		pass, PASS = false, false
	}
	if r != 32 || r_ != 32 {
		fmt.Print("\tFAIL: rune\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func add() {
	pass := true

	if u+1 != 2 || u_+1 != 2 || u8+1 != 9 || u16+1 != 17 || u32+1 != 33 {
		fmt.Print("\tFAIL: uint\n")
		pass, PASS = false, false
	}
	if i+1 != 2 || i_+1 != 2 || i8+1 != 9 || i16+1 != 17 || i32+1 != 33 {
		fmt.Print("\tFAIL: int\n")
		pass, PASS = false, false
	}
	if f32+1 != 4.2 || f32_+1 != 4.2 {
		fmt.Print("\tFAIL: float\n")
		pass, PASS = false, false
	}
	if b+1 != 9 || b_+1 != 9 {
		fmt.Print("\tFAIL: byte\n")
		pass, PASS = false, false
	}
	if r+1 != 33 || r_+1 != 33 {
		fmt.Print("\tFAIL: rune\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func subtract() {
	pass := true

	if u-1 != 0 || u_-1 != 0 || u8-1 != 7 || u16-1 != 15 || u32-1 != 31 {
		fmt.Print("\tFAIL: uint\n")
		pass, PASS = false, false
	}
	if i-1 != 0 || i_-1 != 0 || i8-1 != 7 || i16-1 != 15 || i32-1 != 31 {
		fmt.Print("\tFAIL: int\n")
		pass, PASS = false, false
	}
	if f32-1 != 2.2 || f32_-1 != 2.2 {
		fmt.Print("\tFAIL: float\n")
		pass, PASS = false, false
	}
	if b-1 != 7 || b_-1 != 7 {
		fmt.Print("\tFAIL: byte\n")
		pass, PASS = false, false
	}
	if r-1 != 31 || r_-1 != 31 {
		fmt.Print("\tFAIL: rune\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func multiplication() {
	pass := true

	if u*2 != 2 || u_*2 != 2 || u8*2 != 16 || u16*2 != 32 || u32*2 != 64 {
		fmt.Print("\tFAIL: uint\n")
		pass, PASS = false, false
	}
	if i*2 != 2 || i_*2 != 2 || i8*2 != 16 || i16*2 != 32 || i32*2 != 64 {
		fmt.Print("\tFAIL: int\n")
		pass, PASS = false, false
	}
	if f32*2 != 6.4 || f32_*2 != 6.4 {
		fmt.Print("\tFAIL: float\n")
		pass, PASS = false, false
	}
	if b*2 != 16 || b_*2 != 16 {
		fmt.Print("\tFAIL: byte\n")
		pass, PASS = false, false
	}
	if r*2 != 64 || r_*2 != 64 {
		fmt.Print("\tFAIL: rune\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func division() {
	pass := true

	if u/1 != 1 || u_/1 != 1 || u8/2 != 4 || u16/2 != 8 || u32/2 != 16 {
		fmt.Print("\tFAIL: uint\n")
		pass, PASS = false, false
	}
	if i/1 != 1 || i_/1 != 1 || i8/2 != 4 || i16/2 != 8 || i32/2 != 16 {
		fmt.Print("\tFAIL: int\n")
		pass, PASS = false, false
	}
	if f32/2 != 1.6 || f32_/2 != 1.6 {
		fmt.Print("\tFAIL: float\n")
		pass, PASS = false, false
	}
	if b/2 != 4 || b_/2 != 4 {
		fmt.Print("\tFAIL: byte\n")
		pass, PASS = false, false
	}
	if r/2 != 16 || r_/2 != 16 {
		fmt.Print("\tFAIL: rune\n")
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}
}

func main() {
	fmt.Print("\n\n== Numeric\n\n")

	fmt.Println("=== RUN value")
	value()
	fmt.Println("=== RUN add")
	add()
	fmt.Println("=== RUN subtract")
	subtract()
	fmt.Println("=== RUN multiplication")
	multiplication()
	fmt.Println("=== RUN division")
	division()

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: Numeric")
	}
}

// TODO: remove
var x, y = 14, 9
var (
	and = x & y
	or  = x | y
	xor = x ^ y
	not = !true

	lShift    = 9 << 2
	rShift    = 9 >> 2
	lShiftNeg = -9 << 2
	rShiftNeg = -9 >> 2
)

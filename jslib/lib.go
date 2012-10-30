// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

// Package g handles the features and Go types in JavaScript.

package g

// The specific type that it represents.
const (
	invalidT int = iota
	arrayT
	mapT
	sliceT
)

// == Array
//

func init() {
	// Use the toString() method when Array.isArray isn't implemented:
	// https://developer.mozilla.org/en/JavaScript/Reference/Global_Objects/Array/isArray#Compatibility
	if !Array.isArray {
		Array.isArray = func(arg interface{}) bool {
			return Object.prototype.toString.call(arg) == "[object Array]"
		}
	}
}

// The array can not be compared with nil.
// The capacity is the same than length.

// arrayType represents a fixed array type.
type arrayType struct {
	v []interface{} // array's value

	len_ map[int]int
}

// len returns the length for the given dimension.
func (a arrayType) len(dim int) int {
	if dim == nil {
		return a.len_[0]
	}
	return a.len_[len(arguments)]
}

// cap returns the capacity for the given dimension.
func (a arrayType) cap(dim int) int {
	if dim == nil {
		return a.len_[0]
	}
	return a.len_[len(arguments)]
}

// str returns the array (of bytes or runes) like a string.
func (a arrayType) str() string {
	return a.v.join("")
}

// typ returns the type.
func (a arrayType) typ() int { return arrayT }

// MkArray initializes an array of dimension "dim" to value "zero",
// merging the elements of "data" if any.
func MkArray(dim []int, zero interface{}, data []interface{}) *arrayType {
	a := new(arrayType)

	if data != nil {
		if !equalDim(dim, getDimArray(data)) {
			a.v = initArray(dim, zero)
			mergeArray(a.v, data)
		} else {
			a.v = data
		}
	} else {
		a.v = initArray(dim, zero)
	}

	for i, v := range dim {
		a.len_[i] = v
	}

	return a
}

// * * *

// equalDim reports whether d1 and d2 are equal.
func equalDim(d1, d2 []int) bool {
	if len(d1) != len(d2) {
		return false
	}
	for i, v := range d1 {
		if v != d2[i] {
			return false
		}
	}
	return true
}

// getDimArray returns the dimension of an array.
func getDimArray(a []interface{}) (dim []int) {
	for {
		dim.push(len(a))

		if Array.isArray(a[0]) {
			a = a[0]
		} else {
			break
		}
	}
	return
}

// initArray returns an array of dimension given in "dim" initialized to "zero".
func initArray(dim []int, zero interface{}) (a []interface{}) {
	if len(dim) == 0 {
		return zero
	}
	nextArray := initArray(dim.slice(1), zero)

	for i := 0; i < dim[0]; i++ {
		a[i] = nextArray
	}
	return
}

// mergeArray merges src in array dst.
func mergeArray(dst, src []interface{}) {
	for i, srcVal := range src {
		if Array.isArray(srcVal) {
			mergeArray(dst[i], srcVal)
		} else {
			isHashMap := false

			// The position is into a hash map, if any
			if typeof(srcVal) == "object" {
				for k, v := range srcVal {
					if srcVal.hasOwnProperty(k) { // identify a hashmap
						isHashMap = true
						i = k
						dst[i] = v
					}
				}
			}
			if !isHashMap {
				dst[i] = srcVal
			}
		}
	}
}

// == Slice
//

// sliceType represents a slice type.
type sliceType struct {
	arr interface{}   // the array where data is got or created from scratch using make
	v   []interface{} // elements appended

	low  int // indexes for the array
	high int
	len  int // total of elements
	cap  int

	nil_ bool // for variables declared like slices
}

// typ returns the type.
func (s sliceType) typ() int { return sliceT }

func (s sliceType) isNil() bool {
	if s.len != 0 || s.cap != 0 {
		return false
	}
	return s.nil_
}

// MkSlice initializes a slice with the zero value.
func MkSlice(zero interface{}, len, cap int) *sliceType {
	s := new(sliceType)

	if zero == nil {
		s.nil_ = true
		return s
	}

	arr := new(arrayType)
	arr.len_[0] = len
	for i := 0; i < len; i++ {
		arr.v[i] = zero
	}
	s.arr = arr

	if cap != nil {
		s.cap = cap
	} else {
		s.cap = len
	}

	s.len = len
	s.high = len
	return s
}

// Slice creates a new slice with the elements in "data".
func Slice(zero interface{}, data []interface{}) *sliceType {
	s := new(sliceType)

	if zero == nil {
		s.nil_ = true
		return s
	}

	arr := new(arrayType)
	for i, srcVal := range data {
		isHashMap := false

		// The position is into a hash map, if any
		if typeof(srcVal) == "object" {
			for k, v := range srcVal {
				if srcVal.hasOwnProperty(k) { // identify a hashmap
					isHashMap = true

					for i; i < k; i++ {
						arr.v[i] = zero
					}
					arr.v[i] = v
				}
			}
		}
		if !isHashMap {
			arr.v[i] = srcVal
		}
	}
	s.len = len(arr.v)
	arr.len_[0] = s.len
	s.arr = arr

	s.cap = s.len
	s.high = s.len
	return s
}

// SliceFrom creates a new slice from an array using the indexes low and high.
func SliceFrom(src interface{}, low, high int) *sliceType {
	s := new(sliceType)
	s.set(src, low, high)
	return s
}

// set sets the elements of a slice.
func (s sliceType) set(src interface{}, low, high int) { // TODO: remove
	if low != nil {
		s.low = low
	} else {
		s.low = 0
	}
	if high != nil {
		s.high = high
	} else {
		if src.arr != nil { // slice
			s.high = src.len
		} else {
			s.high = len(src.v)
		}
	}

	s.len = s.high - s.low

	if src.arr != nil { // slice
		s.arr = src.arr
		s.cap = src.cap - s.low
		s.low += src.low
		s.high += src.low
	} else { // array
		s.arr = src
		s.cap = src.cap() - s.low
	}
}

// setv sets a value.
/*func (s sliceType) setv(index int, v interface{}) {
	s.arr.v[index[0]+s.len] = v
}*/

// get gets the slice.
func (s sliceType) get() []interface{} {
	if s.arr != nil {
		arr := s.arr.v.slice(s.low, s.high)

		if len(s.v) != 0 {
			return arr.concat(s.v)
		} else {
			return arr
		}
	}
	return s.v
}

// str returns the slice (of bytes or runes) like a string.
func (s sliceType) str() string {
	_s := s.get()
	return _s.join("")
}

// * * *

// Append implements the function "append".
func Append(src []interface{}, elt ...interface{}) (dst sliceType) {
	// Copy src to the new slice
	dst.low = src.low
	dst.high = src.high
	dst.len = src.len
	dst.cap = src.cap
	dst.nil_ = src.nil_

	arr := new(arrayType)
	arr.len_[0] = src.arr.len_[0]
	for _, v := range src.arr.v {
		arr.v.push(v)
	}
	dst.arr = arr

	for _, v := range src.v {
		dst.v.push(v)
	}
	//==

	// TODO: handle len() in interfaces
	// lastIdxElt := len(elt) - 1

	for _, v := range elt {
		if /*i == lastIdxElt &&*/ Array.isArray(v) { // The last field could be an ellipsis
			for _, vArr := range v {
				dst.v.push(vArr)
				if dst.len == dst.cap {
					dst.cap = dst.len * 2
				}
				dst.len++
			}
			break
		}

		dst.v.push(v)
		if dst.len == dst.cap {
			dst.cap = dst.len * 2
		}
		dst.len++
	}
	return dst
}

// Copy implements the function "copy".
func Copy(dst []interface{}, src interface{}) (n int) {
	// []T to []T
	if src.arr != nil {
		for i := src.low; i < src.high; i++ {
			if n == dst.len {
				return
			}
			dst.arr.v[n] = src.arr.v[i]
			n++
		}
		for _, v := range src.v {
			if n == dst.len {
				return
			}
			dst.v.push(v)
			n++
		}
		return
	}

	// string to []byte
	for ; n < len(src); n++ {
		if n == dst.len {
			break
		}
		dst.arr.v[n] = src[n]
	}
	return
}

// == Map
//

// The length into a map is rarely used so, in JavaScript, I prefer to calculate
// the length instead of use a field.
//
// A map has not built-in function "cap".

// mapType represents a map type.
// The compiler adds the appropriate zero value for the map (which it is work out
// from the map type).
type mapType struct {
	v    map[interface{}]interface{} // map's value
	zero interface{}                 // zero value for the map's value
}

// len returns the number of elements.
func (m mapType) len() int {
	len := 0
	for key, _ := range m.v {
		if m.v.hasOwnProperty(key) {
			len++
		}
	}
	return len
}

// typ returns the type.
func (m mapType) typ() int { return mapT }

// Map creates a map storing its zero value.
func Map(zero interface{}, v map[interface{}]interface{}) *mapType {
	m := &mapType{v, zero}
	return m
}

// get returns the value for the key "k" if it exists and a boolean indicating it.
// If looking some key up in M's map gets you "nil" ("undefined" in JS),
// then return a copy of the zero value.
func (m mapType) get(k interface{}) (interface{}, bool) {
	v := m.v

	// Allow multi-dimensional index (separated by commas)
	for i := 0; i < len(arguments); i++ {
		v = v[arguments[i]]
	}

	if v == nil {
		return m.zero, false
	}
	return v, true
}

// == Utility
//

// Export adds public names from "exported" to the map "pkg".
func Export(pkg map[interface{}]interface{}, exported []interface{}) {
	for _, v := range exported {
		pkg.v = v
	}
}

/*func Len(v interface{}) {
	
}*/

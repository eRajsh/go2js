// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

// Package g handles the features and Go types in JavaScript.

package g

// The specific kind of type that it represents.
const (
	invalid uint = iota
	arrayKind
	mapKind
	sliceKind
)

func init() {
	// Use the toString() method when Array.isArray isn't implemented:
	// https://developer.mozilla.org/en/JavaScript/Reference/Global_Objects/Array/isArray#Compatibility
	if !Array.isArray {
		Array.isArray = func(arg interface{}) {
			return Object.prototype.toString.call(arg) == "[object Array]"
		}
	}
}

// == Array
//

// TODO 1: mergeArray could be integrated in initArray ?

// arrayType represents a fixed array type.
type arrayType struct {
	v []interface{} // array's value

	len uint
	cap uint
	// Note: the array in Go can not be compared with nil
}

// MkArray initializes an array of dimension "dim" to value "zero",
// merging the elements of "elem" if any.
func MkArray(dim []uint, zero interface{}, elem []interface{}) *arrayType {
	a := new(arrayType)

	if elem != nil {
		if !equalDim(dim, getDimArray(elem)) {
			a.v = initArray(dim, zero)
			mergeArray(a.v, elem)
		} else {
			a.v = elem
		}
	} else {
		a.v = initArray(dim, zero)
	}

	a.len = dim[0]
	a.cap = a.len
	return a
}

func (a arrayType) kind() uint { return arrayKind }

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

// TODO: could be done during translation, if it isn't possible TODO 1
// equalDim reports whether d1 and d2 are equal.
func equalDim(d1, d2 []uint) bool {
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

// TODO: could be done during translation, if it isn't possible TODO 1
// getDimArray returns the dimension of an array.
func getDimArray(a []interface{}) (dim []uint) {
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
func initArray(dim []uint, zero interface{}) (a []interface{}) {
	if len(dim) == 0 {
		return zero
	}
	nextArray := initArray(dim.slice(1), zero)

	for i := 0; i < dim[0]; i++ {
		a[i] = nextArray
	}
	return
}

/*
func initArray(dim []uint, zero interface{}) (a []interface{}) {
	if len(dim) > 1 {
		nextArray := initArray(dim.slice(1), zero)

		for i := 0; i < dim[0]; i++ {
			a[i] = nextArray
		}
	} else {
		for i := 0; i < dim[0]; i++ {
			a[i] = zero
		}
	}
	return
}*/

// == Slice
//

// sliceType represents a slice type.
type sliceType struct {
	array interface{}   // the array where this slice is got, if any
	elem  []interface{} // elements from scratch (make) or appended to the array

	low  uint // indexes for the array
	high uint
	len  uint // total of elements
	cap  uint

	isNil bool // for variables declared like slices
}

// NilSlice creates a null slice.
// For variables declared like slices.
func NilSlice() *sliceType {
	s := new(sliceType)
	s.isNil = true
	s.len, s.cap = 0, 0
	return s
}

// MkSlice initializes a slice with the zero value.
func MkSlice(zero interface{}, len, cap uint) *sliceType {
	s := new(sliceType)
	s.len = len

	for i := 0; i < len; i++ {
		s.elem[i] = zero
	}

	if cap != nil {
		s.cap = cap
	} else {
		s.cap = len
	}

	return s
}

// Slice creates a new slice with the elements in "elem".
func Slice(zero interface{}, elem []interface{}) *sliceType {
	s := new(sliceType)

	if len(arguments) == 0 {
		s.isNil = true
		return s
	}

	for i, srcVal := range elem {
		isHashMap := false

		// The position is into a hash map, if any
		if typeof(srcVal) == "object" {
			for k, v := range srcVal {
				if srcVal.hasOwnProperty(k) { // identify a hashmap
					isHashMap = true

					for i; i < k; i++ {
						s.elem[i] = zero
					}
					s.elem[i] = v
				}
			}
		}
		if !isHashMap {
			s.elem[i] = srcVal
		}
	}

	s.len = len(s.elem)
	s.cap = s.len
	return s
}

// SliceFrom creates a new slice from an array using the indexes low and high.
func SliceFrom(a interface{}, low, high uint) *sliceType {
	s := new(sliceType)

	s.array = a
	s.low = low
	s.high = high
	s.len = high - low
	s.cap = a.cap - low
	return s
}

// set sets the elements of a slice.
func (s sliceType) set(i interface{}, low, high uint) {
	s.low, s.high = low, high

	if i.elem != nil { // from make
		s.elem = i.elem.slice(low, high)
		s.cap = i.cap - low
		s.len = len(s.elem)

	} else { // from array
		s.array = i
		s.cap = len(i) - low
		s.len = high - low
	}
}

// get gets the slice.
func (s sliceType) get() []interface{} {
	if len(s.elem) != 0 {
		return s.elem
	}
	//a := s.array
	return s.array.slice(s.low, s.high)
}

// str returns the slice like a string.
func (s sliceType) str() string {
	_s := s.get()
	return _s.join("")
}

func (s sliceType) kind() uint { return sliceKind }

/*
func (s sliceType) setSlice(i interface{}, low, high uint) {
	s.low, s.high = low, high
	s.len = high - low

	if i.array != nil { // from slice
		s.array = i.array
		s.cap = i.cap - low
	} else { // from array
		s.array = i
		s.cap = len(i) - low
	}
}

// Appends an element to the slice.
func (s sliceType) append(elt interface{}) {
	if s.len == s.cap {
		s.cap = s.len * 2
	}
	s.len++
}
*/

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

	//len   uint
}

// Map creates a map storing its zero value.
func Map(zero interface{}, v map[interface{}]interface{}) *mapType {
	m := &mapType{v, zero}

	/*m := &mapType{v, zero, 0}
	for key, _ := range v {
		if v.hasOwnProperty(key) {
			m.len++
		}
	}*/
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

// If it were using the field length.
/*func (m mapType) set(k, v interface{}) {
	m.v[k] = v
	m.len++
}

func (m mapType) del(k interface{}) {
	delete(m, k)
	m.len--
}*/

// == Utility
//

// Export adds public names from "exported" to the map "pkg".
func Export(pkg map[interface{}]interface{}, exported []interface{}) {
	for _, v := range exported {
		pkg.v = v
	}
}

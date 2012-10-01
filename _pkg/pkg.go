// Copyright 2011  The "GoScript" Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Handle Go features.

package g

// Export adds public names from "exported" to the map "pkg".
func Export(pkg map[interface{}]interface{}, exported []interface{}) {
	for _, v := range exported {
		pkg.v = v
	}
}

// == Slice
//

// Slice represents a slice.
type Slice struct {
	array interface{}   // the array where this slice is got, if any
	elts  []interface{} // elements from scratch (make) or appended to the array

	low  uint // indexes for the array
	high uint
	len  uint // total of elements
	cap  uint

	isNil bool
}

// NilSlice creates a null slice.
// For variables declared like slices.
func NilSlice() *Slice {
	s := new(Slice)
	s.isNil = true
	s.len, s.cap = 0, 0
	return s
}

// * * *

// MakeSlice initializes a slice with the zero value.
func MakeSlice(zero interface{}, len, cap uint) *Slice {
	s := new(Slice)
	s.len = len

	for i := 0; i < len; i++ {
		s.elts[i] = zero
	}

	if cap != nil {
		s.cap = cap
	} else {
		s.cap = len
	}

	return s
}

// NewSlice creates a new slice.
func NewSlice(elts []interface{}) *Slice {
	s := new(Slice)
	s.elts = elts
	s.len = len(elts)
	s.cap = s.len
	return s
}

// NewSliceFrom creates a new slice from an array using the indexes low and high.
func NewSliceFrom(a interface{}, low, high uint) *Slice {
	s := new(Slice)
	s.array = a
	s.low = low
	s.high = high
	s.len = high - low
	s.cap = a.cap - low
	return s
}

// * * *

// get gets the slice.
func (s Slice) get() []interface{} {
	if len(s.elts) != 0 {
		return s.elts
	}
	//      a := s.array
	return s.array.slice(s.low, s.high)
}

// str returns the slice like a string.
func (s Slice) str() string {
	_s := s.get()
	return _s.join("")
}

/*
func (s Slice) setSlice(i interface{}, low, high uint) {
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

// Sets the slice from an .
func (s Slice) set(i interface{}, low, high uint) {


        if i.elts != nil { // from make
                s.elts = i.elts.slice(low, high)
                s.cap = i.cap - low
                s.len = len(s.elts)

        } else { // from array
                s.array = i
                s.cap = len(i) - low
                s.len = high - low
        }

        s.low, s.high = low, high
}

// Appends an element to the slice.
func (s Slice) append(elt interface{}) {
        if s.len == s.cap {
                s.cap = s.len * 2
        }
        s.len++
}
*/

// == Map
//

// M represents a map.
// The compiler adds the appropriate zero value for the map (which it is work out
// from the map type).
type Map struct {
	f    map[interface{}]interface{} // map
	zero interface{}                 // zero value for the map
	// TODO add "cap"
}

// Gets the value for the key "k".
// If looking some key up in M's map gets you "nil" ("undefined" in JS),
// then return a copy of the zero value.
func (m Map) get(k interface{}) (interface{}, bool) {
	v := m.f

	// Allow multi-dimensional index (separated by commas)
	for i := 0; i < len(arguments); i++ {
		v = v[arguments[i]]
	}

	if v == nil {
		return m.zero, false
	}
	return v, true
}

/*
function getType(obj){
        if(obj===null)return "[object Null]"; // special case
        return Object.prototype.toString.call(obj);
}

function isArray(o) {
        return Object.prototype.toString.call(o) === '[object Array]';
}

*/

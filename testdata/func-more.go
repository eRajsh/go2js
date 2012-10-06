// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import "fmt"

var PASS = true

// Our struct representing a person.
type person struct {
	name string
	age  int
}

// Returns true and the older person in a group of persons,
// or false and nil if the group is empty.
func getOlder(people ...person) (person, bool) { // Variadic function.
	if len(people) == 0 {
		return person{}, false
	}

	older := people[0] // The first one is the older for now.

	for _, value := range people {
		if value.age > older.age {
			older = value
		}
	}
	return older, true
}

func main() {
	pass := true

	fmt.Print("\n\n== More functions\n")

	var (
		ok    bool
		older person
	)

	// Declare some persons.
	paul := person{"Paul", 23}
	jim := person{"Jim", 24}
	sam := person{"Sam", 84}
	rob := person{"Rob", 54}
	karl := person{"Karl", 19}

	tests := []string{"Jim", "Sam", "Sam", "Karl"}

	older, _ = getOlder(paul, jim)
	if older.name != tests[0] {
		fmt.Printf("\tFAIL: (getOlder paul,jim) => got %v, want %v\n",
			older.name, tests[0])
		pass, PASS = false, false
	}

	older, _ = getOlder(paul, jim, sam)
	if older.name != tests[1] {
		fmt.Printf("\tFAIL: (getOlder paul,jim,sam) => got %v, want %v\n",
			older.name, tests[1])
		pass, PASS = false, false
	}

	older, _ = getOlder(paul, jim, sam, rob)
	if older.name != tests[2] {
		fmt.Printf("\tFAIL: (getOlder paul,jim,sam,rob) => got %v, want %v\n",
			older.name, tests[2])
		pass, PASS = false, false
	}

	older, _ = getOlder(karl)
	if older.name != tests[3] {
		fmt.Printf("\tFAIL: (getOlder karl) => got %v, want %v\n",
			older.name, tests[3])
		pass, PASS = false, false
	}

	// There is no older person in an empty group.
	older, ok = getOlder()
	if ok {
		fmt.Printf("\tFAIL: (getOlder) => got %v, want %v\n", ok, !ok)
		pass, PASS = false, false
	}

	if pass {
		fmt.Println("\tpass")
	}

	if PASS {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
		print("Fail: More functions")
	}
}

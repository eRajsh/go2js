// Copyright 2011 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
// If a copy of the MPL was not distributed with this file, You can obtain one at
// http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

// JavaScript library name.
const LIB_RESERVED_NAME = "g"

// Reserved Words in JavaScript.
//
// https://developer.mozilla.org/en/JavaScript/Reference/Reserved_Words
// http://golang.org/ref/spec#Keywords
var reserved = map[string]struct{}{
	/* These ones are reserved too in Go.
	"break":     void,
	"case":      void,
	"const":     void,
	"continue":  void,
	"default":   void,
	"delete":    void,
	"else":      void,
	"false":     void,
	"for":       void,
	"if":        void,
	"import":    void,
	"interface": void,
	"package":   void,
	"return":    void,
	"switch":    void,
	"true":      void,
	"var":       void,*/

	"catch":      void,
	"debugger":   void,
	"do":         void,
	"finally":    void,
	"function":   void,
	"in":         void,
	"instanceof": void,
	"new":        void,
	"this":       void,
	"throw":      void,
	"try":        void,
	"typeof":     void,
	"void":       void,
	"while":      void,
	"with":       void,

	"class":      void,
	"enum":       void,
	"export":     void,
	"extends":    void,
	"implements": void,
	"let":        void,
	"null":       void,
	"private":    void,
	"protected":  void,
	"public":     void,
	"static":     void,
	"super":      void,
	"yield":      void,
}

var validImport = []string{"fmt", "math", "rand"}

// Constants to transform.
var Constant = map[string]string{
	"math.E":      "Math.E",
	"math.Ln2":    "Math.LN2",
	"math.Log2E":  "Math.LOG2E",
	"math.Ln10":   "Math.LN10",
	"math.Log10E": "Math.LOG10E",
	"math.Pi":     "Math.PI",
	"math.Sqrt2":  "Math.SQRT2",
}

// Functions that can be transformed since JavaScript has an equivalent one.
var Function = map[string]string{
	"fmt.Print":   "document.write",
	"fmt.Println": "document.write",
	"fmt.Printf":  "document.write",
	"fmt.Sprint":  "",
	"fmt.Sprintf": "",

	"math.Abs":   "Math.abs",
	"math.Acos":  "Math.acos",
	"math.Asin":  "Math.asin",
	"math.Atan":  "Math.atan",
	"math.Atan2": "Math.atan2",
	"math.Ceil":  "Math.ceil",
	"math.Cos":   "Math.cos",
	"math.Exp":   "Math.exp",
	"math.Floor": "Math.floor",
	"math.Log":   "Math.log",
	"math.Max":   "Math.max",
	"math.Min":   "Math.min",
	"math.Pow":   "Math.pow",
	"math.Sin":   "Math.sin",
	"math.Sqrt":  "Math.sqrt",
	"math.Tan":   "Math.tan",
	// https://developer.mozilla.org/en/JavaScript/Reference/Global_Objects/Math/round
	//"math.":      "Math.round",

	"rand.Float32": "Math.random",
	"rand.Float64": "Math.random",
}

// Imports
//
// http://golang.org/doc/go_spec.html#Import_declarations
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/import
func (tr *transform) getImport(spec []ast.Spec) {

	// godoc go/ast ImportSpec
	//  Doc     *CommentGroup // associated documentation; or nil
	//  Name    *Ident        // local package name (including "."); or nil
	//  Path    *BasicLit     // import path
	//  Comment *CommentGroup // line comments; or nil
	//  EndPos  token.Pos     // end of spec (overrides Path.Pos if nonzero)
	for _, v := range spec {
		iSpec := v.(*ast.ImportSpec)
		path := strings.Replace(iSpec.Path.Value, "\"", "", -1)

		// Core library
		if !strings.Contains(path, ".") {
			found := false
			for _, v := range validImport {
				if v == path {
					found = true
					break
				}
			}

			if !found {
				tr.addError("%s: import from core library", path)
				continue
			}
		}

		//import objectName.*;
		//fmt.Println(iSpec.Name, pathDir)
	}
}

// Returns the arguments of a Go function, formatted for JS.
func (tr *transform) GetArgs(funcName string, args []ast.Expr) string {
	var jsArgs string

	switch funcName {
	case "print", "fmt.Print", "fmt.Sprint":
		jsArgs = tr.joinArgsPrint(args, false)
	case "println", "fmt.Println":
		jsArgs = tr.joinArgsPrint(args, true)
	case "fmt.Printf", "fmt.Sprintf":
		jsArgs = tr.joinArgsPrintf(args)
	default:
		for i, v := range args {
			if i != 0 {
				jsArgs += "," + SP
			}
			jsArgs += tr.getExpression(v).String()
		}
	}

	return jsArgs
}

//
// === Utility

// validIdent checks if the name is a reserved word in JavaScript, returning a
// safe name adding "_" at the end of the name.
// It checks also if the name is the JavaScript library.
func validIdent(name interface{}) string {
	name_ := fmt.Sprintf("%s", name)

	if _, ok := reserved[name_]; ok {
		return name_ + "_"
	}
	if !Bootstrap && name_ == LIB_RESERVED_NAME {
		return name_ + "_"
	}
	return name_
}

// Returns arguments of Print, Println.
func (tr *transform) joinArgsPrint(args []ast.Expr, addLine bool) string {
	var jsArgs string
	lenArgs := len(args) - 1

	// Appends a character.
	add := func(s, char string) string {
		if strings.HasSuffix(s, "\"") {
			s = s[:len(s)-1] + char + "\""
		} else {
			s += SP + "+" + SP + "\"" + char + "\""
		}
		return s
	}

	for i, v := range args {
		expr := tr.getExpression(v).String()

		if i != 0 {
			jsArgs += SP + "+" + SP + expr
		} else {
			jsArgs = expr
		}

		if addLine {
			if i == lenArgs {
				jsArgs = add(jsArgs, "<br>")
			} else {
				jsArgs = add(jsArgs, " ")
			}
		}
	}

	return jsArgs
}

// Matches verbs for "fmt.Printf"
// http://golang.org/pkg/fmt/
var (
	reVerb      = regexp.MustCompile(`%[+\-# 0]?[bcdefgopqstvxEGTUX]`)
	reVerbWidth = regexp.MustCompile(`%[0-9]+[.]?[0-9]*[bcdefgoqxEGUXsqxX]`)
)

const VERB = "<<V>>"

// Returns arguments of Printf.
func (tr *transform) joinArgsPrintf(args []ast.Expr) string {
	result := ""

	// === Format
	format := tr.getExpression(args[0]).String()

	format = strings.Replace(format, "%%", "%", -1) // literal percent sign
	format = reVerb.ReplaceAllString(format, VERB)

	if reVerbWidth.MatchString(format) {
		format = reVerbWidth.ReplaceAllString(format, VERB)
	}
	// ===

	values := strings.Split(format, VERB)

	for i, v := range args[1:] {
		if i != 0 {
			result += fmt.Sprintf("%s+%s", SP, SP+`"`)
		}
		result += fmt.Sprintf("%s+%s", values[i]+`"`+SP, SP+tr.getExpression(v).String())
	}
	result += fmt.Sprintf("%s+%s", SP, SP+`"`+values[len(values)-1])

	return result
}

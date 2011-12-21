// Copyright 2011  The "GoJscript" Authors
//
// Use of this source code is governed by the BSD 2-Clause License
// that can be found in the LICENSE file.
//
// This software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
// OR CONDITIONS OF ANY KIND, either express or implied. See the License
// for more details.

package gojs

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"
)

// Imports
//
// http://golang.org/doc/go_spec.html#Import_declarations
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/import
func (tr *transform) getImport(spec []ast.Spec) {

	// http://golang.org/pkg/go/ast/#ImportSpec || godoc go/ast ImportSpec
	//  Doc     *CommentGroup // associated documentation; or nil
	//  Name    *Ident        // local package name (including "."); or nil
	//  Path    *BasicLit     // import path
	//  Comment *CommentGroup // line comments; or nil
	//  EndPos  token.Pos     // end of spec (overrides Path.Pos if nonzero)
	for _, v := range spec {
		iSpec := v.(*ast.ImportSpec)
		path := iSpec.Path.Value
		pathDir := strings.SplitN(path, "/", 2)[0]

		if !strings.Contains(pathDir, ".") {
			tr.addError("%s: import from core library", path)
			continue
		}

		//import objectName.*; 
		//fmt.Println(iSpec.Name, pathDir)
	}
}

// Constants
//
// http://golang.org/doc/go_spec.html#Constant_declarations
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/const
func (tr *transform) getConst(spec []ast.Spec) {
	iotaExpr := make([]string, 0) // iota expressions

	// http://golang.org/pkg/go/ast/#ValueSpec || godoc go/ast ValueSpec
	//  Doc     *CommentGroup // associated documentation; or nil
	//  Names   []*Ident      // value names (len(Names) > 0)
	//  Type    Expr          // value type; or nil
	//  Values  []Expr        // initial values; or nil
	//  Comment *CommentGroup // line comments; or nil
	for _, s := range spec {
		vSpec := s.(*ast.ValueSpec)

		// Checking
		if err := tr.CheckAndAddError(vSpec.Type); err != nil {
			continue
		}

		tr.addLine(vSpec.Pos())
		isFirst := true

		for i, ident := range vSpec.Names {
			if ident.Name == "_" {
				iotaExpr = append(iotaExpr, "")
				continue
			}

			value := strconv.Itoa(ident.Obj.Data.(int)) // possible value of iota

			if vSpec.Values != nil {
				v := vSpec.Values[i]

				// Checking
				tr.CheckAndAddError(v)
				if tr.hasError {
					continue
				}

				expr := newExpression("")
				expr.transform(v)
				exprStr := expr.String()

				if expr.useIota {
					value = strings.Replace(exprStr, IOTA, value, -1)
					iotaExpr = append(iotaExpr, exprStr)
				} else {
					value = exprStr
				}
			} else {
				if tr.hasError {
					continue
				}
				value = strings.Replace(iotaExpr[i], IOTA, value, -1)
			}

			tr.addIfExported(ident)

			// === Write
			if isFirst {
				isFirst = false
				tr.WriteString(fmt.Sprintf(
					"const %s%s=%s%s", ident.Name, SP, SP, value))
			} else {
				tr.WriteString(fmt.Sprintf(
					",%s%s%s=%s%s", SP, ident.Name, SP, SP, value))
			}
		}

		// It is possible that there is only a blank identifier
		if !isFirst {
			tr.WriteString(";")
		}
	}
}

// Variables
//
// http://golang.org/doc/go_spec.html#Variable_declarations
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/var
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/let
//
// TODO: use let for local variables
func (tr *transform) getVar(spec []ast.Spec) {
	// http://golang.org/pkg/go/ast/#ValueSpec || godoc go/ast ValueSpec
	for _, s := range spec {
		vSpec := s.(*ast.ValueSpec)

		// Checking
		if err := tr.CheckAndAddError(vSpec.Type); err != nil {
			continue
		}

		tr.addLine(vSpec.Pos())
		isFirst := true

		for i, ident := range vSpec.Names {
			var exprStr string
			var skip bool

			if ident.Name == "_" {
				skip = true
			}

			if vSpec.Values != nil {
				value := vSpec.Values[i]

				// Skip when it is not a function because it could return more
				// than one value.
				if ident.Name == "_" {
					if _, ok := value.(*ast.CallExpr); !ok {
						continue
					}
				}

				// Checking
				tr.CheckAndAddError(value)
				if tr.hasError {
					continue
				}

				src := newExpression(ident.Name)
				src.transform(value)
				exprStr = src.String()  // getExpression(value)

			} else if skip || tr.hasError {
				continue
			}

			tr.addIfExported(ident)

			// === Write
			// TODO: calculate expression using "exp/types"
			if isFirst {
				isFirst = false
				tr.WriteString("var " + ident.Name)
			} else {
				tr.WriteString("," + SP + ident.Name)
			}

			if exprStr != "" && exprStr != EMPTY {
				tr.WriteString(SP + "=" + SP + exprStr)
			}
		}

		last := tr.Bytes()[tr.Len()-1] // last character

		if last != '}' && last != ';' {
			tr.WriteString(";")
		}
	}
}

// Types
//
// http://golang.org/doc/go_spec.html#Type_declarations
func (tr *transform) getType(spec []ast.Spec) {
	// Format fields
	format := func(fields []string) (args, allFields string) {
		for i, f := range fields {
			if i == 0 {
				args = f
			} else {
				args += "," + SP + f
				allFields += SP
			}

			allFields += fmt.Sprintf("this.%s=%s;", f, f)
		}
		return
	}

	// http://golang.org/pkg/go/ast/#TypeSpec || godoc go/ast TypeSpec
	//  Doc     *CommentGroup // associated documentation; or nil
	//  Name    *Ident        // type name
	//  Type    Expr          // *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes
	//  Comment *CommentGroup // line comments; or nil
	for _, s := range spec {
		tSpec := s.(*ast.TypeSpec)
		fields := make([]string, 0) // names of fields
		//!anonField := make([]bool, 0) // anonymous field

		// Checking
		if err := tr.CheckAndAddError(tSpec.Type); err != nil {
			continue
		}

		switch typ := tSpec.Type.(type) {
		default:
			panic(fmt.Sprintf("unimplemented: %T", typ))

		// http://golang.org/pkg/go/ast/#Ident || godoc go/ast Ident
		//  NamePos token.Pos // identifier position
		//  Name    string    // identifier name
		//  Obj     *Object   // denoted object; or nil
		case *ast.Ident:

		// http://golang.org/pkg/go/ast/#StructType || godoc go/ast StructType
		//  Struct     token.Pos  // position of "struct" keyword
		//  Fields     *FieldList // list of field declarations
		//  Incomplete bool       // true if (source) fields are missing in the Fields list
		case *ast.StructType:
			if typ.Incomplete {
				panic("list of fields incomplete ???")
			}

			// http://golang.org/pkg/go/ast/#FieldList || godoc go/ast FieldList
			//  Opening token.Pos // position of opening parenthesis/brace, if any
			//  List    []*Field  // field list; or nil
			//  Closing token.Pos // position of closing parenthesis/brace, if any
			for _, field := range typ.Fields.List {
				if _, ok := field.Type.(*ast.FuncType); ok {
					tr.addError("%s: function type in struct",
						tr.fset.Position(field.Pos()))
					continue
				}

				// http://golang.org/pkg/go/ast/#Field || godoc go/ast Field
				//  Doc     *CommentGroup // associated documentation; or nil
				//  Names   []*Ident      // field/method/parameter names; or nil if anonymous field
				//  Type    Expr          // field/method/parameter type
				//  Tag     *BasicLit     // field tag; or nil
				//  Comment *CommentGroup // line comments; or nil

				// Checking
				if err := tr.CheckAndAddError(field.Type); err != nil {
					continue
				}
				if field.Names == nil {
					tr.addError("%s: anonymous field in struct",
						tr.fset.Position(field.Pos()))
					continue
				}

				for _, n := range field.Names {
					name := n.Name

					if name == "_" {
						continue
					}

					fields = append(fields, name)
					//!anonField = append(anonField, false)
				}
			}
		}

		if tr.hasError {
			continue
		}

		// === Write
		args, allFields := format(fields)

		tr.addIfExported(tSpec.Name)
		tr.addLine(tSpec.Pos())

		tr.WriteString(fmt.Sprintf("function %s(%s)%s{", tSpec.Name, args, SP))

		if len(allFields) != 0 {
			tr.WriteString(allFields)
			tr.WriteString("}")
		} else {
			tr.WriteString("}") //! empty struct
		}
	}
}

// Functions
//
// http://golang.org/doc/go_spec.html#Function_declarations
// https://developer.mozilla.org/en/JavaScript/Reference/Statements/function
func (tr *transform) getFunc(decl *ast.FuncDecl) {
	// http://golang.org/pkg/go/ast/#FuncDecl || godoc go/ast FuncDecl
	//  Doc  *CommentGroup // associated documentation; or nil
	//  Recv *FieldList    // receiver (methods); or nil (functions)
	//  Name *Ident        // function/method name
	//  Type *FuncType     // position of Func keyword, parameters and results
	//  Body *BlockStmt    // function body; or nil (forward declaration)

	// Check empty functions
	if len(decl.Body.List) == 0 {
		return
	}

	tr.addLine(decl.Pos())
	tr.WriteString(fmt.Sprintf("function %s(%s)%s",
		decl.Name, getParams(decl.Type), SP))
	tr.getStatement(decl.Body)
	tr.addIfExported(decl.Name)
}

//
// === Utility

// Gets the parameters.
//
// http://golang.org/pkg/go/ast/#FuncType || godoc go/ast FuncType
//  Func    token.Pos  // position of "func" keyword
//  Params  *FieldList // (incoming) parameters; or nil
//  Results *FieldList // (outgoing) results; or nil
func getParams(f *ast.FuncType) string {
	s := ""

	for i, v := range f.Params.List {
		if i != 0 {
			s += "," + SP
		}
		s += v.Names[0].Name
	}

	return s
}
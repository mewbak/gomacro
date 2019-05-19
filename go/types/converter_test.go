// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"go/token"
	"go/types"
	"testing"
)

func gmkparam(t types.Type) *types.Var {
	return types.NewParam(token.NoPos, nil, "", t)
}

func gmktuple(ts ...types.Type) *types.Tuple {
	vs := make([]*types.Var, len(ts))
	for i := range ts {
		vs[i] = gmkparam(ts[i])
	}
	return types.NewTuple(vs...)
}

func gmksignature(params *types.Tuple, results *types.Tuple) *types.Signature {
	return types.NewSignature(nil, params, results, false)
}

func TestConverterType(t *testing.T) {
	pos := token.NoPos
	gerr := types.Universe.Lookup("error").Type()
	gpkg := types.NewPackage("time", "time")
	gint := types.Typ[types.Int]
	named := types.NewNamed(
		types.NewTypeName(pos, gpkg, "Duration", nil),
		gint,
		nil,
	)
	gs := []types.Type{
		gerr,
		gerr.Underlying(),
		types.NewStruct(
			[]*types.Var{
				types.NewField(pos, nil, "foo", types.Typ[types.String], false),
				types.NewField(pos, nil, "int", gint, true),
			},
			[]string{
				`json:"foo"`,
				`json:"bar"`,
			}),
		named,
		named.Underlying(),
		gmksignature(gmktuple(named), gmktuple(types.Typ[types.Bool])),
		types.NewMap(gint, named),
		types.NewPointer(named),
		types.NewChan(types.RecvOnly, named),
	}
	var c Converter
	for _, g := range gs {
		typ := c.Type(g)
		s1, s2 := typ.String(), g.String()
		if s1 != s2 {
			t.Errorf("conversion mismatch: got %s expecting %s", s1, s2)
		}
	}
}

func TestConverterPackage(t *testing.T) {
	gpkg := types.NewPackage("", "")
	{
		ginscope := types.Universe
		goutscope := gpkg.Scope()
		for _, name := range ginscope.Names() {
			obj := ginscope.Lookup(name)
			goutscope.Insert(obj)
		}
	}
	var c Converter
	pkg := c.Package(gpkg)
	outnames := pkg.Scope().Names()
	names := []string{
		"bool", "byte", "complex128", "complex64", "error", "false", "float32", "float64",
		"int", "int16", "int32", "int64", "int8", "iota", "rune", "string", "true",
		"uint", "uint16", "uint32", "uint64", "uint8", "uintptr",
	}
	if len(outnames) != len(names) {
		t.Errorf("scope contains %d names, expecting %d", len(outnames), len(names))
	}
	for i, name := range names {
		if name != outnames[i] {
			t.Errorf("scope contains %q, expecting %q", outnames[i], name)
		}
	}
}
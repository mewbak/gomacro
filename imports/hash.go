// this file was generated by gomacro command: import "hash"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "hash"
	. "reflect"
)

func Package_hash() (map[string]Value, map[string]Type) {
	return map[string]Value{}, map[string]Type{
			"Hash":   TypeOf((*pkg.Hash)(nil)).Elem(),
			"Hash32": TypeOf((*pkg.Hash32)(nil)).Elem(),
			"Hash64": TypeOf((*pkg.Hash64)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_hash()
	Binds["hash"] = binds
	Types["hash"] = types
}
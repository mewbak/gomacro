// this file was generated by gomacro command: import "hash/adler32"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "hash/adler32"
	. "reflect"
)

func Package_hash_adler32() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Checksum": ValueOf(pkg.Checksum),
			"New":      ValueOf(pkg.New),
			"Size":     ValueOf(pkg.Size),
		}, map[string]Type{}
}

func init() {
	binds, types := Package_hash_adler32()
	Binds["hash/adler32"] = binds
	Types["hash/adler32"] = types
}
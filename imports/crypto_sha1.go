// this file was generated by gomacro command: import "crypto/sha1"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "crypto/sha1"
	. "reflect"
)

func Package_crypto_sha1() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"BlockSize": ValueOf(pkg.BlockSize),
			"New":       ValueOf(pkg.New),
			"Size":      ValueOf(pkg.Size),
			"Sum":       ValueOf(pkg.Sum),
		}, map[string]Type{}
}

func init() {
	binds, types := Package_crypto_sha1()
	Binds["crypto/sha1"] = binds
	Types["crypto/sha1"] = types
}
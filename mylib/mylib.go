package mylib

import "golang.org/x/xerrors"

// Hello for debug output.
func Hello() error {
	return xerrors.New("hello")
}

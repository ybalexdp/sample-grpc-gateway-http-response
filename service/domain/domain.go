package domain

import "golang.org/x/xerrors"

var Errors = struct {
	SampleNotFound error
}{
	SampleNotFound: xerrors.New("sample.not_found"),
}

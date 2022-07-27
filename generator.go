package meta

import "io"

type PkgGen interface {
	Bytes() ([]byte, error)
	Write(writer io.Writer) error
	Print() error
	Generate() error
}

type PkgGenFactory interface {
	Create(absPath, relPath string) (PkgGen, error)
}

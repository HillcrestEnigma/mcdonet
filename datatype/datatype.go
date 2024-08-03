package datatype

import "io"

type reader interface {
	io.Reader
	io.ByteReader
}

type writer interface {
	io.Writer
	io.ByteWriter
}

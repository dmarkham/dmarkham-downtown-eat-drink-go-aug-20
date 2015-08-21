package hexseeker

import (
	"io"
)

type Hexer struct{ io.ReadSeeker }

func (h *Hexer) Read(data []byte) (n int, err error) {
	return h.Read(data)
}

func (h *Hexer) Seek(offset int64, whence int) (int64, error) {
	return h.Seek(offset, whence)
}

func NewHexer(rws io.ReadSeeker) io.ReadSeeker {
	return &Hexer{rws}
}

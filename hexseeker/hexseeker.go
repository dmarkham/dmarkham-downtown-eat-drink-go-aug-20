package hexseeker

import (
	"io"
)

type Hexer struct{ io.ReadSeeker }

func NewHexer(rws io.ReadSeeker) io.ReadSeeker {
       return &Hexer{rws}
}

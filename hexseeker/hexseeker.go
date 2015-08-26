package hexseeker

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

type Hexer struct{ io.ReadSeeker }

func NewHexer(rws io.ReadSeeker) *Hexer {
	return &Hexer{rws}
}

func (h *Hexer) FmtAt(offset, n int64) string {

	grab := int64(16)

	_, err := h.Seek(offset, 0)
	if err != nil {
		return ""
	}
	leftover := n % grab
	var buf = make([]byte, n+leftover)
	_, err = h.Read(buf)

	if err != nil {
		return ""
	}

	st := ""
	i := int64(0)
	for i = int64(0); i < int64(n/grab); i++ {
		stT := hex.Dump(buf[i*grab : i*grab+16])
		os := fmt.Sprintf("%08x", offset+i*grab)
		stT = strings.Replace(stT, "00000000", os, 1)
		st = st + stT
	}
	if leftover > 0 {
		stT := hex.Dump(buf[n-leftover : n])
		os := fmt.Sprintf("%08x", offset+n-leftover)
		stT = strings.Replace(stT, "00000000", os, 1)
		st = st + stT
	}
	return st
}

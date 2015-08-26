package hexseeker

import (
	"bytes"
	"fmt"
	"testing"
	"testing/quick"
)

func TestFmtAt(t *testing.T) {

	f := func(v1 []byte, n uint16, offset uint16) bool {
		hexer := NewHexer(bytes.NewReader(v1))
		n = n % 50
		offset = offset % 50
		if int64(len(v1)) < int64(offset)+int64(n) {
			return true
		}
		if offset < 0 || n <= 0 {
			return true
		}
		y := hexer.FmtAt(int64(offset), int64(n))
		fmt.Println(y)
		if y == "" {
			return false
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

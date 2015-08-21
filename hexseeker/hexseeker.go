package hexseeker

import (
	"io"
  "encoding/hex"
)

type Hexer struct{ io.ReadSeeker }

func NewHexer(rws io.ReadSeeker) *Hexer {
       return &Hexer{rws}
}

func (h *Hexer) FmtAt(offset, n int64) (string){
  _, err :=   h.Seek(offset,0)
  if err != nil{
    return ""
  }
  var buf = make([]byte, n)
  _ , err = h.Read(buf)  
  
  if err != nil{
    return ""
  }
  st := hex.Dump(buf)
  return st
}


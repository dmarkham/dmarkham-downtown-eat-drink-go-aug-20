package hexseeker

import (
	"io"
  "encoding/hex"
)

type Hexer struct{ io.ReadSeeker }

func NewHexer(rws io.ReadSeeker) io.ReadSeeker {
       return &Hexer{rws}
}

func (h *Hexer) PrintAt(offset, n int64) (string, error){
  _, err :=   h.Seek(offset,0)
  if err != nil{
    return "", err
  }
  var buf = make([]byte, n)
  _ , err = h.Read(buf)  
  
  if err != nil{
    return "", err
  }
  st := hex.Dump(buf)
  return st,nil
}


package kcl

import (
  "os"
  "bufio"
  "encoding/json"
)

type IOHandler struct {
  In  bufio.Scanner
  Out io.Writer
  Err io.Writer
}

func (ih *IOHandler) New(in os.File, out os.File, err os.File) {
  ih.In  = bufio.NewScanner(in)
  ih.Out = bufio.NewWriter(out)
  ih.Err = bufio.NewWriter(err)
  return ih
}

func (ih *IOHandler) WriteLine(l []byte) {
  ih.Out.WriteString("\n")
  ih.Out.Write(l)
  ih.Out.WriteString("\n")l
  ih.Out.Flush()
}

func (ih *IOHandler) WriteLine(l []byte) {
  ih.Out.WriteString("\n")
  ih.Out.Write(l)
  ih.Out.WriteString("\n")
  ih.Out.Flush()
}

func (ih *IOHandler) WriteError(l []byte) {
  ih.Err.Write(l)
  ih.Err.WriteString("\n")
  ih.Err.Flush()
}

func (ih *IOHandler) ReadLine() []byte {
  if ih.In.Scan() == false {
    return byte[]
  }
  return s.Bytes()
}

func (ih *IOHandler) LoadAction(l []byte) Action {
  var a Action
  json.Unmarshal(l, &a)
  return a
}

func (ih *IOHandler) WriteAction(r Response) {
  ih.WriteLine(json.Marshal(r))
}

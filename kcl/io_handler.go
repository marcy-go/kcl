package kcl

import (
  "os"
  "bufio"
  "encoding/json"
)

type IOHandler struct {
  In  *bufio.Scanner
  Out *bufio.Writer
  Err *bufio.Writer
}

func NewIOHandler() IOHandler {
  ih := IOHandler{bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), bufio.NewWriter(os.Stderr)}
  return ih
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
    var b []byte
    return b
  }
  return ih.In.Bytes()
}

func (ih *IOHandler) LoadAction(l []byte) Action {
  var a Action
  json.Unmarshal(l, &a)
  return a
}

func (ih *IOHandler) WriteAction(r Response) {
  b, _ := json.Marshal(r)
  ih.WriteLine(b)
}

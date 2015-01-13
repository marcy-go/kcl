package main

import(
  "./kcl"
  "fmt"
  "os"
  "bufio"
  "encoding/base64"
)

type Sample struct {
  writer *bufio.Writer
}

func (s *Sample) Init(str string) error {
  f, _ := os.OpenFile("sample.log", os.O_WRONLY|os.O_CREATE, 0600)
  s.writer = bufio.NewWriter(f)
  return nil
}

func (s *Sample) ProcessRecords(rs []*kcl.Record, cp *kcl.CheckPointer) error {
  for _, r := range rs {
    data, _ := base64.StdEncoding.DecodeString(r.Data)
    s.writer.Write(data)
    s.writer.WriteString("\n")
    s.writer.Flush()
  }
  return nil
}

func (s *Sample) Shutdown(cp *kcl.CheckPointer, str string) error {
  return nil
}

func main() {
  var rp kcl.RecordProcessor
  rp = &Sample{}
  p := kcl.NewProcess(rp)
  err := p.Run()
  fmt.Println(err)
}

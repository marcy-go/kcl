package kcl

import (
  "os"
  "errors"
)

ih = IOHandler.New(os.Stdin, os.Stdout, os.Stderr)

type Process struct {
  rp RecordProcessor
  cp CheckPointer
  ih := ih
}

func Run(p *Process) error {
  p.cp = CheckPointer{&p.ih}
  for {
    l = p.ih.ReadLine()
    if len(l) > 0 {
      handleLine(l)
    } else {
      return errors.New("独自のエラーです。")
    }
  }
}

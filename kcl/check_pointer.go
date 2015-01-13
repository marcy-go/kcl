package kcl

import (
  "fmt"
)

type CheckPointError struct {
  Message string
}

func (e *CheckPointError) Error() string {
  return fmt.Sprintf("%s", e.Message)
}

type CheckPointer struct {
  IOHandler *IOHandler
}

func (cp *CheckPointer) getAction() Action {
  var a Action
  var l []byte
  for a.IsEmpty() {
    l = cp.IOHandler.ReadLine()
    a = cp.IOHandler.LoadAction(l)
  }
  return a
}

func (cp *CheckPointer) Run(seq string) error {
  res := Response{"checkpoint", seq}
  cp.IOHandler.WriteAction(res)
  a := cp.getAction()
  if a.Type != "checkpoint" {
    return &CheckPointError{"Invalid Action"}
  } else if a.IsError() {
    return &CheckPointError{a.Error}
  }
  return nil
}

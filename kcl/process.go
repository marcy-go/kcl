package kcl

import (
  "fmt"
  "os"
  "errors"
  "encoding/json"
)

type MalformedAction struct {
  Action string
  Key string
}

func (e *MalformedAction) Error() string {
  if e.Key == "" {
    err := fmt.Sprintf("Received an action which couldn't be understood. Action was %s", e.Action)
  } else {
    err := fmt.Sprintf("Action %s was expected to have key %s", e.Action, e.Key)
  }
  return err
}

ih := IOHandler.New(os.Stdin, os.Stdout, os.Stderr)

type Process struct {
  rp RecordProcessor
  ih := &ih
  cp := CheckPointer{&ih}
}

func (p *Process) Run() error {
  for {
    l = p.ih.ReadLine()
    if len(l) > 0 {
      err := p.handleLine(l)
      if  err != nil {
        return err
      }
    } else {
      return errors.New("No Input Error")
    }
  }
}

func (p *Process) handleLine(l string) error {
  a := p.ih.LoadAction(l)
  err := p.performAction(&a)
  if err != nil {
    return err
  }
  p.reportDone(a.Type)
  return nil
}

func (p *Process) performAction(a *Action) error {
  err := nil
  if a.Type == "" {
    err = &MalformedAction{"unknown", "action"}
  } else if a.Type == "initialize" {
    if a.ShardId == "" {
      err = &MalformedAction{a.Type, "shardId"}
    } else {
      err = p.rp.Init(a.ShardId)
    }
  } else if a.Type == "processRecords" {
    if len(a.Records) == 0 {
      err = &MalformedAction{a.Type, "records"}
    } else {
      err = p.rp.ProcessRecords(a.&Records, p.&cp)
    }
  } else if a.Type == "shutdown" {
    if a.Reason == "" {
      err = &MalformedAction{a.Type, "reason"}
    } else {
      err = p.rp.Shutdown(p.&cp, a.Reason)
    }
  } else {
    err = &MalformedAction{a.Type}
  }
  if err != nil {
    var buf byte[]
    runtime.Stack(buf, true)
    p.ih.WriteError(buf)
    return err
  }
  return nil
}

func (p *Process) reportDone(r string) {
  res := Response{"status", r}
  p.ih.WriteAction(res)
}

package kcl

import (
  "fmt"
  "errors"
  "runtime"
)

type MalformedAction struct {
  Action string
  Key string
}

func (e *MalformedAction) Error() string {
  var err string
  if e.Key == "" {
    err = fmt.Sprintf("Received an action which couldn't be understood. Action was %s", e.Action)
  } else {
    err = fmt.Sprintf("Action %s was expected to have key %s", e.Action, e.Key)
  }
  return err
}

type Process struct {
  Rp RecordProcessor
  Ih *IOHandler
  Cp *CheckPointer
}

func (p *Process) handleLine(l []byte) error {
  a := p.Ih.LoadAction(l)
  err := p.performAction(&a)
  if err != nil {
    return err
  }
  p.reportDone(a.Type)
  return nil
}

func (p *Process) performAction(a *Action) error {
  var err error

  switch a.Type {
    case "":
      err = &MalformedAction{"unknown", "action"}

    case "initialize":
      if a.ShardId == "" {
        err = &MalformedAction{a.Type, "shardId"}
      } else {
        err = p.Rp.Init(a.ShardId)
      }

    case "processRecords":
      if len(a.Records) == 0 {
        err = &MalformedAction{a.Type, "records"}
      } else {
        err = p.Rp.ProcessRecords(a.Records, p.Cp)
      }

    case "shutdown":
      if a.Reason == "" {
        err = &MalformedAction{a.Type, "reason"}
      } else {
        err = p.Rp.Shutdown(p.Cp, a.Reason)
      }

    default:
      err = &MalformedAction{a.Type, ""}
  }

  if err != nil {
    var buf []byte
    runtime.Stack(buf, true)
    p.Ih.WriteError(buf)
    return err
  }
  return nil
}

func (p *Process) reportDone(r string) {
  res := Response{"status", r}
  p.Ih.WriteAction(res)
}

func Run(rp RecordProcessor) error {
  ih := NewIOHandler()
  p := Process{rp, &ih, &CheckPointer{&ih}}
  for {
    l := p.Ih.ReadLine()
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

package kcl

type RecordProcessor interface {
  Init(string) error
  ProcessRecords([]*Record, *CheckPointer) error
  Shutdown(*CheckPointer, reason string) error
}

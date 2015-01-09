package kcl

type Record struct {
	Data           string `json:"data"`
	PartitionKey   string `json:"partitionKey"`
	SequenceNumber string `json:"sequenceNumber"`
}

type RecordProcessor interface {
  New(string) error
  ProcessRecord([]*Record, *CheckPointer) error
  Shutdown(*CheckPointer, reason string) error
}

package kcl

type Record struct {
  Data           string `json:"data"`
  PartitionKey   string `json:"partitionKey"`
  SequenceNumber string `json:"sequenceNumber"`
}

type Action struct {
  Type    string `json:"action"`
  ShardId string `json:"shardId"`
  Records []*Record `json:"records"`
  Reason  string `json:"reason"`
  Error   string `json:"error"`
}

func (a *Action) IsEmpty() bool {
  var s string
  return a.Type == s && a.ShardId == s
}

func (a *Action) IsError() bool {
  var s string
  return a.Error != s
}

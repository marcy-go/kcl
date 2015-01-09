package kcl

type Action struct {
  Name    string `json:"action"`
  ShardId string `json:"shardId"`
  Error   string `json:"error"`
}

func (a *Action) IsEmpty() bool {
  var s string
  return a.Name == s && a.ShardId == s
}

func (a *Action) IsError() bool {
  var s string
  return a.Error != s
}

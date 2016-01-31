package util

import "encoding/json"

func JsonDecode(input string, value interface{}) {
  if err := json.Unmarshal([]byte(input), value); err != nil {
    panic(err)
  }
}

func ToJson(input interface{}) string {
  data, err := json.MarshalIndent(input, "", "  ")
  if err != nil {
    panic(err)
  }
  return string(data)
}
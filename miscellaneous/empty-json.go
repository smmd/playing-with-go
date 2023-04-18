package main

import (
	"encoding/json"
	"fmt"
)

var (
  rawJSON1 = `{"event": [{"name":"david"},{"name":"sag"}]}`
  rawJSON2 = `{"event": {}}`
)

type event struct {
  Event interface{} `json:"event"`
}

type CustomObject struct {
  Name string `json:"name"`
}

// This was an excercise build with my coworker to manage [] and {} on json files as objects
func main() {
	e1 := event{}
  if err := json.Unmarshal([]byte(rawJSON1), &e1); err != nil {
     panic(err)
  }

  if s, ok := e1.Event.([]interface{}); ok {
    co := []CustomObject{}

    bytes, err := json.Marshal(s)
    if err != nil {
      panic(err)
    }
    
    if err := json.Unmarshal(bytes, &co); err != nil {
      panic(err)
    }
    fmt.Println("custom object: ", co)
  }

  fmt.Println(e1.Event)

  e2 := event{}
  if err := json.Unmarshal([]byte(rawJSON2), &e2); err != nil {
     panic(err)
  }

  fmt.Println(e2.Event)
}

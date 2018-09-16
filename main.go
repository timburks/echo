package main

import (
  "fmt"
  "github.com/golang/protobuf/proto"
  echo "./generated"
)

func main() {
  document := echo.EchoMessage{Text:"hello"}
  bytes, _ := proto.Marshal(&document)
  fmt.Printf("%+v\n", bytes)
}


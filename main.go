package main

import (
  "fmt"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/jsonpb"
  echo "./generated"
)

func main() {
  document := echo.EchoMessage{Text:"hello"}
  document.Sample_100Value = 100
  
  bytes, _ := proto.Marshal(&document)
  fmt.Printf("%+v\n", bytes)

  ma := jsonpb.Marshaler{}
  s, _ := ma.MarshalToString(&document)
  fmt.Printf("%s\n", s)  

  var d2 echo.EchoMessage
  jsonpb.UnmarshalString(s, &d2)
  fmt.Printf("%+v\n", d2)  
  
}


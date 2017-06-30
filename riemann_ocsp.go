package main

import (
  "github.com/amir/raidman"
  "os"
  "fmt"
  "strconv"
)

func main() {

  if len(os.Args) != 8 {
    fmt.Printf("usage: %s <tag> <host> <service> <state> <description> <ttl> <server:port>\n\n", os.Args[0])
    panic("wrong number of arguments")
  }

  c, err := raidman.Dial("tcp", os.Args[7])
  if err != nil {
    panic(err)
  }

  f, err := strconv.ParseFloat(os.Args[6], 32)
  if err != nil {
    panic(err)
  }

  var state string

  switch os.Args[4] {
    case "0": state = "ok"
    case "1": state = "warning"
    case "2": state = "critical"
    case "3": state = "unknown"
    default : state = os.Args[4]
  }

  var event = &raidman.Event{
    Host:         os.Args[2],
    Service:      os.Args[3],
    State:        state,
    Description:  os.Args[5],
    Ttl:          float32(f),
    Tags:         []string{os.Args[1]},
  }

  err = c.Send(event)
  if err != nil {
    panic(err)
  }

  c.Close()
}

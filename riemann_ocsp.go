package main

import (
  "github.com/amir/raidman"
  "os"
  "fmt"
  "strings"
  "strconv"
)

func main() {

  if len(os.Args) != 9 {
    fmt.Printf("usage: %s <tag> <host> <service> <state> <metric> <description> <ttl> <server:port>\n\n", os.Args[0])
    panic("wrong number of arguments")
  }

  c, err := raidman.Dial("tcp", os.Args[8])
  if err != nil {
    panic(err)
  }

  f, err := strconv.ParseFloat(os.Args[7], 32)
  if err != nil {
    panic(err)
  }

  m, err := strconv.ParseFloat(os.Args[5], 32)
  if err != nil {
    panic(err)
  }

  var state string

  switch os.Args[4] {
    case "0": state = "ok"
    case "1": state = "warning"
    case "2": state = "critical"
    case "3": state = "unknown"
    default : state = strings.ToLower(os.Args[4])
  }

  var num_state string

  switch state {
    case "ok": num_state = "0"
    case "warning": num_state = "1"
    case "critical": num_state = "2"
    case "unknown": num_state = "3"
    default : num_state = "3"
  }

  var service = os.Args[1] +string('.')+ strings.Replace(strings.ToLower(os.Args[3]), " ", "_", -1)

  var event = &raidman.Event{
    Host:         os.Args[2],
    Service:      service,
    State:        state,
    Metric:     float32(m),
    Description:  os.Args[6],
    Ttl:          float32(f),
    Tags:         []string{os.Args[1]},
    Attributes: map[string]string{"num_state": num_state},
  }

  err = c.Send(event)
  if err != nil {
    panic(err)
  }

  c.Close()
}

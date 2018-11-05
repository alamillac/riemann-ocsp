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
    fmt.Printf("usage: %s <server:port> <host> <service> <state> <metric> <description> <ttl>\n\n", os.Args[0])
    panic("wrong number of arguments")
  }

  c, err := raidman.Dial("tcp", os.Args[1])
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

  var state string = strings.ToLower(os.Args[4])

  var service = strings.Replace(strings.ToLower(os.Args[3]), " ", "_", -1)

  var event = &raidman.Event{
    Host:         os.Args[2],
    Service:      service,
    State:        state,
    Metric:     float32(m),
    Description:  os.Args[6],
    Ttl:          float32(f),
  }

  err = c.Send(event)
  if err != nil {
    panic(err)
  }

  c.Close()
}

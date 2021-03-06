package main

import (
  "strconv"
  "strings"
  "math"
  "fmt"
  "flag"
  "os/exec"
  "time"
  "net/http"
)

func RoundToHour(minute float64) int {
  // for testing with round-to-min
  //return int(math.Ceil(minute))
  return int(math.Ceil(minute / 60)) * 60
}

var shutdownCmd string

func shutdown() {
  cmd := exec.Command("sh", "-c", shutdownCmd)
  out, _ := cmd.Output()
  fmt.Printf("shutdown: %q => %q\n", shutdownCmd, out)
}

func main() {
  var listenAddress string
  flag.StringVar(&shutdownCmd, "shutdown", "halt", "command to run on shutdown")
  flag.StringVar(&listenAddress, "listen", "localhost:31337", "command to run on shutdown")
  flag.Parse()

  resp, _ := http.Get("http://169.254.169.254/latest/meta-data/local-ipv4")
  fmt.Println(resp.Header["Last-Modified"])

  start, _ := time.Parse(time.RFC1123, resp.Header["Last-Modified"][0])
  fmt.Println("Instance started at", start)
  start = start.Add( time.Duration(-5) * time.Minute )

  var BoomX *time.Timer

  Reset := func(min int) {
    if BoomX != nil {
      BoomX.Stop()
    }
    t := start.Add( time.Duration(min) * time.Minute )
    fmt.Println("reset", min, t)
    BoomX = time.AfterFunc(t.Sub(time.Now()), shutdown)
  }
  Reset( RoundToHour(time.Since(start).Minutes()) )

  http.HandleFunc("/ping/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", r.URL.Path)
    ping, _ := strconv.ParseInt(strings.Split(r.URL.Path, "/")[2], 0, 64)
    Reset( RoundToHour(time.Since(start).Minutes() + float64(ping / 60)) )
  })

  http.ListenAndServe(listenAddress, nil)
}

package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "bytes"
    "io"
    "os" 
    "encoding/json"
    "strconv"
)

func readfile(filename string) string {
  buf := bytes.NewBuffer(nil)
  f, _ := os.Open(filename) // Error handling elided for brevity.
  io.Copy(buf, f)           // Error handling elided for brevity.
  f.Close()
  s := string(buf.Bytes())
  return s
}

func main() {
    
    sips := []string{}
  
    http.HandleFunc("/tropo", func(w http.ResponseWriter, r *http.Request) {
      sipcountstr := strconv.Itoa(len(sips))
      sipm, _ := json.Marshal(sips)     
      res := `{"tropo":[
      {
         "say":[{"value":"There are `+sipcountstr+` available operators. Transferring you now."}]
      },
      {
         "transfer":{
            "to":`+string(sipm)+`,
            "choices":{"terminator":"#"},
            "on":{
               "say":{"value":"http://www.phono.com/audio/holdmusic.mp3"},
               "event":"ring"
            }
         }
      }
   ]}`





        fmt.Fprintf(w, "%s", res)
        println(res)
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        html := readfile("phone.html")
        fmt.Fprintf(w, "%s", html)
    })

    http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        sip_ := r.URL.Query().Get("sip")
        sips_ := strings.Split(sip_,"=")
        sip := "sip:"+sips_[1]
        fmt.Fprintf(w,"Success")
        sips = append(sips,sip)
        println("SIP LIST ______________")
        for _, s := range sips {
            println("-"+s)
        }
        println("")
    })
    println("Let's go!")  
    log.Fatal(http.ListenAndServe(":8000", nil))
}

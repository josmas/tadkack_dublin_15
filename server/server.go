package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "strconv"
    "time"
    "sync"
)

type Reciever struct {
    sip      string
    lastPing time.Time
}

func update (recievers []Reciever, addsip *string) []Reciever {

    ret := []Reciever{}
    now := time.Now()

    for _, s := range recievers {

        if addsip!=nil && s.sip == *addsip {
           ret = append(ret,Reciever{sip:*addsip,lastPing: now})
           addsip = nil
           continue 
        }

        elapsed := time.Now().Sub(s.lastPing).Seconds()
        if elapsed < 5 {
            ret = append(ret,s) 
        }       
    }

    if addsip!=nil {
       ret = append(ret,Reciever{sip:*addsip,lastPing: now}) 
    }

    return ret      	
}

func recievers2json(recievers []Reciever) string {
   list := []string{}
   for _, r := range recievers {
      list = append(list,r.sip)
   }
   sipm, _ := json.Marshal(list)
   return string(sipm)
}

func main() {

    recievers := []Reciever{}    
    var mutex = &sync.Mutex{}
 
    http.HandleFunc("/tropo", func(w http.ResponseWriter, r *http.Request) {
      
      mutex.Lock()
      recievers = update(recievers,nil)
      mutex.Unlock()
      
      sipcountstr := strconv.Itoa(len(recievers))
      sipjson:= recievers2json(recievers)     
      res := `{"tropo":[
      {
         "say":[{"value":"There are `+sipcountstr+` available operators. Transferring you now."}      ]},
      {
         "transfer":{
            "to":`+sipjson+`,
            "choices":{"terminator":"#"},
            "on":{
               "say":{"value":"http://www.phono.com/audio/holdmusic.mp3"},
               "event":"ring"
            }
         }
      }]}`

        fmt.Fprintf(w, "%s", res)
        println(res)
    })

    http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
        sip := "sip:"+r.URL.Query().Get("sip")
        mutex.Lock()
        recievers = update(recievers,&sip)
        mutex.Unlock()
        fmt.Fprintf(w,strconv.Itoa(len(recievers)))
    })

    http.Handle("/", http.FileServer(http.Dir("./static/")))

    println("Let's go!")  
    log.Fatal(http.ListenAndServe(":8000", nil))
}

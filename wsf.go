package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "regexp"
  "time"
  "strconv"
  "os"
)

type Schedule struct {
  ScheduleStart	string
  ScheduleEnd 	string
  ScheduleID 	int
  ScheduleName 	string
  TerminalCombos []struct {
    ArrivingTerminalName    string
    DepartingTerminalName   string
    SailingNotes	    string
    Times []struct {
      ArrivingTime	string
      DepartingTime	string
      VesselName	string
      VesselPositionNum int
    } `json:Times`
  } `json:TerminalCombos`
}


func apidate(date string) time.Time {
      regex := regexp.MustCompile(`/Date\(([0-9]+)000-0[\d]00\)/`)

      time_string := regex.FindStringSubmatch(date)[1]
      time_int, _  := strconv.Atoi(time_string)
      time := time.Unix(int64(time_int), 0)

      return(time)
} 


func main() {

  wsf_api_key  := os.Getenv("WSFAPI")
  wsf_api_url  := "http://www.wsdot.wa.gov/Ferries/API/Schedule/rest/scheduletoday/5/TRUE"

  var wsf_sched_url string

  if len(wsf_api_key) > 1 { 
    wsf_sched_url = fmt.Sprintf("%s?apiaccesscode=%s", wsf_api_url, wsf_api_key)
  } else {
    fmt.Println("To use this tool you must set the WSFAPI environment variable with your WSF API Access Code.")
    fmt.Println("Obtain an access code at http://www.wsdot.wa.gov/traffic/api/")
    os.Exit(1)
  }

  r, _ := http.Get(wsf_sched_url)
  b, _ := ioutil.ReadAll(r.Body)
  
  //fmt.Println(string(b)) 

  var s Schedule
  json.Unmarshal(b, &s)

  fmt.Println(s.ScheduleName)

  for _, v := range s.TerminalCombos { 
    //fmt.Printf("Got: %d -  %+v\n", k, v) 
    fmt.Println(v.DepartingTerminalName, "-->", v.ArrivingTerminalName, v.SailingNotes)

    for _, times := range v.Times {
      time := apidate(times.DepartingTime)
      fmt.Printf("  %2d:%02d   --  %s (%d)\n", time.Hour(), time.Minute(), times.VesselName, times.VesselPositionNum)
    }
  }

}

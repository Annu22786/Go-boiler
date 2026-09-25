package config

import(
  "os"
  "time"
)
//parents structs
type Observibility struct{
  service_name string
  Enviroment string
  Logging Logger
  NewRelic Relic
  Healthcheck
}
//child structs
type Logger struct{
  Level string
  Format string
  SlowQuery time.Duration
}

type Relic struct{
  License_key string
  Applogforwadingenabled bool
  Distributedtracing bool
  Debuglogging bool
}

type Health struct{
  Enabled bool
  Checks []strings
  Interval time.Duration
  Timeout yime.Duration
}
// by default observibilty define

func Defaultobservability() *Observibility{
  return &Observibility {
    service_name : "Boilerplate",
    Enviroment : "Development",
    logging : Logger{
      Level : "Debug",
      Format : "json" ,
      SlowQuery : 100*time.milisecond
    },
    NewRelic : Relic{
      License_key : "" ,
      Applogforwadingenabled : true ,
      Distributedtracing : false ,
      Debuglogging : false 
    },
    Healthcheck : Health {
      Enabled : true ,
      Checks : []strings{Database,Redis},
      Interval : 100*time.milisecond ,
      Timeout : 100*time.milisecond ,
      }
   }
}
// validation 
func (c *observibilty) validate () {
  if c.service_name == ""{
    return fmt.Error{"service name cannot be null"}}
  loglevel := map[string]bool{
    "Fatal" : true , "Error" : true , "Dedug" : true , "info" : true , "warn" : true
  }
  if !loglevel (c.Logging.level){
    r3turn fmt.Error("logging level error")
  }
  if c.Logging.SlowQuery < 0{
    return fmt.Error("logging query cannot be negative")
  }
}


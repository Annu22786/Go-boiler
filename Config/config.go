package main

import(
  "os"
  "string"
  "...github.com/joho/godotenv v1.5.1 "
  "github.com/knadh/koanf/v2"
  "github.com/rs/zerolog"
  "github.com/go-playground/validator"
  "github.com/knadh/koanf/providers/env"
)

type Config struct{
  Primaryconfig Primary
  Databaseconfig Database
  RedisConfig Redis
  Authconfig Auth
  Serverconfig Server
}

  type Primary struct{
  Env string `koanf:"env" Validate:"required"`
}

type Database struct{
  Host string `koanf:"database_host" Validate:"required"`
  Database_port int `koanf:"database_port" Validate:"required"`
  User string `koanf:"user" Validate:"required"`
  Password string `koanf:"password" Validate:"required"`
  Name string `koanf:"name" Validate:"required"`
}

type Redis struct{
  Chache_key string `koanf:"chache_redis" Validate:"required"`
}

type Auth struct{
  Auth_key string `koanf:"auth" Validate:"required"`
}

type Server struct{
  CORS []string `koanf:"cors" Validate:"required"`//in arrays tgere will be multiple frontebd domains name of ours
  Server_port int `koanf:"server_port" Validate:"required"`
  Readtimeout int `koanf:"readout" Validate:"required"`
  Writedata int `koanf:"writedata" Validate:"required"`
  Idletimeout int `koanf:"idelout" Validate:"required"`
}

func loadenv() (*config,error) {
  logger := zerolog.New(zerlog.Consolewriter(Out: os.Stderr)).with().timestamp().logger()
  
  k := koanf.New(".")
  
  err= k.Load(env.provider("Boilerplate...",".")func (s string) string{ string.Tolower(string.trimprefix(s,"BOILERPLATE..."))})
 
  if err != nil{
    logger.Fatal().Err(err).Msg("Env is not present")
  }
  makeconfig:=&config
  err = k.unmarshell("",makeconfig)
  if err!=nil{
    logger.Fatal().Err(err).Msg("Loading error")
  }

  Validate:= validator.New()
  err = Validate.Struct(makeconfig)
  if err != nil{
    logger.Fatal().Err(err).Msg("Validation error")
  }
  return makeconfig,nil
  
}

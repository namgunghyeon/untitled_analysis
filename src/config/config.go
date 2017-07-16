package config

import (
  "fmt"
  "github.com/BurntSushi/toml"
)

func LoadIgnoreDirs() {}

func LoadCouchbase() Config{
  var conf Config
  if _, err := toml.DecodeFile("src/config/couchbase.toml", &conf); err != nil {
    fmt.Println("error", err)
    return conf
  }
  return conf
}

package config

import (
  "log"
  "github.com/spf13/viper"
  _ "github.com/lib/pq"
  "database/sql"
)

var (
  db *sql.DB
)
func Connect(){
  // dsn := "host=127.23.25.219 user=postgres port=5432 password=r0ZNpOLIC9bCm0fdAI dbname=esc "
  // dsn:="postgres://postgres:r0ZNpOLIC9bCm0fdAI@127.23.25.219:5432/esc"
  // dsn:= "host=localhost user=postgres password=postgrespw port=55001 dbname=esc"
  viper.SetConfigFile(".env")

  err := viper.ReadInConfig()
  if err!=nil{
    log.Fatalf("Error while reading config %s",err)
  }
  dsn,ok := viper.Get("PSQL_LINK").(string)
  if !ok{
    log.Fatalf("Invalid type assertion")
  }
  

  d,err := sql.Open("postgres",dsn)
  if err != nil{
    panic(err)
  }
  db = d
}

func GetDB()*sql.DB{
  return db
}

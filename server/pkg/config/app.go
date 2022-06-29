package config

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var (
  db *gorm.DB
)
func Connect(){
  // dsn := "host=127.23.25.219 user=postgres port=5432 password=r0ZNpOLIC9bCm0fdAI dbname=esc "
  // dsn:="postgres://postgres:r0ZNpOLIC9bCm0fdAI@127.23.25.219:5432/esc"
  dsn:= "host=localhost user=postgres password=postgrespw port=55001 dbname=esc"

  d,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
  if err != nil{
    panic(err)
  }
  db = d
}

func GetDB()*gorm.DB{
  return db
}

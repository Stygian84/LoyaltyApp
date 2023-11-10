package utils

import (
  "math/rand"
  "strings"
  "time"
)

func init(){
  rand.Seed(time.Now().UnixNano())
}

func RandomInt(min,max int64) int64{
  return min+rand.Int63n(max-min+1)
}
func RandomFloat(n float64)float64{
  return rand.Float64()*n
}

func RandomString(n int)string{
  strCharacters := "abcdefghijklmnopqrstuvwxyz"
  var sb strings.Builder
  k:= len(strCharacters)
  for i:=0; i<n;i++{
    index := rand.Intn(k)
    c:= strCharacters[index]
    sb.WriteByte(c)
  }
  return sb.String()
}

func RandomSelect(items []string)string{
  size := len(items)
  index := rand.Intn(size)
  return items[index]
}



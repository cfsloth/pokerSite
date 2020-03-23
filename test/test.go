package main

import (
  "../domain/entity"
  "fmt"
  "testing"
)

func TestHello(*testing.T)  {
  var a entity.Card
  a.ID = 12
  a.Name = "s"
  a.Symbol = "s"
  fmt.Printf("Id: %d \n Name: %s \nSymbol: %s",a.ID,a.Name,a.Symbol)
}

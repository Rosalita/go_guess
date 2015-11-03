package main

import (
  "fmt"
  "math/rand"
  "time"
  )

func main(){
    guesses:= 0
    var name string
    var guess int
    rand.Seed(time.Now().Unix())
    randnum:=rand.Intn(100)
    randnum += 1
    fmt.Println("Hello! What is your name?")
    fmt.Scanln(&name)
    fmt.Printf("Hello %s\n", name)
    fmt.Println("I'm thinking of a number between 1 and 100 ")
    fmt.Println("What is the number I am thinking of?")
    fmt.Scanln(&guess)
    guesses ++
    fmt.Printf("Your guess was %d\n", guess)
    fmt.Printf("The random number was %d\n", randnum)
}

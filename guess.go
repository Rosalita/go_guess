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

    for {
      fmt.Scanln(&guess)
      guesses ++
      if guess < randnum {
          fmt.Println("Your guess was too low, try again.")
      } else if guess > randnum{
          fmt.Println("Your guess was too high, try again.")
      } else {
          fmt.Printf("Congratulations %s! you guessed the number was %d\n", name,randnum)
          break
      }
    }
    fmt.Printf("It took you %d guesses to do this\n",guesses)

}

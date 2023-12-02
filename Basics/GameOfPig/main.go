package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

const rounds = 10

func main() {
	if len(os.Args) != 3 {
		fmt.Println("--usage: ./GameOfPig [holdcount1] [holdcount2]")
		return
	}

	holdcount1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	holdcount2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	player1 := Player{strategy: getHoldAtValueStrategy(holdcount1)}
	player2 := Player{strategy: getHoldAtValueStrategy(holdcount2)}
	conductor := Conductor{player1: player1, player2: player2}
	conductor.Conduct(rounds)
}

func RollDie() int {
	return rand.Intn(6) + 1
}

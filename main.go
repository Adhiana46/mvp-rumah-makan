package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	Point      int
	DiceCount  int
	RollResult [][]int
}

func (p *Player) AddPoint() {
	p.Point += 1
}

func (p *Player) DecrDice() {
	p.DiceCount -= 1
}

func (p *Player) RollDice() []int {
	result := []int{}

	for i := 0; i < p.DiceCount; i++ {
		result = append(result, rand.Int()%6+1)
	}

	p.RollResult = append(p.RollResult, result)

	return result
}

func (p *Player) GetLastRollStrings() []string {
	result := []string{}

	n := len(p.RollResult)
	if len(p.RollResult) == 0 {
		return result
	}

	lastResult := p.RollResult[n-1]
	for _, v := range lastResult {
		result = append(result, strconv.Itoa(v))
	}

	return result
}

func (p *Player) IsStopPlaying() bool {
	return p.DiceCount == 0
}

func main() {
	var inN, inM string
	n := 0
	m := 0

	fmt.Print("Masukan Jumlah Pemain: ")
	fmt.Scan(&inN)

	fmt.Print("Masukan Jumlah Dadu: ")
	fmt.Scan(&inM)

	// convert input string to int
	n, err := strconv.Atoi(inN)
	if err != nil {
		panic("Tolong masukan angka untuk jumlah pemain")
	}

	m, err = strconv.Atoi(inM)
	if err != nil {
		panic("Tolong masukan angka untuk jumlah dadu")
	}

	players := make([]Player, n)
	for i, _ := range players {
		players[i] = Player{
			Point:     0,
			DiceCount: m,
		}
	}

	fmt.Println("\n\n*** GAME DADU ***")
	fmt.Printf("Pemain = %v, Dadu = %v\n", n, m)
	fmt.Println("====================")

	fmt.Println(players)

	count := 1
	nPlayer := len(players)
	for x := 0; x < 5; x++ {
		fmt.Printf("Giliran %v lempar dadu:\n", count)

		// Roll the dice
		for i := 0; i < nPlayer; i++ {
			if players[i].DiceCount == 0 {
				fmt.Printf("\tPemain #%v (%v): _(Berhenti bermain karena tidak memiliki dadu)\n", i+1, players[i].Point)
			} else {
				players[i].RollDice()
				rollResult := players[i].GetLastRollStrings()
				fmt.Printf("\tPemain #%v (%v): %s\n", i+1, players[i].Point, strings.Join(rollResult, ", "))
			}
		}

		// Evaluasi
		for i := 0; i < nPlayer; i++ {
			// TODO: -_-
		}

		fmt.Println("====================")
		time.Sleep(2 * time.Second)
		count++
	}
}

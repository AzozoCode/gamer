package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	kickBall()
}

type FootballPlayer struct {
	power   int
	stamina int
}

type CR7 struct {
	power   int
	stamina int
	SUI     int
}

func (c CR7) kickBall() {
	shot := c.stamina + c.power*c.SUI
	fmt.Printf("CR 7 is kicking the ball shot:%d\n", shot)

}

func (p FootballPlayer) kickBall() {
	shot := p.stamina + p.power

	fmt.Printf("I'm kicking the ball shot:%d\n", shot)
}

func main() {

	team := make([]Player, 11)

	for i := 0; i < len(team)-1; i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10) + 1,
			power:   rand.Intn(10) + 1,
		}
	}

	team[len(team)-1] = CR7{
		stamina: 10,
		power:   10,
		SUI:     10,
	}

	for _, player := range team {
		player.kickBall()
	}

}

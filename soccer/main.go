package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	kickBall() int
	getName() string
}

type FootballPlayer struct {
	power   int
	stamina int
	name    string
}

type CR7 struct {
	power   int
	stamina int
	SUI     int
	name    string
}

type Messi struct {
	power   int
	stamina int
	SUI     int
	name    string
}

func (c CR7) kickBall() int {
	return c.stamina + c.power*c.SUI

}

func (p CR7) getName() string {
	return p.name
}

func (m Messi) kickBall() int {
	return m.stamina + m.power*m.SUI
}

func (p Messi) getName() string {
	return p.name
}

func (p FootballPlayer) kickBall() int {
	return p.stamina + p.power
}

func (p FootballPlayer) getName() string {
	return p.name
}

func main() {

	team := make([]Player, 11)

	for i := 0; i < len(team)-2; i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10) + 1,
			power:   rand.Intn(10) + 1,
			name:    "Random",
		}
	}

	team[len(team)-1] = CR7{
		stamina: 10,
		power:   10,
		SUI:     10,
		name:    "CR7",
	}

	team[len(team)-2] = Messi{
		stamina: 10,
		power:   10,
		SUI:     8,
		name:    "Messi",
	}

	for i := range team {

		fmt.Printf("%s is kicking the ball with shot: %d\n", team[i].getName(), team[i].kickBall())
	}
	fmt.Println("---------------------------------------------------------")

	for _, p := range team {

		fmt.Printf("%s is kicking the ball with shot: %d\n", p.getName(), p.kickBall())
	}

}

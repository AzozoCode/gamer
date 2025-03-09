package main

import (
	"fmt"
	"time"
)

type Game struct {
	isRunning  bool
	players    map[string]*Player
	isPausedCh chan bool
	quitCh     chan bool
	doneCh     chan bool
}

func NewGame() *Game {

	return &Game{
		isRunning:  false,
		players:    make(map[string]*Player),
		isPausedCh: make(chan bool),
		quitCh:     make(chan bool),
		doneCh:     make(chan bool),
	}
}

func (g *Game) Start() {
	g.isRunning = true

	go gameLoop(g)

}

func (g *Game) quitGame() {
	fmt.Println("Quiting game...")
	g.quitCh <- true
	<-g.doneCh //block until game loop signals it's done
	fmt.Println("Game has exited")

}

func gameLoop(g *Game) {

	interval := 1 * time.Second
	timer := time.NewTicker(interval)
	defer timer.Stop()

	for {
		select {
		case <-g.quitCh:
			g.doneCh <- true
		case <-timer.C:
			fmt.Println("Game is running....")
		}
	}
}

type Player struct {
	name   string
	health int
	power  int
}

func addPlayer(name string, hp, power int) *Player {

	fmt.Printf("adding player with name: %s | health: %d | power: %d\n", name, hp, power)
	return &Player{
		power:  power,
		name:   name,
		health: hp,
	}
}

func main() {

	game := NewGame()

	playerA := addPlayer("Bob", 100, 100)
	playerB := addPlayer("Amy", 100, 120)

	game.players[playerA.name] = playerA
	game.players[playerB.name] = playerB

	game.Start()

	//simulate quitting after 3 seconds
	time.Sleep(3 * time.Second)
	game.quitGame()

	fmt.Println("Game stopped")
}

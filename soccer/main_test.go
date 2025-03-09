package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_KickBall(t *testing.T) {
	team := make([]FootballPlayer, 11)

	for i := 0; i < len(team); i++ {
		team[i] = FootballPlayer{
			power:   rand.Intn(10) + 1,
			stamina: rand.Intn(10) + 1,
			name:    "test player",
		}
	}

	for i := range team {
		t.Run(team[i].name, func(t *testing.T) {

			res := team[i].kickBall()
			require.NotZero(t, res)
			require.EqualValues(t, team[i].power+team[i].stamina, res)

		})
	}

}

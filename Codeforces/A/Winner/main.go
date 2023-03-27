package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n       int
	players []Player
	record  = make(map[string]int)
	winners = make(map[string]int)
)

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &n)

	for i := 0; i < n; i++ {
		var player Player
		fmt.Fscan(gin, &player.Name, &player.Score)
		players = append(players, player)
		record[player.Name] += player.Score
	}

	maxScore := 0
	for _, score := range record {
		if maxScore < score {
			maxScore = score
		}
	}

	for _, player := range players {
		winners[player.Name] += player.Score
		if record[player.Name] == maxScore && winners[player.Name] >= maxScore {
			fmt.Fprintln(gout, player.Name)
			return
		}
	}
}

type Player struct {
	Name  string
	Score int
}

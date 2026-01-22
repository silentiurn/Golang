package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	players := [4][3]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			players[i][j] = deck[j+i*3]
		}
	}
	for i, p := range players {
		fmt.Printf("Player %d%s", i+1, ": ")
		for j, card := range p {
			fmt.Printf("%d", card)
			if j != 2 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
	}
}

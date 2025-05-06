package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type res struct {
	g string
	m int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	secr := rand.Intn(900) + 100
	secrS := strconv.Itoa(secr)
	sD := make(map[rune]int)
	for _, d := range secrS {
		sD[d]++
	}

	fmt.Println("I picked a number. Try to guess.")
	fmt.Println("Give up? Type 'quit'")

	att := 0
	var hist []res

	for {
		var g string
		fmt.Print("Your guess: ")
		fmt.Scanln(&g)
		if strings.ToLower(g) == "quit" {
			fmt.Printf("Gave up, huh? It was %s\n", secrS)
			break
		}
		att++
		m := 0
		gD := make(map[rune]int)
		for _, d := range g {
			gD[d]++
		}

		count := make(map[rune]bool)
		for sDgt, sC := range sD {
			if gC, ok := gD[sDgt]; ok {
				minC := sC
				if gC < minC {
					minC = gC
				}
				m += minC
				count[sDgt] = true
			}
		}

		hist = append(hist, res{g: g, m: m})
		fmt.Printf("%d matched\n", m)
		if g == secrS {
			fmt.Printf("You got it! It's %s in %d tries!\n", secrS, att)
			fmt.Println("\nYour guesses:")
			for _, r := range hist {
				fmt.Printf("%s - %d\n", r.g, r.m)
			}
			break
		}
	}
}

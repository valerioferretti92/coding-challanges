package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no string to sequence has been provided as cmd arg")
	}

	yellow := color.New(color.FgYellow, color.Bold)
	green := color.New(color.FgGreen, color.Bold)

	film := os.Args[1]
	lastseen := make(map[rune]int)
	var slenghts []int

	// record last occurrance of each char in a map
	for i, r := range film {
		lastseen[r] = i
	}

	var sleft = 0
	var slenght = 1
	for k, v := range film {
		yellow.Print("[DEBUG]")
		fmt.Printf(": k:%d, v:%c, sleft:%d, slenght:%d\n", k, v, sleft, slenght)

		// if reading last of occurance of a rune AND there are no more runes to inspect
		// then printout scene lenght && reinitialize
		if lastseen[v] == k && sleft+slenght-1 == k {
			slenghts = append(slenghts, slenght)
			sleft = sleft + slenght
			slenght = 1
			continue
		}

		// if reading rune that appears again after current possible scene end
		// then extend scene end hypothesis
		if lastseen[v] > sleft+slenght-1 {
			slenght = lastseen[v] - sleft + 1
			continue
		}
	}

	for k, v := range slenghts {
		green.Printf("[SCENE %d]", k)
		fmt.Printf(": %d runes\n", v)
	}
}

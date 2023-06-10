package main

import "fmt"

func main() {
	fmt.Println(getHint("1807", "7810"))
}

func getHint(secret string, guess string) string {
	nonBullGuesses := make(map[rune]int)
	nonBullSecrets := make(map[rune]int)

	cows := 0
	bulls := 0
	for i, rg := range guess {
		rs := rune(secret[i])

		if rs == rg {
			bulls++
			continue
		}

		gcounter, ok := nonBullGuesses[rg]
		if !ok {
			nonBullGuesses[rg] = 1
		} else {
			nonBullGuesses[rg] = gcounter + 1
		}

		scounter, ok := nonBullSecrets[rs]
		if !ok {
			nonBullSecrets[rs] = 1
		} else {
			nonBullSecrets[rs] = scounter + 1
		}
	}

	for s, cg := range nonBullGuesses {
		cs, ok := nonBullSecrets[s]
		if !ok {
			continue
		}

		if cg <= cs {
			cows = cows + cg
		} else {
			cows = cows + cs
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

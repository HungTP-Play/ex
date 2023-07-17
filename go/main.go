package main

import (
	"fmt"

	"github.com/HungTP-Play/ex/go/ex"
)

func main() {
	input := "bbaaggbagbgabasdbfasbfbasdfbbbbgabfgdbsdbfbsbababsbsbfdbfbbababababababsfbsdbfasbdfbsdfahsadfgasdfgasdgfggggggggggggggagagagagagabbbbbgagagaga"
	count := ex.Ex1_2(input, "bag")
	fmt.Printf("count: %d\n", count)
}

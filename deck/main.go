package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/robsantossilva/deck/deck"
)

func main() {
	cards := deck.New(deck.Sort(deck.Most))
	//ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	fmt.Println(perm)
}

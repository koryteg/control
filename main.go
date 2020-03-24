package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there are you ready for control!!!??!")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// deck := NewDeck()
	// Shuffle(deck)

	// for _, card := range deck {
	// 	fmt.Printf(" : %s", card.Name )
	// }

}


// a Game needs 2-4 players.
// when we start

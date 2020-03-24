package main

import (
	"math/rand"
)


type Card struct {
	Name string
	Type string // Sliver, Bronze
	Rule string
	CellCharge int
}

var cards = []Card{
	Card{
		Name: "Rift",
		Type: "Silver",
		CellCharge: 1,
		Rule: "Draw a card or destroy a Nova in Play",
	},
	Card{
		Name: "Exotic Matter",
		Type: "Silver",
		CellCharge: 2,
		Rule: "You may install another fuel cell with a carge of six or less, a chain cannot exceed a charge of 10",
	},
	Card{
		Name: "deflector",
		Type: "Silver",
		CellCharge: 3,
		Rule: "if destroyed by an anomaly's ability, draw a card or make a player discard a random card.",
	},
	Card{
		Name: "Reactor",
		Type: "Silver",
		CellCharge: 5,
		Rule: "This Fuel Cell is Immune to defusing once installed it is not immune to abilities",
	},
	Card{
		Name: "Anomaly",
		Type: "Bronze",
		CellCharge: 4,
		Rule: "Destroy all Silver fuel cells in play (including your own)",
	},
	Card{
		Name: "Wormhole",
		Type: "Bronze",
		CellCharge: 4,
		Rule: "look through the discard pile. Take one card from it and add it to your hand",
	},
	Card{
		Name: "Rewind",
		Type: "Bronze",
		CellCharge: 5,
		Rule: "Return an installed fuel cell to the bottom of the deck",
	},
	Card{
		Name: "Future Shift",
		Type: "Bronze",
		CellCharge: 6,
		Rule: "Look at the top two cards of the deck, return one to the top, then play the other however you want",
	},
	Card{
		Name: "Dark Energy",
		Type: "Bronze",
		CellCharge: 6,
		Rule: "Transfer control of a silver Fuel Cell in play. The ability is triggered for the new owner.",
	},
	Card{
		Name: "Singularity",
		Type: "Bronze",
		CellCharge: 7,
		Rule: "Destroy all Bronze Fuel Cells in play (including your own).",
	},
	Card{
		Name: "Antimatter",
		Type: "Bronze",
		CellCharge: 8,
		Rule: "Force an opponent to discard a random card from their hand, then one of their choice",
	},
	Card{
		Name: "Time Stop",
		Type: "Bronze",
		CellCharge: 9,
		Rule: "Play any time a Bronze Fuel Cell is burned to cancel its ability. the stopped player gets another turn",
	},
	Card{
		Name: "Nova",
		Type: "Bronze",
		CellCharge: 10,
		Rule: "no ability",
	},
}

type Deck []Card

func NewDeck()(deck Deck) {

	for i := 0; i < len(cards) ; i++ {
		for n := 0; n < 4; n++ {
			card := cards[i]
			deck = append(deck, card)
		}
	}

	return
}

func Shuffle(d Deck) Deck {
	for i := 1; i < len(d); i++ {
		r := rand.Intn(i + 1)

		if i != r {
			d[r], d[i] = d[i], d[r]
		}

	}
	return d
}


// rift := Card{
// 	Name: "Rift",
// 	Type: "Silver",
// 	CellCharge: 1,
// 	Rule: "Draw a card or destroy a Nova in Play",
// }
// exotic_matter := Card{
// 	Name: "Exotic Matter",
// 	Type: "Silver",
// 	CellCharge: 2,
// 	Rule: "You may install another fuel cell with a carge of six or less, a chain cannot exceed a charge of 10",
// }
// deflector := Card{
// 	Name: "deflector",
// 	Type: "Silver",
// 	CellCharge: 3,
// 	Rule: "if destroyed by an anomaly's ability, draw a card or make a player discard a random card.",
// }
// reactor := Card{
// 	Name: "Reactor",
// 	Type: "Silver",
// 	CellCharge: 5,
// 	Rule: "This Fuel Cell is Immune to defusing once installed it is not immune to abilities",
// }
// anomaly := Card{
// 	Name: "Anomaly",
// 	Type: "Bronze",
// 	CellCharge: 4,
// 	Rule: "Destroy all Silver fuel cells in play (including your own)",
// }
// wormhole := Card{
// 	Name: "Wormhole",
// 	Type: "Bronze",
// 	CellCharge: 4,
// 	Rule: "look through the discard pile. Take one card from it and add it to your hand",
// }
// rewind := Card{
// 	Name: "Rewind",
// 	Type: "Bronze",
// 	CellCharge: 5,
// 	Rule: "Return an installed fuel cell to the bottom of the deck",
// }
// future_shift := Card{
// 	Name: "Future Shift",
// 	Type: "Bronze",
// 	CellCharge: 6,
// 	Rule: "Look at the top two cards of the deck, return one to the top, then play the other however you want",
// }

// dark_energy := Card{
// 	Name: "Dark Energy",
// 	Type: "Bronze",
// 	CellCharge: 6,
// 	Rule: "Transfer control of a silver Fuel Cell in play. The ability is triggered for the new owner.",
// }

// singularity := Card{
// 	Name: "Singularity",
// 	Type: "Bronze",
// 	CellCharge: 7,
// 	Rule: "Destroy all Bronze Fuel Cells in play (including your own).",
// }

// antimatter := Card{
// 	Name: "Antimatter",
// 	Type: "Bronze",
// 	CellCharge: 8,
// 	Rule: "Force an opponent to discard a random card from their hand, then one of their choice",
// }

// timestop := Card{
// 	Name: "Time Stop",
// 	Type: "Bronze",
// 	CellCharge: 9,
// 	Rule: "Play any time a Bronze Fuel Cell is burned to cancel its ability. the stopped player gets another turn",
// }

// nova := Card{
// 	Name: "Nova",
// 	Type: "Bronze",
// 	CellCharge: 10,
// 	Rule: "no ability",
// }

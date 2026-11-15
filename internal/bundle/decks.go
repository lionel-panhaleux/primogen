package bundle

import (
	"context"
	"fmt"
	"maps"
	"slices"
)

// Deck is a deck of the tournament winning deck archive (TWDA).
type Deck struct {
	ID      string
	Name    string
	Player  string
	Author  string
	Event   string
	Date    string
	Place   string
	Players int
	Format  string
	URL     string // the event's report
	Comment string
	Cards   []DeckCard
}

type DeckCard struct {
	ID      int
	Count   int
	Comment string
}

// twdaDeck is a deck of krcg 5's twda.json, an object keyed by deck id.
type twdaDeck struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Author  string `json:"author"`
	Player  string `json:"player"`
	Comment string `json:"comment"`
	Event   struct {
		Name    string `json:"name"`
		Date    string `json:"date"`
		Format  string `json:"format"`
		Players int    `json:"players_count"`
		Place   string `json:"place"`
		URL     string `json:"url"`
	} `json:"event"`
	Cards []struct {
		ID      int    `json:"id"`
		Count   int    `json:"count"`
		Comment string `json:"comment"`
	} `json:"cards"`
}

// fetchDecks reads the TWDA, leaving out a deck naming a card the card data
// lacks.
func fetchDecks(ctx context.Context, u string, cards []Card, warn func(string, ...any)) ([]Deck, string, error) {
	var raw map[string]twdaDeck
	version, err := fetchJSON(ctx, u, &raw)
	if err != nil {
		return nil, "", err
	}
	if len(raw) == 0 {
		return nil, "", fmt.Errorf("%s: no decks", u)
	}
	known := map[int]bool{}
	for _, c := range cards {
		known[c.ID] = true
	}
	var decks []Deck
	var left []string
	for _, id := range slices.Sorted(maps.Keys(raw)) {
		t := raw[id]
		d := Deck{
			ID: id, Name: t.Name, Player: t.Player, Author: t.Author, Comment: t.Comment,
			Event: t.Event.Name, Date: t.Event.Date, Place: t.Event.Place, Players: t.Event.Players,
			Format: t.Event.Format, URL: t.Event.URL,
		}
		for _, c := range t.Cards {
			if !known[c.ID] {
				left = append(left, fmt.Sprintf("%s (card %d)", id, c.ID))
				d.Cards = nil
				break
			}
			d.Cards = append(d.Cards, DeckCard{c.ID, c.Count, c.Comment})
		}
		if d.Cards != nil {
			decks = append(decks, d)
		}
	}
	if len(left) > 0 {
		warn("%d decks of the TWDA name a card the card data lacks, left out: %s", len(left), left[:min(len(left), 5)])
	}
	return decks, version, nil
}

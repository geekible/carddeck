package carddeck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Heart's
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Errorf("wrong number of cards in the deck. expected %d, got %d", 13*4, len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expect := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expect {
		t.Errorf("expected Ace of Spade's got %+v", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expect := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expect {
		t.Errorf("expected Ace of Spade's got %+v", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	// shuffleRand deterministic
	// first call to shuffleRand.Perm(52) should be
	// [40 35 50 0 ...]
	shuffleRand = rand.New(rand.NewSource(0))

	originalDeck := New()
	first := originalDeck[40]
	second := originalDeck[35]
	cards := New(Shuffle)

	if cards[0] != first {
		t.Errorf("expected the first card to be %s, got %s", first, cards[0])
	}

	if cards[1] != second {
		t.Errorf("expected the first card to be %s, got %s", second, cards[1])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Errorf("expected 3 jokers, got %d", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("expected all twos and threes to be filtered out, got %+v", c)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("expected %d cards got %d", 13*4*3, len(cards))
	}
}

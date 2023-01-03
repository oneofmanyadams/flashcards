package flashcards

type CardDeck struct {
	Cards []Flashcard
	DrawPile []Flashcard
	DiscardPile []Flashcard
	MinLevel int
	MaxLevel int
}

func (s *CardDeck) SetLevels(min, max int) {
	s.MinLevel = min
	s.MaxLevel = max
}

func (s *CardDeck) AddCard(card FlashCard) {
	s.Cards = append(s.Cards, card)
}

func (s *CardDeck) PullCard() (card FlashCard) {
	// Retrieve card from Draw Pile.
	// Place it in Discard Pile (with recorded Answer?)
}

func (s *CardDeck) ShuffleDeck() {
	// rearrange card deck in random order.
}

func (s *CardDeck) SaveDeck() {
	// Save deck in current state (including card order).
	// That way progress can be picked up at a alter date.
}
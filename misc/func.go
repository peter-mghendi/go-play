package misc

// This file contains all central functions

// GetCard returns a string representation of a Card var
func GetCard(cardStruct *Card) string {
	return cardStruct.Face + " of " + cardStruct.Suit
}

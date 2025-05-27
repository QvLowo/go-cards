package main

func main() {
	cards := newDeck()
	// 因為都在同一個package,所以可以直接使用deal方法,不用額外import等等
	hand, remainingCards := deal(cards, 5)
	// hand & remainingCards的型態都是deck
	hand.print()           // deal第一個回傳的值
	remainingCards.print() // deal第二個回傳的值

}

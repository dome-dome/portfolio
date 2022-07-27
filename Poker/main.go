package main

func main() {
	cards := make([]*card, 0, 53)
	hand := make([]*card, 0, 10)
	suit0, suit1, suit2, suit3 := "♠", "♣", "◆", "♥"

	all := &(all{
		cards: cards,
		hand:  hand,
		suits: []*string{&suit0, &suit1, &suit2, &suit3}, //煩雑だがcard型の定義を満たすためにsuit0~3はポインタ型にする
	})

	all.set()

	all.start()
}

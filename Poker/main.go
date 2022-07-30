package main

//パッケージ変数にすることで他のファイル内でも使用可能になる.
var suit0, suit1, suit2, suit3 string = "♠", "♣", "◆", "♥"

func main() {
	cards := make([]*card, 0, 53)
	hand := make([]*card, 0, 10)

	all := &(all{
		cards: cards,
		hand:  hand,
		suits: []*string{&suit0, &suit1, &suit2, &suit3}, //煩雑だがcard型の定義を満たすためにsuit0~3はポインタ型にする
	})

	all.set()

	all.start()
}

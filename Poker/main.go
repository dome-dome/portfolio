package main

var roles = &([]string{
	"1.ロイヤルストレートフラッシュ",
	"2.ストレートフラッシュ",
	"3.フォーカード",
	"4.フルハウス",
	"5.フラッシュ",
	"6.ストレート",
	"7.スリーカード",
	"8.ツーペア",
	"9.ワンペア",
})

var suit0, suit1, suit2, suit3 string = "♠", "♣", "◆", "♥"
var suits = &([]*string{&suit0, &suit1, &suit2, &suit3})	//煩雑だがcard型の定義を満たすためにsuit0~3はポインタ型にする

var cards = make([]*card, 0, 53)
var hand = make([]*card, 0, 10)


func main() {
	
	for _, suit := range *suits {
		for i := 1; i <= 13; i++ {
			cards = append(cards, &(card{	//&(値)　と同じ
				suit : suit,
				number : i,
			}))
		}
	}

var n int = 5
for i := 0; i < 3; i++ {
	hand, cards = drawHand(cards, hand, n)

	hand, ok_royalst, countRoyalst := outputHand(hand)

	bit, countFla, countStr := judgeHand(hand, ok_royalst)

	outputRole(bit, roles)

	switch {
	case countRoyalst == 4 && countFla == 4:
		println("1枚交換すればロイヤルストレートフラッシュになるかも？")
	case countStr == 4:
		println("1枚交換すればストレートになるかも？")
	case countFla == 4:
		println("1枚交換すればフラッシュになるかも？")
	}
	hand, cards, n = selfChange(hand, cards)
}
}
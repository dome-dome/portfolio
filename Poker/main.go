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

var cs = make([]*card, 0, 53)


func main() {
	
	for _, suit := range *suits {
		for i := 1; i <= 13; i++ {
			cs = append(cs, &(card{	//&(値)　と同じ
				suit : suit,
				number : i,
			}))
		}
	}

var n int
for i := 0; i < 3; i++ {
	hand, cards := drawHand(cs, 5)

	hand, ok_royalst, countRoyal := outputHand(hand)

	bit, countFla, countStr := judgeHand(hand, ok_royalst)

	outputRole(bit, roles)

	hand, cards, n = change(hand, cards)
}
}
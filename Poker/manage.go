package main

import "fmt"

type card struct {
	suit   *string
	number int
}

type all struct { //allをポインタにして, allからcardsやhandにアクセスする.
	cards []*card
	hand  []*card
	suits []*string
}

func (a *all) set() {

	for _, suit := range a.suits {
		for i := 1; i <= 13; i++ {
			a.cards = append(a.cards, &(card{ //&(値)　と同じ
				suit:   suit,
				number: i,
			}))
		}
	}

}

var last = 0
var cnt = 0

func (a *all) start() {
	for i := 0; i < 3; i++ {

		a.drawHand(cnt, last)

		ok_RylSt, cntRylSt := a.outputHand()

		bit, cntFl := a.judgeHand(ok_RylSt)

		outputRole(bit)

		switch {
		case cntRylSt == 4 && cntFl == 4:
			println("1枚交換すればロイヤルストレートフラッシュになるかも？")
		case cntFl == 4:
			println("1枚交換すればフラッシュになるかも？")
		}
		fmt.Printf("%d回目の交換\n", i+1)

		last = a.selfChange()

		cnt++

		//3回目の交換の結果
		if i == 2 {
			a.drawHand(cnt, last)

			ok_RylSt, cntRylSt = a.outputHand()

			bit, cntFl = a.judgeHand(ok_RylSt)

			outputRole(bit)
		}
	}
}

// パッケージ変数なのでoutputRole()内で引数として渡さなくても使える
var roles = &([]string{
	"No.1: ロイヤルストレートフラッシュ",
	"No.2: ストレートフラッシュ",
	"No.3: フォーカード",
	"No.4: フルハウス",
	"No.5: フラッシュ",
	"No.6: ストレート",
	"No.7: スリーカード",
	"No.8: ツーペア",
	"No.9: ワンペア",
})

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	//"golang.org/x/exp/slices"
)

func (a *all) drawHand(cnt int, last int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a.cards), func(i, j int) { a.cards[i], a.cards[j] = a.cards[j], a.cards[i] })

	if cnt == 0 {
		a.hand = a.cards[:5]
		a.cards = a.cards[5:]
	} else if last > 0 {
		a.hand = append(a.hand[last:], a.cards[:last]...)
		a.cards = a.cards[last:]
	}
}

//S
//シード値の設定
//山札のシャッフル
//手札を引き, 山札を決める

func (a *all) outputHand() (bool, int) {
	ok_RylSt := true
	var cntRylSt int

	//ユーザ定義型のスライスなのでslices.Sort(hand)はできない
	sort.Slice(a.hand, func(i, j int) bool { return a.hand[i].number < a.hand[j].number })

	for i, c := range a.hand {
		switch c.number {
		case 1:
			cntRylSt++
			fmt.Printf("%d. %s A\n", i+1, *(c.suit)) //c.suitは*string型なのでデリファレンスする.
		case 10:
			cntRylSt++
			fmt.Printf("%d. %s 10\n", i+1, *(c.suit))
		case 11:
			cntRylSt++
			fmt.Printf("%d. %s J\n", i+1, *(c.suit))
		case 12:
			cntRylSt++
			fmt.Printf("%d. %s Q\n", i+1, *(c.suit))
		case 13:
			cntRylSt++
			fmt.Printf("%d. %s K\n", i+1, *(c.suit))
		default:
			fmt.Printf("%d. %s %d\n", i+1, *(c.suit), c.number)
			ok_RylSt = false
		}
	}

	return ok_RylSt, cntRylSt
}

//A
//手札のソート
//手札の出力(特別な数字は柄で出力)

func (a *all) judgeHand(ok_RylSt bool) (int, int, int) {

	var check = make([]int, 14) //c.number == 13のときのために注意して14にする.

	for _, c := range a.hand {
		check[c.number]++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(check)))

	max, nmax := check[0], check[1]
	var bit int
	switch max {
	case 4:
		bit |= (1 << 2) //フォーカード

	case 3:
		bit |= (1 << 6) //スリーカード

		if nmax == 2 {
			bit |= (1 << 3) //フルハウス
		}

	case 2:
		bit |= (1 << 8) //ワンペア

		if nmax == 2 {
			bit |= (1 << 7) //ツーペア
		}

	}

	var ok_Fl, ok_St = true, true
	cntFl, cntSt := 1, 1
	for i := 0; i+1 < 5; i++ {
		if *(a.hand[i].suit) != *(a.hand[i+1].suit) { //suit *string
			ok_Fl = false
		} else {
			cntFl++
		}

		if a.hand[i].number != a.hand[i+1].number {
			ok_St = false
		} else {
			cntSt++
		}
	}

	switch {
	case ok_Fl:
		bit |= (1 << 4) //フラッシュ

	case ok_St:
		bit |= (1 << 5) //ストレート

	case ok_Fl && ok_St:
		bit |= (1 << 1) //ストフラ

	case ok_Fl && *&ok_RylSt:
		bit |= (1 << 0) //ロイヤルストフラ
	}

	return bit, cntFl, cntSt
}

//B
//sliceを利用した役の判定と出力

func outputRole(bit int, roles *[]string) {
	for i := 0; i < 9; i++ {
		if bit&(1<<i) == (1 << i) { //if bit & (1<<i) == 1	としてはいけない. 1 = (1<<0) だから.
			println("現在の役: ", (*roles)[i])
			println("------------------------------------------")
			return //役があれば終了.
		}
	}
	println("現在役なし")
	println("------------------------------------------")
}

/*
ファイブカード
1.ロイヤルストレートフラッシュ1
2.ストレートフラッシュ10
3.フォーカード100
4.フルハウス1000
5.フラッシュ10000
6.ストレート100000
7.スリーカード1000000
8.ツーペア10000000
9.ワンペア100000000
*/

func (a *all) selfChange() int {
	var memory = make([]bool, 5)

	for {
		print("入れ替えたいカードの番号を選んで下さい, 0を押したら交換終了, 6を押したら変更をもとに戻します. > ")

		var n int
		fmt.Scan(&n)
		if n < 0 || n > 6 {
			for n < 0 || n > 6 {
				println("0~6の番号を入力して下さい.")
				fmt.Scan(&n)
			}
		}

		switch {
		case n >= 1 && n <= 5:
			memory[n-1] = true
			fmt.Printf("-------%d番目のカードを交換しました-------\n", n)

		case n == 0:
		LOOP1:
			for {
				var witch string
				fmt.Printf("交換を終わりますか？ [yes/no (y/n)] > ")
				fmt.Scan(&witch)

				switch witch {
				case "yes", "y":
					var last int
					for i, ok := range memory {
						if ok {
							a.cards = append(a.cards, a.hand[i])

							//入れ替えるカードの数字を0にしてソートする.
							a.hand[i] = zerocard
							last++
						}
					}

					println("-------交換を終わります-------")
					sort.Slice(a.hand, func(i, j int) bool { return a.hand[i].number < a.hand[j].number })
					return last

				case "no", "n":
					println("続けます.")
					break LOOP1

				default:

				}
			}

		case n == 6:
		LOOP2:
			for {
				var witch string
				fmt.Printf("交換をもとに戻しますか？ [yes/no (y/n)] > ")
				fmt.Scan(&witch)

				switch witch {
				case "yes", "y":
					memory = make([]bool, 5)
					println("-------戻しました-------")
					break LOOP2
				case "no", "n":
					println("続けます.")
					break LOOP2
				default:

				}
			}
		}

	}
}

//C
//self change
//手札の交換.山札へ戻しSへ

var suit0 = "♠"
var zerocard = &(card{
	suit:   &suit0,
	number: 0,
})

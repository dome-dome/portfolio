package main
//sliceはポインタにすると操作が面倒だと思う
//だからユーザ定義型の変数をポインタにしたものにスライスを入れてるのでは

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	//"golang.org/x/exp/slices"
)

type card struct {
	suit   *string
	number int
}

func drawHand(cards *[]*card, hand *[]*card) (*[]*card, *[]*card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*cards), func(i, j int) { (*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i] })

	for i := 0; i < len(*hand); i++ {
		//println(hand[i].number) ///////test
		if (*hand)[i].number != 0 {
			*hand = append((*hand)[i:], (*cards)[:i]...) //...必要っぽい
			(*cards) = (*cards)[i:]
		}
	}

	return hand, cards
}

//S
//シード値の設定
//山札のシャッフル
//手札をn枚引き, 山札を決める

func outputHand(hand *[]*card, ok_royalst *bool, countRoyalst *int) (*[]*card, *bool, *int) {
	//ユーザ定義型のスライスなのでslices.Sort(hand)はできない
	sort.Slice((*hand), func(i, j int) bool { return (*hand)[i].number < (*hand)[j].number })

	
	for i, c := range (*hand) {
		switch c.number {
		case 1:
			*countRoyalst++
			fmt.Printf("%d. %s A\n", i+1, *(c.suit)) ////!!!!!!!
		case 10:
			*countRoyalst++
			fmt.Printf("%d. %s 10\n", i+1, *(c.suit))
		case 11:
			*countRoyalst++
			fmt.Printf("%d. %s J\n", i+1, *(c.suit))
		case 12:
			*countRoyalst++
			fmt.Printf("%d. %s Q\n", i+1, *(c.suit))
		case 13:
			*countRoyalst++
			fmt.Printf("%d. %s K\n", i+1, *(c.suit))
		default:
			fmt.Printf("%d. %s %d\n", i+1, *(c.suit), c.number)
			*ok_royalst = false
		}
	}

	return hand, ok_royalst, countRoyalst
}

//A
//手札のソート
//手札の出力(特別な数字は柄で出力)

func judgeHand(hand *[]*card, ok_royalst *bool) (int, int, int) {

	var check = make([]int, 14) //!!!!!!!!!!! 13にしてたせいでc.number == 13のときバグ発生

	for _, c := range *hand {
		check[c.number]++ //????????????アクセスしてint型を表示しているのになんの問題？
	}

	//sort.Ints(check)だけで添字を逆順にしても良い
	sort.Sort(sort.Reverse(sort.IntSlice(check)))
	/* sortパッケージを使わないでcheckを逆順ソート
	for left, right := 0, len(check)-1; left < right; left, right = left+1, right-1 {
		check[left], check[right] = check[right], check[left]
	}
	*/

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

	var ok_flash, ok_straight = true, true
	var countFla, countStr int
	for i := 0; i+1 < 5; i++ {
		if *((*hand)[i].suit) != *((*hand)[i+1].suit) { //hand[i]もポインタだがhand[i]経由でsuit *stringにアクセスしている.
			ok_flash = false
		} else {
			countFla++
		}

		if (*hand)[i].number != (*hand)[i+1].number {
			ok_straight = false
		} else {
			countStr++
		}
	}

	switch {
	case ok_flash:
		bit |= (1 << 4) //フラッシュ

	case ok_straight:
		bit |= (1 << 5) //ストレート

	case ok_flash && ok_straight:
		bit |= (1 << 1) //ストフラ

	case ok_flash && *ok_royalst:
		bit |= (1 << 0) //ロイヤルストフラ
	}

	return bit, countFla, countStr
}

//B
//mapを利用した役の判定と出力	//sliceの方が良い
//役によってコインのオッズも変わる

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

func selfChange(hand *[]*card, cards *[]*card) (*[]*card, *[]*card) {
	var memory = make([]bool, 5) // handがポインタ型だったらこれをやっても無駄.newhandが書き換えられたらhandも書き換えられる.

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
					for i, ok := range memory {
						if ok {
							*cards = append(*cards, (*hand)[i])

							/*"golang.org/x/exp/slices"パッケージを使う場合
							hand = slices.Delete(hand, n-1, n)
							*/
							//hand = append(hand[:i], hand[i:]...)
							//!!!! ...が重要
							//これをやると	handの長さと添字がおかしくなる

							//入れ替えるカードの数字を0にしてソートする.
							(*hand)[i] = zerocard
						}
					}

					println("-------交換を終わります-------")
					sort.Slice(*hand, func(i, j int) bool { return (*hand)[i].number < (*hand)[j].number })
					return hand, cards

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

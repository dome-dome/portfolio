package main

import "fmt"

func (qslice *qslice) start() {
	for {
		//0-index
		num := qslice.inputQuizNumber()

		if num == -1 {
			return
		}

		Q := qslice.qs[num]
		fmt.Printf("%d. %s\n", num+1, Q.name)
		for i, c := range Q.choices {
			fmt.Printf("%d. %s\n", i+1, c)
		}

		for count := 1; count < 3; count++ {
			choice := Q.inputChoice() //0-index

			if choice == Q.ans { //0-index, 0-index

				println("正解！")
				println("--------------------------")
				break
			} else {
				switch count {
				case 1:
					println("もう一度！")

				case 2:
					println("残念、不正解！")
					println("--------------------------")
				}
			}
		}
	}
}

func (Q *quiz) inputChoice() int {
	var choice int
	for choice < 1 || choice > len(Q.choices) {
		print("回答 > ")
		fmt.Scanln(&choice)
	}

	return choice - 1 //0-index
}

func (qslice *qslice) inputQuizNumber() int {

	for i, Q := range qslice.qs {
		fmt.Printf("クイズ %d: %s\n", i+1, Q.name)
	}
	fmt.Printf("%d: 終了する\n", len(qslice.qs)+1)

	var num int
	for num < 1 || num > len(qslice.qs)+1 {
		print("番号を入力して下さい > ")
		fmt.Scan(&num)
	}

	if num == len(qslice.qs)+1 {
		println("終了")
		return -1
	}

	return num - 1 //0-index
}

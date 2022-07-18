package main

type quiz struct {
	name    string
	choices []string
	ans     int
}

type qslice struct {
	qs []*quiz
}

func question01() *quiz {
	return &quiz{
		name:    "偶数を選ぼう",
		choices: []string{"1", "7", "11", "4", "19"},
		ans:     3, //0-index
	}
}

func question02() *quiz {
	return &quiz{
		name:    "数字の5は奇数であり素数でもあるか？",
		choices: []string{"はい", "いいえ"},
		ans:     0,
	}
}

func main() {
	qslice := &qslice{
		qs: []*quiz{
			question01(), question02(),
		},
	}

	qslice.start()
}

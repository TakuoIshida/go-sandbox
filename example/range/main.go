package main

// range ループ構文ないの値要素はコピーである
// ポインタを使う or 別の要素に代入する or　rangeのインデックスを使う
func main() {
	type account struct {
		balance float64
	}
	accounts := []account{
		{balance: 1.0},
		{balance: 1.0},
		{balance: 1.0},
	}

	for _, a := range accounts {
		a.balance = 2.0
	}
	println(accounts[0].balance) // 1.0のまま

	for i := range accounts {
		accounts[i].balance = 2.0
	}
	println(accounts[0].balance) // 2.0に変更される

	accounts2 := []*account{
		{balance: 1.0},
		{balance: 1.0},
	}

	for _, a := range accounts2 {
		println(a.balance) // 1.0のまま
		a.balance = 2.0
		println(a.balance) // 1.0のまま
	}

	// 終了しないroop
	s := []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		println(s[i])
		if i == 10 {
			break
		}
		s = append(s, i+1)
	}
	println("break")
}

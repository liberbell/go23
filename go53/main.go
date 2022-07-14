package main

func main() {
	spy, _ := quoto.NewQuotoFromYahoo(
		"spy", "2022-07-01", "2022-07-14", quote.Daily, true,
	)
}

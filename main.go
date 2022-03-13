package main

import (
	"fmt"

	dotago "github.com/sepehr500/dota-go/dota"
)

var key = ""

func debugPrint(str interface{}) {
	fmt.Printf("%+v\n", str)
}

func main() {
	client := dotago.New(key)
	params := &dotago.MatchHistoryParams{AccountID: "41051979"}
	result, _ := client.GetMatchHistory(params)
	debugPrint(result)
}

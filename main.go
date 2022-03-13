package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	dotago "github.com/sepehr500/dota-go/dota"
)

func debugPrint(str interface{}) {
	fmt.Printf("%+v\n", str)
}

func main() {
	godotenv.Load()

	var key = os.Getenv("DOTA_KEY")
	client := dotago.New(key)
	params := &dotago.MatchHistoryParams{AccountID: "41051979"}
	result, _ := client.GetMatchHistory(params)
	debugPrint(result)
}

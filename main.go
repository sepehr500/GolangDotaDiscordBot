package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/joho/godotenv"
	dotago "github.com/sepehr500/dota-go/dota"
)

var heroData dotago.HeroData

func debugPrint(str interface{}) {
	fmt.Printf("%+v\n", str)
}

type GetMatchData struct {
	CleanHeroName   string
	IsRadiantWin    bool
	IsPlayerRadiant bool
	Kills           int
	Deaths          int
	Assists         int
	EndTime         int
}

func getMatchData(matchData *dotago.MatchDetailsResult, accountId int) GetMatchData {
	startTime := matchData.Result.StartTime
	duration := matchData.Result.Duration
	endTime := startTime + duration
	var currentPlayer *dotago.Player = &dotago.Player{}
	for i, s := range matchData.Result.Players {
		if s.AccountID == accountId {
			currentPlayer = &matchData.Result.Players[i]
			break
		}
	}
	isRadiantWin := matchData.Result.RadiantWin
	isPlayerRadiant := currentPlayer.PlayerSlot < 128
	cleanHeroName := heroData[fmt.Sprint(currentPlayer.HeroID)].Name
	return GetMatchData{
		CleanHeroName:   cleanHeroName,
		IsRadiantWin:    isRadiantWin,
		IsPlayerRadiant: isPlayerRadiant,
		Kills:           currentPlayer.Kills,
		Deaths:          currentPlayer.Deaths,
		Assists:         currentPlayer.Assists,
		EndTime:         endTime,
	}
}

type GameSummaryResult struct {
	TotalWins   int
	TotalLosses int
	TotalGames  int
	WinRate     int
}

func getWeekGameSummery(accountId int, client *dotago.Client) GameSummaryResult {
	matchHistory, _ := client.GetMatchHistory(&dotago.MatchHistoryParams{AccountID: accountId})
	sunday := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, -1*int(time.Now().Weekday()))
	thisWeeksMatches := []dotago.MatchHistory{}
	for i, match := range matchHistory.Result.Matches {
		if time.Unix(int64(match.StartTime), 0).After(sunday) {
			thisWeeksMatches = append(thisWeeksMatches, matchHistory.Result.Matches[i])
		}
	}
	totalWins := 0
	totalLosses := 0
	totalGames := len(thisWeeksMatches)
	for _, match := range thisWeeksMatches {
		matchData, _ := client.GetMatchDetails(&dotago.MatchDetailsParams{MatchID: match.MatchID})
		matchSummary := getMatchData(matchData, accountId)
		if matchData.Result.RadiantWin == matchSummary.IsPlayerRadiant {
			totalWins += 1
		} else {
			totalLosses += 1
		}
	}
	winRate := int((float64(totalWins) / float64(totalGames)) * 100)
	return GameSummaryResult{
		TotalWins:   totalWins,
		TotalLosses: totalLosses,
		TotalGames:  totalGames,
		WinRate:     winRate,
	}
}

func main() {
	println("RUNNING")
	godotenv.Load()

	jsonFile, _ := os.Open("herodata.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &heroData)
	defer jsonFile.Close()

	var key = os.Getenv("DOTA_KEY")
	client := dotago.New(key)
	result := getWeekGameSummery(41051979, client)
	debugPrint(result)
	// params := &dotago.MatchHistoryParams{AccountID: "41051979"}
	// result, _ := client.GetMatchHistory(params)
	// for i, s := range result.Result.Matches {
	// 	fmt.Println(i, time.Unix(int64(s.StartTime), 0))
	// }
	// debugPrint(result)
}

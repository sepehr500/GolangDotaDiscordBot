package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	dotago "github.com/sepehr500/dota-go/dota"
)

const CHANNEL_ID = "801267428830085144"

var EmojiDictionary = map[string]string{
	"WIN":      "✅",
	"LOSS":     "❌",
	"TERRIBLE": "🤮",
	"BAD":      "🤕",
	"ALERT":    "🚨",
}

// tracks the most recent game time for each player
var latestGameTimeMap = map[int]time.Time{}

// Convert to map?
var playerArray = []dotago.PlayerData{
	{
		ID:   83516914,
		Name: "XANNY",
	},
	{
		ID:   106795090,
		Name: "zanerang",
	},
	{
		ID:   253318253,
		Name: "phil",
	},
	{
		ID:   41051979,
		Name: "DependencyInjection",
	},
	{
		ID:   41121344,
		Name: "Shyan",
	},
	{
		ID:   114907302,
		Name: "YahBoyChoi",
	},
}

var heroData dotago.HeroData

func debugPrint(str interface{}) {
	fmt.Printf("%+v\n", str)
}

type GetMatchData struct {
	GameID          int64
	AccountID       int
	CleanHeroName   string
	IsRadiantWin    bool
	IsPlayerRadiant bool
	Kills           int
	Deaths          int
	Assists         int
	EndTime         int
	IsWinner        bool
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
	cleanHeroName := strings.Title(strings.Split(heroData[fmt.Sprint(currentPlayer.HeroID)].Name, "npc_dota_hero_")[1])
	return GetMatchData{
		CleanHeroName:   cleanHeroName,
		IsRadiantWin:    isRadiantWin,
		IsPlayerRadiant: isPlayerRadiant,
		Kills:           currentPlayer.Kills,
		Deaths:          currentPlayer.Deaths,
		Assists:         currentPlayer.Assists,
		EndTime:         endTime,
		IsWinner:        isPlayerRadiant == isRadiantWin,
		AccountID:       accountId,
		GameID:          matchData.Result.MatchID,
	}
}

type GameSummaryResult struct {
	TotalWins   int
	TotalLosses int
	TotalGames  int
	WinRate     int
	AccountID   int
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
	for _, match := range thisWeeksMatches {
		matchData, _ := client.GetMatchDetails(&dotago.MatchDetailsParams{MatchID: match.MatchID})
		// 7 is ranked
		if matchData.Result.LobbyType != 7 {
			continue
		}
		matchSummary := getMatchData(matchData, accountId)
		if matchData.Result.RadiantWin == matchSummary.IsPlayerRadiant {
			totalWins += 1
		} else {
			totalLosses += 1
		}
	}
	totalGames := totalWins + totalLosses
	winRate := int((float64(totalWins) / float64(totalGames)) * 100)
	if totalGames == 0 {
		winRate = 0
	}
	return GameSummaryResult{
		TotalWins:   totalWins,
		TotalLosses: totalLosses,
		TotalGames:  totalGames,
		WinRate:     winRate,
		AccountID:   accountId,
	}
}

func getAllPlayerStatsForWeek(client *dotago.Client) string {
	summaries := []GameSummaryResult{}
	for _, player := range playerArray {
		summaries = append(summaries, getWeekGameSummery(player.ID, client))
	}
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].WinRate > summaries[j].WinRate
	})
	message := "Weekly Game Summary\n\n"
	for i, summary := range summaries {
		name := ""
		for _, player := range playerArray {
			if player.ID == summary.AccountID {
				name = player.Name
				break
			}
		}
		message = message + fmt.Sprintf("%d. %s (%d games) - %d%%\n", i+1, name, summary.TotalGames, summary.WinRate)

	}
	return message
}

func getMostRecentGame(accountId int, client *dotago.Client) (GetMatchData, error) {
	matchHistory, err := client.GetMatchHistory(&dotago.MatchHistoryParams{AccountID: accountId})
	if err != nil {
		return GetMatchData{}, err
	}
	match := matchHistory.Result.Matches[0]
	matchData, err := client.GetMatchDetails(&dotago.MatchDetailsParams{MatchID: match.MatchID})
	if err != nil {
		return GetMatchData{}, err
	}
	matchSummary := getMatchData(matchData, accountId)
	return matchSummary, nil
}

func parsedMostRecentGame(matchData GetMatchData) string {
	feedMessage := ""
	wonEmoji := ""
	wonString := ""
	userName := ""
	for _, player := range playerArray {
		if player.ID == matchData.AccountID {
			userName = player.Name
			break
		}
	}
	if !matchData.IsWinner && matchData.Deaths > matchData.Kills+2 {
		feedMessage = EmojiDictionary["ALERT"] + " FEED ALERT " + EmojiDictionary["ALERT"]
	}
	if matchData.IsWinner {
		wonEmoji = EmojiDictionary["WIN"]
		wonString = "won"
	}
	if !matchData.IsWinner {
		wonEmoji = EmojiDictionary["LOSS"]
		wonString = "lost"
	}
	dotaWebsiteLink := fmt.Sprintf("https://www.opendota.com/matches/%d", matchData.GameID)
	return fmt.Sprintf("%s %s %s %s has %s with %d kills, %d deaths and %d assists.\n%s", feedMessage, wonEmoji, userName, matchData.CleanHeroName, wonString, matchData.Kills, matchData.Deaths, matchData.Assists, dotaWebsiteLink)
}

func pollForMostRecentGames(client *dotago.Client, discord *discordgo.Session) {
	for {
		log.Println("Polling for most recent games")
		for _, player := range playerArray {
			mostRecentGame, err := getMostRecentGame(player.ID, client)
			if err != nil {
				log.Println("ERROR")
				continue
			}
			gameEndTime := time.Unix(int64(mostRecentGame.EndTime), 0)
			latestGameTime, ok := latestGameTimeMap[player.ID]
			if !ok {
				latestGameTimeMap[player.ID] = gameEndTime
				continue
			}
			if gameEndTime.After(latestGameTime) {
				latestGameTimeMap[player.ID] = gameEndTime
				mostRecentGameString := parsedMostRecentGame(mostRecentGame)
				discord.ChannelMessageSend(CHANNEL_ID, mostRecentGameString)
				log.Println("Sent message:", mostRecentGameString)
			}
		}
		time.Sleep(time.Minute)
	}
}

func weeklySummaryPolling(client *dotago.Client, discord *discordgo.Session) {
	time.Sleep(time.Hour * 12)
	for {
		log.Println("Polling for weekly summary")
		message := getAllPlayerStatsForWeek(client)
		discord.ChannelMessageSend(CHANNEL_ID, message)
		log.Println("Sent message:", message)
		time.Sleep(time.Hour * 12)
	}
}

func messageCreate(client *dotago.Client) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == "!weekly-summary" {
			message := getAllPlayerStatsForWeek(client)
			println("MESSAGE", message)
			s.ChannelMessageSend(CHANNEL_ID, message)
		}
	}
}

func main() {
	println("RUNNING")
	godotenv.Load()
	var key = os.Getenv("DOTA_KEY")
	client := dotago.New(key)

	discord, _ := discordgo.New("Bot " + os.Getenv("DISCORD_KEY"))
	discord.AddHandler(messageCreate(client))
	discord.Open()
	defer discord.Close()

	jsonFile, _ := os.Open("herodata.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &heroData)
	defer jsonFile.Close()

	go pollForMostRecentGames(client, discord)
	go weeklySummaryPolling(client, discord)
	select {}
}

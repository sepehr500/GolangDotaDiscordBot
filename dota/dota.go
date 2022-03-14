package dotago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

const API_URL = "https://api.steampowered.com/"

var matchCache map[int]*MatchDetailsResult = make(map[int]*MatchDetailsResult)

func New(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) buildURL(resourcename string, params interface{}) string {
	// Figure out how to assign key to this param object
	query, _ := query.Values(params)
	return API_URL + fmt.Sprintf("IDOTA2Match_570/%s/V001?key=%s&", resourcename, c.Token) + query.Encode()
}

func (c *Client) GetMatchHistory(params *MatchHistoryParams) (*MatchHistoryResult, error) {
	resp, getErr := http.Get(c.buildURL("GetMatchHistory", params))
	if getErr != nil {
		return nil, getErr
	}
	defer resp.Body.Close()
	matchHistory := &MatchHistoryResult{}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	err := json.Unmarshal(body, matchHistory)
	if err != nil {
		return nil, err
	}
	return matchHistory, nil
}

func (c *Client) GetMatchDetails(params *MatchDetailsParams) (*MatchDetailsResult, error) {
	if matchCache[params.MatchID] != nil {
		return matchCache[params.MatchID], nil
	}
	resp, getErr := http.Get(c.buildURL("GetMatchDetails", params))
	if getErr != nil {
		log.Fatal(getErr)
	}
	defer resp.Body.Close()
	matchDetails := &MatchDetailsResult{}
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, matchDetails)
	if err != nil {
		println(err.Error())
	}
	matchCache[params.MatchID] = matchDetails
	return matchDetails, err
}

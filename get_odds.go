package main

import (
	"fmt"
)

func GetActiveSports() []string {
	oddsApiKey := getOddsApiKey()
	formattedEndpoint := fmt.Sprintf("/v3/sports/?apiKey=%s", oddsApiKey)
	respBodyByteArr := makeApiRequest(formattedEndpoint)
	decodedResp := processActiveSportsResponse(respBodyByteArr)
	return getListOfSportsFromActiveResp(&decodedResp)
}

func GetNflTotalsOdds() []string {
	oddsApiKey := getOddsApiKey()
	formattedEndpoint := fmt.Sprintf("/v3/odds/?apiKey=%s&sport=americanfootball_nfl&region=%s&mkt=totals&oddsFormat=american", oddsApiKey, region)
	respBodyByteArr := makeApiRequest(formattedEndpoint)
	decodedResp := processNflTotalsResponse(respBodyByteArr)
	formattedNflTotalsOdds := formatNflTotalsResp(decodedResp)
	var retArr []string
	for _, odds := range formattedNflTotalsOdds.Odds {
		var formattedString string = ""
		formattedString += fmt.Sprintf("%sgame. Over Under: %0.1f, Over odds: %d; Under Odds: %d", odds.Teams, odds.Over, odds.OverOdds, odds.UnderOdds)
		retArr = append(retArr, formattedString)
	}
	return retArr
}

func GetNflH2hOdds() { // TODO - IMPLEMENT []string {
	oddsApiKey := getOddsApiKey()
	formattedEndpoint := fmt.Sprintf("/v3/odds/?apiKey=%s&sport=americanfootball_nfl&region=%s&mkt=h2h&oddsFormat=american", oddsApiKey, region)
	respBodyByteArr := makeApiRequest(formattedEndpoint)
	fmt.Printf("%T", respBodyByteArr)
}

func GetNflSpreadsOdds() { // TODO - IMPLEMENT []string {
	oddsApiKey := getOddsApiKey()
	formattedEndpoint := fmt.Sprintf("/v3/odds/?apiKey=%s&sport=americanfootball_nfl&region=%s&mkt=spreads&oddsFormat=american", oddsApiKey, region)
	respBodyByteArr := makeApiRequest(formattedEndpoint)
	fmt.Printf("%T", respBodyByteArr)
}

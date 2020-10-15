package main

import (
	"fmt"
)

// GetActiveSports retrieves the sports currently accepting Bets
// according to the-odds-api
func GetActiveSports() []string {
	oddsAPIKey := getOddsAPIKey()
	formattedEndpoint := fmt.Sprintf("/v3/sports/?apiKey=%s", oddsAPIKey)
	respBodyByteArr := makeAPIRequest(formattedEndpoint)
	decodedResp := processActiveSportsResponse(respBodyByteArr)
	return getListOfSportsFromActiveResp(&decodedResp)
}

// GetNflTotalsOdds retrieves an array of strings containing a message
// that can be user-read regarding NFL Over/Unders for the upcoming week
func GetNflTotalsOdds() []string {
	oddsAPIKey := getOddsAPIKey()
	formattedEndpoint := fmt.Sprintf("/v3/odds/?apiKey=%s&sport=americanfootball_nfl&region=%s&mkt=totals&oddsFormat=american", oddsAPIKey, region)
	respBodyByteArr := makeAPIRequest(formattedEndpoint)
	decodedResp := processNflTotalsResponse(respBodyByteArr)
	formattedNflTotalsOdds := formatNflTotalsResp(decodedResp)
	var retArr []string
	for _, odds := range formattedNflTotalsOdds.Odds {
		formattedString := formatNflTotalsMessageString(odds)
		retArr = append(retArr, formattedString)
	}
	return retArr
}

// func GetNflH2hOdds() { // TODO - IMPLEMENT []string {
// 	oddsAPIKey := getOddsAPIKey()
// 	formattedEndpoint := fmt.Sprintf("/v3/odds/?apiKey=%s&sport=americanfootball_nfl&region=%s&mkt=h2h&oddsFormat=american", oddsAPIKey, region)
// 	respBodyByteArr := makeAPIRequest(formattedEndpoint)
// 	fmt.Printf("%T", respBodyByteArr)
// }

// func GetNflSpreadsOdds() { // TODO - IMPLEMENT []string {
// 	oddsAPIKey := getOddsAPIKey()
// 	formattedEndpoint := fmt.Sprintf("/v3/odds/?apiKey=%s&sport=americanfootball_nfl&region=%s&mkt=spreads&oddsFormat=american", oddsAPIKey, region)
// 	respBodyByteArr := makeAPIRequest(formattedEndpoint)
// 	fmt.Printf("%T", respBodyByteArr)
// }

package odds_briefing

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const baseApiUrl string = "https://api.the-odds-api.com"
const region string = "us" // only bet on USA!

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func h2hToAmericanOdds(h2hOdds string) int {
	floatOdds, err := strconv.ParseFloat(h2hOdds, 64)
	if err != nil {
		ErrorHelper(err)
	}
	floatOdds = round(floatOdds, 0.01)
	if floatOdds >= 2.0 {
		return int((floatOdds - 1) * 100)
	}
	return int(-100 / (floatOdds - 1))
}

func makeApiRequest(endpoint string) *io.ReadCloser {
	finalUrl := baseApiUrl + endpoint

	resp, err := http.Get(finalUrl)
	if err != nil {
		ErrorHelper(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("non 200 error code for: %s", finalUrl)
		ErrorHelper(err)
	}

	return &resp.Body
}

func getActiveSports() []string {
	var credFile credentials
	credFile.loadCredentials()
	oddsApiKey := credFile.OddsApiKey
	formattedEndpoint := fmt.Sprintf("/v3/sports/?apiKey=%s", oddsApiKey)
	body := makeApiRequest(formattedEndpoint)
	decoder := json.NewDecoder(*body)

	var ActiveSportsResponse ActiveSportsResponse
	var listOfSports []string
	err := decoder.Decode(&ActiveSportsResponse)
	if err != nil {
		ErrorHelper(err)
	}

	for _, entry := range ActiveSportsResponse.Data {
		listOfSports = append(listOfSports, entry.Key)
	}

	return listOfSports
}

func GetNflOdds() { // []string {
	// get the latest NFL odds
}

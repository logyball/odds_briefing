package oddsapi

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/loganballard/odds_briefing/logger"
	"github.com/stretchr/testify/assert"
)

func getJSONFilepath(filename string) string {
	curDir, err := os.Getwd()
	if err != nil {
		logger.ErrorHelper(err)
	}
	return filepath.Join(curDir, "example_responses", filename)
}

/*
	Helper functions
*/
func TestGetOddsApiKey(t *testing.T) {
	apiKey := getOddsAPIKey()
	assert.NotEmpty(t, apiKey, "odds api key not found")
}

func TestH2hToAmericanOdds(t *testing.T) {
	assert.Equal(t, h2hToAmericanOdds("1.20"), -500)
	assert.Equal(t, h2hToAmericanOdds("1.22"), -454)
	assert.Equal(t, h2hToAmericanOdds("1.25"), -400)
	assert.Equal(t, h2hToAmericanOdds("1.29"), -344)
	assert.Equal(t, h2hToAmericanOdds("2.25"), 125)
	assert.Equal(t, h2hToAmericanOdds("10"), 900)

}

/*
	Active Sports API call
*/

func TestProcessActiveSportsResponse(t *testing.T) {
	activeSportsJSONFilepath := getJSONFilepath("active_sports.json")
	activeSportsByteArr, err := ioutil.ReadFile(activeSportsJSONFilepath)
	if err != nil {
		fmt.Printf("failed reading file: %s\n", err.Error())
		t.FailNow()
	}
	activeSportsResp := processActiveSportsResponse(activeSportsByteArr)
	assert.NotNil(t, activeSportsResp, "active sports resp is nil")
	assert.IsType(t, ActiveSportsResponse{}, activeSportsResp, "active sports resp is wrong type")
	assert.NotEmpty(t, activeSportsResp.Data)
	assert.True(t, activeSportsResp.Success)
}

func TestGetListOfSportsFromActiveResp(t *testing.T) {
	activeSportsJSONFilepath := getJSONFilepath("active_sports.json")
	activeSportsByteArr, err := ioutil.ReadFile(activeSportsJSONFilepath)
	if err != nil {
		fmt.Printf("failed reading file: %s\n", err.Error())
		t.FailNow()
	}
	activeSportsResp := processActiveSportsResponse(activeSportsByteArr)
	listOfSports := getListOfSportsFromActiveResp(&activeSportsResp)
	assert.NotEmpty(t, listOfSports, "sports list is empty")
	assert.Greater(t, len(listOfSports), 0, "sports list is too short")
	for _, sport := range listOfSports {
		assert.NotEmpty(t, sport, "empty sport")
	}
}

/*
	NFL Totals API Call
*/

func TestProcessNFLTotalsResponse(t *testing.T) {
	nflTotalsJSONFilepath := getJSONFilepath("totals_nfl.json")
	nflTotalsByteArr, err := ioutil.ReadFile(nflTotalsJSONFilepath)
	if err != nil {
		fmt.Printf("failed reading file: %s\n", err.Error())
		t.FailNow()
	}
	nflTotalsResp := processNflTotalsResponse(nflTotalsByteArr)
	assert.NotNil(t, nflTotalsResp, "nfl totals  resp is nil")
	assert.IsType(t, totalsOddsResponse{}, nflTotalsResp, "active sports resp is wrong type")
	assert.NotEmpty(t, nflTotalsResp.Games)
	assert.True(t, nflTotalsResp.Success)
	for _, entry := range nflTotalsResp.Games {
		for _, sites := range entry.Sites {
			assert.Greater(t, len(sites.Odds.Totals.Points), 1)
			assert.Greater(t, len(sites.Odds.Totals.Odds), 1)
			assert.Greater(t, len(sites.Odds.Totals.Position), 1)
			assert.Equal(t, len(sites.Odds.Totals.Points), len(sites.Odds.Totals.Odds), len(sites.Odds.Totals.Position))
			for _, pts := range sites.Odds.Totals.Points {
				assert.NotEmpty(t, pts, "points was empty")
			}
			for _, ptsStr := range sites.Odds.Totals.PointsStr {
				assert.NotEmpty(t, ptsStr, "points string was empty")
			}
			for _, ptsFl := range sites.Odds.Totals.PointsFloat {
				assert.NotEmpty(t, ptsFl, "points float was empty")
			}
			for _, odds := range sites.Odds.Totals.Odds {
				assert.NotEmpty(t, odds, "odds was empty")
			}
			for _, pos := range sites.Odds.Totals.Position {
				assert.NotEmpty(t, pos, "pos was empty")
			}
		}
	}
}

func TestFormattingProcessedNflTotalsResp(t *testing.T) {
	nflTotalsJSONFilepath := getJSONFilepath("totals_nfl.json")
	nflTotalsByteArr, err := ioutil.ReadFile(nflTotalsJSONFilepath)
	if err != nil {
		fmt.Printf("failed reading file: %s\n", err.Error())
		t.FailNow()
	}
	nflTotalsResp := processNflTotalsResponse(nflTotalsByteArr)
	totalOdds := formatNflTotalsResp(nflTotalsResp)
	assert.NotEmpty(t, totalOdds.OddsType)
	assert.NotEmpty(t, totalOdds.Sport)
	assert.NotEmpty(t, totalOdds.Odds)
	for _, odd := range totalOdds.Odds {
		assert.NotEmpty(t, odd.Teams)
		assert.NotEmpty(t, odd.Gametime)
		assert.NotEmpty(t, odd.Over)
		assert.NotEmpty(t, odd.Under)
		assert.NotEmpty(t, odd.OverOdds)
		assert.NotEmpty(t, odd.UnderOdds)
	}
}

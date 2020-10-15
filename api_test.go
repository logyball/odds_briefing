// +build api_tests

//
// Only run these tests sometimes to not destroy rate limited APIs
//
// Set the TEST_HARMFUL_STUFF environment variable to "true" if you want to test really dangerous stuff
//
package main

import (
	"os"
	"testing"

	odds "github.com/loganballard/odds_briefing/odds_api"
	"github.com/stretchr/testify/assert"
)

const HARMFUL_TEST_ENV_VAR string = "TWILIO_TESTS"

////// Odds API - Free but limited to 500 req/mo
func TestMakeApiRequestCanReachBaseUrl(t *testing.T) {
	body := odds.MakeAPIRequest("")
	assert.NotNil(t, body, "body of response was nil")
}

// Active Sports
func TestGetActiveSports(t *testing.T) {
	body := odds.GetActiveSports()
	assert.NotNil(t, body, "body of response from get active sports was nil")
}

func TestGetActiveSportsHasAtLeastOneSportWeCareAbout(t *testing.T) {
	body := odds.GetActiveSports()
	assert.NotNil(t, body, "body of response from get active sports was nil")
}

// NFL Totals
func TestGetNflTotalsOdds(t *testing.T) {
	totalOdds := odds.GetNflTotalsOdds()
	assert.NotEmpty(t, totalOdds)
}

////// Twilio API - pay as you go
func TestTwilioApiWrapper(t *testing.T) {
	if os.Getenv(HARMFUL_TEST_ENV_VAR) == "true" {
		err := sendTwilioMsgFromGeneratedOddsData("gambling data!")
		assert.Nil(t, err, "error when sending message")
	}
	t.Skip()
}

func TestSendingFirstTotalsOddsAsMessage(t *testing.T) {
	if os.Getenv(HARMFUL_TEST_ENV_VAR) == "true" {
		err := sendFirstTotalsOddsAsMessage()
		assert.Nil(t, err, "error when sending message")
	}
	t.Skip()
}

func TestSendingFirstFiveTotalsOddsAsMessage(t *testing.T) {
	if os.Getenv(HARMFUL_TEST_ENV_VAR) == "true" {
		err := sendFirstXTotalsOddsAsMessage(5)
		assert.Nil(t, err, "error when sending message")
	}
	t.Skip()
}

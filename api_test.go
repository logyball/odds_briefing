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

	"github.com/stretchr/testify/assert"
)

////// Odds API - Free but limited to 500 req/mo
func TestMakeApiRequestCanReachBaseUrl(t *testing.T) {
	body := makeApiRequest("")
	assert.NotNil(t, body, "body of response was nil")
}

// Active Sports
func TestGetActiveSports(t *testing.T) {
	body := GetActiveSports()
	assert.NotNil(t, body, "body of response from get active sports was nil")
}

func TestGetActiveSportsHasAtLeastOneSportWeCareAbout(t *testing.T) {
	body := GetActiveSports()
	assert.NotNil(t, body, "body of response from get active sports was nil")
}

// NFL Totals
func TestGetNflTotalsOdds(t *testing.T) {
	totalOdds := GetNflTotalsOdds()
	assert.NotEmpty(t, totalOdds)
}

////// Twilio API - pay as you go
func TestTwilioApiWrapper(t *testing.T) {
	err := sendTwilioMsgFromGeneratedOddsData("gambling data!")
	assert.Nil(t, err, "error when sending message")
}

func TestSendingFirstTotalsOddsAsMessage(t *testing.T) {
	err := SendFirstTotalsOddsAsMessage()
	assert.Nil(t, err, "error when sending message")
}

func TestSendingFirstFiveTotalsOddsAsMessage(t *testing.T) {
	if os.Getenv("TEST_HARMFUL_STUFF") == "true" {
		err := SendFirstXTotalsOddsAsMessage(5)
		assert.Nil(t, err, "error when sending message")
	}
	t.Skip()
}

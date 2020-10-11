package odds_briefing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestH2hToAmericanOdds(t *testing.T) {
	assert.Equal(t, h2hToAmericanOdds("1.20"), -500)
	assert.Equal(t, h2hToAmericanOdds("1.22"), -454)
	assert.Equal(t, h2hToAmericanOdds("1.25"), -400)
	assert.Equal(t, h2hToAmericanOdds("1.29"), -344)
	assert.Equal(t, h2hToAmericanOdds("2.25"), 125)
	assert.Equal(t, h2hToAmericanOdds("10"), 900)

}

func TestMakeApiRequestCanReachBaseUrl(t *testing.T) {
	body := makeApiRequest("")
	assert.NotNil(t, body, "body of response was nil")
}

func TestGetActiveSports(t *testing.T) {
	body := getActiveSports()
	assert.NotNil(t, body, "body of response from get active sports was nil")
}

func TestGetActiveSportsHasAtLeastOneSportWeCareAbout(t *testing.T) {
	body := getActiveSports()
	assert.NotNil(t, body, "body of response from get active sports was nil")
}

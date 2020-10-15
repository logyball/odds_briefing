package oddsapi

import "time"

// FormattedTotalsOdds is an array of TotalOdds
type FormattedTotalsOdds struct {
	Sport    string
	OddsType string
	Odds     []TotalOdds
}

// TotalOdds represents the API data return from the-odds-api with over/unders for the NFL
type TotalOdds struct {
	Teams     string
	Gametime  time.Time
	Over      float64
	Under     float64
	OverOdds  int
	UnderOdds int
}

// FormattedH2hOdds is an arr of H2hOdds
type FormattedH2hOdds struct {
	Sport    string
	OddsType string
	Odds     []H2hOdds
}

// H2hOdds represents the moneyline odds return from the-odds-api
type H2hOdds struct {
	Home     string
	Away     string
	HomeOdds int
	AwayOdds int
}

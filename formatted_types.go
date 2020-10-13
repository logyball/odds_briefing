package main

import "time"

type FormattedTotalsOdds struct {
	Sport    string
	OddsType string
	Odds     []TotalOdds
}

type TotalOdds struct {
	Teams     string
	Gametime  time.Time
	Over      float64
	Under     float64
	OverOdds  int
	UnderOdds int
}

type FormattedH2hOdds struct {
	Sport    string
	OddsType string
	Odds     H2hOdds
}

type H2hOdds struct {
	Home     string
	Away     string
	HomeOdds int
	AwayOdds int
}

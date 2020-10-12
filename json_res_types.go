package main

type ActiveSportsResponse struct {
	Success bool                `json:"success"`
	Data    []ActiveSportsEntry `json:"data"`
}

type ActiveSportsEntry struct {
	Key     string `json:"key"`
	Active  bool   `json:"active"`
	Group   string `json:"group"`
	Details string `json:"details"`
	Title   string `json:"title"`
}

type TotalsOddsResponse struct {
	Success bool              `json:"success"`
	Games   []TotalsOddsEntry `json:"data"`
}

type TotalsOddsEntry struct {
	Teams      []string              `json:"teams"`
	Gametime   int64                 `json:"commence_time"`
	Sites      []OddsTotalsSiteEntry `json:"sites"`
	SitesCount int                   `json:"sites_count"`
}

type OddsTotalsSiteEntry struct {
	Site            string     `json:"site_nice"`
	UpdateTimestamp int64      `json:"last_update"`
	Odds            TotalsOdds `json:"odds"`
}

type TotalsOdds struct {
	Totals TotalObj `json:"totals"`
}

type TotalObj struct {
	Points      []interface{} `json:"points"`
	PointsStr   []string
	PointsFloat []float64
	Odds        []int    `json:"odds"`
	Position    []string `json:"position"`
}

type oddsH2hSiteEntry struct {
	site_key    string
	Site_nice   string
	Last_update int64
	Odds        h2hOdds
}

type h2hOdds struct {
	H2h []float64
}

type H2hOddsResponse struct {
	Success bool
	Data    []h2hOddsEntry
}

type h2hOddsEntry struct {
	sport_key     string
	sport_nice    string
	Teams         []string
	commence_time int64
	Home_team     string
	Sites         []oddsH2hSiteEntry
	Sites_count   int
}

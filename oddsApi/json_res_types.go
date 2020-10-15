package oddsApi

// ActiveSportsResponse is an array of sports that are currently gamble-able
type ActiveSportsResponse struct {
	Success bool                `json:"success"`
	Data    []activeSportsEntry `json:"data"`
}

type activeSportsEntry struct {
	Key     string `json:"key"`
	Active  bool   `json:"active"`
	Group   string `json:"group"`
	Details string `json:"details"`
	Title   string `json:"title"`
}

type totalsOddsResponse struct {
	Success bool              `json:"success"`
	Games   []totalsOddsEntry `json:"data"`
}

type totalsOddsEntry struct {
	Teams      []string              `json:"teams"`
	Gametime   int64                 `json:"commence_time"`
	Sites      []oddsTotalsSiteEntry `json:"sites"`
	SitesCount int                   `json:"sites_count"`
}

type oddsTotalsSiteEntry struct {
	Site            string     `json:"site_nice"`
	UpdateTimestamp int64      `json:"last_update"`
	Odds            totalsOdds `json:"odds"`
}

type totalsOdds struct {
	Totals totalObj `json:"totals"`
}

type totalObj struct {
	Points      []interface{} `json:"points"`
	PointsStr   []string
	PointsFloat []float64
	Odds        []int    `json:"odds"`
	Position    []string `json:"position"`
}

type oddsH2hSiteEntry struct {
	SiteKey    string  `json:"site_key"`
	SiteNice   string  `json:"site_nice"`
	LastUpdate int64   `json:"last_update"`
	Odds       h2hOdds `json:"odds"`
}

type h2hOdds struct {
	H2h []float64
}

type h2hOddsResponse struct {
	Success bool
	Data    []h2hOddsEntry
}

type h2hOddsEntry struct {
	SportKey     string             `json:"sport_key"`
	SportNice    string             `json:"sport_nice"`
	Teams        []string           `json:"teams"`
	CommenceTime int64              `json:"commence_time"`
	HomeTeam     string             `json:"home_team"`
	Sites        []oddsH2hSiteEntry `json:"sites"`
	SitesCount   int                `json:"sites_count"`
}

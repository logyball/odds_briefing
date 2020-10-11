package odds_briefing

type ActiveSportsResponse struct {
	Success bool
	Data    []activeSportsEntry
}

type activeSportsEntry struct {
	Key           string
	active        bool
	group         string
	details       string
	title         string
	has_outrights bool
}

type OddsResponse struct {
	Success bool
	data    []oddsEntry
}

type oddsEntry struct {
	sport_key   string
	sport_nice  string
	Teams       []string
	Sites       []oddsSiteEntry
	Sites_count int
}

type oddsSiteEntry struct {
	site_key    string
	Site_nice   string
	Last_update string
	Odds        interface{}
}

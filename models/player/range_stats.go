package player

import "time"

type PlayerRangeStats struct {
	Items []Item `json:"items"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type Item struct {
	RangeStats RangeStats `json:"stats"`
}

type RangeStats struct {
	Rounds           string    `json:"Rounds"`
	GameMode         string    `json:"Game Mode"`
	HeadshotsPercent string    `json:"Headshots %"`
	Score            string    `json:"Score"`
	UpdatedAt        time.Time `json:"Updated At"`
	SecondHalfScore  string    `json:"Second Half Score"`
	Nickname         string    `json:"Nickname"`
	Kills            string    `json:"Kills"`
	FirstHalfScore   string    `json:"First Half Score"`
	PlayerID         string    `json:"Player Id"`
	OvertimeScore    string    `json:"Overtime score"`
	Headshots        string    `json:"Headshots"`
	Game             string    `json:"Game"`
	Assists          string    `json:"Assists"`
	QuadroKills      string    `json:"Quadro Kills"`
	MVPs             string    `json:"MVPs"`
	MatchID          string    `json:"Match Id"`
	Deaths           string    `json:"Deaths"`
	DoubleKills      string    `json:"Double Kills"`
	ADR              string    `json:"ADR"`
	CreatedAt        time.Time `json:"Created At"`
	TripleKills      string    `json:"Triple Kills"`
	CompetitionID    string    `json:"Competition Id"`
	KDRatio          string    `json:"K/D Ratio"`
	BestOf           string    `json:"Best Of"`
	MatchRound       string    `json:"Match Round"`
	Map              string    `json:"Map"`
	Team             string    `json:"Team"`
	FinalScore       string    `json:"Final Score"`
	KRRatio          string    `json:"K/R Ratio"`
	PentaKills       string    `json:"Penta Kills"`
	Region           string    `json:"Region"`
	MatchFinishedAt  int64     `json:"Match Finished At"`
	Result           string    `json:"Result"`
	Winner           string    `json:"Winner"`
}

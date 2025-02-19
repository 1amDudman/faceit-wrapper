package player

type PlayerStats struct {
	PlayerID string    `json:"player_id"`
	GameID   string    `json:"game_id"`
	Lifetime Lifetime  `json:"lifetime"`
	Segments []Segment `json:"segments"`
}

type Lifetime struct {
	OneVOneWinRate               string   `json:"1v1 Win Rate"`
	EnemiesFlashedPerRound       string   `json:"Enemies Flashed per Round"`
	SniperKillRate               string   `json:"Sniper Kill Rate"`
	AverageHeadshotsPercent      string   `json:"Average Headshots %"`
	TotalUtilityCount            string   `json:"Total Utility Count"`
	UtilityDamageSuccessRate     string   `json:"Utility Damage Success Rate"`
	UtilityDamagePerRound        string   `json:"Utility Damage per Round"`
	TotalFlashCount              string   `json:"Total Flash Count"`
	UtilitySuccessRate           string   `json:"Utility Success Rate"`
	TotalRoundsWithExtendedStats string   `json:"Total Rounds with extended stats"`
	TotalOneVOneWins             string   `json:"Total 1v1 Wins"`
	FlashesPerRound              string   `json:"Flashes per Round"`
	TotalMatches                 string   `json:"Total Matches"`
	EntrySuccessRate             string   `json:"Entry Success Rate"`
	TotalUtilitySuccesses        string   `json:"Total Utility Successes"`
	TotalOneVTwoWins             string   `json:"Total 1v2 Wins"`
	Wins                         string   `json:"Wins"`
	AverageKDRatio               string   `json:"Average K/D Ratio"`
	WinRatePercent               string   `json:"Win Rate %"`
	ADR                          string   `json:"ADR"`
	KDRatio                      string   `json:"K/D Ratio"`
	SniperKillRatePerRound       string   `json:"Sniper Kill Rate per Round"`
	TotalKillsWithExtendedStats  string   `json:"Total Kills with extended stats"`
	TotalEntryWins               string   `json:"Total Entry Wins"`
	TotalDamage                  string   `json:"Total Damage"`
	Matches                      string   `json:"Matches"`
	TotalOneVTwoCount            string   `json:"Total 1v2 Count"`
	EntryRate                    string   `json:"Entry Rate"`
	FlashSuccessRate             string   `json:"Flash Success Rate"`
	RecentResults                []string `json:"Recent Results"`
	TotalEntryCount              string   `json:"Total Entry Count"`
	CurrentWinStreak             string   `json:"Current Win Streak"`
	OneVTwoWinRate               string   `json:"1v2 Win Rate"`
	LongestWinStreak             string   `json:"Longest Win Streak"`
	TotalUtilityDamage           string   `json:"Total Utility Damage"`
	TotalSniperKills             string   `json:"Total Sniper Kills"`
	TotalEnemiesFlashed          string   `json:"Total Enemies Flashed"`
	TotalFlashSuccesses          string   `json:"Total Flash Successes"`
	TotalOneVOneCount            string   `json:"Total 1v1 Count"`
	UtilityUsagePerRound         string   `json:"Utility Usage per Round"`
	TotalHeadshotsPercent        string   `json:"Total Headshots %"`
}

type Segment struct {
	Stats      Stats  `json:"stats"`
	Type       string `json:"type"`
	Mode       string `json:"mode"`
	Label      string `json:"label"`
	ImgSmall   string `json:"img_small,omitempty"`
	ImgRegular string `json:"img_regular,omitempty"`
}

type Stats struct {
	TotalEnemiesFlashed          string `json:"Total Enemies Flashed"`
	SniperKillRatePerRound       string `json:"Sniper Kill Rate per Round"`
	TotalFlashSuccesses          string `json:"Total Flash Successes"`
	Headshots                    string `json:"Headshots"`
	AverageTripleKills           string `json:"Average Triple Kills"`
	TotalMatches                 string `json:"Total Matches"`
	FlashesPerRound              string `json:"Flashes per Round"`
	AverageAssists               string `json:"Average Assists"`
	MVPs                         string `json:"MVPs"`
	UtilitySuccessRate           string `json:"Utility Success Rate"`
	EntryRate                    string `json:"Entry Rate"`
	QuadroKills                  string `json:"Quadro Kills"`
	TotalOneVOneWins             string `json:"Total 1v1 Wins"`
	TotalEntryCount              string `json:"Total Entry Count"`
	FlashSuccessRate             string `json:"Flash Success Rate"`
	TotalSniperKills             string `json:"Total Sniper Kills"`
	TotalKillsWithExtendedStats  string `json:"Total Kills with extended stats"`
	AverageQuadroKills           string `json:"Average Quadro Kills"`
	Matches                      string `json:"Matches"`
	AverageKills                 string `json:"Average Kills"`
	TotalUtilitySuccesses        string `json:"Total Utility Successes"`
	UtilityUsagePerRound         string `json:"Utility Usage per Round"`
	AverageMVPs                  string `json:"Average MVPs"`
	TotalUtilityCount            string `json:"Total Utility Count"`
	AverageDeaths                string `json:"Average Deaths"`
	KDRatio                      string `json:"K/D Ratio"`
	HeadshotsPerMatch            string `json:"Headshots per Match"`
	TripleKills                  string `json:"Triple Kills"`
	PentaKills                   string `json:"Penta Kills"`
	TotalFlashCount              string `json:"Total Flash Count"`
	Wins                         string `json:"Wins"`
	AverageHeadshotsPercent      string `json:"Average Headshots %"`
	AverageKDRatio               string `json:"Average K/D Ratio"`
	TotalUtilityDamage           string `json:"Total Utility Damage"`
	OneVOneWinRate               string `json:"1v1 Win Rate"`
	TotalOneVTwoWins             string `json:"Total 1v2 Wins"`
	TotalOneVOneCount            string `json:"Total 1v1 Count"`
	TotalOneVTwoCount            string `json:"Total 1v2 Count"`
	Deaths                       string `json:"Deaths"`
	TotalRoundsWithExtendedStats string `json:"Total Rounds with extended stats"`
	ADR                          string `json:"ADR"`
	Kills                        string `json:"Kills"`
	TotalDamage                  string `json:"Total Damage"`
	Rounds                       string `json:"Rounds"`
	SniperKillRate               string `json:"Sniper Kill Rate"`
	UtilityDamagePerRound        string `json:"Utility Damage per Round"`
	AveragePentaKills            string `json:"Average Penta Kills"`
	Assists                      string `json:"Assists"`
	TotalEntryWins               string `json:"Total Entry Wins"`
	TotalHeadshotsPercent        string `json:"Total Headshots %"`
	KRRatio                      string `json:"K/R Ratio"`
	OneVTwoWinRate               string `json:"1v2 Win Rate"`
	AverageKRRatio               string `json:"Average K/R Ratio"`
	UtilityDamageSuccessRate     string `json:"Utility Damage Success Rate"`
	WinRatePercent               string `json:"Win Rate %"`
	EnemiesFlashedPerRound       string `json:"Enemies Flashed per Round"`
	EntrySuccessRate             string `json:"Entry Success Rate"`
}

package match

type MatchStats struct {
	Rounds []Round `json:"rounds"`
}

type Round struct {
	BestOf        string     `json:"best_of"`
	CompetitionID *string    `json:"competition_id"`
	GameID        string     `json:"game_id"`
	GameMode      string     `json:"game_mode"`
	MatchID       string     `json:"match_id"`
	MatchRound    string     `json:"match_round"`
	Played        string     `json:"played"`
	RoundStats    RoundStats `json:"round_stats"`
	Teams         []Team     `json:"teams"`
}

type RoundStats struct {
	Score  string `json:"Score"`
	Map    string `json:"Map"`
	Region string `json:"Region"`
	Winner string `json:"Winner"`
	Rounds string `json:"Rounds"`
}

type Team struct {
	TeamID    string    `json:"team_id"`
	Premade   bool      `json:"premade"`
	TeamStats TeamStats `json:"team_stats"`
	Players   []Player  `json:"players"`
}

type TeamStats struct {
	OvertimeScore   string `json:"Overtime score"`
	FirstHalfScore  string `json:"First Half Score"`
	SecondHalfScore string `json:"Second Half Score"`
	Team            string `json:"Team"`
	TeamWin         string `json:"Team Win"`
	TeamHeadshots   string `json:"Team Headshots"`
	FinalScore      string `json:"Final Score"`
}

type Player struct {
	PlayerID    string      `json:"player_id"`
	Nickname    string      `json:"nickname"`
	PlayerStats PlayerStats `json:"player_stats"`
}

type PlayerStats struct {
	QuadroKills                      string `json:"Quadro Kills"`
	FirstKills                       string `json:"First Kills"`
	UtilityUsagePerRound             string `json:"Utility Usage per Round"`
	UtilitySuccesses                 string `json:"Utility Successes"`
	ClutchKills                      string `json:"Clutch Kills"`
	KnifeKills                       string `json:"Knife Kills"`
	ADR                              string `json:"ADR"`
	PentaKills                       string `json:"Penta Kills"`
	KDRatio                          string `json:"K/D Ratio"`
	TripleKills                      string `json:"Triple Kills"`
	EnemiesFlashedPerRoundInAMatch   string `json:"Enemies Flashed per Round in a Match"`
	UtilityEnemies                   string `json:"Utility Enemies"`
	MatchEntrySuccessRate            string `json:"Match Entry Success Rate"`
	Assists                          string `json:"Assists"`
	Result                           string `json:"Result"`
	Damage                           string `json:"Damage"`
	HeadshotsPercent                 string `json:"Headshots %"`
	Deaths                           string `json:"Deaths"`
	PistolKills                      string `json:"Pistol Kills"`
	FlashSuccesses                   string `json:"Flash Successes"`
	OneVOneCount                     string `json:"1v1Count"`
	OneVTwoCount                     string `json:"1v2Count"`
	ZeusKills                        string `json:"Zeus Kills"`
	FlashSuccessRatePerMatch         string `json:"Flash Success Rate per Match"`
	FlashesPerRoundInAMatch          string `json:"Flashes per Round in a Match"`
	UtilityDamagePerRoundInAMatch    string `json:"Utility Damage per Round in a Match"`
	MatchOneVTwoWinRate              string `json:"Match 1v2 Win Rate"`
	DoubleKills                      string `json:"Double Kills"`
	FlashCount                       string `json:"Flash Count"`
	Kills                            string `json:"Kills"`
	MVPs                             string `json:"MVPs"`
	KRRatio                          string `json:"K/R Ratio"`
	UtilityDamage                    string `json:"Utility Damage"`
	EnemiesFlashed                   string `json:"Enemies Flashed"`
	SniperKills                      string `json:"Sniper Kills"`
	EntryWins                        string `json:"Entry Wins"`
	MatchOneVOneWinRate              string `json:"Match 1v1 Win Rate"`
	UtilitySuccessRatePerMatch       string `json:"Utility Success Rate per Match"`
	SniperKillRatePerMatch           string `json:"Sniper Kill Rate per Match"`
	UtilityCount                     string `json:"Utility Count"`
	MatchEntryRate                   string `json:"Match Entry Rate"`
	UtilityDamageSuccessRatePerMatch string `json:"Utility Damage Success Rate per Match"`
	EntryCount                       string `json:"Entry Count"`
	SniperKillRatePerRound           string `json:"Sniper Kill Rate per Round"`
	OneVOneWins                      string `json:"1v1Wins"`
	Headshots                        string `json:"Headshots"`
	OneVTwoWins                      string `json:"1v2Wins"`
}

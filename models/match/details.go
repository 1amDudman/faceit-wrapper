package match

type MatchDetails struct {
	MatchID         string           `json:"match_id"`
	Version         int              `json:"version"`
	Game            string           `json:"game"`
	Region          string           `json:"region"`
	CompetitionID   string           `json:"competition_id"`
	CompetitionType string           `json:"competition_type"`
	CompetitionName string           `json:"competition_name"`
	OrganizerID     string           `json:"organizer_id"`
	Teams           map[string]Team  `json:"teams"`
	Voting          Voting           `json:"voting"`
	CalculateElo    bool             `json:"calculate_elo"`
	ConfiguredAt    int64            `json:"configured_at"`
	StartedAt       int64            `json:"started_at"`
	FinishedAt      int64            `json:"finished_at"`
	DemoURL         []string         `json:"demo_url"`
	ChatRoomID      string           `json:"chat_room_id"`
	BestOf          int              `json:"best_of"`
	Results         Results          `json:"results"`
	DetailedResults []DetailedResult `json:"detailed_results"`
	Status          string           `json:"status"`
	FaceitURL       string           `json:"faceit_url"`
}

type Team struct {
	FactionID   string        `json:"faction_id"`
	Leader      string        `json:"leader"`
	Avatar      string        `json:"avatar"`
	Roster      []MatchPlayer `json:"roster"`
	Stats       TeamStats     `json:"stats"`
	Substituted bool          `json:"substituted"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
}

type MatchPlayer struct {
	PlayerID          string `json:"player_id"`
	Nickname          string `json:"nickname"`
	Avatar            string `json:"avatar"`
	Membership        string `json:"membership"`
	GamePlayerID      string `json:"game_player_id"`
	GamePlayerName    string `json:"game_player_name"`
	GameSkillLevel    int    `json:"game_skill_level"`
	AnticheatRequired bool   `json:"anticheat_required"`
}

type TeamStats struct {
	WinProbability float64    `json:"winProbability"`
	SkillLevel     SkillRange `json:"skillLevel"`
	Rating         int        `json:"rating"`
}

type SkillRange struct {
	Average int `json:"average"`
	Range   struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"range"`
}

type Voting struct {
	VotedEntityTypes []string `json:"voted_entity_types"`
	Location         Location `json:"location"`
	Map              Map      `json:"map"`
}

type Location struct {
	Pick     []string         `json:"pick"`
	Entities []LocationEntity `json:"entities"`
}

type LocationEntity struct {
	ClassName      string `json:"class_name"`
	GameLocationID string `json:"game_location_id"`
	GUID           string `json:"guid"`
	ImageLG        string `json:"image_lg"`
	ImageSM        string `json:"image_sm"`
	Name           string `json:"name"`
}

type Map struct {
	Entities []MapEntity `json:"entities"`
	Pick     []string    `json:"pick"`
}

type MapEntity struct {
	ClassName string `json:"class_name"`
	GameMapID string `json:"game_map_id"`
	GUID      string `json:"guid"`
	ImageLG   string `json:"image_lg"`
	ImageSM   string `json:"image_sm"`
	Name      string `json:"name"`
}

type Results struct {
	Winner string `json:"winner"`
	Score  Score  `json:"score"`
}

type Score struct {
	Faction1 int `json:"faction1"`
	Faction2 int `json:"faction2"`
}

type DetailedResult struct {
	AscScore bool     `json:"asc_score"`
	Winner   string   `json:"winner"`
	Factions Factions `json:"factions"`
}

type Factions struct {
	Faction1 ScoreDetails `json:"faction1"`
	Faction2 ScoreDetails `json:"faction2"`
}

type ScoreDetails struct {
	Score int `json:"score"`
}

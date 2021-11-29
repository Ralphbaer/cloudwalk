package usecase

// QuakeLog represents the characteristics of the device used
type QuakeLog struct {
	GameNum int    `json:"game_num"`
	Games   []Game `json:"games"`
}

// Game represents the characteristics of the device used
type Game struct {
	GameNum    int               `json:"gameNum"`
	TotalKills int               `json:"totalKills"`
	Players    map[string]string `json:"players"`
	Kills      map[string]int    `json:"kills"`
	DeathCause map[string]int    `json:"deathCause"`
}

// GroupedInformationReport represents the characteristics of the device used
type GroupedInformationReport struct {
	GameNum    int            `json:"gameNum"`
	TotalKills int            `json:"totalKills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

func (ql *QuakeLog) newGame() {
	ql.GameNum += 1
	ql.Games = append(ql.Games, Game{
		GameNum:    ql.GameNum,
		TotalKills: 0,
		Players:    nil,
		Kills:      nil,
	})
}

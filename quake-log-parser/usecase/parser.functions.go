package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func (g *Game) pushPlayer(playerId, playerNickname string) {
	if g.Players == nil {
		g.Players = make(map[string]string)
	}

	if g.Kills == nil {
		g.Kills = make(map[string]int)
	}

	if g.DeathCause == nil {
		g.DeathCause = make(map[string]int)
	}

	if g.Players[playerId] == playerNickname {
		return
	}

	if g.Players[playerId] != playerNickname {
		g.Players[playerId] = playerNickname
		return
	}

	g.Players[playerId] = playerNickname
	g.Kills[playerNickname] = 0
}

func (g *Game) pushKill(killer string) {
	if killer != "" {
		g.Kills[killer] += 1
	}
}

func (g *Game) pushDeathCause(deathCause string) {
	g.DeathCause[deathCause] += 1
}

func (g *Game) pullKill(killed string) {
	g.Kills[killed] -= 1
}

// PrintGroupedInformationReport prints the grouped information about kills and players of each match
func (q *QuakeLog) PrintGroupedInformationReport() {
	log.Println("-----------Grouped Information Report-----------")
	gir := make([]*GroupedInformationReport, 0)

	for _, v := range q.Games {
		gir = append(gir, &GroupedInformationReport{
			GameNum: v.GameNum,
			TotalKills: v.TotalKills,
			Players: func() []string{
				players := make([]string, 0)
				for _, p := range v.Players {
					players = append(players, p)
				}
				return players
			}(),
			Kills:   v.Kills,
		})
	}

	j, _ := json.MarshalIndent(gir, "", "	")

	log.Println(string(j))
}

// PrintGlobalRanking prints the global ranking of players
// e.g.: 
// -----------Global Ranking:-----------
// Isgalamido: 144
// Zeh: 123
// Assasinu Credi: 111
func (q *QuakeLog) PrintGlobalRanking() {
	ranking := q.getGlobalRanking()

	fmt.Println("\n-----------Global Ranking:-----------")
	fmt.Println("")

	for _, v := range ranking {
		fmt.Printf("%s: %d\n", v.Key, v.Value)
	}
}

func (q *QuakeLog) getGlobalRanking() Ranking {
	ranking := make(map[string]int)

	for _, v := range q.Games {
		for k := range v.Kills {
			killer := k
			points := v.Kills[killer]

			if _, ok := ranking[killer]; !ok {
				ranking[killer] = points

				continue
			}
			ranking[killer] += points
		}
	}

	p := make(Ranking, len(ranking))
	i := 0
	for k, v := range ranking {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	return p
}

// Pair is a struct that holds a key and a value
type Pair struct {
	Key   string
	Value int
}

// Ranking is a type that represents Pair as a slice
type Ranking []Pair

// Len returns the length of the Ranking
func (p Ranking) Len() int           { return len(p) }
// Swap swaps the values of two pairs
func (p Ranking) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// Less returns true if the value of the first pair is less than the value of the second pair
func (p Ranking) Less(i, j int) bool { return p[i].Value > p[j].Value }

type ReportOfDeath struct {
	GameNum     int            `json:"gameNum"`
	KillByMeans map[string]int `json:"killByMeans"`
}

// PrintDeathCausesReport prints the report of death causes.
// e.g.: [{
//  "game_num": 1,	
//  "killByMeans": {
//    "MOD_SHOTGUN": 10,
//    "MOD_RAILGUN": 2,
//    "MOD_GAUNTLET": 1,
//  },
//  ...
//}]
func (q *QuakeLog) PrintDeathCausesReport() {
	fmt.Println("\n-----------Death Causes Report:-----------")

	rd := make([]ReportOfDeath, 0)

	for _, v := range q.Games {
		rd = append(rd, ReportOfDeath{v.GameNum, v.DeathCause})
	}

	j, _ := json.MarshalIndent(rd, "", "	")

	fmt.Println(string(j))
}

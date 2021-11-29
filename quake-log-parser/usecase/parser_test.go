package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/Ralphbaer/cloudwalk/quake-log-parser/repository"
	"github.com/golang/mock/gomock"
)

func TestCartUseCase_GroupGamesInformation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := repository.NewQuakeLogFileRepository("./qgames-test.log")
	uc := QuakeLogParserUseCase{
		Repository: repo,
	}

	got, err := uc.groupGamesInformation(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}

	m := struct {
		want *QuakeLog
		got  *QuakeLog
	}{
		want: &QuakeLog{
			GameNum: 1,
			Games: []Game{
				{
					GameNum:    1,
					TotalKills: 130,
					Players: map[string]string{
						"2": "Oootsimo",
						"3": "Isgalamido",
						"4": "Zeh",
						"5": "Dono da Bola",
						"6": "Mal",
						"7": "Assasinu Credi",
						"8": "Chessus",
					},
					Kills: map[string]int{
						"Assasinu Credi": 16,
						"Dono da Bola":   8,
						"Isgalamido":     12,
						"Mal":            -3,
						"Oootsimo":       20,
						"Zeh":            7,
					},
					DeathCause: map[string]int{
						"MOD_FALLING":       7,
						"MOD_MACHINEGUN":    9,
						"MOD_RAILGUN":       9,
						"MOD_ROCKET":        29,
						"MOD_ROCKET_SPLASH": 49,
						"MOD_SHOTGUN":       7,
						"MOD_TRIGGER_HURT":  20,
					},
				},
			},
		},
		got: &QuakeLog{
			GameNum: got.GameNum,
			Games:   got.Games,
		},
	}

	if !reflect.DeepEqual(m.want, m.got) {
		t.Error("Got:", m.got, "Want:", m.want)
		return
	}
}

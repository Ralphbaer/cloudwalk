package usecase

import (
	"reflect"
	"testing"
)

func TestGame_pushPlayer(t *testing.T) {
	type fields struct {
		GameNum    int
		TotalKills int
		Players    map[string]string
		Kills      map[string]int
		DeathCause map[string]int
	}
	type args struct {
		playerId       string
		playerNickname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "PushPlayer1",
			fields: fields{
				GameNum:    1,
				TotalKills: 0,
				Players:    map[string]string{},
				Kills:      map[string]int{},
				DeathCause: map[string]int{},
			},
			args: args{
				playerId:       "1",
				playerNickname: "player1",
			},
		},
		{
			name: "PushPlayer2",
			fields: fields{
				GameNum:    2,
				TotalKills: 0,
				Players:    map[string]string{},
				Kills:      map[string]int{},
				DeathCause: map[string]int{},
			},
			args: args{
				playerId:       "2",
				playerNickname: "player2",
			},
		},
		{
			name: "PushPlayer3",
			fields: fields{
				GameNum:    3,
				TotalKills: 0,
				Players: map[string]string{
					"1": "player1",
					"2": "player2",
				},
				Kills:      map[string]int{},
				DeathCause: map[string]int{},
			},
			args: args{
				playerId:       "3",
				playerNickname: "player3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				GameNum:    tt.fields.GameNum,
				TotalKills: tt.fields.TotalKills,
				Players:    tt.fields.Players,
				Kills:      tt.fields.Kills,
				DeathCause: tt.fields.DeathCause,
			}
			g.pushPlayer(tt.args.playerId, tt.args.playerNickname)
			if g.Players[tt.args.playerId] != tt.args.playerNickname {
				t.Errorf("Game.pushPlayer() = %v, want %v", g.Players[tt.args.playerId], tt.args.playerNickname)
			}
		})
	}
}

func TestGame_pushKill(t *testing.T) {
	type fields struct {
		GameNum    int
		TotalKills int
		Players    map[string]string
		Kills      map[string]int
		DeathCause map[string]int
	}
	type args struct {
		killer string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int
	}{
		{
			name: "PushKill1",
			fields: fields{
				GameNum:    1,
				TotalKills: 0,
				Players: map[string]string{
					"1": "player1",
					"2": "player2",
				},
				Kills:      map[string]int{},
				DeathCause: map[string]int{},
			},
			args: args{
				killer: "1",
			},
			want: map[string]int{
				"1": 1,
			},
		},
		{
			name: "PushKill2",
			fields: fields{
				GameNum:    2,
				TotalKills: 2,
				Players: map[string]string{
					"1": "player1",
					"2": "player2",
				},
				Kills: map[string]int{
					"2": 1,
				},
				DeathCause: map[string]int{},
			},
			args: args{
				killer: "2",
			},
			want: map[string]int{
				"2": 2,
			},
		},
		{
			name: "PushKill3",
			fields: fields{
				GameNum:    3,
				TotalKills: 3,
				Players: map[string]string{
					"1": "player1",
					"2": "player2",
					"3": "player3",
				},
				Kills: map[string]int{
					"1": 1,
					"2": 1,
					"3": 1,
				},
				DeathCause: map[string]int{},
			},
			args: args{
				killer: "3",
			},
			want: map[string]int{
				"1": 1,
				"2": 1,
				"3": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				GameNum:    tt.fields.GameNum,
				TotalKills: tt.fields.TotalKills,
				Players:    tt.fields.Players,
				Kills:      tt.fields.Kills,
				DeathCause: tt.fields.DeathCause,
			}
			g.pushKill(tt.args.killer)
			if !reflect.DeepEqual(g.Kills, tt.want) {
				t.Error("Got:", g.Kills, "Want:", tt.want)
				return
			}
		})
	}
}

func TestGame_pushDeathCause(t *testing.T) {
	type fields struct {
		GameNum    int
		TotalKills int
		Players    map[string]string
		Kills      map[string]int
		DeathCause map[string]int
	}
	type args struct {
		deathCause string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int
	}{
		{
			name: "PushDeathCause1",
			fields: fields{
				GameNum:    1,
				TotalKills: 1,
				Players: map[string]string{
					"1": "player1",
				},
				Kills:      map[string]int{},
				DeathCause: map[string]int{},
			},
			args: args{
				deathCause: "MOD_FALLING",
			},
			want: map[string]int{
				"MOD_FALLING": 1,
			},
		},
		{
			name: "PushDeathCause2",
			fields: fields{
				GameNum:    1,
				TotalKills: 3,
				Players: map[string]string{
					"1": "player1",
					"2": "player2",
				},
				Kills: map[string]int{},
				DeathCause: map[string]int{
					"MOD_FALLING": 1,
					"MOD_ROCKET":  1,
				},
			},
			args: args{
				deathCause: "MOD_FALLING",
			},
			want: map[string]int{
				"MOD_FALLING": 2,
				"MOD_ROCKET":  1,
			},
		},
		{
			name: "PushDeathCause3",
			fields: fields{
				GameNum:    1,
				TotalKills: 6,
				Players: map[string]string{
					"1": "player1",
					"2": "player2",
					"3": "player3",
				},
				Kills: map[string]int{
					"1": 1,
				},
				DeathCause: map[string]int{
					"MOD_FALLING": 3,
					"MOD_ROCKET":  1,
				},
			},
			args: args{
				deathCause: "MOD_FALLING",
			},
			want: map[string]int{
				"MOD_FALLING": 4,
				"MOD_ROCKET":  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				GameNum:    tt.fields.GameNum,
				TotalKills: tt.fields.TotalKills,
				Players:    tt.fields.Players,
				Kills:      tt.fields.Kills,
				DeathCause: tt.fields.DeathCause,
			}
			g.pushDeathCause(tt.args.deathCause)
			if !reflect.DeepEqual(g.DeathCause, tt.want) {
				t.Error("Got:", g.DeathCause, "Want:", tt.want)
				return
			}
		})
	}
}

func TestGame_pullKill(t *testing.T) {
	type fields struct {
		GameNum    int
		TotalKills int
		Players    map[string]string
		Kills      map[string]int
		DeathCause map[string]int
	}
	type args struct {
		killed string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int
	}{
		{
			name: "PullKill1",
			fields: fields{
				GameNum:    1,
				TotalKills: 1,
				Players: map[string]string{
					"1": "player1",
				},
				Kills: map[string]int{
					"1": 1,
				},
				DeathCause: map[string]int{
					"MOD_FALLING": 1,
				},
			},
			args: args{
				killed: "1",
			},
			want: map[string]int{
				"1": 0,
			},
		},
		{
			name: "PullKill2",
			fields: fields{
				GameNum:    2,
				TotalKills: 1,
				Players: map[string]string{
					"1": "player1",
				},
				Kills: map[string]int{},
				DeathCause: map[string]int{
					"MOD_FALLING": 1,
				},
			},
			args: args{
				killed: "1",
			},
			want: map[string]int{
				"1": -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				GameNum:    tt.fields.GameNum,
				TotalKills: tt.fields.TotalKills,
				Players:    tt.fields.Players,
				Kills:      tt.fields.Kills,
				DeathCause: tt.fields.DeathCause,
			}
			g.pullKill(tt.args.killed)
			if !reflect.DeepEqual(g.Kills, tt.want) {
				t.Error("Got:", g.Kills, "Want:", tt.want)
				return
			}
		})
	}
}

func TestQuakeLog_getGlobalRanking(t *testing.T) {
	type fields struct {
		GameNum int
		Games   []Game
	}
	tests := []struct {
		name   string
		fields fields
		want   Ranking
	}{
		{
			name: "GetGlobalRanking1",
			fields: fields{
				GameNum: 1,
				Games: []Game{
					{
						GameNum:    1,
						TotalKills: 1,
						Players: map[string]string{
							"1": "Isgalamido",
							"2": "Zeh",
						},
						Kills:      map[string]int{
							"1": 10,
							"2": 5,
						},
						DeathCause: map[string]int{},
					},
				},
			},
			want: Ranking{
				Pair{
					Key: "1",
					Value: 10,
				},
				Pair{
					Key: "2",
					Value: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QuakeLog{
				GameNum: tt.fields.GameNum,
				Games:   tt.fields.Games,
			}
			if got := q.getGlobalRanking(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuakeLog.getGlobalRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}

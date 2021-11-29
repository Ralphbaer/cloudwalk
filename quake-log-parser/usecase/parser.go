package usecase

import (
	"bufio"
	"bytes"
	"context"
	"strconv"
	"strings"

	r "github.com/Ralphbaer/cloudwalk/quake-log-parser/repository"
)

// QuakeLogParserUseCase represents a collection of use cases for QuakeLogParser operations
type QuakeLogParserUseCase struct {
	Repository r.QuakeLogRepository
}

// Run just orchestrate the functions
func (uc *QuakeLogParserUseCase) Run(ctx context.Context) error {
	ql, err := uc.groupGamesInformation(ctx)
	if err != nil {
		return err
	}
	//ql.PrintGroupedInformationReport()
	ql.PrintGlobalRanking()
	//ql.PrintDeathCausesReport()

	return nil
}

const WORLD = 1022

func (uc *QuakeLogParserUseCase) groupGamesInformation(ctx context.Context) (*QuakeLog, error) {
	f, err := uc.Repository.GetFile(ctx)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ql := &QuakeLog{}
	for scanner.Scan() {
		line := scanner.Bytes()
		lineStr := scanner.Text()

		if bytes.Contains(line, []byte("InitGame")) {
			ql.newGame()
		}

		if strings.HasPrefix(lineStr, " ") {
			lineStr = strings.TrimSpace(lineStr)
		}

		if bytes.Contains(line, []byte("ClientUserinfoChanged")) {
			ql.Games[ql.GameNum-1].pushPlayer(strings.Split(lineStr, " ")[2], strings.Split(lineStr, "\\")[1])
		}

		if bytes.Contains(line, []byte("Kill")) {
			ql.Games[ql.GameNum-1].TotalKills += 1
			killer := strings.Split(lineStr, " ")[2]
			killed := strings.Split(lineStr, " ")[3]
			deathCause := strings.Split(string(lineStr), "by ")[1]

			ql.Games[ql.GameNum-1].pushDeathCause(deathCause)
			i, err := strconv.Atoi(killer)
			if err != nil {
				return nil, err
			}

			if i == WORLD || killer == killed {
				ql.Games[ql.GameNum-1].pullKill(ql.Games[ql.GameNum-1].Players[killed])
				continue
			}

			ql.Games[ql.GameNum-1].pushKill(ql.Games[ql.GameNum-1].Players[killer])
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ql, nil
}

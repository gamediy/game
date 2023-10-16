package logic

import (
	"context"
	"errors"
	"fmt"
	"game/db/dao"
	"game/db/model/entity"
	"time"
)

type GameIssueRespone struct {
	Issue          int64     `json:"issue"`
	OpenTime       time.Time `json:"openTime"`
	CloseTime      time.Time `json:"closeTime"`
	StartTime      time.Time `json:"startTime"`
	OpenTimeStr    string    `json:"openTimeStr"`
	CloseTimeStr   string    `json:"closeTimeStr"`
	StartTimeStr   string    `json:"startTimeStr"`
	CloseCountdown int64     `json:"closeCountdown"`
	OpenCountdown  int64     `json:"openCountdown"`
	Status         int32     `json:"status"` //1可以投注，2：结束投注。封盘
	StatusStr      string    `json:"statusStr"`
	TimeNow        time.Time `json:"timeNow"`
	Date           string    `json:"date"`
	Name           string    `json:"name"`
	GameCode       int       `json:"gameCode"`
}

func GetIssue(ctx context.Context, gameCode int32) (GameIssueRespone, error) {

	game := &entity.Game{}
	dao.Game.Ctx(ctx).Where("code", gameCode)

	var currentIssue GameIssueRespone
	if game.Status != 1 {
		return currentIssue, errors.New("Status Off")
	}
	currentIssue.Name = game.Name
	currentIssue.GameCode = game.Code
	//2006-01-02 15:04:05
	issueList := make([]GameIssueRespone, 0)
	format := time.Now().Format("2006-01-02")
	timeStart, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", format, game.StartTime), time.Local)
	timeEnd, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", format, game.EndTime), time.Local)
	var issueIndex int64 = 1

	for timeStart.Before(timeEnd) {
		duration, _ := time.ParseDuration(fmt.Sprintf("%ds", game.IntervalSeconds))
		closeDur, _ := time.ParseDuration(fmt.Sprintf("-%ds", game.CloseSeconds))
		m := GameIssueRespone{}
		m.StartTime = timeStart
		timeStart = timeStart.Add(duration)
		m.Issue = issueIndex
		issueIndex++
		m.CloseTime = timeStart.Add(closeDur)
		m.Date = timeStart.Format("2006-01-02")
		m.OpenTime = timeStart
		issueList = append(issueList, m)
	}
	timeNow := time.Now()
	for _, v := range issueList {
		if timeNow.After(v.StartTime) && timeNow.Before(v.OpenTime) {
			v.CloseCountdown = v.CloseTime.Unix() - timeNow.Unix()
			v.OpenCountdown = v.OpenTime.Unix() - timeNow.Unix()

			v.TimeNow = timeNow
			v.Status = 1
			v.StatusStr = "Betting"
			if v.CloseCountdown <= 0 {
				v.Status = 2
				v.StatusStr = "Drawing"
			}
			v.Name = game.Name
			v.GameCode = game.Code
			currentIssue = v
			continue
		}
	}
	currentIssue.CloseTimeStr = currentIssue.CloseTime.Format("2006-01-02 15:04:05")
	currentIssue.OpenTimeStr = currentIssue.OpenTime.Format("2006-01-02 15:04:05")
	currentIssue.StartTimeStr = currentIssue.StartTime.Format("2006-01-02 15:04:05")

	return currentIssue, nil
}

package service

import (
	"errors"
	"fmt"
	"time"

	_ "time/tzdata"

	"github.com/fatahnuram/learn-go-kasir-api/internal/dto"
	"github.com/fatahnuram/learn-go-kasir-api/internal/repository"
)

type ReportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return ReportService{
		repo: repo,
	}
}

func (s ReportService) GetReportToday() (*dto.ReportResp, error) {
	now := time.Now()
	wib, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}
	timewib := now.In(wib)

	startwib := time.Date(
		timewib.Year(),
		timewib.Month(),
		timewib.Day(),
		0,
		0,
		0,
		0,
		wib,
	)
	endwib := time.Date(
		timewib.Year(),
		timewib.Month(),
		timewib.Day(),
		23,
		59,
		59,
		999999,
		wib,
	)

	return s.repo.GetReportByTimeRange(startwib, endwib)
}

func (s ReportService) GetReportByTimeRange(start, end string) (*dto.ReportResp, error) {
	timelayout := "2006-01-02"

	starttime, err := time.Parse(timelayout, start)
	if err != nil {
		return nil, fmt.Errorf("parse start date error: %s", err.Error())
	}
	endtime, err := time.Parse(timelayout, end)
	if err != nil {
		return nil, fmt.Errorf("parse end date error: %s", err.Error())
	}

	wib, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	startwib := starttime.In(wib)
	endwib := endtime.In(wib)

	starttz := time.Date(
		startwib.Year(),
		startwib.Month(),
		startwib.Day(),
		0,
		0,
		0,
		0,
		wib,
	)
	endtz := time.Date(
		endwib.Year(),
		endwib.Month(),
		endwib.Day(),
		23,
		59,
		59,
		999999,
		wib,
	)

	if starttz.After(endtz) {
		return nil, errors.New("end date is earlier than start date")
	}

	return s.repo.GetReportByTimeRange(starttz, endtz)
}

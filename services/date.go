package services

import (
  "time"
)

type DateService struct {
  currentDate time.Time
}

func NewDateService() *DateService {
  return &DateService{time.Now()}
}

func (d *DateService) DaysLeft() int {
  until := time.Until(time.Date(2025, time.Month(1), 1, 0, 0, 0, 0, time.UTC))
  
  return int(until.Hours() / 24)
}

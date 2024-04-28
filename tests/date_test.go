package tests

import (
  "time"
  "testing"

  "gitlab.com/croc3/umbrellacorptask/src/services"
)

var janFirst = time.Date(2025, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

func TestDateService(t *testing.T) {
  expected := int(time.Until(janFirst).Hours() / 24)

  d := services.NewDateService()
  actual := d.DaysLeft()

  if actual != expected {
    t.Errorf("expected %d but got %d", expected, actual)
  }
}

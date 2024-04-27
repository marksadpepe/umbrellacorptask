package main

import (
  "net/http"

  "github.com/labstack/echo/v4"
)

type ResponseMessage struct {
  DaysLeft int `json:"daysLeft"`
}

type Service interface {
  DaysLeft() int
}

type Endpoint struct {
  service Service
}

func NewEndpoint(s Service) Endpoint {
  return Endpoint{s}
}

func (e *Endpoint) HandleRequest(c echo.Context) error {
  daysLeft := e.service.DaysLeft()
  res := ResponseMessage{daysLeft}

  err := c.JSON(http.StatusOK, res)
  if err != nil {
    return err
  }

  return nil
}

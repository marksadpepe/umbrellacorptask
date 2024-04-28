package controllers

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

type Controller struct {
  service Service
}

func NewController(s Service) Controller {
  return Controller{s}
}

func (c *Controller) HandleRequest(ctx echo.Context) error {
  daysLeft := c.service.DaysLeft()
  res := ResponseMessage{daysLeft}

  err := ctx.JSON(http.StatusOK, res)
  if err != nil {
    return err
  }

  return nil
}

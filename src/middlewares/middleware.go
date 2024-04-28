package middlewares

import (
  "log"
  "strings"

  "github.com/labstack/echo/v4"
)

const adminRole = "admin"

func IsUserAdmin(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    userRole := strings.TrimSpace(strings.ToLower((c.Request().Header.Get("User-Role"))))

    if strings.Contains(userRole, adminRole) {
      log.Printf("red button user detected\n")
    }

    err := next(c)
    if err != nil {
      return err
    }

    return nil
  }
}

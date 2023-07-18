package mw

import (
	"log"
	"strings"

	"github.com/labstack/echo"
)

const roleAdmin = "admin"

func Rolecheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		val := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(val, roleAdmin) {
			log.Println("Red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}

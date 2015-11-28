package middleware

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

func TxMiddleware(session *dbr.Session) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) (err error) {
			tx, err := session.Begin()

			if err != nil {
				logrus.Error(err)
			} else {
				c.Set("Tx", tx)

				if err := h(c); err != nil {
					tx.Rollback()
					logrus.Error("Rollback")
					c.Error(err)
				} else {
					logrus.Debug("Commit")
					tx.Commit()
				}
			}
			return
		}
	}
}

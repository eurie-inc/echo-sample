package middleware

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *dbr.Session) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			tx, _ := db.Begin()

			c.Set(TxKey, tx)

			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Debug("Transction Rollback: ", err)
				return err
			}
			logrus.Debug("Transaction Commit")
			tx.Commit()

			return nil
		})
	}
}

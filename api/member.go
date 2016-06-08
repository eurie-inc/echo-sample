package api

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/eurie-inc/echo-sample/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func PostMember() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := new(model.Member)
		c.Bind(&m)

		tx := c.Get("Tx").(*dbr.Tx)

		member := model.NewMember(m.Number, m.Name, m.Position)

		if err := member.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		return c.JSON(fasthttp.StatusCreated, member)
	}
}

func GetMember() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		number, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		member := new(model.Member)
		if err := member.Load(tx, number); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "Member does not exists.")
		}
		return c.JSON(fasthttp.StatusOK, member)
	}
}

func GetMembers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)

		position := c.QueryParam("position")
		members := new(model.Members)
		if err = members.Load(tx, position); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "Member does not exists.")
		}

		return c.JSON(fasthttp.StatusOK, members)
	}
}

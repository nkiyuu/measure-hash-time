package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/sha3"
)

type responseModel struct {
	Times    int     `json:"times"`
	Digest   string  `json:"digest"`
	HashTime float64 `json:"hash_time"`
}

func main() {
	e := echo.New()

	e.GET("/measure-hash", func(c echo.Context) error {
		testPhoneNumber := "0000000000"
		timesQuery := c.QueryParam("times")
		times, err := strconv.Atoi(timesQuery)
		if err != nil {
			return err
		}

		var res responseModel
		st := time.Now()
		tp := testPhoneNumber
		for i := 0; i < times; i++ {
			tp = fmt.Sprintf("%x", sha3.Sum256([]byte(tp)))
		}
		en := time.Now()

		res.Times = times
		res.HashTime = en.Sub(st).Seconds()
		res.Digest = fmt.Sprintf("%v", tp)

		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

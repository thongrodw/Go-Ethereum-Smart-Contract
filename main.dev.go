package main

import (

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoint
	e.GET("/greet", func(c echo.Context) error {
		// reply, err := conn.Greet(&bind.CallOpts{})
		reply = "Success";
		// if err != nil {
		// 	return err
		// }
		return c.JSON(http.StatusOK,{{"data":"success"}})
	})
	
	e.GET("/greet/:_greeting", func(c echo.Context) error {
		// _greeting := c.Param("_greeting")
		// reply, err := conn.SetGreeting(auth, _greeting)
		reply = "Success";
		// if err != nil {
		// 	return err
		// }
		return c.JSON(http.StatusOK,{{"data":"success"}})
	})

	// Start server
	e.Logger.Fatal(e.Start(":1324"))
}
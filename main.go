package main

import (
	"context"
	"greetergo/api" // your generated smart contract bindings
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA("3d1fdf4e80b90480758df55827b4af449f2ade80899abfbd3943682453bdeaf6")
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	conn, err := api.NewApi(common.HexToAddress("0x8c8f88821280353aa02f395bd7c6ea68b68c772b"), client)
	if err != nil {
		panic(err)
	}
	
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoint
	e.GET("/greet", func(c echo.Context) error {
		reply, err := conn.Greet(&bind.CallOpts{})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})
	
	e.GET("/greet/:_greeting", func(c echo.Context) error {
		_greeting := c.Param("_greeting")
		reply, err := conn.SetGreeting(auth, _greeting)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1324"))
}
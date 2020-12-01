package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

type M map[string]interface{}

/*
	Penggunaan echo.WrapHandler untuk routing handler bertipa func(http.ResponseWriter, *http.Requests)
	atau http.HanderFunc
*/
var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("from action index"))
	if err != nil {
		log.Fatal(err)
	}
}

var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("from action home"))
		if err != nil {
			log.Fatal(err)
		}
	},
)

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("from action about"))
			if err != nil {
				log.Fatal(err)
			}
		},
	),
)

func main() {
	r := echo.New()

	// Method .String()
	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	// Method .JSON()
	r.GET("/json", func(ctx echo.Context) error {
		data := M{"Message": "Hello", "Counter": 2}
		return ctx.JSON(http.StatusOK, data)
	})

	// Parsing Query String
	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	// Parsing URL Path Param
	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	// Parsing URL Path Param and so forth
	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("Hello %s, I have a messages for you: %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	// Parsing Form Data
	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf(
			"Hello %s, I have a message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/home", echo.WrapHandler(ActionHome))
	r.GET("/about", ActionAbout)

	// Routing Static Assets
	r.Static("/static", "assets")

	if err := r.Start(":9000"); err != nil {
		return
	}
}

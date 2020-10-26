package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rakyll/statik/fs"

	_ "github.com/erikh/betl/statik"
)

var staticFilesystem http.FileSystem

func defaultEnv(env, dflt string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}

	return dflt
}

func main() {
	var err error
	staticFilesystem, err = fs.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/*", serveStatic)

	e.Logger.Fatal(e.Start(defaultEnv("BETL_LISTEN", ":8000")))
}

// Handler
func serveStatic(c echo.Context) error {
	path := strings.TrimLeft(c.Request().RequestURI, "/")
	if path == "" {
		path = "/index.html"
	}

	f, err := staticFilesystem.Open(path)
	if err != nil {
		return err
	}

	if _, err := io.Copy(c.Response().Writer, f); err != nil {
		return err
	}

	return nil
}

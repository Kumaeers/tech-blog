package main

import (
	"log"
	"net/http"
	"os"

	"app/repository"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB
var e = createMux()

func main () {
	db = connectDB()
	repository.SetDB(db)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/public", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello public!")
	})
	e.GET("/private", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello private!")
	})
	// Start server
	e.Logger.Fatal(e.Start(":8082"))
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
			 e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
			 e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

func createMux() *echo.Echo {
	// アプリケーションインスタンスを生成
	e := echo.New()

	// アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// アプリケーションインスタンスを返却
return e
}

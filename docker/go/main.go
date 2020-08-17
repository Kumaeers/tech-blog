package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var db *sqlx.DB
var e = createMux()

func main() {
	// db = connectDB()
	// repository.SetDB(db)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/public", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello public!")
	})
	e.GET("/private", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello private!")
	}, firebaseMiddleware())
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

// JWTを検証する
func firebaseMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Firebase SDK のセットアップ
			opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
			app, err := firebase.NewApp(context.Background(), nil, opt)
			if err != nil {
				return err
			}

			client, err := app.Auth(context.Background())
			if err != nil {
				return err
			}

			// クライアントから送られてきた JWT 取得
			auth := c.Request().Header.Get("Authorization")
			idToken := strings.Replace(auth, "Bearer ", "", 1)

			// JWT の検証
			token, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return err
			}

			c.Set("token", token)
			return next(c)
		}
	}
}

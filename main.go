package main

import (
	"log"
	"os"

	"go_blog/handler"
	"go_blog/repository"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	// ルーティング設定
	// TOP ページに記事の一覧を表示します。
	e.GET("/", handler.ArticleIndex)

	// 記事に関するページは "/articles" で開始するようにします。
	// 記事一覧画面には "/" と "/articles" の両方でアクセスできるようにします。
	// パスパラメータの ":id" も ":articleID" と明確にしています。
	e.GET("/articles", handler.ArticleIndex)                // 一覧画面
	e.GET("/articles/new", handler.ArticleNew)              // 新規作成画面
	e.GET("/articles/:articleID", handler.ArticleShow)      // 詳細画面
	e.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面

	// HTML ではなく JSON を返却する処理は "/api" で開始するようにします。
	// 記事に関する処理なので "/articles" を続けます。
	e.GET("/api/articles", handler.ArticleList)                    // 一覧
	e.POST("/articles/create", handler.ArticleCreate)              // 作成
	e.DELETE("/articles/delete/:articleID", handler.ArticleDelete) // 削除
	e.PATCH("/articles/update/:articleID", handler.ArticleUpdate)  // 更新

	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
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
	e.Use(middleware.CSRF())

	// `/src` 配下のファイルに `/css,/js` のパスでアクセスできるようにする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	e.Validator = &CustomValidator{validator: validator.New()}

	// アプリケーションインスタンスを返却
	return e
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kbavi/urlshortner/url"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	service url.Service
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	Init()

	app.Get("/:id", RedirectHandler)

	app.Post("/shorten", SaveURLHandler)

	app.Listen(":8080")
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUrl := os.Getenv("TURSO_DATABASE_URL")
	dbAuthToken := os.Getenv("TURSO_AUTH_TOKEN")
	uri := fmt.Sprintf("%s?authToken=%s", dbUrl, dbAuthToken)

	conn, err := sql.Open("libsql", uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", uri, err)
		os.Exit(1)
	}

	db, err := gorm.Open(sqlite.New(sqlite.Config{Conn: conn}), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	db.AutoMigrate(&url.UrlDbEntity{})

	// inMemoryRepo := url.NewInMemoryRepository()
	sqliteRepo := url.NewSqliteRepo(db)
	service = url.NewService(sqliteRepo)
}

type SaveURLRequest struct {
	LongURL string `json:"long_url"`
}

func SaveURLHandler(c *fiber.Ctx) error {
	request := new(SaveURLRequest)

	if err := c.BodyParser(request); err != nil {
		fmt.Println(err)
		return err
	}

	url, err := service.Create(request.LongURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"message": "Something went wrong",
		})
	}
	return c.JSON(url)
}

func RedirectHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	url, err := service.FindByShortID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]any{
			"message": "URL not found",
		})
	}
	return c.Redirect(url.OriginalUrl)
}

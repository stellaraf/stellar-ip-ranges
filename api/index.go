package handler

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stellaraf/stellar-ip-ranges/lib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	config := fiber.Config{
		AppName: "stellar-ip-ranges",
		Network: "tcp",
	}

	app := fiber.New(config)
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(cache.New(cache.Config{
		Expiration:   15 * time.Minute,
		CacheControl: true,
	}))

	// app.Get("*", func(ctx *fiber.Ctx) error {
	// 	log.Println(ctx.AllParams())
	// 	return ctx.SendStatus(200)
	// })

	app.Get("/json/:scope", lib.JSONHandler)

	app.Get("/json", lib.JSONHandler)

	app.Get("/:type/json", lib.JSONHandler)
	app.Get("/:scope/json", lib.JSONHandler)

	app.Get("/:scope/:type/json", lib.JSONHandler)
	app.Get("/:scope/:type", lib.TextHandler)
	app.Get("/:scope", lib.TextHandler)
	app.Get("/:type", lib.TextHandler)

	app.Get("/", lib.TextHandler)

	return adaptor.FiberApp(app)
}

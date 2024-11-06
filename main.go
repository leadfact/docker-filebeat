package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"go.elastic.co/ecszerolog"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"./logs/out.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := ecszerolog.New(file).With().CallerWithSkipFrameCount(2).Logger()
	logger.WithLevel(zerolog.TraceLevel)

	logger.Info().Msg("Starting code")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		logger.Info().Msg(fmt.Sprintf("root access: %s", string(c.Body())))
		return c.SendString("Hello, World!")
	})

	app.Get("/adr", func(c *fiber.Ctx) error {

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("GET")
		req.Header.SetContentType("application/json")
		req.SetRequestURI("https://ifconfig.me")

		res := fasthttp.AcquireResponse()
		err = fasthttp.Do(req, res)
		logger.Info().Msg(log(req, res))

		return c.SendString(string(res.Body()))
	})

	app.Get("/logs", func(c *fiber.Ctx) error {
		logger.Info().Msg("/logs access")
		return c.SendString("/logs")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		logger.Info().Msg("/hello access")
		return c.SendString("/hello")
	})

	app.Get("/hello/world", func(c *fiber.Ctx) error {
		logger.Info().Msg("/hello/world access")
		return c.SendString("/hello/world")
	})

	logger.Info().Msg("Server started on port 8080")

	err = app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

func log(req *fasthttp.Request, resp *fasthttp.Response) string {
	return fmt.Sprintf("\n------------------NEW p2p_trade REQ------------------\n%s\n-------------------------------------------\n%s\n", req.String(), resp.String())
}

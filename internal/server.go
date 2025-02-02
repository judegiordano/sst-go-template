package internal

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/judegiordano/gogetem/pkg/fibererrors"
	"github.com/judegiordano/sst_template/api/dev"
	"github.com/judegiordano/sst_template/api/health"
	"github.com/judegiordano/sst_template/api/metrics"
	"github.com/judegiordano/sst_template/middleware"
)

func Server() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler:      fibererrors.ErrorHandler,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		EnablePrintRoutes: false,
	})
	// middleware
	metrics.Router(app)
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        60,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return fibererrors.TooManyRequests(c, errors.New("too many requests"))
		},
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(cache.New(cache.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
			ms := middleware.CACHE_DEFAULT_TIME.Milliseconds()
			mss := strconv.Itoa(int(ms))
			newCacheTime, _ := strconv.Atoi(c.GetRespHeader(middleware.CACHE_KEY, mss))
			return time.Millisecond * time.Duration(newCacheTime)
		},
		CacheHeader:  "X-Cache",
		CacheControl: true,
		MaxBytes:     1_000_000_000, // 1gb
		Methods:      []string{fiber.MethodGet},
	}))
	// routes
	dev.Router(app)
	health.Router(app)
	return app
}

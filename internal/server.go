package internal

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	FiberLambdaAdapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
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
	"github.com/judegiordano/gogetem/pkg/logger"
	"github.com/judegiordano/sst_template/api/dev"
	"github.com/judegiordano/sst_template/api/health"
	"github.com/judegiordano/sst_template/middleware"
)

var App *fiber.App
var Lambda *FiberLambdaAdapter.FiberLambda

func init() {
	logger.Debug("[ENV]", Env)
	App = fiber.New(fiber.Config{
		ErrorHandler:      fibererrors.ErrorHandler,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		EnablePrintRoutes: false,
	})
	// middleware
	App.Use(compress.New())
	App.Use(recover.New())
	App.Use(cors.New())
	App.Use(etag.New())
	App.Use(helmet.New())
	App.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        100,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return fibererrors.TooManyRequests(c, errors.New("too many requests"))
		},
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	App.Use(cache.New(cache.Config{
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
	dev.Router(App)
	health.Router(App)
	// lambda adaptor
	Lambda = FiberLambdaAdapter.New(App)
}

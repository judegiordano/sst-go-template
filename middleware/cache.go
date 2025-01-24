package middleware

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const CACHE_KEY = "Cache-Time"
const CACHE_DEFAULT_TIME = time.Minute

func Cache(c *fiber.Ctx, exp time.Duration) {
	ms := exp.Milliseconds()
	mss := strconv.Itoa(int(ms))
	c.Response().Header.Add(CACHE_KEY, mss)
}

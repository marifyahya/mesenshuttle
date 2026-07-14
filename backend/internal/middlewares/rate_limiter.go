package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"mesenshuttle-backend/pkg/utils"
)

func RateLimiter(redisClient *redis.Client, maxRequests int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if redisClient == nil {
			c.Next()
			return
		}

		ip := c.ClientIP()
		userAgent := c.Request.UserAgent()
		key := "rate_limit:" + ip + ":" + userAgent

		ctx := context.Background()
		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			c.Next()
			return
		}

		if count == 1 {
			redisClient.Expire(ctx, key, window)
		}

		if count > int64(maxRequests) {
			utils.ErrorResponse(c, http.StatusTooManyRequests, "Too many requests, please try again later.")
			c.Abort()
			return
		}

		c.Next()
	}
}

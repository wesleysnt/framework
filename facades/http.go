package facades

import (
	"github.com/wesleysnt/framework/contracts/http"
)

func RateLimiter() http.RateLimiter {
	return App().MakeRateLimiter()
}

func View() http.View {
	return App().MakeView()
}

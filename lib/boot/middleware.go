// Package boot handles the initialization of the web components.
package boot

import (
	"net/http"

	"github.com/pcieslar/goforge/middleware/logrequest"
	"github.com/pcieslar/goforge/middleware/rest"
	"github.com/pcieslar/goforge/core/router"
	"github.com/gorilla/context"
)

// SetUpMiddleware contains the middleware that applies to every request.
func SetUpMiddleware(h http.Handler) http.Handler {
	return router.ChainHandler( // Chain middleware, top middleware runs first
		h,                    // Handler to wrap
		setUpCSRF,            // Prevent CSRF
		rest.Handler,         // Support changing HTTP method sent via query string
		logrequest.Handler,   // Log every request
		context.ClearHandler, // Prevent memory leak with gorilla.sessions
	)
}

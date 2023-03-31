package route

import (
	"net/http"

	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/router"
)

func Handler() http.Handler {
	router.Handle(http.MethodGet, "/_health", health)

	router.Handle(http.MethodPost, "/", index)

	return router.Handler()
}

func health(c *router.Context) {
	c.OK(map[string]any{"version": constant.VERSION, "bidtime": constant.BIDTIME})
}

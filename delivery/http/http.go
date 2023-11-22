package http

import (
	"net/http"
	"regexp"

	echoSentry "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/teq-quocbang/arrows/config"
	"github.com/teq-quocbang/arrows/delivery/http/account"
	"github.com/teq-quocbang/arrows/delivery/http/auth"
	"github.com/teq-quocbang/arrows/delivery/http/checkout"
	"github.com/teq-quocbang/arrows/delivery/http/example"
	"github.com/teq-quocbang/arrows/delivery/http/healthcheck"
	"github.com/teq-quocbang/arrows/delivery/http/post"
	"github.com/teq-quocbang/arrows/delivery/http/producer"
	"github.com/teq-quocbang/arrows/delivery/http/product"
	"github.com/teq-quocbang/arrows/delivery/http/statistics"
	"github.com/teq-quocbang/arrows/delivery/http/storage"
	"github.com/teq-quocbang/arrows/usecase"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	var (
		e         = echo.New()
		loggerCfg = middleware.DefaultLoggerConfig
	)

	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/health-check"
	}

	e.Use(middleware.LoggerWithConfig(loggerCfg))
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echoSentry.New(echoSentry.Options{Repanic: true}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOriginFunc: func(origin string) (bool, error) {
			return regexp.MatchString(
				`^https:\/\/(|[a-zA-Z0-9]+[a-zA-Z0-9-._]*[a-zA-Z0-9]+\.)teqnological.asia$`,
				origin,
			)
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))

	// Health check
	healthcheck.Init(e.Group("/health-check"))

	// API docs
	if !config.GetConfig().Stage.IsProd() {
		e.GET("/docs/*", echoSwagger.WrapHandler)
	}

	// APIs
	api := e.Group("/api")
	example.Init(api.Group("/examples"), useCase)
	account.Init(api.Group("/user"), useCase)
	product.Init(api.Group("/product", auth.Auth), useCase)
	producer.Init(api.Group("/producer", auth.Auth), useCase)
	product.ProductsInit(api.Group("/products", auth.Auth), useCase)
	storage.Init(api.Group("/storage", auth.Auth), useCase)
	checkout.Init(api.Group("/checkout", auth.Auth), useCase)
	statistics.Init(api.Group("/statistics", auth.Auth), useCase)

	// arrows
	post.Init(api.Group("/post", auth.Auth), useCase)

	return e
}

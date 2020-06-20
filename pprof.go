package pprof_echo

import (
	"net/http/pprof"
	"strings"

	"github.com/labstack/echo/v4"
)

// Wrap wraps default pprof endpoints into Echo compliant way
func Wrap(e *echo.Echo) {
	wrapGroup(e.Group(""))
}

func wrapGroup(g *echo.Group) {
	routers := []struct {
		Method  string
		Path    string
		Handler echo.HandlerFunc
	}{
		{"GET", "/debug/pprof", indexHandler()},
		{"GET", "/debug/heap", heapHandler()},
		{"GET", "/debug/goroutine", goroutineHandler()},
		{"GET", "/debug/block", blockHandler()},
		{"GET", "/debug/threadcreate", threadCreateHandler()},
		{"GET", "/debug/cmdline", cmdlineHandler()},
		{"GET", "/debug/profile", profileHandler()},
		{"GET", "/debug/symbol", symbolHandler()},
		{"POST", "/debug/symbol", symbolHandler()},
		{"GET", "/debug/trace", traceHandler()},
		{"GET", "/debug/mutex", mutexHandler()},
		{"GET", "/debug/allocs", allocsHandler()},
	}

	for _, r := range routers {
		switch r.Method {
		case "GET":
			g.GET(strings.TrimPrefix(r.Path, ""), r.Handler)
		case "POST":
			g.POST(strings.TrimPrefix(r.Path, ""), r.Handler)
		}
	}
}

func indexHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Index(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func allocsHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("allocs").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func heapHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("heap").ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	}
}

func goroutineHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("goroutine").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func blockHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("block").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func threadCreateHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func cmdlineHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Cmdline(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func profileHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Profile(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func symbolHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Symbol(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func traceHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Trace(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

func mutexHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("mutex").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

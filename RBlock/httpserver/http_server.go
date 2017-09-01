package httpserver

import (

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
	"github.com/artjoma/flog"
	"log"
	"io/ioutil"
)


func ServeFiles(r *fasthttprouter.Router, path string, rootPath string) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("Path must end with /*filepath in path '" + path + "'")
	}
	prefix := path[:len(path)-10]

	fileHandler := FSHandler(rootPath, strings.Count(prefix, "/"))

	r.GET(path, func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Cache-Control", "no-cache")
		fileHandler(ctx)
	})
}
func FSHandler(root string, stripSlashes int) fasthttp.RequestHandler {
	fs := &fasthttp.FS{
		Root:               root,
		IndexNames:         []string{"index.htm"},
		GenerateIndexPages: false,
		AcceptByteRange:    true,
		CacheDuration:      1 * time.Second,
	}
	if stripSlashes > 0 {
		fs.PathRewrite = fasthttp.NewPathSlashesStripper(stripSlashes)
	}
	return fs.NewRequestHandler()
}

func CreateHttpServer(logger *flog.Logger, address string ,port string, router *fasthttprouter.Router) {
	//Server
	server := &fasthttp.Server{Handler: router.Handler}
	server.Name = "#"
	server.Concurrency = 8192
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.ReadBufferSize = 4096
	server.WriteBufferSize = 4096
	server.MaxRequestBodySize = 1024 * 256
	//of fasthttp logger
	log := log.New(ioutil.Discard, "", log.LstdFlags)
	server.Logger = fasthttp.Logger(log)
	logger.Info("Trying to start HTTP server async. Address: "+ address)
	go func() {
		errListen := server.ListenAndServe(port)
		if errListen != nil {
			logger.Err("Server shutdown ! " + errListen.Error())
		}
	}()
}

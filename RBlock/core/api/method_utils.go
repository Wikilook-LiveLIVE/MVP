package api

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"encoding/json"
	"RBlock/errors"
)

const(
	TOKEN_HEADER = "token"
)

func writeErrResponse(ctx *fasthttp.RequestCtx, err *errors.AppErr) {
	var code int = 200
	switch err {
	case errors.ERR_METHOD_UNAVAILABLE:
		code = fasthttp.StatusInternalServerError
	case errors.ERR_INVALID_REQUEST:
		code = fasthttp.StatusBadRequest
	case errors.ERR_SECURITY_TOKEN_INVALID:
		code = fasthttp.StatusUnauthorized
	default:
		code = 200
	}
	ctx.Response.SetStatusCode(code)

	ctx.Response.Header.Set("Cache-Control", "no-cache")
	ctx.Response.Header.SetContentType("application/json; charset=utf-8")
	ctx.Response.AppendBodyString("{\"err\":{\"id\":" +
		strconv.Itoa(err.Id) + ",\"code\":\"" + err.Code + "\"}}")
}
//jsonFields plain object ot map
func  writeResponse(ctx *fasthttp.RequestCtx, jsonFields interface{}) {
	ctx.Response.SetStatusCode(200)
	ctx.Response.Header.Set("Cache-Control", "no-cache")
	ctx.Response.Header.SetContentType("application/json; charset=utf-8")
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(jsonFields)
}




package api

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"path/filepath"
	"runtime"
	"RBlock/errors"
	"RBlock/httpserver"
	"RBlock/core/types"
)

var(
	RESPONSE_EMPTY_OK = map[string]interface{}{"resp": "ok"}
)
type HttpApiController struct {
}

//free resources
func (self *HttpApiController) DestroyHttpController() {
	//TODO shutdown
}

func NewHttpApiController(address string, port string, appFolder string) *HttpApiController {
	router := fasthttprouter.New()
	//static pages
	httpserver.ServeFiles(router, "/ui/*filepath", filepath.Join(appFolder, "/ui/"))
	controller := &HttpApiController{}
	router.PanicHandler = controller.PanicHandler
	/*
		API
	*/
	//TODO: SignUp

	//create server
	httpserver.CreateHttpServer(log, address, port, router)

	return controller
}

func (self *HttpApiController) PanicHandler(ctx *fasthttp.RequestCtx, error interface{}) {
	stack := make([]byte, 2512)
	runtime.Stack(stack[:], false)
	log.Err(fmt.Sprintf( "%q\n%s\n", error, stack[:]))
	writeErrResponse(ctx, errors.ERR_METHOD_UNAVAILABLE)
}
/*
Sign Up
 */
func (self *HttpApiController) SignUp(ctx *fasthttp.RequestCtx) {
	log.Debug("sign-up")
	signUp := &types.SignUp{}
	if signUp.FromJSON(ctx.PostBody()) == nil {
		if err := customerService.SignUp(signUp); err == nil {
			writeResponse(ctx, RESPONSE_EMPTY_OK)
		} else {
			writeErrResponse(ctx, err)
		}
	} else {
		writeErrResponse(ctx, errors.ERR_INVALID_REQUEST)
	}
}

/*
Activate Customer after SignUp
 */
func (self *HttpApiController) SignUpActivate(ctx *fasthttp.RequestCtx) {
	//sign-up-activate?activ-code=aAsdF783HjfaR53eDfcF
	code := ctx.FormValue("activ-code")
	//check is empty ?
	if code == nil {
		ctx.Redirect("/ui/htm/#sign-up-activ-err", 200)
	} else {
		activationCode := string(code)
		log.Info(fmt.Sprint(activationCode))
		customerService.SignUpActivation(activationCode)

		ctx.Redirect("/ui/htm/#sign-up-activ-ok", 200)
	}
}
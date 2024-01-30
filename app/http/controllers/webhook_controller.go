package controllers

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/json"
)

type WebhookController struct {
	//Dependent services
}

func NewWebhookController() *WebhookController {
	return &WebhookController{
		//Inject services
	}
}

func (r *WebhookController) Index(ctx http.Context) http.Response {
	header, err := json.Marshal(ctx.Request().Headers())
	all, err2 := json.Marshal(ctx.Request().All())

	var tmp string
	if err == nil {
		tmp = fmt.Sprint("Header : ", string(header))
	}

	if err2 == nil {
		tmp = fmt.Sprint(tmp, "\nAll: ", string(all))
	}

	err3 := facades.Storage().Put("log.txt", tmp)
	if err3 != nil {

	}
	return ctx.Response().String(200, tmp)
}

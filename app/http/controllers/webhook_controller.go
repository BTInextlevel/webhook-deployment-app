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
	//all, err2 := json.Marshal(ctx.Request().All())

	var tmp string
	if err == nil {
		tmp = fmt.Sprint("Header  isi: ", string(header))
	}

	//if err2 == nil {
	//	tmp = fmt.Sprint(tmp, "\nAll: ", "") //  string(all))
	//}

	var data map[string][]interface{}
	errx := ctx.Request().Bind(&data)

	old, _ := facades.Storage().Get("log.txt")
	if errx == nil {

		payload, errun := json.MarshalString(data)
		fmt.Printf("")
		if errun == nil {
			tmp = fmt.Sprint(tmp, "\nPayloads", payload, "\n\n")
		} else {
			fmt.Printf("error marshal ", errun.Error())
		}
	} else {
		fmt.Printf("error errx get log : ", errx.Error())
	}

	tmp = fmt.Sprint(old, "\n", tmp)
	err3 := facades.Storage().Put("log.txt", tmp)
	if err3 != nil {

	}
	return ctx.Response().String(200, tmp)
}

package controllers

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/json"
	"os/exec"
	"strings"
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

	var data map[string]interface{}
	errx := ctx.Request().Bind(&data)

	old, _ := facades.Storage().Get("log.txt")
	if errx == nil {

		payload, errun := json.MarshalString(data)
		fmt.Printf("")
		if errun == nil {
			tmp = fmt.Sprint(tmp, "\nPayloads", payload, "\n\n")
		} else {
			fmt.Printf("error di marshal ", errun.Error())
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

func ListenWebHook(ctx http.Context) http.Response {
	var payload map[string]interface{}
	err := ctx.Request().Bind(&payload)
	if err != nil {
		return ctx.Response().Json(401, http.Json{
			"message": err.Error(),
		})
	}

	fileconfig, errconfig := facades.Storage().Get("config.js")
	if errconfig != nil {
		return ctx.Response().Json(401, http.Json{
			"message": "Tidak ada configurasi di tentukan",
		})
	}

	var js map[string]interface{}
	errUm := json.Unmarshal([]byte(fileconfig), &js)
	if errUm != nil {
		return ctx.Response().Json(401, http.Json{
			"message": "Gagal membuka file config",
		})
	}

	elRepository, ok := js["repositories"].([]interface{})
	if !ok {
		return ctx.Response().Json(401, http.Json{
			"message": "tidak ada repository diatur pada config",
		})
	}

	elRef, ok := js["ref"].(string)

	namePush := getNamePush(ctx)

	for _, repo := range elRepository {
		repoMap, ok := repo.(map[string]interface{})
		if !ok {
			continue
		}

		nameRef, ok := repoMap["ref"].(string)
		nameRepo, ok := repoMap["name"].(string)

		if nameRepo == *namePush {
			if elRef == nameRef {
				direktori, _ := repoMap["directory"].(string)
				cmds := repoMap["command"].([]interface{})
				eksekusi(direktori, cmds)
			}
		}
	}

	return ctx.Response().Json(200, http.Json{
		"message": "OK",
	})
}

func eksekusi(direktori string, cmds []interface{}) {
	for _, perintah := range cmds {
		pp := strings.Split(perintah.(string), " ")
		cmd := exec.Command(pp[0], pp[1:]...)
		cmd.Dir = direktori
		_, er := cmd.CombinedOutput()
		if er != nil {
			continue
		}
	}
}

func getNamePush(ctx http.Context) *string {
	var js map[string]interface{}

	err := ctx.Request().Bind(&js)
	if err != nil {
		return nil
	}

	repo, ok := js["repository"].(map[string]interface{})
	if !ok {
		return nil
	}

	name, ok := repo["name"].(string)
	if !ok {
		return nil
	}
	return &name
}

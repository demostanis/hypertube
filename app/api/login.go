package api

import (
	"fmt"
	"net/http"

	"github.com/demostanis/hypertube/pages"
	ghttp "maragu.dev/gomponents/http"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func APILoginHandler(w http.ResponseWriter, r *http.Request) {
	var params LoginParams

	paramsInto(&params, w, r)

	token, err := auth("crocotube-auth", params.Username, params.Password, "default")
	if err != nil {
		apiError(w, r, pages.Login, err.Error())
		return
	}

	fmt.Printf("access token: %s\n", token)
	ghttp.Adapt(pages.LoginHandler)(w, r)
}

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

	err := paramsInto(&params, w, r)
	if err != nil {
		return
	}

	token, err := auth("crocotube-auth", params.Username, params.Password, "default")
	if err != nil {
		_, _ = apiError(w, r, pages.Login, err.Error())
		return
	}

	fmt.Printf("access token: %s\n", token)
	ghttp.Adapt(pages.LoginHandler)(w, r)
}

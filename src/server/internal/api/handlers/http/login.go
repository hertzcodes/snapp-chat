package http

import (
	"encoding/json"
	"net/http"

	"github.com/hertzcodes/snapp-chat/server/internal/api/handlers/common"
	"github.com/hertzcodes/snapp-chat/server/internal/app"
)

func Login(appContainer app.App) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var req common.LoginRequest
		defer r.Body.Close()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		service := appContainer.UserService()
		id, err := service.SignIn(req)

		if err != nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		}
		_ = id

	}
}

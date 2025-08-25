package users

import (
	"apigo/usecase/users"
	"encoding/json"
	"net/http"
)

type UserGetHandler struct {
	usergetUsecase users.UserGetUsecase
}

func NewUserHandler(usecase users.UserGetUsecase) UserGetHandler {
	return UserGetHandler{
		usergetUsecase: usecase,
	}
}

func (us *UserGetHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := us.usergetUsecase.GetUser()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

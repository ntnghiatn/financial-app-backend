package v1

import (
	"encoding/json"
	"net/http"

	"github.com/ntnghiatn/financial-app-backend/internal/api/utils"
	"github.com/ntnghiatn/financial-app-backend/internal/database"
	"github.com/ntnghiatn/financial-app-backend/internal/model"
	"github.com/sirupsen/logrus"
)

type UserAPI struct {
	DB database.Database //will represent all database interface
}

type UserParameters struct {
	model.User
	Password string `json:"password"`
}

func (api *UserAPI) Create(w http.ResponseWriter, r *http.Request) {
	// show func name in logs to track error faster
	log := logrus.WithField("func", "user.go -> Create()")

	//Load Parameters
	var userParameters UserParameters
	if err := json.NewDecoder(r.Body).Decode(&userParameters); err != nil {
		log.WithError(err).Warn("could not decode parametes")
		utils.WriteError(w, http.StatusBadRequest, "could not decode parameters", map[string]string{"error": err.Error()})
		return
	}

	//
	log = log.WithFields(logrus.Fields{
		"email": userParameters.Email,
	})

	//
	if err := userParameters.Verify(); err != nil {
		log.WithError(err).Warn("Not all fields found")
		utils.WriteError(w, http.StatusBadRequest, "Not all fields found", map[string]string{
			"error": err.Error(),
		})
	}

	//
	hashed, err := model.HashPassword(userParameters.Password)
	if err != nil {
		log.WithError(err).Warn("could not hash password")
		utils.WriteError(w, http.StatusInternalServerError, "could not hash password", nil)
		return
	}

	newUser := &model.User{
		Email:        userParameters.Email,
		PasswordHash: &hashed,
	}

	log.Info(newUser)

}

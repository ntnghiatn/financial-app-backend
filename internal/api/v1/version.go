package v1

import (
	"encoding/json"
	"net/http"

	"github.com/ntnghiatn/financial-app-backend/internal/config"
	"github.com/sirupsen/logrus"
)

//API for return version

//When server starts, we set version and then use it necessery .

// Server version represents the server version
type ServerVersion struct {
	Version string `json:"version"`
}

// Machaled JSON
var versionJSON []byte

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})
	if err != nil {
		panic(err)
	}

	//
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("Testtttt")

	students := []string{"hs1", "hs2", "hs3"}

	seed := Find(students)

	logrus.Infoln("Ok con do", seed("hello"))

	w.WriteHeader(200)

	if _, err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error writing version.")
	}
}

func Find(inps []string) func(s string) bool {
	return func(s string) bool {
		isExist := false
		for i := 0; i < len(inps); i++ {
			if inps[i] == s {
				isExist = true
			}
		}
		return isExist
	}

}

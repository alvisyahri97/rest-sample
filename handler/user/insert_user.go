package user

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/alvisyahri97/rest-sample/repo/users"

	"github.com/go-http-utils/headers"
)

func (h *Handler) InsertUser(w http.ResponseWriter, r *http.Request) {
	var newUser users.User

	switch r.Header.Get(headers.ContentType) {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			WriteResp(w, r, err, http.StatusBadRequest, nil)
			return
		}
	case "application/xml":
		err := xml.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			WriteResp(w, r, err, http.StatusBadRequest, nil)
			return
		}
	}
	if newUser == (users.User{}) {
		WriteResp(w, r, errors.New("parameter not valid"), http.StatusBadRequest, nil)
		return
	}

	err := h.userRepo.InsertUser(&newUser)
	if err != nil {
		WriteResp(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	WriteResp(w, r, nil, http.StatusOK, newUser)
}

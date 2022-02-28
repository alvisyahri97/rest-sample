package user

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strconv"

	"github.com/alvisyahri97/rest-sample/repo/users"

	"github.com/go-http-utils/headers"

	"github.com/go-chi/chi"
)

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	if err != nil {
		WriteResp(w, r, err, http.StatusBadRequest, nil)
		return
	}

	var updatedUser users.User
	switch r.Header.Get(headers.ContentType) {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			WriteResp(w, r, err, http.StatusBadRequest, nil)
			return
		}
	case "application/xml":
		err := xml.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			WriteResp(w, r, err, http.StatusBadRequest, nil)
			return
		}
	}
	if updatedUser == (users.User{}) {
		WriteResp(w, r, errors.New("parameter not valid"), http.StatusBadRequest, nil)
		return
	}
	updatedUser.ID = userid

	err = h.userRepo.UpdateByID(userid, updatedUser)
	if err != nil {
		WriteResp(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	WriteResp(w, r, nil, http.StatusOK, updatedUser)
}

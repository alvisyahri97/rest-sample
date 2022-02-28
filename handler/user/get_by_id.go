package user

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	if err != nil {
		WriteResp(w, r, err, http.StatusBadRequest, nil)
		return
	}

	user, err := h.userRepo.GetByID(userid)
	if err != nil {
		WriteResp(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	WriteResp(w, r, nil, http.StatusOK, user)
}

package user

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	if err != nil {
		WriteResp(w, r, err, http.StatusBadRequest, nil)
		return
	}

	err = h.userRepo.DeleteByID(userid)
	if err != nil {
		WriteResp(w, r, err, http.StatusInternalServerError, nil)
		return
	}
	WriteResp(w, r, nil, http.StatusOK, resp{
		Message: "Success",
	})
}

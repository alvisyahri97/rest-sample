package user

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/go-http-utils/headers"
)

func WriteResp(w http.ResponseWriter, r *http.Request, err error, httpStatus int, data interface{}) {
	if err != nil {
		WriteErrResp(w, r, httpStatus, err)
		return
	}
	ParseResp(w, r, data, httpStatus)
}

func WriteErrResp(w http.ResponseWriter, r *http.Request, httpStatus int, err error) {
	if err == nil {
		return
	}

	ParseResp(w, r, resp{
		Message: err.Error(),
	}, httpStatus)
}

func ParseResp(w http.ResponseWriter, r *http.Request, data interface{}, httpStatus int) {
	var resp []byte
	switch r.Header.Get(headers.Accept) {
	case "application/json":
		resp, _ = json.Marshal(data)
		w.Header().Add(headers.ContentType, "application/json")
	case "application/xml":
		resp, _ = xml.Marshal(data)
		w.Header().Add(headers.ContentType, "application/xml")
	default:
		WriteErrResp(w, r, http.StatusNotAcceptable, errors.New("request not acceptable"))
	}

	w.WriteHeader(httpStatus)
	w.Write(resp)
}

package render

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/user"
	"path"
	"text/template"

	"github.com/go-chi/chi"
)

var baseURL = "http://localhost:8000"

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var client = &http.Client{}
	var res user.User
	request, err := http.NewRequest("GET", baseURL+"/user/"+chi.URLParam(r, "userid"), nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "profile",
		"name":  res.Name,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Routes map[string]func(w http.ResponseWriter, r *http.Request)

type Data struct {
	Row []Row `json:"row"`
}

type Row struct {
	Project string   `json:"project"`
	CVEs    []string `json:"cves"`
}

func CustomRoutes(data *Data) Routes {
	r := make(Routes)

	for _, d := range data.Row {
		func(d Row) {
			key := fmt.Sprintf("/%s", d.Project)
			r[key] = routeFactory(&d)
		}(d)
	}

	return r
}

func routeFactory(d *Row) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(d)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func DefaultRoutes() Routes {
	r := make(Routes)
	r["/"] = home
	r["/about"] = about

	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`HOME`))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is what we are about!"))
}

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func encode[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func decode[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode json: %w", err)
	}
	return data, nil
}

func (r Router) group(patter string, fn func(Router)) {
	internalRouter := NewRouter()
	fn(internalRouter)
	// fmt.Printf("pattern: '%s'\n", patter)
	r.Mux.Handle(patter+"/", http.StripPrefix(patter, internalRouter.Mux))
	// http.StripPrefix("/", internalRouter.Mux)
}

func (r Router) get(patter string, handler func(http.ResponseWriter, *http.Request)) {
	r.Mux.HandleFunc(fmt.Sprintf("GET %s", patter), handler)
}

func (r Router) put(patter string, handler func(http.ResponseWriter, *http.Request)) {
	r.Mux.HandleFunc(fmt.Sprintf("PUT %s", patter), handler)
}

func (r Router) patch(patter string, handler func(http.ResponseWriter, *http.Request)) {
	r.Mux.HandleFunc(fmt.Sprintf("PATCH %s", patter), handler)
}

func (r Router) post(patter string, handler func(http.ResponseWriter, *http.Request)) {
	r.Mux.HandleFunc(fmt.Sprintf("POST %s", patter), handler)
}

func (r Router) delete(
	patter string,
	handler func(http.ResponseWriter, *http.Request),
) {
	r.Mux.HandleFunc(fmt.Sprintf("DELETE %s", patter), handler)
}

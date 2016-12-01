package generator

import (
	"encoding/json"
	"log"
	"net/http"
)

type idGenResponseGet struct {
	ID int32 `json:"id"`
}

type idGenHandler struct {
	idgen idGenerator
}

func (h *idGenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)
	encoder := json.NewEncoder(w)
	encoder.Encode(idGenResponseGet{ID: h.idgen.Generate()})
	//w.Write([]byte("{\"id\":" + strconv.Itoa(int(h.idgen.Generate())) + "}"))
}

func NewIdGenHandler(idgen idGenerator) *idGenHandler {
	return &idGenHandler{idgen: idgen}
}

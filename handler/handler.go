package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fmusayev/golang-excel/types"
	util "github.com/fmusayev/golang-excel/utils"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Handler struct{}

func NewHandler(router *mux.Router) *Handler {
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	h := &Handler{}

	router.HandleFunc("/readiness", h.Health)
	router.HandleFunc("/health", h.Health)
	router.HandleFunc("/v1/generate-excel", h.DownloadExcel).Methods("POST")

	return h
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DownloadExcel(w http.ResponseWriter, r *http.Request) {
	log.Debug("DownloadExcel.start")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Error("Exception.DownloadExcel.ReadBody", err)
		http.Error(w, err.Error(), 400)
		return
	}

	var input []types.ExcelForm
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Error("Exception.DownloadExcel.Unmarshal", err)
		http.Error(w, err.Error(), 400)
		return
	}

	xlsx := util.OpenExcel("resources/main.xlsx")
	// first indx is column headers
	indx := 2
	for _, v := range input {
		util.WriteData(v, indx, xlsx)
		indx++
	}

	w.Header().Add("Content-Disposition", "attachment; filename=main.xlsx")
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.WriteHeader(http.StatusOK)
	xlsx.Write(w)
	log.Debug("DownloadExcel.end")
}

package handler

import (
	"net/http"

	"github.com/fatahnuram/learn-go-kasir-api/internal/helpers"
	"github.com/fatahnuram/learn-go-kasir-api/internal/service"
)

type ReportHandler struct {
	service service.ReportService
}

func NewReportHandler(service service.ReportService) ReportHandler {
	return ReportHandler{
		service: service,
	}
}

func (h ReportHandler) GetReportToday() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := h.service.GetReportToday()
		if err != nil {
			helpers.RespondJson(w, r, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, result)
	})
}

func (h ReportHandler) GetReportByTimeRange() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startdate := r.URL.Query().Get("start_date")
		enddate := r.URL.Query().Get("end_date")

		if startdate == "" || enddate == "" {
			helpers.RespondJson(w, r, http.StatusBadRequest, map[string]string{
				"error": "either start date or end date are missing",
			})
			return
		}

		result, err := h.service.GetReportByTimeRange(startdate, enddate)
		if err != nil {
			helpers.RespondJson(w, r, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		helpers.RespondJson(w, r, http.StatusOK, result)
	})
}

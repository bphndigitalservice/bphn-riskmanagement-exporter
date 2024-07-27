package server

import (
	"bphn.go.id/mr-report/report/builder"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	excelBuilder *builder.ExcelBuilder
}

func NewHandler(excelBuilder *builder.ExcelBuilder) *Handler {
	return &Handler{
		excelBuilder: excelBuilder,
	}
}

func (h *Handler) GenerateReport(w http.ResponseWriter, r *http.Request) {

	syear := r.URL.Query().Get("year")

	year, err := strconv.ParseInt(syear, 10, 32)
	if err != nil {
		panic(err)
	}

	f, err := h.excelBuilder.Export(int(year))

	if err != nil {
		fmt.Errorf("Error generating report: %v", err)
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Path))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	if err := f.Write(w); err != nil {
		fmt.Fprint(w, err.Error())
	}
}

package builder

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func (ex *ExcelBuilder) getDataEvaluasiRisiko(year int) (DataEvaluasiRisiko, error) {
	var data DataEvaluasiRisiko
	risks, err := ex.repository.GetRiskAnalysisByYear(year)
	if err != nil {
		return DataEvaluasiRisiko{}, fmt.Errorf("get risk evaluation for year %d: %v", year, err)
	}

	data = DataEvaluasiRisiko{
		Risks: risks,
	}

	return data, nil
}

func (ex *ExcelBuilder) createEvaluasiRisikoHeader(f *excelize.File, period int) {
	f.SetCellValue(SheetEvaluasiRisiko, SheetHeader_EvaluasiRisiko_valueRangeStart, SheetHeader_EvaluasiRisiko)
	f.MergeCell(SheetEvaluasiRisiko, SheetHeader_EvaluasiRisiko_valueRangeStart, SheetHeader_EvaluasiRisiko_ValueRangeEnd)

	style, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 14,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})

	sheetHeaderStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 14,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
	})

	f.SetCellStyle(SheetEvaluasiRisiko, "A2", "H2", sheetHeaderStyle)
	f.SetCellStyle(SheetEvaluasiRisiko, "A4", "C5", style)
	f.SetCellValue(SheetEvaluasiRisiko, "A4", "Unit Pemilik Risiko")
	f.SetCellValue(SheetEvaluasiRisiko, "C4", ": BADAN PEMBINAAN HUKUM NASIONAL")
	f.SetCellValue(SheetEvaluasiRisiko, "A5", "Periode Penerapan")
	f.SetCellValue(SheetEvaluasiRisiko, "C5", fmt.Sprintf(": %d", period))

}

func (ex *ExcelBuilder) createEvaluasiRisikoTable(f *excelize.File) {

	style, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "gradient", Color: []string{"95b3d7", "95b3d7"}, Shading: 1},
		Font: &excelize.Font{
			Bold: true,
			Size: 14,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})

	f.SetCellValue(SheetEvaluasiRisiko, "A7", "No.")
	f.MergeCell(SheetEvaluasiRisiko, "A7", "A8")

	f.SetCellValue(SheetEvaluasiRisiko, "B7", "Sisa Risiko")
	f.MergeCell(SheetEvaluasiRisiko, "B7", "B8")

	f.SetCellValue(SheetEvaluasiRisiko, "C7", "Tingkat Risiko")
	f.MergeCell(SheetEvaluasiRisiko, "C7", "C8")

	f.SetCellValue(SheetEvaluasiRisiko, "D7", "Tingkat Prioritas Risiko")
	f.MergeCell(SheetEvaluasiRisiko, "D7", "D8")

	f.SetCellValue(SheetEvaluasiRisiko, "E7", "Tingkat Toleransi Risiko")
	f.MergeCell(SheetEvaluasiRisiko, "E7", "E8")

	f.SetCellValue(SheetEvaluasiRisiko, "F7", "Indikator Risiko")
	f.MergeCell(SheetEvaluasiRisiko, "F7", "H7")

	f.SetCellValue(SheetEvaluasiRisiko, "F8", "Indikasi")
	f.SetCellValue(SheetEvaluasiRisiko, "G8", "Penjelasan")
	f.SetCellValue(SheetEvaluasiRisiko, "H8", "Batas Aman")

	cells := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for i, cell := range cells {
		f.SetCellValue(SheetEvaluasiRisiko, fmt.Sprintf("%s9", cell), i+1)
	}

	f.SetCellStyle(SheetEvaluasiRisiko, "A7", "H9", style)

}

func (ex *ExcelBuilder) fillEvaluasiRisikoData(f *excelize.File, report Report) {
	// Data Start from this row
	startRowNum := SheetEvaluasiRisiko_RowStart
	for i, risk := range report.SheetEvaluasiRisiko.Risks {

		NoCell := fmt.Sprintf("A%d", startRowNum)

		f.SetCellValue(SheetEvaluasiRisiko, NoCell, i+1)
		f.SetColWidth(SheetEvaluasiRisiko, "A", "A", 5)

		// Risk Residual Cell
		RiskResidualCell := fmt.Sprintf("B%d", startRowNum)
		f.SetCellValue(SheetEvaluasiRisiko, RiskResidualCell, risk.SisaRisiko)

		// Risk Level Cell
		KemungkinanUraianCell := fmt.Sprintf("C%d", startRowNum)
		riskLevel := ex.riskLevel(risk.KemungkinanNilai, risk.NilaiPetaRisiko)
		f.SetCellValue(SheetEvaluasiRisiko, KemungkinanUraianCell, riskLevel.Value)

		// Prioritas Risiko Cell
		PrioritasRisikoCell := fmt.Sprintf("D%d", startRowNum)
		f.SetCellValue(SheetEvaluasiRisiko, PrioritasRisikoCell, risk.PrioritasRisiko)

		// Toleransi Risiko Cell
		ToleransiCell := fmt.Sprintf("E%d", startRowNum)
		f.SetCellValue(SheetEvaluasiRisiko, ToleransiCell, risk.Alasan)

		// Indikasi Indikator Cell
		IndikasiIndikatorCell := fmt.Sprintf("F%d", startRowNum)
		f.SetCellValue(SheetEvaluasiRisiko, IndikasiIndikatorCell, risk.IndikasiIndikator)

		PenjelasanIndikatorCell := fmt.Sprintf("G%d", startRowNum)
		f.SetCellValue(SheetEvaluasiRisiko, PenjelasanIndikatorCell, risk.PenjelasanIndikator)

		BatasAmanIndikatorCell := fmt.Sprintf("H%d", startRowNum)
		f.SetCellValue(SheetEvaluasiRisiko, BatasAmanIndikatorCell, risk.BatasAmanIndikator)

		startRowNum++
	}
	f.SetColWidth(SheetEvaluasiRisiko, "B", "G", 25)
	ex.signPlaceholder(f, SheetEvaluasiRisiko, startRowNum+3, "F")
}

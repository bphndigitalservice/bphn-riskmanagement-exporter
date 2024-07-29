package builder

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func (ex *ExcelBuilder) getDataAnalisisRisiko(year int) (DataAnalisisRisiko, error) {
	var data DataAnalisisRisiko
	risks, err := ex.repository.GetRiskAnalysisByYear(year)
	if err != nil {
		return DataAnalisisRisiko{}, fmt.Errorf("get risk analysis for year %d: %v", year, err)
	}

	data = DataAnalisisRisiko{
		Risks: risks,
	}

	return data, nil
}

func (ex *ExcelBuilder) createAnalisisRisikoHeader(f *excelize.File, period int) {
	f.SetCellValue(SheetAnalisisRisiko, SheetHeader_IdentifikasiRisiko_valueRangeStart, SheetHeader_AnalisisRisiko)
	f.MergeCell(SheetAnalisisRisiko, SheetHeader_AnalisisRisiko_valueRangeStart, SheetHeader_AnalisisRisiko_ValueRangeEnd)

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

	f.SetCellStyle(SheetAnalisisRisiko, "A2", "F2", sheetHeaderStyle)
	f.SetCellStyle(SheetHeader_AnalisisRisiko, "A4", "C5", style)
	f.SetCellValue(SheetAnalisisRisiko, "A4", "Unit Pemilik Risiko")
	f.SetCellValue(SheetAnalisisRisiko, "C4", ": BADAN PEMBINAAN HUKUM NASIONAL")
	f.SetCellValue(SheetAnalisisRisiko, "A5", "Periode Penerapan")
	f.SetCellValue(SheetAnalisisRisiko, "C5", fmt.Sprintf(": %d", period))

}

func (ex *ExcelBuilder) createAnalisisRisikoTable(f *excelize.File) {

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

	f.SetCellValue(SheetAnalisisRisiko, "A7", "No.")
	f.MergeCell(SheetAnalisisRisiko, "A7", "A8")

	f.SetCellValue(SheetAnalisisRisiko, "B7", "Sisa Risiko")
	f.MergeCell(SheetAnalisisRisiko, "B7", "B8")

	f.SetCellValue(SheetAnalisisRisiko, "C7", "Kemungkinan")
	f.MergeCell(SheetAnalisisRisiko, "C7", "D7")
	f.SetCellValue(SheetAnalisisRisiko, "C8", "Uraian")
	f.SetCellValue(SheetAnalisisRisiko, "D8", "Nilai")

	f.SetCellValue(SheetAnalisisRisiko, "E7", "Alasan")
	f.MergeCell(SheetAnalisisRisiko, "E7", "E8")

	f.SetCellValue(SheetAnalisisRisiko, "F7", "Dampak")
	f.MergeCell(SheetAnalisisRisiko, "F7", "G7")
	f.SetCellValue(SheetAnalisisRisiko, "F8", "Uraian")
	f.SetCellValue(SheetAnalisisRisiko, "G8", "Nilai")

	f.SetCellValue(SheetAnalisisRisiko, "H7", "Tingkat Risiko")
	f.MergeCell(SheetAnalisisRisiko, "H7", "H8")

	f.SetCellValue(SheetAnalisisRisiko, "I7", "Profil Risiko")
	f.MergeCell(SheetAnalisisRisiko, "I7", "I8")

	cells := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	for i, cell := range cells {
		f.SetCellValue(SheetAnalisisRisiko, fmt.Sprintf("%s9", cell), i+1)
	}

	f.SetCellStyle(SheetAnalisisRisiko, "A7", "I9", style)

}

func (ex *ExcelBuilder) fillAnalisisRisikoData(f *excelize.File, report Report) {
	// Data Start from this row
	startRowNum := SheetAnalisisRisiko_RowStart
	for i, risk := range report.SheetAnalisisRisiko.Risks {

		NoCell := fmt.Sprintf("A%d", startRowNum)

		f.SetCellValue(SheetAnalisisRisiko, NoCell, i+1)
		f.SetColWidth(SheetAnalisisRisiko, "A", "A", 5)

		// Risk Residual Cell
		RiskResidualCell := fmt.Sprintf("B%d", startRowNum)
		f.SetCellValue(SheetAnalisisRisiko, RiskResidualCell, risk.SisaRisiko)

		// Kemungkinan Uraian Cell
		KemungkinanUraianCell := fmt.Sprintf("C%d", startRowNum)
		f.SetCellValue(SheetAnalisisRisiko, KemungkinanUraianCell, risk.KemungkinanUraian)

		// Kemungkinan Nilai Cell
		KemungkinanNilaiCell := fmt.Sprintf("D%d", startRowNum)
		f.SetCellValue(SheetAnalisisRisiko, KemungkinanNilaiCell, risk.KemungkinanNilai)

		// Alasan Cell
		AlasanCell := fmt.Sprintf("E%d", startRowNum)
		f.SetCellValue(SheetAnalisisRisiko, AlasanCell, risk.Alasan)

		// Dampak Uraian Peta Risiko Cell
		DampakUraianPetaRisikoCell := fmt.Sprintf("F%d", startRowNum)
		f.SetCellValue(SheetAnalisisRisiko, DampakUraianPetaRisikoCell, risk.DampakUraianPetaRisiko)

		// Nilai Peta Risiko Cell
		NilaiPetaRisikoCell := fmt.Sprintf("G%d", startRowNum)
		f.SetCellValue(SheetAnalisisRisiko, NilaiPetaRisikoCell, risk.NilaiPetaRisiko)

		// Tingkat Risiko Cell

		riskLevelValue := ex.riskLevel(risk.KemungkinanNilai, risk.NilaiPetaRisiko)

		TingkatRisikoCell := fmt.Sprintf("H%d", startRowNum)
		f.SetCellValue(SheetAnalisisRisiko, TingkatRisikoCell, riskLevelValue.Value)

		ProfilRisikoCell := fmt.Sprintf("I%d", startRowNum)
		f.SetCellValue(SheetAnalisisRisiko, ProfilRisikoCell, "")
		style, err := f.NewStyle(&excelize.Style{
			Fill: excelize.Fill{
				Type: "gradient", Color: []string{riskLevelValue.Color, riskLevelValue.Color}, Shading: 1,
			},
		})
		if err != nil {
			log.Printf("Error creating new sheet analisisRisiko: %v", err)
		}

		f.SetCellStyle(SheetAnalisisRisiko, ProfilRisikoCell, ProfilRisikoCell, style)

		f.SetColWidth(SheetAnalisisRisiko, "B", "I", 45)

		startRowNum++
	}
}

func (ex *ExcelBuilder) riskLevel(estimatedRiskValue int, riskMapValue int) RiskLevel {
	result := estimatedRiskValue * riskMapValue
	if result >= 20 && result <= 25 {
		return RiskLevel{
			Value: result,
			Color: "FF0000",
		}
	} else if result >= 16 && result <= 19 {
		return RiskLevel{
			Value: result,
			Color: "f78c00",
		}
	} else if result >= 12 && result <= 15 {
		return RiskLevel{
			Value: result,
			Color: "eff700",
		}
	} else if result >= 6 && result <= 11 {
		return RiskLevel{
			Value: result,
			Color: "0000f7",
		}
	} else if result >= 1 && result <= 5 {
		return RiskLevel{
			Value: result,
			Color: "49b1ff",
		}
	}

	return RiskLevel{
		Value: result,
		Color: "FFFFFF",
	}
}

package builder

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func (ex *ExcelBuilder) getDataPenangananRisiko(year int) (DataPenangananRisiko, error) {
	var data DataPenangananRisiko
	risks, err := ex.repository.GetRiskTreatmentByYear(year)
	if err != nil {
		return DataPenangananRisiko{}, fmt.Errorf("get risk evaluation for year %d: %v", year, err)
	}

	data = DataPenangananRisiko{
		Risks: risks,
	}

	return data, nil
}

func (ex *ExcelBuilder) createPenangananRisikoHeader(f *excelize.File, period int) {
	f.SetCellValue(SheetPenangananRisiko, SheetHeader_PenangananRisiko_valueRangeStart, SheetHeader_PenangananRisiko)
	f.MergeCell(SheetPenangananRisiko, SheetHeader_PenangananRisiko_valueRangeStart, SheetHeader_PenangananRisiko_ValueRangeEnd)

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

	f.SetCellStyle(SheetPenangananRisiko, "A2", "J2", sheetHeaderStyle)
	f.SetCellStyle(SheetPenangananRisiko, "A4", "C5", style)
	f.SetCellValue(SheetPenangananRisiko, "A4", "Unit Pemilik Risiko")
	f.SetCellValue(SheetPenangananRisiko, "C4", ": BADAN PEMBINAAN HUKUM NASIONAL")
	f.SetCellValue(SheetPenangananRisiko, "A5", "Periode Penerapan")
	f.SetCellValue(SheetPenangananRisiko, "C5", fmt.Sprintf(": %d", period))

}

func (ex *ExcelBuilder) createPenangananRisikoTable(f *excelize.File) {

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

	f.SetCellValue(SheetPenangananRisiko, "A7", "No.")
	f.MergeCell(SheetPenangananRisiko, "A7", "A8")

	f.SetCellValue(SheetPenangananRisiko, "B7", "Indikator Risiko")
	f.MergeCell(SheetPenangananRisiko, "B7", "C7")

	f.SetCellValue(SheetPenangananRisiko, "B8", "Indikasi")
	f.SetCellValue(SheetPenangananRisiko, "C8", "Batas Aman")

	f.SetCellValue(SheetPenangananRisiko, "D7", "Opsi Penanganan")
	f.MergeCell(SheetPenangananRisiko, "D7", "D8")

	f.SetCellValue(SheetPenangananRisiko, "E7", "Kegiatan Pengendalian")
	f.MergeCell(SheetPenangananRisiko, "E7", "E8")

	f.SetCellValue(SheetPenangananRisiko, "F7", "Indikator Pengendalian")
	f.MergeCell(SheetPenangananRisiko, "F7", "G7")
	f.SetCellValue(SheetPenangananRisiko, "F8", "Output")
	f.SetCellValue(SheetPenangananRisiko, "G8", "Target")

	f.SetCellValue(SheetPenangananRisiko, "H7", "Jadwal")
	f.MergeCell(SheetPenangananRisiko, "H7", "H8")

	f.SetCellValue(SheetPenangananRisiko, "I7", "Penanggung Jawab")
	f.MergeCell(SheetPenangananRisiko, "I7", "I8")
	f.SetCellValue(SheetPenangananRisiko, "J7", "Cadangan Risiko")
	f.MergeCell(SheetPenangananRisiko, "J7", "J8")

	cells := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i, cell := range cells {
		f.SetCellValue(SheetPenangananRisiko, fmt.Sprintf("%s9", cell), i+1)
	}

	f.SetCellStyle(SheetPenangananRisiko, "A7", "J9", style)

}

func (ex *ExcelBuilder) fillPenangananRisikoData(f *excelize.File, report Report) {
	// Data Start from this row
	startRowNum := SheetPenangananRisiko_RowStart
	for i, risk := range report.SheetPenangananRisiko.Risks {

		NoCell := fmt.Sprintf("A%d", startRowNum)

		f.SetCellValue(SheetPenangananRisiko, NoCell, i+1)
		f.SetColWidth(SheetPenangananRisiko, "A", "A", 5)

		// Indikasi Indikator Cell
		IndikasiIndikatorCell := fmt.Sprintf("B%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, IndikasiIndikatorCell, risk.IndikasiIndikator)

		// Batas Aman Cell
		BatasAmanCell := fmt.Sprintf("C%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, BatasAmanCell, risk.BatasAmanIndikator)

		// Opsi Penanganan Cell
		OpsiPenangananCell := fmt.Sprintf("D%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, OpsiPenangananCell, risk.OpsiPenanganan)

		// Kegiatan Pengendalian Cell
		KegiatanPendendalianCell := fmt.Sprintf("E%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, KegiatanPendendalianCell, risk.IndikasiIndikator)

		OutputIndikatorCell := fmt.Sprintf("F%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, OutputIndikatorCell, risk.OutputIndikator)

		TargetIndikatorCell := fmt.Sprintf("G%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, TargetIndikatorCell, risk.TargetIndikator)

		JadwalCell := fmt.Sprintf("H%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, JadwalCell, risk.Jadwal)

		PenanggungJawabCell := fmt.Sprintf("I%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, PenanggungJawabCell, risk.PenanggungJawab)

		CadanganRisikoCell := fmt.Sprintf("J%d", startRowNum)
		f.SetCellValue(SheetPenangananRisiko, CadanganRisikoCell, risk.PenanggungJawab)

		f.SetColWidth(SheetPenangananRisiko, "B", "J", 30)

		startRowNum++
	}
	ex.signPlaceholder(f, SheetPenangananRisiko, startRowNum+3, "G")
}

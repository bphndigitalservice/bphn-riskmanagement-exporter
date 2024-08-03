package builder

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func (ex *ExcelBuilder) getDataMonitoringRisiko(year int) (DataPemantauanRisiko, error) {
	var data DataPemantauanRisiko
	risks, err := ex.repository.GetRiskAnalysisByYear(year)
	if err != nil {
		return DataPemantauanRisiko{}, fmt.Errorf("get risk monitoring for year %d: %v", year, err)
	}

	data = DataPemantauanRisiko{
		Risks: risks,
	}

	return data, nil
}

func (ex *ExcelBuilder) createPemantauanRisikoHeader(f *excelize.File, period int) {
	f.SetCellValue(SheetPemantauanRisiko, SheetHeader_PemantaunRisiko_valueRangeStart, SheetHeader_PemantaunRisiko)
	f.MergeCell(SheetPemantauanRisiko, SheetHeader_PemantaunRisiko_valueRangeStart, SheetHeader_PemantaunRisiko_ValueRangeEnd)

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

	f.SetCellStyle(SheetPemantauanRisiko, "A2", "H2", sheetHeaderStyle)
	f.SetCellStyle(SheetHeader_PemantaunRisiko, "A4", "C5", style)
	f.SetCellValue(SheetPemantauanRisiko, "A4", "Unit Pemilik Risiko")
	f.SetCellValue(SheetPemantauanRisiko, "C4", ": BADAN PEMBINAAN HUKUM NASIONAL")
	f.SetCellValue(SheetPemantauanRisiko, "A5", "Periode Penerapan")
	f.SetCellValue(SheetPemantauanRisiko, "C5", fmt.Sprintf(": %d", period))

}

func (ex *ExcelBuilder) createPemantauanRisikoTable(f *excelize.File) {

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

	f.SetCellValue(SheetPemantauanRisiko, "A7", "No.")
	f.MergeCell(SheetPemantauanRisiko, "A7", "A8")

	f.SetCellValue(SheetPemantauanRisiko, "B7", "Kegiatan Pengendalian")
	f.MergeCell(SheetPemantauanRisiko, "B7", "B8")

	f.SetCellValue(SheetPemantauanRisiko, "C7", "Indikator Pengendalian")
	f.MergeCell(SheetPemantauanRisiko, "C7", "F7")
	f.SetCellValue(SheetPemantauanRisiko, "C8", "Output")
	f.SetCellValue(SheetPemantauanRisiko, "D8", "Target")
	f.SetCellValue(SheetPemantauanRisiko, "E8", "Realisasi")
	f.SetCellValue(SheetPemantauanRisiko, "F8", "%")

	f.SetCellValue(SheetPemantauanRisiko, "G7", "Indikator Risiko")
	f.MergeCell(SheetPemantauanRisiko, "G7", "J7")
	f.SetCellValue(SheetPemantauanRisiko, "G8", "Risiko")
	f.SetCellValue(SheetPemantauanRisiko, "H8", "Batas Aman")
	f.SetCellValue(SheetPemantauanRisiko, "I8", "Realisasi")
	f.SetCellValue(SheetPemantauanRisiko, "J8", "%")

	f.SetCellValue(SheetPemantauanRisiko, "K7", "Residu Risiko")
	f.MergeCell(SheetPemantauanRisiko, "K7", "K8")

	f.SetCellValue(SheetPemantauanRisiko, "L7", "Progres sampai dengan semester <replace with your semester>")
	f.MergeCell(SheetPemantauanRisiko, "L7", "L8")

	cells := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
	for i, cell := range cells {
		if cell == "F" {
			f.SetCellValue(SheetPemantauanRisiko, fmt.Sprintf("%s9", cell), "6=(5/4)x100")
			continue
		}
		if cell == "J" {
			f.SetCellValue(SheetPemantauanRisiko, fmt.Sprintf("%s9", cell), "10=(9/8)x100")
			continue
		}
		f.SetCellValue(SheetPemantauanRisiko, fmt.Sprintf("%s9", cell), i+1)
	}

	f.SetCellStyle(SheetPemantauanRisiko, "A7", "L9", style)

}

func (ex *ExcelBuilder) fillPemantauanRisikoData(f *excelize.File, report Report) {
	// Data Start from this row
	startRowNum := SheetPemantauanRisiko_RowStart
	for i, risk := range report.SheetPemantauanRisiko.Risks {

		NoCell := fmt.Sprintf("A%d", startRowNum)

		f.SetCellValue(SheetPemantauanRisiko, NoCell, i+1)
		f.SetColWidth(SheetPemantauanRisiko, "A", "A", 5)

		// Indikasi Indikator Cell
		IndikasiIndikatorCell := fmt.Sprintf("B%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, IndikasiIndikatorCell, risk.IndikasiIndikator)

		// Batas Aman Cell
		BatasAmanCell := fmt.Sprintf("C%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, BatasAmanCell, risk.BatasAmanIndikator)

		// Opsi Penanganan Cell
		OpsiPenangananCell := fmt.Sprintf("D%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, OpsiPenangananCell, risk.OpsiPenanganan)

		// Kegiatan Pengendalian Cell
		KegiatanPendendalianCell := fmt.Sprintf("E%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, KegiatanPendendalianCell, risk.IndikasiIndikator)

		OutputIndikatorCell := fmt.Sprintf("F%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, OutputIndikatorCell, risk.OutputIndikator)

		TargetIndikatorCell := fmt.Sprintf("G%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, TargetIndikatorCell, risk.TargetIndikator)

		JadwalCell := fmt.Sprintf("H%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, JadwalCell, risk.Jadwal)

		PenanggungJawabCell := fmt.Sprintf("I%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, PenanggungJawabCell, risk.PenanggungJawab)

		CadanganRisikoCell := fmt.Sprintf("J%d", startRowNum)
		f.SetCellValue(SheetPemantauanRisiko, CadanganRisikoCell, risk.PenanggungJawab)

		f.SetColWidth(SheetPemantauanRisiko, "B", "L", 30)

		startRowNum++
	}
	ex.signPlaceholder(f, SheetPemantauanRisiko, startRowNum+3, "H")
}

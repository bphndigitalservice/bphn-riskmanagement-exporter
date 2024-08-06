package builder

import (
	"bphn.go.id/mr-report/report/repository"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func (ex *ExcelBuilder) getDataIdentifikasiRisiko(year int) ([]DataIdentifikasiRisiko, error) {
	var data []DataIdentifikasiRisiko
	indicators, err := ex.repository.GetIndicatorsByYear(year)
	if err != nil {
		return nil, fmt.Errorf("get indicators by year %d: %w", year, err)
	}

	for _, indicator := range indicators {
		identifikasiRisiko := DataIdentifikasiRisiko{
			Indicator: indicator,
			Problems:  map[uint64][]repository.Problem{},
			Risks:     map[uint64][]repository.Risk{},
		}

		problems, pErr := ex.repository.GetProblemsByIndicator(indicator.ID)
		if pErr != nil {
			return nil, fmt.Errorf("get problem ids by indicator %d: %w", indicator.ID, pErr)
		}

		identifikasiRisiko.Problems[indicator.ID] = problems

		for _, problem := range problems {
			risks, err := ex.repository.GetRisksByProblem(problem.ID)
			if err != nil {
				return nil, fmt.Errorf("get risks by year %d: %w", year, err)
			}
			identifikasiRisiko.Risks[problem.ID] = risks
		}

		data = append(data, identifikasiRisiko)
	}

	return data, nil
}

func (ex *ExcelBuilder) createIdentifikasiRisikoHeader(f *excelize.File, period int) {
	f.SetCellValue(SheetIdentifikasirisiko, SheetHeader_IdentifikasiRisiko_valueRangeStart, SheetHeader_IdentifikasiRisiko)
	f.MergeCell(SheetIdentifikasirisiko, SheetHeader_IdentifikasiRisiko_valueRangeStart, SheetHeader_IdentifikasiRisiko_ValueRangeEnd)

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

	f.SetCellStyle(SheetIdentifikasirisiko, "A2", "M2", sheetHeaderStyle)
	f.SetCellStyle(SheetIdentifikasirisiko, "A4", "C5", style)
	f.SetCellValue(SheetIdentifikasirisiko, "A4", "Unit Pemilik Risiko")
	f.SetCellValue(SheetIdentifikasirisiko, "C4", ": BADAN PEMBINAAN HUKUM NASIONAL")
	f.SetCellValue(SheetIdentifikasirisiko, "A5", "Periode Penerapan")
	f.SetCellValue(SheetIdentifikasirisiko, "C5", fmt.Sprintf(": %d", period))
}

func (ex *ExcelBuilder) createIdentifikasiRisikoTable(f *excelize.File) {

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

	f.SetCellValue(SheetIdentifikasirisiko, "A7", "No.")
	f.MergeCell(SheetIdentifikasirisiko, "A7", "A8")

	f.SetCellValue(SheetIdentifikasirisiko, "B7", "Indikator Kinerja")
	f.MergeCell(SheetIdentifikasirisiko, "B7", "B8")

	f.SetCellValue(SheetIdentifikasirisiko, "C7", "Permasalahan")
	f.MergeCell(SheetIdentifikasirisiko, "C7", "C8")

	// Risiko
	f.SetCellValue(SheetIdentifikasirisiko, "D7", "Risiko")
	f.MergeCell(SheetIdentifikasirisiko, "D7", "E7")
	f.SetCellValue(SheetIdentifikasirisiko, "D8", "Pernyataan")
	f.SetCellValue(SheetIdentifikasirisiko, "E8", "Pemilik")

	// Penyebab
	f.SetCellValue(SheetIdentifikasirisiko, "F7", "Penyebab")
	f.MergeCell(SheetIdentifikasirisiko, "F7", "H7")
	f.SetCellValue(SheetIdentifikasirisiko, "F8", "Uraian")
	f.SetCellValue(SheetIdentifikasirisiko, "G8", "Sumber")
	f.SetCellValue(SheetIdentifikasirisiko, "H8", "C/UC")

	// Dampak
	f.SetCellValue(SheetIdentifikasirisiko, "I7", "Dampak")
	f.MergeCell(SheetIdentifikasirisiko, "I7", "J7")
	f.SetCellValue(SheetIdentifikasirisiko, "I8", "Uraian")
	f.SetCellValue(SheetIdentifikasirisiko, "J8", "Pihak yang Terkena")

	// Pengendalian Intern yang Ada
	f.SetCellValue(SheetIdentifikasirisiko, "K7", "Pengendalian Intern yang Ada")
	f.MergeCell(SheetIdentifikasirisiko, "K7", "K8")

	// Sisa Risiko
	f.SetCellValue(SheetIdentifikasirisiko, "L7", "Sisa Risiko")
	f.MergeCell(SheetIdentifikasirisiko, "L7", "L8")

	// Kriteria Risiko
	f.SetCellValue(SheetIdentifikasirisiko, "M7", "Kriteria Risiko")
	f.MergeCell(SheetIdentifikasirisiko, "M7", "M8")

	cells := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"}
	for i, cell := range cells {
		f.SetCellValue(SheetIdentifikasirisiko, fmt.Sprintf("%s9", cell), i+1)
	}

	f.SetCellStyle(SheetIdentifikasirisiko, "A7", "M9", style)

}

func (ex *ExcelBuilder) fillIdentifikasiRisikoData(f *excelize.File, report Report) {
	// Data Start from this row
	startRowNum := SheetIdentifikasiRisiko_RowStart

	for i, data := range report.SheetIdentifikasiRisiko {
		// Number cell
		NoCell := fmt.Sprintf("A%d", startRowNum)

		f.SetCellValue(SheetIdentifikasirisiko, NoCell, i+1)
		f.SetColWidth(SheetIdentifikasirisiko, "A", "A", 5)

		// Indikator cell
		IndicatorCell := fmt.Sprintf("B%d", startRowNum)
		f.SetCellValue(SheetIdentifikasirisiko, IndicatorCell, data.Indicator.IndicatorDefinition)
		f.SetColWidth(SheetIdentifikasirisiko, "B", "M", 45)

		ProblemStartRowNum := startRowNum
		for _, problem := range data.Problems[data.Indicator.ID] {
			// Problem Cell
			ProblemCell := fmt.Sprintf("C%d", ProblemStartRowNum)
			f.SetCellValue(SheetIdentifikasirisiko, ProblemCell, problem.ProblemDefinition)

			RiskStartRowNum := ProblemStartRowNum
			for _, risk := range data.Risks[problem.ID] {
				// Risk Statement Cell
				RiskStatementCell := fmt.Sprintf("D%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, RiskStatementCell, risk.RisikoPernyataan)

				// Risk Owner Cell
				RiskOwnerCell := fmt.Sprintf("E%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, RiskOwnerCell, risk.Owner)

				// Risk Detail Cell
				RiskDetailCell := fmt.Sprintf("F%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, RiskDetailCell, risk.PenyebabUraian)

				// Risk Source Cell
				RiskCauseSourceCell := fmt.Sprintf("G%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, RiskCauseSourceCell, risk.PenyebabSumber)

				// Risk C/UC Cell
				RiskCUCCell := fmt.Sprintf("H%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, RiskCUCCell, risk.PenyebabCUC)

				// Dampak Uraian Daftar Risiko
				ElaboratedImpactListCell := fmt.Sprintf("I%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, ElaboratedImpactListCell, risk.DampakUraianDaftarRisiko)

				// Dampak Uraian Pihak Daftar Risiko
				ElaboratedImpactStakeholderCell := fmt.Sprintf("J%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, ElaboratedImpactStakeholderCell, risk.DampakPihakDaftarRisiko)

				// Internal Control
				InternalControllCell := fmt.Sprintf("K%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, InternalControllCell, risk.PengendalianIntern)

				// Risk Residual
				RiskResidualCell := fmt.Sprintf("L%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, RiskResidualCell, risk.SisaRisiko)

				// Risk Criteria
				RiskCriteriaCell := fmt.Sprintf("M%d", RiskStartRowNum)
				f.SetCellValue(SheetIdentifikasirisiko, RiskCriteriaCell, risk.SisaRisiko)

				RiskStartRowNum++
			}

			if len(data.Risks[problem.ID]) > 0 {
				ProblemStartRowNum = RiskStartRowNum
			} else {
				ProblemStartRowNum++
			}
		}

		startRowNum = ProblemStartRowNum
	}

	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "top",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// numbering style
	numberCellStyle, _ := f.NewStyle(&excelize.Style{Border: []excelize.Border{
		{Type: "left", Color: "000000", Style: 1},
		{Type: "top", Color: "000000", Style: 1},
		{Type: "bottom", Color: "000000", Style: 1},
		{Type: "right", Color: "000000", Style: 1},
	}})
	f.SetCellStyle(SheetIdentifikasirisiko, fmt.Sprintf("A%d", SheetIdentifikasiRisiko_RowStart), fmt.Sprintf("A%d", startRowNum), numberCellStyle)

	// text styling
	topCell := fmt.Sprintf("B%d", SheetIdentifikasiRisiko_RowStart)
	bottomCell := fmt.Sprintf("M%d", startRowNum)
	err = f.SetCellStyle(SheetIdentifikasirisiko, topCell, bottomCell, style)
	if err != nil {
		log.Fatal(err)
	}

	ex.signPlaceholder(f, SheetIdentifikasirisiko, startRowNum+3, "K")
}

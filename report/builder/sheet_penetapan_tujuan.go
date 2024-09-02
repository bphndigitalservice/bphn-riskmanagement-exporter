package builder

import (
	"bphn.go.id/mr-report/report/repository"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func (ex *ExcelBuilder) getDataPenetapanTujuan(year int) ([]DataPenetapanTujuan, error) {
	strategies, sErr := ex.repository.GetStrategiesByYear(year)

	if sErr != nil {
		return nil, fmt.Errorf("Error fetching strategies data")
	}

	var data []DataPenetapanTujuan

	for _, strategy := range strategies {

		d := DataPenetapanTujuan{
			Strategy:   strategy,
			Objectives: map[uint64][]repository.Objective{},
			Indicators: map[uint64][]repository.Indicator{},
			Problems:   map[uint64][]repository.Problem{},
		}

		objectives, oErr := ex.repository.GetObjectivesByStrategy(strategy.ID)
		if oErr != nil {
			return nil, fmt.Errorf("Error fetching objectives data")
		}

		d.Objectives[strategy.ID] = objectives

		for _, objective := range objectives {
			indicators, iErr := ex.repository.GetIndicatorsByObjective(objective.ID)
			if iErr != nil {
				return nil, fmt.Errorf("Error fetching indicators data")
			}
			d.Indicators[objective.ID] = indicators

			for _, indicator := range indicators {
				problems, pErr := ex.repository.GetProblemsByIndicator(indicator.ID)
				if pErr != nil {
					return nil, fmt.Errorf("Error fetching problems data")
				}
				d.Problems[indicator.ID] = problems
			}

		}

		data = append(data, d)
	}

	return data, nil
}

func (ex *ExcelBuilder) createPenetapanTujuanHeader(f *excelize.File, period int) {
	f.SetCellValue(SheetPenetapantujuan, SheetHeaderPenetapanTujuan_ValueRangeStart, SheetHeaderPenetapanTujuan)
	f.MergeCell(SheetPenetapantujuan, SheetHeaderPenetapanTujuan_ValueRangeStart, SheetHeaderPenetapanTujuan_ValueRangeEnd)

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
			Horizontal: "center",
		},
	})

	f.SetCellStyle(SheetPenetapantujuan, "A2", "L2", sheetHeaderStyle)
	f.SetCellStyle(SheetPenetapantujuan, "A4", "C5", style)
	f.SetCellValue(SheetPenetapantujuan, "A4", "Unit Pemilik Risiko")
	f.SetCellValue(SheetPenetapantujuan, "C4", ": BADAN PEMBINAAN HUKUM NASIONAL")
	f.SetCellValue(SheetPenetapantujuan, "A5", "Periode Penerapan")
	f.SetCellValue(SheetPenetapantujuan, "C5", fmt.Sprintf(": %d", period))

}

func (ex *ExcelBuilder) createPenetapanTujuanTable(f *excelize.File) {

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

	f.SetCellValue(SheetPenetapantujuan, "A7", "No.")
	f.SetCellValue(SheetPenetapantujuan, "B7", "Strategi/Program/Kegiatan")
	f.SetCellValue(SheetPenetapantujuan, "C7", "Sasaran")
	f.SetCellValue(SheetPenetapantujuan, "D7", "Indikator Kinerja")
	f.SetCellValue(SheetPenetapantujuan, "E7", "Permasalahan")
	f.SetCellValue(SheetPenetapantujuan, "F7", "UPR")

	cells := []string{"A", "B", "C", "D", "E", "F"}
	for i, cell := range cells {
		f.SetCellValue(SheetPenetapantujuan, fmt.Sprintf("%s8", cell), i+1)
	}

	f.SetCellStyle(SheetPenetapantujuan, "A7", "F8", style)

}

func (ex *ExcelBuilder) fillPenetapanTujuanData(f *excelize.File, report Report) {

	// Data Start from this row
	startRowNum := SheetPenetapanTujuan_RowStart

	for i, data := range report.SheetPenetapanTujuan {
		// Number cell
		NoCell := fmt.Sprintf("A%d", startRowNum)

		f.SetCellValue(SheetPenetapantujuan, NoCell, i+1)
		f.SetColWidth(SheetPenetapantujuan, "A", "A", 5)

		// Strategy Definition cell
		StrategyCell := fmt.Sprintf("B%d", startRowNum)
		f.SetCellValue(SheetPenetapantujuan, StrategyCell, data.Strategy.Definition)
		f.SetColWidth(SheetPenetapantujuan, "B", "E", 45)

		// Objective Cells
		ObjectiveCellRowNum := startRowNum
		for _, objective := range data.Objectives[data.Strategy.ID] {
			ObjectiveCell := fmt.Sprintf("C%d", ObjectiveCellRowNum)
			f.SetCellValue(SheetPenetapantujuan, ObjectiveCell, objective.ObjectiveDefinition)
			//ObjectiveCellRowNum++

			IndicatorCellRowNum := ObjectiveCellRowNum
			indicatorCount := 0

			for _, indicator := range data.Indicators[objective.ID] {
				IndicatorCell := fmt.Sprintf("D%d", IndicatorCellRowNum)
				f.SetCellValue(SheetPenetapantujuan, IndicatorCell, indicator.IndicatorDefinition)

				ProblemsCellRowNum := IndicatorCellRowNum
				problemsCount := 0

				for _, problem := range data.Problems[indicator.ID] {
					ProblemCell := fmt.Sprintf("E%d", ProblemsCellRowNum)
					f.SetCellValue(SheetPenetapantujuan, ProblemCell, problem.ProblemDefinition)
					UPRCell := fmt.Sprintf("F%d", ProblemsCellRowNum)
					f.SetCellValue(SheetPenetapantujuan, UPRCell, problem.OwnerNickname)
					ProblemsCellRowNum++
					problemsCount++

				}

				if problemsCount > 0 {
					IndicatorCellRowNum = ProblemsCellRowNum
					indicatorCount = problemsCount
				} else {
					IndicatorCellRowNum++
					indicatorCount++
				}
			}

			if indicatorCount > 0 {
				ObjectiveCellRowNum = IndicatorCellRowNum
			} else {
				ObjectiveCellRowNum++
			}

		}

		if len(data.Objectives[data.Strategy.ID]) > 0 {
			startRowNum = ObjectiveCellRowNum
		} else {
			startRowNum++
		}

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
	f.SetCellStyle(SheetPenetapantujuan, fmt.Sprintf("A%d", SheetPenetapanTujuan_RowStart), fmt.Sprintf("A%d", startRowNum), numberCellStyle)

	// text styling
	topCell := fmt.Sprintf("B%d", SheetPenetapanTujuan_RowStart)
	bottomCell := fmt.Sprintf("F%d", startRowNum)
	err = f.SetCellStyle(SheetPenetapantujuan, topCell, bottomCell, style)
	if err != nil {
		log.Fatal(err)
	}

	ex.signPlaceholder(f, SheetPenetapantujuan, startRowNum+3, "E")
}

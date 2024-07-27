package builder

import (
	"bphn.go.id/mr-report/report/repository"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

var sheets []string = []string{
	SheetIdentifikasirisiko,
	SheetKriteriadanskala,
	SheetAnalisisrisiko,
	SheetEvaluasirisiko,
	SheetPenangananrisiko,
	SheetPemantauanrisiko,
}

type ExcelBuilder struct {
	repository repository.RiskRepository
}

func NewExcelBuilder(repository repository.RiskRepository) *ExcelBuilder {
	return &ExcelBuilder{
		repository: repository,
	}
}

func (ex *ExcelBuilder) Export(year int) (*excelize.File, error) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	data, err := ex.getDataPenetapanTujuan(year)
	if err != nil {
		log.Printf("tai %v", err.Error())
		return nil, err
	}

	builder := Report{
		SheetPenetapanTujuan: data,
		Period:               uint64(year),
	}

	ex.prepareSheets(f)

	// 1. Penetapan Tujuan
	ex.setPenetapanTujuanHeader(f, year)
	ex.setPenetapanTujuanTable(f)
	ex.fillPenetapanTujuan(f, builder)

	// Save spreadsheet by the given path.
	if err := f.SaveAs("Manajemen Risiko_Export.xlsx"); err != nil {
		fmt.Println(err)
		f.Close()
		return nil, err
	}

	return f, nil
}

func (ex *ExcelBuilder) prepareSheets(f *excelize.File) {

	err := f.SetSheetName("Sheet1", SheetPenetapantujuan)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sheet := range sheets {
		_, err := f.NewSheet(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

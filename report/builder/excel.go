package builder

import (
	"bphn.go.id/mr-report/report/repository"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

var sheets []string = []string{
	SheetIdentifikasirisiko,
	SheetKriteriaDanSkala,
	SheetAnalisisRisiko,
	SheetEvaluasiRisiko,
	SheetPenangananRisiko,
	SheetPemantauanRisiko,
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

	dataPenetapanTujuan, err := ex.getDataPenetapanTujuan(year)
	if err != nil {
		log.Printf("Error %v", err.Error())
		return nil, err
	}

	dataIdentifikasiRisiko, err := ex.getDataIdentifikasiRisiko(year)
	if err != nil {
		log.Printf("Error %v", err.Error())
		return nil, err
	}

	dataAnalisisRisiko, err := ex.getDataAnalisisRisiko(year)
	if err != nil {
		log.Printf("Error %v", err.Error())
	}

	dataEvaluasiRisiko, err := ex.getDataEvaluasiRisiko(year)
	if err != nil {
		log.Printf("Error %v", err.Error())
	}

	dataPenangananRisiko, err := ex.getDataPenangananRisiko(year)
	if err != nil {
		log.Printf("Error %v", err.Error())
	}

	dataPemantauanRisiko, err := ex.getDataMonitoringRisiko(year)
	if err != nil {
		log.Printf("Error %v", err.Error())
	}

	builder := Report{
		SheetPenetapanTujuan:    dataPenetapanTujuan,
		SheetIdentifikasiRisiko: dataIdentifikasiRisiko,
		SheetAnalisisRisiko:     dataAnalisisRisiko,
		SheetEvaluasiRisiko:     dataEvaluasiRisiko,
		SheetPenangananRisiko:   dataPenangananRisiko,
		SheetPemantauanRisiko:   dataPemantauanRisiko,
		Period:                  uint64(year),
	}

	ex.prepareSheets(f)

	// 1. Penetapan Tujuan
	ex.createPenetapanTujuanHeader(f, year)
	ex.createPenetapanTujuanTable(f)
	ex.fillPenetapanTujuanData(f, builder)

	// 2. Identifikasi Risiko
	ex.createIdentifikasiRisikoHeader(f, year)
	ex.createIdentifikasiRisikoTable(f)
	ex.fillIdentifikasiRisikoData(f, builder)

	// 3. Kriteria dan Skala
	ex.fillKriteriaDanSkala(f)

	// 4. Analisis Risiko
	ex.createAnalisisRisikoHeader(f, year)
	ex.createAnalisisRisikoTable(f)
	ex.fillAnalisisRisikoData(f, builder)

	// 5. Evaluasi Risiko
	ex.createEvaluasiRisikoHeader(f, year)
	ex.createEvaluasiRisikoTable(f)
	ex.fillEvaluasiRisikoData(f, builder)

	// 6. Penanganan Risiko
	ex.createPenangananRisikoHeader(f, year)
	ex.createPenangananRisikoTable(f)
	ex.fillPenangananRisikoData(f, builder)

	// 7. Pemantauan Risiko
	ex.createPemantauanRisikoHeader(f, year)
	ex.createPemantauanRisikoTable(f)
	ex.fillPemantauanRisikoData(f, builder)

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

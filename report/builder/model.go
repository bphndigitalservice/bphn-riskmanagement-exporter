package builder

import "bphn.go.id/mr-report/report/repository"

type Report struct {
	SheetPenetapanTujuan    []DataPenetapanTujuan
	SheetIdentifikasiRisiko []DataIdentifikasiRisiko
	SheetAnalisisRisiko     DataAnalisisRisiko
	Period                  uint64
}

type DataPenetapanTujuan struct {
	Strategy   repository.Strategy
	Objectives map[uint64][]repository.Objective
	Indicators map[uint64][]repository.Indicator
	Problems   map[uint64][]repository.Problem
}

type DataIdentifikasiRisiko struct {
	Indicator repository.Indicator
	Problems  map[uint64][]repository.Problem // Indicator ID as key
	Risks     map[uint64][]repository.Risk
}

type DataAnalisisRisiko struct {
	Risks []repository.Risk
}

type RiskLevel struct {
	Color string
	Value int
	Label string
}

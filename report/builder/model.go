package builder

import "bphn.go.id/mr-report/report/repository"

type Report struct {
	SheetPenetapanTujuan []DataPenetapanTujuan
	Period               uint64
}

type DataPenetapanTujuan struct {
	Strategy   repository.Strategy
	Objectives map[uint64][]repository.Objective
	Indicators map[uint64][]repository.Indicator
	Problems   map[uint64][]repository.Problem
}

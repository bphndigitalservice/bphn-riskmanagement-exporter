package repository

import (
	"database/sql"
)

type RiskRepository interface {
	GetStrategiesByYear(year int) ([]Strategy, error)
	GetObjectivesByStrategy(id uint64) ([]Objective, error)
	GetIndicatorsByObjective(id uint64) ([]Indicator, error)
	GetProblemsByIndicator(id uint64) ([]Problem, error)
	GetObjectivesByYear(year int) ([]Objective, error)
	GetIndicatorsByYear(year int) ([]Indicator, error)
	GetProblemsByYear(year int) ([]Problem, error)
}

type riskRepository struct {
	db *sql.DB
}

func (receiver *riskRepository) GetIndicatorsByObjective(id uint64) ([]Indicator, error) {
	rows, err := receiver.db.Query("SELECT * FROM master_indikator WHERE id_sasaran = ? ORDER BY id_indikator", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var indicators []Indicator

	for rows.Next() {
		var indicator Indicator
		if err := rows.Scan(&indicator.ID, &indicator.ObjectiveID, &indicator.IndicatorDefinition, &indicator.CreatedAt, &indicator.CreatedBy); err != nil {
			return nil, err
		}
		indicators = append(indicators, indicator)
	}
	return indicators, nil
}

func (receiver *riskRepository) GetProblemsByIndicator(id uint64) ([]Problem, error) {
	rows, err := receiver.db.Query("SELECT id_permasalahan,id_indikator,judul_permasalahan,tahun,tgl_post,id_user_post,id_pusat FROM master_permasalahan WHERE id_indikator = ? ORDER BY id_permasalahan", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []Problem

	for rows.Next() {
		var problem Problem
		if err := rows.Scan(&problem.ID, &problem.IndicatorID, &problem.ProblemDefinition, &problem.Year, &problem.CreatedAt, &problem.OwnerID, &problem.OwnerDepartmentID); err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}
	return problems, nil
}

func (receiver *riskRepository) GetStrategiesByYear(year int) ([]Strategy, error) {

	rows, err := receiver.db.Query("SELECT id_strategi,judul_strategi,tahun,tgl_post FROM master_strategi WHERE tahun = ?", year)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var strategies []Strategy

	for rows.Next() {
		var strategy Strategy
		if err := rows.Scan(&strategy.ID, &strategy.Definition, &strategy.Year, &strategy.CreatedAt); err != nil {
			return nil, err
		}
		strategies = append(strategies, strategy)
	}
	return strategies, nil
}

func (receiver *riskRepository) GetObjectivesByStrategy(id uint64) ([]Objective, error) {
	rows, err := receiver.db.Query("SELECT id_sasaran,id_strategi,judul_sasaran,id_user_post,tgl_post FROM master_sasaran WHERE id_strategi = ? ORDER BY id_sasaran;", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var objectives []Objective

	for rows.Next() {
		var objective Objective
		if err := rows.Scan(&objective.ID, &objective.StrategyID, &objective.ObjectiveDefinition, &objective.CreatedBy, &objective.CreatedAt); err != nil {
			return nil, err
		}
		objectives = append(objectives, objective)
	}
	return objectives, nil
}

func (receiver *riskRepository) GetObjectivesByYear(year int) ([]Objective, error) {
	query := "select master_sasaran.* from master_sasaran left join master_strategi on master_strategi.id_strategi = master_sasaran.id_strategi where master_strategi.tahun = ?"
	rows, err := receiver.db.Query(query, year)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var objectives []Objective

	for rows.Next() {
		var objective Objective
		if err := rows.Scan(&objective.ID, &objective.StrategyID, &objective.ObjectiveDefinition, &objective.CreatedBy, &objective.CreatedAt); err != nil {
			return nil, err
		}
		objectives = append(objectives, objective)
	}
	return objectives, nil
}

func (receiver *riskRepository) GetIndicatorsByYear(year int) ([]Indicator, error) {
	query := `select master_indikator.*
from master_indikator
         left join master_sasaran on master_indikator.id_sasaran = master_sasaran.id_sasaran
         left join master_strategi on master_strategi.id_strategi = master_sasaran.id_strategi
where master_strategi.tahun = ?`
	rows, err := receiver.db.Query(query, year)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var indicators []Indicator

	for rows.Next() {
		var indicator Indicator
		if err := rows.Scan(&indicator.ID, &indicator.ObjectiveID, &indicator.IndicatorDefinition, &indicator.CreatedAt, &indicator.CreatedBy); err != nil {
			return nil, err
		}
		indicators = append(indicators, indicator)
	}
	return indicators, nil
}

func (receiver *riskRepository) GetProblemsByYear(year int) ([]Problem, error) {

	query := `select master_permasalahan.*
from master_permasalahan
        left join master_indikator on master_indikator.id_indikator = master_permasalahan.id_indikator
         left join master_sasaran on master_indikator.id_sasaran = master_sasaran.id_sasaran
         left join master_strategi on master_strategi.id_strategi = master_sasaran.id_strategi
where master_strategi.tahun = ?`

	rows, err := receiver.db.Query(query, year)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []Problem

	for rows.Next() {
		var problem Problem
		if err := rows.Scan(&problem.ID, &problem.IndicatorID, &problem.ProblemDefinition, &problem.Year, &problem.CreatedAt, &problem.OwnerID, &problem.OwnerDepartmentID); err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}
	return problems, nil
}

func NewRiskRepository(db *sql.DB) RiskRepository {
	return &riskRepository{
		db: db,
	}
}

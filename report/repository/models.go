package repository

import "time"

type Strategy struct {
	ID         uint64    `json:"id"`
	Definition string    `json:"definition"`
	Year       int       `json:"year"`
	CreatedAt  time.Time `json:"created_at"`
}

type Objective struct {
	ID                  uint64    `json:"id"`
	StrategyID          uint64    `json:"strategy_id"`
	ObjectiveDefinition string    `json:"objective_definition"`
	CreatedAt           time.Time `json:"created_at"`
	CreatedBy           uint64    `json:"created_by"`
}

type Indicator struct {
	ID                  uint64    `json:"id"`
	ObjectiveID         uint64    `json:"objective_id"`
	IndicatorDefinition string    `json:"indicator_definition"`
	CreatedAt           time.Time `json:"created_at"`
	CreatedBy           uint64    `json:"created_by"`
}

type Problem struct {
	ID                uint64    `json:"id"`
	IndicatorID       uint64    `json:"indicator_id"`
	ProblemDefinition string    `json:"problem_definition"`
	Year              int       `json:"year"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         uint64    `json:"created_by"`
	OwnerID           uint64    `json:"owner_id"`
	OwnerDepartmentID uint64    `json:"owner_department_id"`
}

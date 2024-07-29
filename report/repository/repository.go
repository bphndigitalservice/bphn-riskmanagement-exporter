package repository

import (
	"database/sql"
	"log"
)

type RiskRepository interface {
	GetStrategiesByYear(year int) ([]Strategy, error)
	GetObjectivesByStrategy(id uint64) ([]Objective, error)
	GetIndicatorsByObjective(id uint64) ([]Indicator, error)
	GetProblemsByIndicator(id uint64) ([]Problem, error)
	GetObjectivesByYear(year int) ([]Objective, error)
	GetIndicatorsByYear(year int) ([]Indicator, error)
	GetProblemsByYear(year int) ([]Problem, error)
	GetRisksByProblem(problem uint64) ([]Risk, error)
	GetRiskAnalysisByYear(year int) ([]Risk, error)
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
		log.Printf("Error getting strategies by year: %v", err)
		return nil, err
	}
	defer rows.Close()

	var strategies []Strategy

	for rows.Next() {
		var strategy Strategy
		if err := rows.Scan(&strategy.ID, &strategy.Definition, &strategy.Year, &strategy.CreatedAt); err != nil {
			log.Printf("Error scanning row: %s", err)
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

func (receiver *riskRepository) GetRisksByProblem(problem uint64) ([]Risk, error) {
	query := `select manajemen_risiko.*,pusat.nickname as owner
from manajemen_risiko
         left join master_permasalahan on master_permasalahan.id_permasalahan = manajemen_risiko.id_permasalahan
         left join master_indikator on master_indikator.id_indikator = master_permasalahan.id_indikator
         left join master_sasaran on master_indikator.id_sasaran = master_sasaran.id_sasaran
         left join master_strategi on master_strategi.id_strategi = master_sasaran.id_strategi
		 left join pusat on pusat.id_pusat=master_permasalahan.id_pusat
where manajemen_risiko.id_permasalahan = ?`

	rows, err := receiver.db.Query(query, problem)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var risks []Risk

	for rows.Next() {
		var mr Risk
		var (
			idPermasalahan           sql.NullInt32
			tahun                    sql.NullInt32
			penyebabUraian           sql.NullString
			penyebabSumber           sql.NullString
			penyebabCUC              sql.NullString
			dampakUraianDaftarRisiko sql.NullString
			dampakPihakDaftarRisiko  sql.NullString
			pengendalianIntern       sql.NullString
			sisaRisiko               sql.NullString
			kriteriaRisiko           sql.NullString
			kemungkinanUraian        sql.NullString
			kemungkinanNilai         sql.NullInt32
			alasan                   sql.NullString
			dampakUraianPetaRisiko   sql.NullString
			nilaiPetaRisiko          sql.NullInt32
			prioritasRisiko          sql.NullString
			toleransiRisiko          sql.NullInt32
			indikasiIndikator        sql.NullString
			penjelasanIndikator      sql.NullString
			batasAmanIndikator       sql.NullString
			opsiPenanganan           sql.NullString
			kegiatanPengendalian     sql.NullString
			outputIndikator          sql.NullString
			targetIndikator          sql.NullString
			jadwal                   sql.NullString
			penanggungJawab          sql.NullString
			cadanganRisiko           sql.NullString
			realisasiPengendalian    sql.NullString
			realisasiRisiko          sql.NullString
			risikoResidu             sql.NullString
			progress                 sql.NullString
			tglPost                  sql.NullTime
			idUserPost               sql.NullInt32
			Owner                    sql.NullString
		)

		err := rows.Scan(
			&mr.IdRisiko,
			&idPermasalahan,
			&tahun,
			&mr.RisikoPernyataan,
			&penyebabUraian,
			&penyebabSumber,
			&penyebabCUC,
			&dampakUraianDaftarRisiko,
			&dampakPihakDaftarRisiko,
			&pengendalianIntern,
			&sisaRisiko,
			&kriteriaRisiko,
			&kemungkinanUraian,
			&kemungkinanNilai,
			&alasan,
			&dampakUraianPetaRisiko,
			&nilaiPetaRisiko,
			&prioritasRisiko,
			&toleransiRisiko,
			&indikasiIndikator,
			&penjelasanIndikator,
			&batasAmanIndikator,
			&opsiPenanganan,
			&kegiatanPengendalian,
			&outputIndikator,
			&targetIndikator,
			&jadwal,
			&penanggungJawab,
			&cadanganRisiko,
			&realisasiPengendalian,
			&realisasiRisiko,
			&risikoResidu,
			&progress,
			&tglPost,
			&idUserPost,
			&Owner,
		)
		if err != nil {
			log.Fatal(err)
		}

		// Convert nullable fields to pointers
		if idPermasalahan.Valid {
			mr.IdPermasalahan = int(idPermasalahan.Int32)
		}
		if tahun.Valid {
			mr.Tahun = int(tahun.Int32)
		}
		if penyebabUraian.Valid {
			mr.PenyebabUraian = penyebabUraian.String
		}
		if penyebabSumber.Valid {
			mr.PenyebabSumber = penyebabSumber.String
		}
		if penyebabCUC.Valid {
			mr.PenyebabCUC = penyebabCUC.String
		}
		if dampakUraianDaftarRisiko.Valid {
			mr.DampakUraianDaftarRisiko = dampakUraianDaftarRisiko.String
		}
		if dampakPihakDaftarRisiko.Valid {
			mr.DampakPihakDaftarRisiko = dampakPihakDaftarRisiko.String
		}
		if pengendalianIntern.Valid {
			mr.PengendalianIntern = pengendalianIntern.String
		}
		if sisaRisiko.Valid {
			mr.SisaRisiko = sisaRisiko.String
		}
		if kriteriaRisiko.Valid {
			mr.KriteriaRisiko = kriteriaRisiko.String
		}
		if kemungkinanUraian.Valid {
			mr.KemungkinanUraian = kemungkinanUraian.String
		}
		if kemungkinanNilai.Valid {
			mr.KemungkinanNilai = int(kemungkinanNilai.Int32)
		}
		if alasan.Valid {
			mr.Alasan = alasan.String
		}
		if dampakUraianPetaRisiko.Valid {
			mr.DampakUraianPetaRisiko = dampakUraianPetaRisiko.String
		}
		if nilaiPetaRisiko.Valid {
			mr.NilaiPetaRisiko = int(nilaiPetaRisiko.Int32)
		}
		if prioritasRisiko.Valid {
			mr.PrioritasRisiko = prioritasRisiko.String
		}
		if toleransiRisiko.Valid {
			mr.ToleransiRisiko = int(toleransiRisiko.Int32)
		}
		if indikasiIndikator.Valid {
			mr.IndikasiIndikator = indikasiIndikator.String
		}
		if penjelasanIndikator.Valid {
			mr.PenjelasanIndikator = penjelasanIndikator.String
		}
		if batasAmanIndikator.Valid {
			mr.BatasAmanIndikator = batasAmanIndikator.String
		}
		if opsiPenanganan.Valid {
			mr.OpsiPenanganan = opsiPenanganan.String
		}
		if kegiatanPengendalian.Valid {
			mr.KegiatanPengendalian = kegiatanPengendalian.String
		}
		if outputIndikator.Valid {
			mr.OutputIndikator = outputIndikator.String
		}
		if targetIndikator.Valid {
			mr.TargetIndikator = targetIndikator.String
		}
		if jadwal.Valid {
			mr.Jadwal = jadwal.String
		}
		if penanggungJawab.Valid {
			mr.PenanggungJawab = penanggungJawab.String
		}
		if cadanganRisiko.Valid {
			mr.CadanganRisiko = cadanganRisiko.String
		}
		if realisasiPengendalian.Valid {
			mr.RealisasiPengendalian = realisasiPengendalian.String
		}
		if realisasiRisiko.Valid {
			mr.RealisasiRisiko = realisasiRisiko.String
		}
		if risikoResidu.Valid {
			mr.RisikoResidu = risikoResidu.String
		}
		if progress.Valid {
			mr.Progress = progress.String
		}
		if tglPost.Valid {
			mr.TglPost = tglPost.Time
		}
		if idUserPost.Valid {
			mr.IdUserPost = int(idUserPost.Int32)
		}
		if Owner.Valid {
			mr.Owner = Owner.String
		}

		risks = append(risks, mr)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return risks, nil
}

func (receiver *riskRepository) GetRiskAnalysisByYear(year int) ([]Risk, error) {
	query := `select manajemen_risiko.* from manajemen_risiko
         left join master_permasalahan on master_permasalahan.id_permasalahan = manajemen_risiko.id_permasalahan
         left join master_indikator on master_indikator.id_indikator = master_permasalahan.id_indikator
         left join master_sasaran on master_indikator.id_sasaran = master_sasaran.id_sasaran
         left join master_strategi on master_strategi.id_strategi = master_sasaran.id_strategi
         where master_strategi.tahun = ? and manajemen_risiko.sisa_risiko is not null ORDER BY master_strategi.id_strategi,master_sasaran.id_sasaran,master_indikator.id_indikator,master_permasalahan.id_permasalahan, manajemen_risiko.id_risiko`

	rows, err := receiver.db.Query(query, year)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var risks []Risk

	for rows.Next() {
		var mr Risk
		var (
			idPermasalahan           sql.NullInt32
			tahun                    sql.NullInt32
			penyebabUraian           sql.NullString
			penyebabSumber           sql.NullString
			penyebabCUC              sql.NullString
			dampakUraianDaftarRisiko sql.NullString
			dampakPihakDaftarRisiko  sql.NullString
			pengendalianIntern       sql.NullString
			sisaRisiko               sql.NullString
			kriteriaRisiko           sql.NullString
			kemungkinanUraian        sql.NullString
			kemungkinanNilai         sql.NullInt32
			alasan                   sql.NullString
			dampakUraianPetaRisiko   sql.NullString
			nilaiPetaRisiko          sql.NullInt32
			prioritasRisiko          sql.NullString
			toleransiRisiko          sql.NullInt32
			indikasiIndikator        sql.NullString
			penjelasanIndikator      sql.NullString
			batasAmanIndikator       sql.NullString
			opsiPenanganan           sql.NullString
			kegiatanPengendalian     sql.NullString
			outputIndikator          sql.NullString
			targetIndikator          sql.NullString
			jadwal                   sql.NullString
			penanggungJawab          sql.NullString
			cadanganRisiko           sql.NullString
			realisasiPengendalian    sql.NullString
			realisasiRisiko          sql.NullString
			risikoResidu             sql.NullString
			progress                 sql.NullString
			tglPost                  sql.NullTime
			idUserPost               sql.NullInt32
		)

		err := rows.Scan(
			&mr.IdRisiko,
			&idPermasalahan,
			&tahun,
			&mr.RisikoPernyataan,
			&penyebabUraian,
			&penyebabSumber,
			&penyebabCUC,
			&dampakUraianDaftarRisiko,
			&dampakPihakDaftarRisiko,
			&pengendalianIntern,
			&sisaRisiko,
			&kriteriaRisiko,
			&kemungkinanUraian,
			&kemungkinanNilai,
			&alasan,
			&dampakUraianPetaRisiko,
			&nilaiPetaRisiko,
			&prioritasRisiko,
			&toleransiRisiko,
			&indikasiIndikator,
			&penjelasanIndikator,
			&batasAmanIndikator,
			&opsiPenanganan,
			&kegiatanPengendalian,
			&outputIndikator,
			&targetIndikator,
			&jadwal,
			&penanggungJawab,
			&cadanganRisiko,
			&realisasiPengendalian,
			&realisasiRisiko,
			&risikoResidu,
			&progress,
			&tglPost,
			&idUserPost,
		)
		if err != nil {
			log.Fatal(err)
		}

		// Convert nullable fields to pointers
		if idPermasalahan.Valid {
			mr.IdPermasalahan = int(idPermasalahan.Int32)
		}
		if tahun.Valid {
			mr.Tahun = int(tahun.Int32)
		}
		if penyebabUraian.Valid {
			mr.PenyebabUraian = penyebabUraian.String
		}
		if penyebabSumber.Valid {
			mr.PenyebabSumber = penyebabSumber.String
		}
		if penyebabCUC.Valid {
			mr.PenyebabCUC = penyebabCUC.String
		}
		if dampakUraianDaftarRisiko.Valid {
			mr.DampakUraianDaftarRisiko = dampakUraianDaftarRisiko.String
		}
		if dampakPihakDaftarRisiko.Valid {
			mr.DampakPihakDaftarRisiko = dampakPihakDaftarRisiko.String
		}
		if pengendalianIntern.Valid {
			mr.PengendalianIntern = pengendalianIntern.String
		}
		if sisaRisiko.Valid {
			mr.SisaRisiko = sisaRisiko.String
		}
		if kriteriaRisiko.Valid {
			mr.KriteriaRisiko = kriteriaRisiko.String
		}
		if kemungkinanUraian.Valid {
			mr.KemungkinanUraian = kemungkinanUraian.String
		}
		if kemungkinanNilai.Valid {
			mr.KemungkinanNilai = int(kemungkinanNilai.Int32)
		}
		if alasan.Valid {
			mr.Alasan = alasan.String
		}
		if dampakUraianPetaRisiko.Valid {
			mr.DampakUraianPetaRisiko = dampakUraianPetaRisiko.String
		}
		if nilaiPetaRisiko.Valid {
			mr.NilaiPetaRisiko = int(nilaiPetaRisiko.Int32)
		}
		if prioritasRisiko.Valid {
			mr.PrioritasRisiko = prioritasRisiko.String
		}
		if toleransiRisiko.Valid {
			mr.ToleransiRisiko = int(toleransiRisiko.Int32)
		}
		if indikasiIndikator.Valid {
			mr.IndikasiIndikator = indikasiIndikator.String
		}
		if penjelasanIndikator.Valid {
			mr.PenjelasanIndikator = penjelasanIndikator.String
		}
		if batasAmanIndikator.Valid {
			mr.BatasAmanIndikator = batasAmanIndikator.String
		}
		if opsiPenanganan.Valid {
			mr.OpsiPenanganan = opsiPenanganan.String
		}
		if kegiatanPengendalian.Valid {
			mr.KegiatanPengendalian = kegiatanPengendalian.String
		}
		if outputIndikator.Valid {
			mr.OutputIndikator = outputIndikator.String
		}
		if targetIndikator.Valid {
			mr.TargetIndikator = targetIndikator.String
		}
		if jadwal.Valid {
			mr.Jadwal = jadwal.String
		}
		if penanggungJawab.Valid {
			mr.PenanggungJawab = penanggungJawab.String
		}
		if cadanganRisiko.Valid {
			mr.CadanganRisiko = cadanganRisiko.String
		}
		if realisasiPengendalian.Valid {
			mr.RealisasiPengendalian = realisasiPengendalian.String
		}
		if realisasiRisiko.Valid {
			mr.RealisasiRisiko = realisasiRisiko.String
		}
		if risikoResidu.Valid {
			mr.RisikoResidu = risikoResidu.String
		}
		if progress.Valid {
			mr.Progress = progress.String
		}
		if tglPost.Valid {
			mr.TglPost = tglPost.Time
		}
		if idUserPost.Valid {
			mr.IdUserPost = int(idUserPost.Int32)
		}

		risks = append(risks, mr)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return risks, nil
}

func NewRiskRepository(db *sql.DB) RiskRepository {
	return &riskRepository{
		db: db,
	}
}

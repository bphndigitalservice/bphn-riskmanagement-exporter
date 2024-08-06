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

type Risk struct {
	IdRisiko                 uint      `json:"id_risiko"`
	IdPermasalahan           int       `json:"id_permasalahan,omitempty"`
	Tahun                    int       `json:"tahun,omitempty"`
	RisikoPernyataan         string    `json:"risiko_pernyataan"`
	PenyebabUraian           string    `json:"penyebab_uraian,omitempty"`
	PenyebabSumber           string    `json:"penyebab_sumber,omitempty"`
	PenyebabCUC              string    `json:"penyebab_c_uc,omitempty"`
	DampakUraianDaftarRisiko string    `json:"dampak_uraian_daftar_risiko,omitempty"`
	DampakPihakDaftarRisiko  string    `json:"dampak_pihak_daftar_risiko,omitempty"`
	PengendalianIntern       string    `json:"pengendalian_intern,omitempty"`
	SisaRisiko               string    `json:"sisa_risiko,omitempty"`
	KriteriaRisiko           string    `json:"kriteria_risiko,omitempty"`
	KemungkinanUraian        string    `json:"kemungkinan_uraian,omitempty"`
	KemungkinanNilai         int       `json:"kemungkinan_nilai,omitempty"`
	Alasan                   string    `json:"alasan,omitempty"`
	DampakUraianPetaRisiko   string    `json:"dampak_uraian_peta_risiko,omitempty"`
	NilaiPetaRisiko          int       `json:"nilai_peta_risiko,omitempty"`
	PrioritasRisiko          string    `json:"prioritas_risiko,omitempty"`
	ToleransiRisiko          int       `json:"toleransi_risiko,omitempty"`
	IndikasiIndikator        string    `json:"indikasi_indikator,omitempty"`
	PenjelasanIndikator      string    `json:"penjelasan_indikator,omitempty"`
	BatasAmanIndikator       string    `json:"batas_aman_indikator,omitempty"`
	OpsiPenanganan           string    `json:"opsi_penanganan,omitempty"`
	KegiatanPengendalian     string    `json:"kegiatan_pengendalian,omitempty"`
	OutputIndikator          string    `json:"output_indikator,omitempty"`
	TargetIndikator          float64   `json:"target_indikator,omitempty"`
	Jadwal                   string    `json:"jadwal,omitempty"`
	PenanggungJawab          string    `json:"penanggung_jawab,omitempty"`
	CadanganRisiko           string    `json:"cadangan_risiko,omitempty"`
	RealisasiPengendalian    float64   `json:"realisasi_pengendalian,omitempty"`
	RealisasiRisiko          float64   `json:"realisasi_risiko,omitempty"`
	RisikoResidu             string    `json:"risiko_residu,omitempty"`
	Progress                 string    `json:"progress,omitempty"`
	TglPost                  time.Time `json:"tgl_post,omitempty"`
	IdUserPost               int       `json:"id_user_post,omitempty"`
	Owner                    string    `json:"owner,omitempty"`
}

package model

type PerancanganObe struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Nama        string `json:"nama"`
	Status      string `json:"status"`
	KurikulumID int    `json:"kurikulum_id"`
}

type PerancanganObeKurikulum struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Nama      string `json:"nama"`
	Status    string `json:"status"`
	Kurikulum string `json:"kurikulum"`
}

func (PerancanganObe) TableName() string {
	return "perancangan_obe"
}

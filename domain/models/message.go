package models

import "time"

type Message struct {
	ID_MESSAGE               uint       `gorm:"primaryKey;autoIncrement" json:"id_message"`
	JUDUL                    string     `gorm:"type:varchar(255);column:judul;not null" json:"judul"`
	ISI_PESAN                string     `gorm:"type:text;column:isi_pesan;not null" json:"isi_pesan"`
	STATUS_PESAN             string     `gorm:"type:smallint;column:status_pesan;not null" json:"status_pesan"`
	CREATED_AT               *time.Time `gorm:"type:timestamp with time zone;column:created_at" json:"created_at"`
	UPDATE_TGL_STATUS_PESAN  *time.Time `gorm:"type:timestamp with time zone;column:update_tgl_status_pesan" json:"update_tgl_status_pesan"`
}

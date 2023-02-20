package models

type Notes struct {
	NoteID  uint   `gorm:"primaryKey" json:"note_id"`
	Title   string `json:"title"`
	Content string `json:"content" gorm:"type:text"`
	UserID  uint   `json:"user_id"`
}

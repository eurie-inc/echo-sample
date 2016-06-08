package model

import (
	"time"

	"github.com/gocraft/dbr"
)

type Member struct {
	Number    int64  `json:"number"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	CreatedAt int64  `json:"createdAt"`
}

func NewMember(member int64, name, position string) *Member {
	return &Member{
		Number:    member,
		Name:      name,
		Position:  position,
		CreatedAt: time.Now().Unix(),
	}
}

func (m *Member) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("member").
		Columns("number", "name", "position", "created_at").
		Record(m).
		Exec()

	return err
}

func (m *Member) Load(tx *dbr.Tx, number int64) error {

	return tx.Select("*").
		From("member").
		Where("number = ?", number).
		LoadStruct(m)
}

type Members []Member

func (m *Members) Load(tx *dbr.Tx, position string) error {

	var condition dbr.Condition
	if position != "" {
		condition = dbr.Eq("position", position)
	}

	return tx.Select("*").
		From("member").
		Where(condition).
		LoadStruct(m)
}

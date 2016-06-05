package model

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
)

type Member struct {
	Number    int64  `json:"number"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
}

func NewMember(member int64, name string) *Member {
	return &Member{
		Number:    member,
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}
}

func (m *Member) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("member").
		Columns("number", "name", "date_created").
		Record(m).
		Exec()

	if err != nil {
		logrus.Error("Error")
	}

	return err
}

func (m *Member) Load(tx *dbr.Tx, number int64) error {

	_, err := tx.Select("*").
		From("member").
		Where("number = ?", number).
		Load(m)

	if err != nil {
		logrus.Error("Error")
	}
	return err
}

type Members []Member

func (m *Members) Load(tx *dbr.Tx) error {

	_, err := tx.Select("*").
		From("member").
		Load(m)

	if err != nil {
		logrus.Error("Error")
	}
	return err
}

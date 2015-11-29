package model

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"time"
)

type Member struct {
	Number      int64  `db:"number" json:"number"`
	Name        string `db:"name" json:"name"`
	DateCreated int64  `db:"date_created" json:"createdAt"`
}

func NewMember(member int64, name string) *Member {
	return &Member{
		Number:      member,
		Name:        name,
		DateCreated: time.Now().UnixNano(),
	}
}

func (m *Member) SaveMember(tx *dbr.Tx) error {

	_, err := tx.InsertInto("member").
		Columns("number", "name", "date_created").
		Record(m).
		Exec()

	if err != nil {
		logrus.Error("Error")
	}

	return err
}

func (m *Member) LoadMember(tx *dbr.Tx, number int64) error {

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

func (m *Members) LoadMembers(tx *dbr.Tx) error {

	_, err := tx.Select("*").
		From("member").
		Load(m)

	if err != nil {
		logrus.Error("Error")
	}
	return err
}

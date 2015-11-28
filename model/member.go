package model

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"gopkg.in/go-playground/validator.v5"
	"time"
)

func init() {
	validator.New("validate", validator.BakedInValidators)
}

type Member struct {
	Number      int64  `db:"number" json:"number" validate:"required"`
	Name        string `db:"name" json:"name" validate:"required"`
	DateCreated int64  `db:"date_created" json:"createdAt" validate:"-"`
}

func NewMember(member int64, name string) *Member {
	return &Member{
		Number:      member,
		Name:        name,
		DateCreated: time.Now().UnixNano(),
	}
}

func (m *Member) CreateMember(tx *dbr.Tx) error {

	_, err := tx.InsertInto("member").
		Columns("number", "name", "date_created").
		Record(m).
		Exec()

	if err != nil {
		logrus.Error("Error")
	}

	return err
}

func (m *Member) GetMember(tx *dbr.Tx, number int64) error {

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

func (m *Members) GetMembers(tx *dbr.Tx) error {

	_, err := tx.Select("*").
		From("member").
		Load(m)

	if err != nil {
		logrus.Error("Error")
	}
	return err
}

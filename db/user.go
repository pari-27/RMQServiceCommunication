package db

import (
	"context"
	"fmt"

	logger "github.com/sirupsen/logrus"
)

type User struct {
	Name string `db:"name" json:"full_name"`
	Age  int    `db:"age" json:"age"`
}

func (s *pgStore) ListUsers(ctx context.Context) (users []User, err error) {
	err = s.db.Select(&users, "SELECT * FROM users ORDER BY name ASC")
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing users")
		return
	}

	return
}

func (s *pgStore) StoreComputedUsers(msg string) (err error) {
	query := "insert into computeUsers (ComputedValue) values ($1)"
	output, err := s.db.Exec(query, msg)
	fmt.Printf("inserted %v", output)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing users")
		return
	}
	return
}

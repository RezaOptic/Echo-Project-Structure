package logic

import (
	"github.com/Echo-Project-Structure/db"
	"github.com/Echo-Project-Structure/models"
)

func Ping(ping models.Ping) (bool, error) {
	_, err := db.Ping(ping)
	if err != nil {
		return false, err
	}
	return true, nil
}

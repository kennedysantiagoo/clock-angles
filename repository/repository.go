package repository

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/kennedysantiagoo/dbConfig"
	"github.com/kennedysantiagoo/model"
)

func GetByHourAndMinute(hour, minute int) (angle model.Angle, err error) {
	db, err := db.GetConnection()

	if err != nil {
		log.Println("An error ocurred while opening connection")
		return angle, err
	}

	defer db.Close()
	query := fmt.Sprintf(`SELECT id, angle from geometry.clock_angles 
	                     WHERE hour = %d AND minute = %d`, hour, minute)

	err = db.QueryRow(query).Scan(&angle.ID, &angle.Angle)

	if err == sql.ErrNoRows {
		return angle, nil
	}

	return angle, err
}

func Insert(angle model.Angle) (err error) {
	db, err := db.GetConnection()
	if err != nil {
		log.Println("An error ocurred while opening connection")
		return err
	}

	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO geometry.clock_angles
	                      (hour, minute, angle) 
	                      VALUES(%d, %d, %d)`, angle.Hour, angle.Minute, angle.Angle)

	_, err = db.Exec(query)

	if err != nil {
		return err
	}

	return
}

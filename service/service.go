package service

import (
	"log"

	"github.com/kennedysantiagoo/model"
	r "github.com/kennedysantiagoo/repository"

	"github.com/kennedysantiagoo/helper"
)

func GetAngle(hour, minute int) (int, error) {
	angle, err := r.GetByHourAndMinute(hour, minute)

	if err != nil {
		return 0, err
	}

	if angle.ID != 0 {
		return angle.Angle, nil
	}

	angle.Angle = calculateAngle(hour, minute)
	angle.Hour = hour
	angle.Minute = minute

	saveAngle(angle)
	return angle.Angle, nil
}

func calculateAngle(hour, minute int) int {
	var angle int

	if minute == 0 && hour > 6 {
		angle = (hour - OCLOCK_HOUR_MINUTE_ANGLE) * HOUR_ANGLE
	} else {
		angle = (hour * HOUR_ANGLE) -
			(minute * MINUTE_ANGLE) +
			(minute / 2)
	}

	angle = helper.NegativeToPositive(angle)
	return angle
}

func saveAngle(angle model.Angle) error {
	err := r.Insert(angle)

	if err != nil {
		log.Println("An error ocurred while saving angle ", err)
		return err
	}

	return nil
}

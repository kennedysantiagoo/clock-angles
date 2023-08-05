package model

import "time"

type Angle struct {
	ID        int
	Hour      int
	Minute    int
	Angle     int
	CreatedAt time.Time
}

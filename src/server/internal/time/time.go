package time

import (
	"log"
	"time"
)

var TehranLoc *time.Location

func init() {
	l, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		log.Fatal(err)
	}
	TehranLoc = l
}

func GetTime(isTehran bool) string {
	now := time.Now()
	if isTehran {
		now = now.In(TehranLoc)
	}
	return now.Format(time.Kitchen)
}

func GetDay(isTehran bool) string {
	now := time.Now()
	if isTehran {
		now = now.In(TehranLoc)
	}
	return now.Format(time.DateOnly)
}

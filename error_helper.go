package main

import (
	"fmt"

	logger "github.com/loganballard/odds_briefing/logger"
)

func ErrorHelper(err error) {
	logger.Error(fmt.Sprintf("%T\n %s\n %#v\n", err, err, err))
}

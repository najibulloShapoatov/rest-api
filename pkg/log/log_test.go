package mylog_test

import (
	mylog "business/pkg/log"
	"testing"
)

var log = mylog.Log

func TestInit(t *testing.T) {

	mylog.Init("../../log/api.log", "INFO")

	log.Info("fgdfgd", 6, 4, 6)
	log.Warn("warning!!!!", 1, 8, 7)
	log.Error("error!__!", 6, 4, 6)
	log.Warn("warning!!!!", 1, 8, 7)

	

}

package notifications

import (
	"sync"

	"github.com/DreamBridgeNetwork/Go-Utils/pkg/queueutils/fifoqueue"
)

const FifoSize = 100

var notificationsFifo *fifoqueue.FifoQueue

const numMaxJobs = 10

var numJobsRunning int
var numJobsRunningMu sync.Mutex
var mainJobRunning bool

type notificationChannel int

const (
	TextEmail notificationChannel = iota
)

type notification struct {
	notificationChannel notificationChannel
	data                interface{}
}

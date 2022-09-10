package notifications

import (
	"sync"

	"github.com/DreamBridgeNetwork/Go-Utils/pkg/queueutils/fifoqueue"
)

// notificationsConfig - Struct with all notifications configuration
type notificationsConfig struct {
	FifoSize   *uint `json:"fifoSize,omitempty"`
	NumMaxJobs *int  `json:"numMaxJobs,omitempty"`
}

var config notificationsConfig

var notificationsFifo *fifoqueue.FifoQueue

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

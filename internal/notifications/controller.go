package notifications

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/DreamBridgeNetwork/Go-Utils/pkg/emailutils"
	"github.com/DreamBridgeNetwork/Go-Utils/pkg/jsonfile"
	"github.com/DreamBridgeNetwork/Go-Utils/pkg/queueutils/blockbucket"
	"github.com/DreamBridgeNetwork/Go-Utils/pkg/queueutils/fifoqueue"
)

func loadConfiguration() error {
	log.Println("notifications.loadConfiguration")

	log.Println("Loading Notifications configuration.")

	err := jsonfile.ReadJSONFile2("../../config/", "notificationsconfig.json", &config)

	if err != nil {
		log.Println("notifications.loadConfiguration - Error reading Notifications configuration file.")
		return err
	}

	confJson, err := json.MarshalIndent(config, "", "    ")

	if err != nil {
		log.Println("notifications.loadConfiguration - Error prointing Json.")
		return err
	}

	log.Println("Notifications configuration loaded:\n", string(confJson))

	return nil
}

// InitiNotifications - Initialize the notifications routine
func InitiNotifications() error {
	log.Println("notifications.InitiNotifications")

	err := loadConfiguration()
	if err != nil {
		log.Println("notifications.InitiNotifications - Error loading configuration.")
		return err
	}

	notificationsFifo = fifoqueue.NewFifo(*config.FifoSize)
	numJobsRunning = 0
	setMainJobRunning(false)

	return nil
}

// AddNewNotification - Add a new notification do be done. May return some errors.
func AddNewNotification(channel notificationChannel, data interface{}) error {
	log.Println("notifications.AddNewNotification")

	if notificationsFifo.IsFull() {
		return errors.New("notifications.AddNewNotification - Can´t add new notitication. Fifo is full")
	}

	newNotification := notification{channel, data}

	err := notificationsFifo.AddBlock(blockbucket.GetNewBlock(newNotification))

	if err != nil {
		log.Println("notifications.AddNewNotification - Error adding new data to fifo.")
		return err
	}

	if !mainJobRunning {
		go sendNotificationsJob()
	}
	return nil
}

// sendNotificationsJob - Job that keeps verifying if there is any notification to be sent.
func sendNotificationsJob() {
	log.Println("notifications.sendNotificationsJob - Send notifications job initialized.")

	setMainJobRunning(true)
	defer setMainJobRunning(false)

	for !notificationsFifo.IsEmpty() {
		if numJobsRunning < *config.NumMaxJobs {
			addNumJobsRunning()

			log.Println("notifications.sendNotificationsJob - Number of notifications to send: ", notificationsFifo.Size())

			newNotification := blockbucket.StoreBlock(notificationsFifo.RemoveBlock()).(notification)

			switch newNotification.notificationChannel {
			case TextEmail:
				go sendEmail(newNotification.data.(emailutils.TextEmail))
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
	log.Println("notifications.sendNotificationsJob - Send notifications job finished.")
}

func sendEmail(email emailutils.TextEmail) {
	defer subtractNumJobsRunning()

	log.Println("notifications.sendEmail - Sending email notification.")

	err := email.EnviaEmailSMTP()

	if err != nil {
		log.Println("notifications.sendEmail - Error sending email: ", err)
	} else {
		log.Println("notifications.sendEmail - Email sent.")
	}
}

func subtractNumJobsRunning() {
	numJobsRunningMu.Lock()
	defer numJobsRunningMu.Unlock()
	numJobsRunning--
}

func addNumJobsRunning() {
	numJobsRunningMu.Lock()
	defer numJobsRunningMu.Unlock()
	numJobsRunning++
}

func setMainJobRunning(state bool) {
	mainJobRunning = state
}

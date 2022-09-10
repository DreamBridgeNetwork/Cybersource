package webhookreceiver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DreamBridgeNetwork/Go-Cybersource/internal/notifications"
	"github.com/DreamBridgeNetwork/Go-Utils/pkg/emailutils"
	"github.com/DreamBridgeNetwork/Go-Utils/pkg/httputils"
	"github.com/DreamBridgeNetwork/Go-Utils/pkg/jsonfile"
	"github.com/DreamBridgeNetwork/Go-Utils/pkg/stringutils"
)

// LoadWebhookReceiverConfig - Load the configurations for the WebhhokReceiver
func LoadWebhookReceiverConfig() error {
	log.Println("webhookreceiver.LoadWebhookReceiverConfig")

	err := jsonfile.ReadJSONFile2("../../config/", "webhookreceiverconfig.json", &webhookConfig)

	if err != nil {
		log.Println("webhookreceiver.LoadWebhookReceiverConfig - Error reading configuration file.")
		return err
	}

	confJson, err := json.MarshalIndent(webhookConfig, "", "    ")

	if err != nil {
		log.Println("webhookreceiver.LoadWebhookReceiverConfig - Error printing Json.")
		return err
	}

	log.Println("WebhookReceiver configuration loaded:\n", string(confJson))

	return nil
}

func Webhoook(w http.ResponseWriter, req *http.Request) {
	log.Println("webhookreceiver.Webhoook")

	body, err := httputils.RequestBodyToString(req)

	if err != nil {
		log.Printf("webhook.Webhoook - Error reading body: %v", err)
		http.Error(w, "webhook.Webhoook - Error reading body.", http.StatusBadRequest)
		return
	}

	request := "Method: " + req.Method
	request += "\nHost: " + req.Host
	request += "\nPath: " + req.URL.Path
	request += "\nRemoteAddress: " + req.RemoteAddr
	request += "\n\nHeaders:\n" + stringutils.MapToString((*map[string][]string)(&req.Header))
	request += "\nBody:\n" + body
	request += "\n\nTrailer:\n" + stringutils.MapToString((*map[string][]string)(&req.Trailer))

	//log.Println(request)

	var email emailutils.TextEmail

	email.From = webhookConfig.EmailConfig.From
	email.Password = webhookConfig.EmailConfig.Password
	email.Subject = "New webhook received"
	email.Body = request

	adjustRecipients(&email, req.Header.Get("v-c-merchant-id"))

	//log.Printf("Email: %+v", email)

	notifications.AddNewNotification(notifications.TextEmail, email)

	fmt.Fprintln(w, request)
}

func adjustRecipients(email *emailutils.TextEmail, headerMerchantID string) {
	//log.Println("HeaderMerchantID: ", headerMerchantID)

	email.To = append(email.To, webhookConfig.EmailConfig.To...)
	email.Co = append(email.Co, webhookConfig.EmailConfig.Co...)
	email.Cco = append(email.Cco, webhookConfig.EmailConfig.Cco...)

	if headerMerchantID != "" {
		for _, merchant := range webhookConfig.Merchant {
			if merchant.MerchantID == headerMerchantID {
				email.To = append(email.To, merchant.To...)
				email.Co = append(email.Co, merchant.Co...)
				email.Cco = append(email.Cco, merchant.Cco...)
			}
		}
	}
}

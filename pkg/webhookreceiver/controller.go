package webhook

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akayna/Go-dreamBridgeRESTAPI/notifications"
	"github.com/akayna/Go-dreamBridgeUtils/emailutils"
	"github.com/akayna/Go-dreamBridgeUtils/httputils"
	"github.com/akayna/Go-dreamBridgeUtils/stringutils"
)

func LoadConfig() error {
	/*log.Println("webhook.LoadConfig - Loading configuration file", configFileName)

	err := jsonfile.ReadJSONFile2("./configurations/", configFileName, &webhookConfig)

	if err != nil {
		log.Println("webhook.LoadConfig -Error loading configuration file: ", configFileName)
		return err
	}

	log.Println("webhook.LoadConfig - Configuration file loaded.")
	log.Printf("\n%+v", webhookConfig)*/
	webhookConfig.EmailConfig.From = "system@dreambridge.net"
	webhookConfig.EmailConfig.Password = "ekveskkmbuheptzy"
	webhookConfig.EmailConfig.To = append(webhookConfig.EmailConfig.To, "rafael.cunha@visa.com")
	webhookConfig.EmailConfig.To = append(webhookConfig.EmailConfig.To, "aquinocunha@gmail.com")
	return nil
}

func Webhoook(w http.ResponseWriter, req *http.Request) {

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
		for _, merchantID := range webhookConfig.MerchantID {
			if merchantID.MerchantID == headerMerchantID {
				email.To = append(email.To, merchantID.To...)
				email.Co = append(email.Co, merchantID.Co...)
				email.Cco = append(email.Cco, merchantID.Cco...)
			}
		}
	}
}

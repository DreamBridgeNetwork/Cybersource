package webhook

import "github.com/akayna/Go-dreamBridgeUtils/emailutils"

//var configFileName = "webhookConfig.json"
var webhookConfig WebhookConfigurations

type WebhookConfigurations struct {
	EmailConfig emailutils.TextEmail `json:"emailConfig"`
	MerchantID  []struct {
		MerchantID string   `json:"merchantID"`
		To         []string `json:"to"`
		Co         []string `json:"co"`
		Cco        []string `json:"cco"`
	} `json:"merchantID"`
}

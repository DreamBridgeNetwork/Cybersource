package webhookreceiver

import "github.com/DreamBridgeNetwork/Go-Utils/pkg/emailutils"

//var configFileName = "webhookConfig.json"
var webhookConfig WebhookConfigurations

type WebhookConfigurations struct {
	EmailConfig emailutils.TextEmail `json:"emailConfig,omitempty"`
	MerchantID  []struct {
		MerchantID string   `json:"merchantID,omitempty"`
		To         []string `json:"to,omitempty"`
		Co         []string `json:"co,omitempty"`
		Cco        []string `json:"cco,omitempty"`
	} `json:"merchantID"`
}

package notifications

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yendric/backup-manager/config"
)

func sendWebhook(content string) error {
	webhookEnabled := config.Configuration.Notification.Webhook.Enabled
	if !webhookEnabled {
		return errors.New("webhook is not enabled")
	}

	webhookURL := config.Configuration.Notification.Webhook.Url
	contentField := config.Configuration.Notification.Webhook.ContentField

	data := map[string]string{contentField: content}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

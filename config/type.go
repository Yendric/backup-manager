package config

type Config struct {
	Notification struct {
		Email   EmailNotification
		Webhook WebhookNotification
	}
	Backups []struct {
		Name         string
		CreateScript string
		BackupDir    string
		Actions      []struct {
			Name   string
			Script string
		}
	}
}

type EmailNotification struct {
	Enabled      bool
	FromAddress  string
	SmtpHost     string
	SmtpPort     int
	SmtpUsername string
	SmtpPassword string
	To           []string
}

type WebhookNotification struct {
	Enabled      bool
	ContentField string
	Url          string
}

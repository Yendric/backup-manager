package notifications

func SendNotifications(backupName string, actionName string, actionSuccess bool, logFile string) error {
	var actionString string
	var emoji string

	if actionSuccess {
		actionString = "success"
		emoji = "✅"
	} else {
		actionString = "failure"
		emoji = "❌"
	}

	err := sendMail(emoji+" "+backupName+" "+actionName+" action: "+actionString, actionName+": action "+actionString, logFile)
	if err != nil {
		return err
	}

	err = sendWebhook(emoji + " " + backupName + " " + actionName + " action: " + actionString)
	if err != nil {
		return err
	}

	return nil
}

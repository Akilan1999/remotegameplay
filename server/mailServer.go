package server

// MailerSend This function sends an automated email to the
// person requesting it
func MailerSend(subject, email, message string) error {
	// Load email credentials from the configuration file
	//Config, err := config.ConfigInit()
	//if err != nil {
	//	return err
	//}

	// sender configuration.
	//config := mailer.Config{
	//	Host:     Config.EmailHost,
	//	Username: Config.EmailID,
	//	Password: Config.EmailPassword,
	//	FromAddr: Config.EmailID,
	//	Port:     Config.EmailPort,
	//	// Enable UseCommand to support sendmail unix command,
	//	// if this field is true then Host, Username, Password and Port are not required,
	//	// because this info already exists in your local sendmail configuration.
	//	//
	//	// Defaults to false.
	//	UseCommand: false,
	//}

	// initialize a new mail sender service.
	//sender := mailer.New(config)
	//
	//// the rich message body.
	//content := message
	//
	//// the recipient(s).
	//to := []string{email}
	//
	//// send the e-mail.
	//err = sender.Send(subject, content, to...)
	//
	//if err != nil {
	//	return err
	//}

	return nil
}

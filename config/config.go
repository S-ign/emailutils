package config

// Auth .
type Auth struct {
	From     string
	Password string
}

// CreateAuth Creates authentcation object for sending email, requries that
// passwords not be directly passed to this function, rather pass in the
// environment variable string.
func CreateAuth(from, password string) (*Auth, error) {
	return &Auth{
		From:     from,
		Password: password,
	}, nil
}

// SMTPConfig .
type SMTPConfig struct {
	SMTPHost string
	SMTPPort int
}

// NewSMTPServer .
func NewSMTPServer(smtpHost string, smtpPort int) (*SMTPConfig, error) {
	return &SMTPConfig{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
	}, nil
}

var (
	// OFFICE365 SMTP Settings
	OFFICE365, _ = NewSMTPServer("smtp.office365.com", 587)
	// GMAIL SMTP Settings
	GMAIL, _ = NewSMTPServer("smtp.gmail.com", 587)
	// ZOHO_PERSONAL SMTP Settings
	ZOHO_PERSONAL, _ = NewSMTPServer("smtp.zoho.com", 587)
	// ZOHO_ORGANIZATION SMTP Settings
	ZOHO_ORGANIZATION, _ = NewSMTPServer("smtppro.zoho.com", 587)
)

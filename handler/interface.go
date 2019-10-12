package handler

// EmailStore defines the interface for sending the email
// If you want to use Sendgrid, Mailchimp or any other mail client then implement
// the following method for the mail client
type EmailStore interface {
	SendEmail(e *Email) error
}

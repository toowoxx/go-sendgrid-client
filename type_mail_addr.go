package sendgrid

type MailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

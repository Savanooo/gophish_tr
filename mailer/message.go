package mailer

import "github.com/gophish/gomail"

const (
	defaultMessageCharset = "UTF-8"
)

// NewMessage creates a message configured for UTF-8 content so non-ASCII
// characters such as Turkish letters are preserved across headers and bodies.
func NewMessage() *gomail.Message {
	return gomail.NewMessage(
		gomail.SetCharset(defaultMessageCharset),
		gomail.SetEncoding(gomail.Base64),
	)
}

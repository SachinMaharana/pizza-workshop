package notify

import (
	"fmt"
	"os"
	"time"
)

// type, contact, contacttype,

// NewEmail ...
func NewEmail(r, s, m string) *Email {
	return &Email{r, s, m}
}

// Email ...
type Email struct {
	receiver string
	emailid  string
	message  string
}

// SendNotification is email implemntation. Now it shares the same code as sms type
func (e *Email) SendNotification() bool {
	name := fmt.Sprintf("Emailed %s", e.emailid)
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error in file")
		return false
	}
	m := fmt.Sprintf("%s Delieverd %s Pizza to %s. Email: %s\n", time.Now().Format("2006.01.02 15:04:05"), e.message, e.receiver, e.emailid)

	if _, err := f.Write([]byte(m)); err != nil {
		fmt.Println("Error in writitng")
		return false
	}
	defer f.Close()
	return true
}

package notify

import (
	"fmt"
	"os"
	"time"
)

// type, contact, contacttype,

// NewSMS ...
func NewSMS(r, s, m string) *SMS {
	return &SMS{r, s, m}
}

// SMS ...
type SMS struct {
	receiver string
	phone    string
	message  string
}

// SendNotification is sms implemntation. Now it shares the same code as email type
func (s SMS) SendNotification() bool {
	name := fmt.Sprintf("SMS %s", s.phone)
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error in file")
		return false
	}
	m := fmt.Sprintf("%s Delieverd %s Pizza to %s. SMS: %s\n", time.Now().Format("2006.01.02 15:04:05"), s.message, s.receiver, s.phone)

	if _, err := f.Write([]byte(m)); err != nil {
		fmt.Println("Error in writitng")
		return false
	}
	defer f.Close()
	return true
}

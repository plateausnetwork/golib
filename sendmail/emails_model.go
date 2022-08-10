package sendmail

import "github.com/sendgrid/sendgrid-go/helpers/mail"

type Email struct {
	Address string
	Name    string
}

type Emails []Email

func (e Emails) List() (list []*mail.Email) {
	for _, val := range e {
		if val.Name == "" {
			val.Name = val.Address
		}
		list = append(list, mail.NewEmail(val.Name, val.Address))
	}
	return list
}

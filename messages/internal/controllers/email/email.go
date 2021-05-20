package email

import (
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"

	"github.com/ZupIT/horusec-platform/messages/internal/enums/templates"
	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer"
)

type IController interface {
	SendEmail(message *emailEntities.Message) error
}

type Controller struct {
	tpl           *template.Template
	mailerService mailer.IService
}

func NewEmailController(mailerService mailer.IService) IController {
	tpl := template.Must(template.New(emailEnums.AccountConfirmation.ToString()).Parse(templates.EmailConfirmationTpl))
	tpl = template.Must(tpl.New(emailEnums.ResetPassword.ToString()).Parse(templates.ResetPasswordTpl))
	tpl = template.Must(tpl.New(emailEnums.OrganizationInvite.ToString()).Parse(templates.OrganizationInviteTpl))

	return &Controller{
		tpl:           tpl,
		mailerService: mailerService,
	}
}

func (c *Controller) SendEmail(data *emailEntities.Message) error {
	body := new(bytes.Buffer)
	if err := c.tpl.ExecuteTemplate(body, data.TemplateName.ToString(), data.Data); err != nil {
		return err
	}

	return c.mailerService.SendEmail(c.createMessage(data, body.String()))
}

func (c *Controller) createMessage(data *emailEntities.Message, body string) *gomail.Message {
	message := gomail.NewMessage()

	message.SetHeader("From", c.mailerService.GetFromHeader())
	message.SetHeader("Subject", data.Subject)
	message.SetHeader("To", data.To)
	message.SetBody("text/html", body)

	return message
}

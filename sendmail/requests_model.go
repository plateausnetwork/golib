package sendmail

const (
	TEMPLATE_LANGUAGE_ENUS = "en-us"
	TEMPLATE_LANGUAGE_PTBR = "pt-br"
)

type SendRequest struct {
	Template string
	To       Emails
	Data     DynamicData
}

func (s *SendRequest) SetDefaultTemplateLanguage() {
	switch s.Data["lang"] {
	case TEMPLATE_LANGUAGE_ENUS:
	case TEMPLATE_LANGUAGE_PTBR:
	default:
		s.Data["lang"] = TEMPLATE_LANGUAGE_ENUS
	}
}

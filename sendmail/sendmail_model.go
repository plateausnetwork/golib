package sendmail

type SendmailOptions struct {
	ApiKey            string
	From              Email
	GlobalDynamicData DynamicData
	Templates         Templates
}

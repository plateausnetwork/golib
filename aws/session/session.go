//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package session

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

type (
	Options struct {
		AWSOptions *session.Options
	}
	Session interface {
		GetSession() *session.Session
	}
	implSession struct {
		awsSession *session.Session
	}
)

func New(opts Options) (Session, error) {
	opts.SetDefaults()
	sess, err := session.NewSessionWithOptions(*opts.AWSOptions)
	if err != nil {
		return nil, err
	}
	return implSession{awsSession: sess}, nil
}

func (o *Options) SetDefaults() {
	if o.AWSOptions == nil {
		o.AWSOptions = &session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}
	}
}

func (i implSession) GetSession() *session.Session {
	return i.awsSession
}

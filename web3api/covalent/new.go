package covalent

import (
	"errors"
	"strings"
	"time"

	"github.com/rhizomplatform/golib/httphelper"
	"github.com/rhizomplatform/golib/web3api/web3mod"
)

const (
	defaultTimeout = 5 * time.Minute
	defaultAPIURL  = "https://api.covalenthq.com/v1/"
)

type (
	implConvalent struct {
		apiURL     string
		apiKey     string
		httpHelper httphelper.Client
	}
	Options struct {
		APIURL     string
		APIKey     string
		HTTPHelper httphelper.Client
	}
)

func New(opt Options) (web3mod.Web3, error) {
	opt.SetDefault()
	if err := opt.Validate(); err != nil {
		return nil, err
	}
	return &implConvalent{
		apiURL: opt.APIURL,
		apiKey: opt.APIKey,
	}, nil
}

func (o *Options) SetDefault() {
	if o.APIURL == "" {
		o.APIURL = defaultAPIURL
	} else if !strings.HasSuffix(o.APIURL, "/") {
		o.APIURL = o.APIURL + "/"
	}
	if o.HTTPHelper == nil {
		o.HTTPHelper = httphelper.New(httphelper.Options{
			BaseURL: o.APIURL,
			Timeout: defaultTimeout,
		})
		o.HTTPHelper.SetAuthBasicToHeader(o.APIKey, "")
	}
}

func (o *Options) Validate() error {
	if o.APIURL == "" {
		return errors.New("invalid api url")
	}
	if o.APIKey == "" {
		return errors.New("invalid api key")
	}
	return nil
}

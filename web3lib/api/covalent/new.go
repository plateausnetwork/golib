package covalent

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/rhizomplatform/golib/httpclient"
	"github.com/rhizomplatform/golib/web3lib/api/web3mod"
)

type (
	implConvalent struct {
		apiURL     string
		apiKey     string
		httpClient *http.Client
	}
	Options struct {
		ApiURL     string
		ApiKey     string
		HttpClient *http.Client
	}
)

func New(opt Options) (web3mod.Web3, error) {
	opt.SetDefault()
	if err := opt.Validate(); err != nil {
		return nil, err
	}
	return &implConvalent{
		apiURL:     opt.ApiURL,
		apiKey:     opt.ApiKey,
		httpClient: opt.HttpClient,
	}, nil
}

func (o *Options) SetDefault() {
	if o.ApiURL == "" {
		o.ApiURL = "https://api.covalenthq.com/v1/"
	} else if !strings.HasSuffix(o.ApiURL, "/") {
		o.ApiURL = o.ApiURL + "/"
	}
	if o.HttpClient == nil {
		header := http.Header{}
		httpclient.SetAuthBasicToHeader(&header, o.ApiKey, "")
		httpclient.SetAppJSONToHeader(&header)
		o.HttpClient = httpclient.New(httpclient.Options{
			Timeout: 5 * time.Minute, // temporarily timeout
			Header:  header,
		})
	}
}

func (o *Options) Validate() error {
	if o.ApiURL == "" {
		return errors.New("invalid api url")
	}
	if o.ApiKey == "" {
		return errors.New("invalid api key")
	}
	if o.HttpClient == nil {
		return errors.New("invalid http client")
	}
	return nil
}

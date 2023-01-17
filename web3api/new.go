package web3api

import (
	"errors"

	"github.com/rhizomplatform/golib/web3api/covalent"
	"github.com/rhizomplatform/golib/web3api/web3mod"
)

const (
	ProviderCovalent = "covalent"
	ConfigAPIKey     = "api_key"
	ConfigAPIURL     = "api_url"
)

type (
	Options struct {
		ProviderType   string
		ProviderConfig map[string]interface{}
	}
)

func New(opts Options) (web3mod.Web3, error) {
	switch opts.ProviderType {
	case ProviderCovalent:
		apiKey, ok := opts.ProviderConfig[ConfigAPIKey].(string)
		if !ok {
			apiKey = ""
		}
		apiURL, ok := opts.ProviderConfig[ConfigAPIURL].(string)
		if !ok {
			apiURL = ""
		}
		return covalent.New(covalent.Options{
			APIKey: apiKey,
			APIURL: apiURL,
		})
	default:
		return nil, errors.New("invalid web3 provider type")
	}
}

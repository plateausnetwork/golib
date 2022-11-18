package web3lib

import (
	"errors"
	"net/http"

	"github.com/rhizomplatform/golib/web3lib/api/covalent"
	"github.com/rhizomplatform/golib/web3lib/api/web3mod"
)

const (
	ProviderCovalent = "covalent"
	ConfigAPIKey     = "api_key"
	ConfigAPIURL     = "api_url"
	ConfigHTTPClient = "http_client"
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
		client, ok := opts.ProviderConfig[ConfigHTTPClient].(*http.Client)
		if !ok {
			client = nil
		}
		return covalent.New(covalent.Options{
			ApiKey:     apiKey,
			ApiURL:     apiURL,
			HttpClient: client,
		})
	default:
		return nil, errors.New("invalid web3 provider type")
	}
}

package covalent

type (
	NFTData struct {
		TokenID      string `json:"token_id"`
		TokenBalance string `json:"token_balance"`
		TokenURL     string `json:"token_url"`
	}
	BalanceItem struct {
		ContractName         string    `json:"contract_name"`
		ContractTickerSymbol string    `json:"contract_ticker_symbol"`
		ContractAddress      string    `json:"contract_address"`
		LogoURL              string    `json:"logo_url"`
		Type                 string    `json:"type"`
		SupportsERC          []string  `json:"supports_erc"`
		NFTData              []NFTData `json:"nft_data"`
	}
	BalanceData struct {
		Items []BalanceItem `json:"items"`
	}
	GetBalanceOut struct {
		Data         *BalanceData `json:"data"`
		Error        bool         `json:"error"`
		ErrorMessage string       `json:"error_message"`
		ErrorCode    string       `json:"error_code"`
	}
)

type (
	GetExternalMetadataIn struct {
		ContractAddress string
		NFTID           string
		ChainID         int
	}
	GetExternalMetadataOut struct {
		Data         *ExternalData `json:"data"`
		Error        bool          `json:"error"`
		ErrorMessage string        `json:"error_message"`
		ErrorCode    string        `json:"error_code"`
	}
	ExternalData struct {
		Items []ExternalItem `json:"items"`
	}
	ExternalItem struct {
		NFTData []NFTData `json:"nft_data"`
	}
	ExternalNFTData struct {
		TokenID  string `json:"token_id"`
		TokenURL string `json:"token_url"`
	}
)

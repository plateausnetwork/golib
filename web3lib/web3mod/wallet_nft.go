package web3mod

const (
	ContractERC721Type  = "erc721"
	ContractERC1155Type = "erc1155"
	ContractOtherType   = "other"
)

type (
	GetWalletNFTsIn struct {
		ChainID          int
		Wallet           string
		NFTAddressFilter string
	}
	GetWalletNFTsOut struct {
		List  []WalletNFT `json:"list"`
		Error error       `json:"error"`
	}
	WalletNFT struct {
		ContractAddress string    `json:"contract_address"`
		ContractName    string    `json:"contract_name"`
		ContractSymbol  string    `json:"contract_symbol"`
		ContractType    string    `json:"contract_type"`
		LogoURL         string    `json:"logo_url"`
		NFTData         []NFTData `json:"nft_data"`
	}
	NFTData struct {
		TokenID      string `json:"token_id"`
		TokenBalance string `json:"token_balance"`
		TokenURL     string `json:"token_url"`
	}
)

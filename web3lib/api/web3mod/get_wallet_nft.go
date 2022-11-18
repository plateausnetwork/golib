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
		List  *WalletNFTs `json:"list"`
		Error error       `json:"error"`
	}
	WalletNFTs []WalletNFT
	WalletNFT  struct {
		ContractAddress string      `json:"contract_address"`
		ContractName    string      `json:"contract_name"`
		ContractSymbol  string      `json:"contract_symbol"`
		ContractType    string      `json:"contract_type"`
		LogoURL         string      `json:"logo_url"`
		NFTData         NFTDataList `json:"nft_data"`
	}
	NFTDataList []NFTData
	NFTData     struct {
		TokenID      string `json:"token_id"`
		TokenBalance string `json:"token_balance"`
		TokenURL     string `json:"token_url"`
	}
)

func (w WalletNFTs) IsEmpty() bool {
	return w == nil || len(w) == 0
}

func (w WalletNFTs) FilterByAddress(address string) WalletNFTs {
	if address == "" {
		return w
	}
	for _, nft := range w {
		if nft.ContractAddress == address {
			return WalletNFTs{
				nft,
			}
		}
	}
	return w
}

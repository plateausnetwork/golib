package covalent

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/rhizomplatform/golib/web3lib/web3mod"
)

const (
	BalanceItemType = "nft"
	ERC721Type      = "erc721"
	ERC1155Type     = "erc1155"
)

func (i implConvalent) GetWalletNFTs(in web3mod.GetWalletNFTsIn) (out web3mod.GetWalletNFTsOut) {
	balance, err := i.getBalanceFromHttp(in)
	if err != nil {
		return web3mod.GetWalletNFTsOut{
			List:  nil,
			Error: err,
		}
	}
	nftList := i.extractNFTListFromBalance(balance)
	nftList = i.filterNFTListByContractAddress(nftList, in.NFTAddressFilter)
	return web3mod.GetWalletNFTsOut{
		List:  nftList,
		Error: nil,
	}
}

func (i implConvalent) getBalanceFromHttp(in web3mod.GetWalletNFTsIn) (*GetBalanceOut, error) {
	requestURL := fmt.Sprintf("%s%d/address/%s/balances_v2/?nft=true&no-nft-fetch=true", i.apiURL, in.ChainID, in.Wallet)
	res, err := i.httpClient.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf(
			"failed to get balance from %s, status code: %d, body: %s",
			requestURL, res.StatusCode, string(body),
		)
	}
	balance := &GetBalanceOut{}
	if err = json.Unmarshal(body, balance); err != nil {
		return nil, err
	}
	return balance, nil
}

func (i implConvalent) extractNFTListFromBalance(balance *GetBalanceOut) []web3mod.WalletNFT {
	if balance == nil || balance.Data == nil {
		return []web3mod.WalletNFT{}
	}
	var (
		conType string
		nftData []web3mod.NFTData
		nftItem web3mod.WalletNFT
		nftList = make([]web3mod.WalletNFT, 0, len(balance.Data.Items))
	)
	for _, item := range balance.Data.Items {
		if item.Type == BalanceItemType {
			conType = web3mod.ContractOtherType
			for _, ercType := range item.SupportsERC {
				switch strings.ToLower(ercType) {
				case ERC721Type:
					conType = web3mod.ContractERC721Type
					break
				case ERC1155Type:
					conType = web3mod.ContractERC1155Type
					break
				}
			}
			nftData = make([]web3mod.NFTData, len(item.NFTData))
			for k := 0; k < len(item.NFTData); k++ {
				nftData[k] = web3mod.NFTData{
					TokenID:      item.NFTData[k].TokenID,
					TokenBalance: item.NFTData[k].TokenBalance,
					TokenURL:     item.NFTData[k].TokenURL,
				}
			}
			nftItem = web3mod.WalletNFT{
				ContractAddress: item.ContractAddress,
				ContractName:    item.ContractName,
				ContractSymbol:  item.ContractTickerSymbol,
				ContractType:    conType,
				LogoURL:         item.LogoURL,
				NFTData:         nftData,
			}
			nftList = append(nftList, nftItem)
		}
	}
	return nftList
}

func (i implConvalent) filterNFTListByContractAddress(
	nftList []web3mod.WalletNFT, contractAddress string,
) (filtered []web3mod.WalletNFT) {
	if contractAddress == "" || len(nftList) == 0 {
		return nftList
	}
	filtered = make([]web3mod.WalletNFT, 0, 1)
	for _, nft := range nftList {
		if nft.ContractAddress == contractAddress {
			filtered = append(filtered, nft)
			break
		}
	}
	return filtered
}

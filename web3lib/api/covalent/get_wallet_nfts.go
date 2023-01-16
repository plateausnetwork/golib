package covalent

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/rhizomplatform/golib/httphelper"
	"github.com/rhizomplatform/golib/logger"
	"github.com/rhizomplatform/golib/web3lib/api/web3mod"
)

const (
	BalanceItemType = "nft"
	ERC721Type      = "erc721"
	ERC1155Type     = "erc1155"
)

func (i *implConvalent) GetWalletNFTs(in web3mod.GetWalletNFTsIn) (out web3mod.GetWalletNFTsOut) {
	balance, err := i.getBalanceFromHttp(in)
	if err != nil {
		return web3mod.GetWalletNFTsOut{
			List:  nil,
			Error: err,
		}
	}
	nftList := i.extractNFTListFromBalance(balance)
	nftList = nftList.FilterByAddress(in.NFTAddressFilter)
	nftList = i.getNFTListMetadata(in.ChainID, nftList)
	return web3mod.GetWalletNFTsOut{
		List:  &nftList,
		Error: nil,
	}
}

func (i *implConvalent) getNFTListMetadata(chainID int, nftList web3mod.WalletNFTs) web3mod.WalletNFTs {
	var wg sync.WaitGroup
	for k := 0; k < len(nftList); k++ {
		nft := &nftList[k]
		for j := 0; j < len(nft.NFTData); j++ {
			nftData := &nft.NFTData[j]
			wg.Add(1)
			go func(wg *sync.WaitGroup, nftData *web3mod.NFTData, i *implConvalent) {
				defer wg.Done()
				request := GetExternalMetadataIn{
					ContractAddress: nft.ContractAddress,
					ChainID:         chainID,
					NFTID:           nftData.TokenID,
				}
				httpResponse, err := i.getExternalMetadataFromHttp(request)
				if err != nil {
					logger.Error("failed to get external, request: ", request, " metadata: ", err)
				}
				validResponse := httpResponse != nil && httpResponse.Data != nil && httpResponse.Error == false && len(httpResponse.Data.Items) > 0
				if validResponse {
					metadata := httpResponse.Data.Items[0].NFTData
					if len(metadata) > 0 {
						nftData.TokenURL = metadata[0].TokenURL
					}
				}
			}(&wg, nftData, i)
		}
	}
	wg.Wait()
	return nftList
}

func (i *implConvalent) getBalanceFromHttp(in web3mod.GetWalletNFTsIn) (*GetBalanceOut, error) {
	var (
		balance = &GetBalanceOut{}
		fail    = i.httpHelper.Get(httphelper.Request{
			Context:     context.Background(),
			Endpoint:    fmt.Sprintf("%d/address/%s/balances_v2/?nft=true&no-nft-fetch=true", in.ChainID, in.Wallet),
			Destination: balance,
		})
	)
	if fail != nil && fail.Err != nil {
		return nil, fail.Err
	}
	return balance, nil
}

func (i *implConvalent) getExternalMetadataFromHttp(in GetExternalMetadataIn) (*GetExternalMetadataOut, error) {
	var (
		nftMetadata = &GetExternalMetadataOut{}
		fail        = i.httpHelper.Get(httphelper.Request{
			Context:     context.Background(),
			Endpoint:    fmt.Sprintf("%d/tokens/%s/nft_metadata/%s/", in.ChainID, in.ContractAddress, in.NFTID),
			Destination: nftMetadata,
		})
	)
	if fail != nil && fail.Err != nil {
		return nil, fail.Err
	}
	return nftMetadata, nil
}

func (i *implConvalent) extractNFTListFromBalance(balance *GetBalanceOut) web3mod.WalletNFTs {
	if balance == nil || balance.Data == nil {
		return web3mod.WalletNFTs{}
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

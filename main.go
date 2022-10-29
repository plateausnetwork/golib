package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rhizomplatform/golib/smartcontracts/verifyinterfaces"
	"github.com/rhizomplatform/golib/web3lib/eth/ethclient"
)

const (
	keyERC20Std = iota
	keyERC20Det
	keyERC721
	keyERC1155Std
	keyERC1155TokenReceiver
	keyERC1155Accepted
	keyERC1155BatchAccepted
)

var (
	iERC20Std             = [4]byte{54, 55, 43, 7}     // 0x36372b07
	iERC20Det             = [4]byte{162, 25, 160, 37}  // 0xa219a025
	iERC721               = [4]byte{128, 172, 88, 205} // 0x80ac58cd
	iERC1155Std           = [4]byte{217, 182, 122, 38} // 0xd9b67a26
	iERC1155TokenReceiver = [4]byte{78, 35, 18, 224}   // 0x4e2312e0
	iERC1155Accepted      = [4]byte{242, 58, 110, 97}  // 0xf23a6e61
	iERC1155BathAccepted  = [4]byte{188, 25, 124, 129} // 0xbc197c81
)

func main() {
	url := "wss://maximum-radial-friday.bsc-testnet.discover.quiknode.pro/2641ec7a970e079be3abe314939955117ad6546c/"
	client, err := ethclient.NewClient(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			currentBlock := header.Number
			checkBlock := big.NewInt(currentBlock.Int64() - 5)
			BlockChecker(ctx, client, checkBlock)
		}
	}
	//checkBlock := big.NewInt(24125456)
	//BlockChecker(ctx, client, checkBlock)
}

func BlockChecker(ctx context.Context, client ethclient.Client, blockNumber *big.Int) {
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	fmt.Println("BLOCK: ", block.Hash())
	for _, tx := range block.Transactions() {
		if tx.To() != nil {
			continue
		}
		from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
		if err != nil {
			from, err = types.Sender(types.HomesteadSigner{}, tx)
			if err != nil {
				fmt.Println("ERR: ", err)
			}
		}
		fmt.Println("TX: ", strings.ToLower(tx.Hash().Hex()))
		fmt.Println("FROM: ", strings.ToLower(from.Hex()))
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			fmt.Println("ERR: ", err)
		}
		if receipt.ContractAddress.String() != "0x0000000000000000000000000000000000000000" {
			contractType := CheckInterfaces(client, receipt.ContractAddress)
			fmt.Println("CONTRACT: ", receipt.ContractAddress.String())
			fmt.Println("TYPE: ", contractType)
		}
	}
	//event := ethevent.NewEvent(ethevent.Options{Client: client})
	//event.SubscribingLogsBlocks(blockFrom)
}

func CheckInterfaces(client ethclient.Client, address common.Address) string {
	checkInterfaces := [][4]byte{
		iERC20Std,
		iERC20Det,
		iERC721,
		iERC1155Std,
		iERC1155TokenReceiver,
		iERC1155Accepted,
		iERC1155BathAccepted,
	}
	conCheck := common.HexToAddress("0x9EeA70E441fbDC34CDEB7ccB01952bBcb3B1000a")
	verifyInterfaces, err := verifyinterfaces.NewVerifyinterfaces(conCheck, client)
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	results, err := verifyInterfaces.GetSupportedInterfaces(nil, address, checkInterfaces)
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	for k, result := range results {
		if result {
			switch k {
			case keyERC20Std:
				return "ERC20STD"
			case keyERC20Det:
				return "ERC20DET"
			case keyERC721:
				return "ERC721"
			case keyERC1155Std:
				return "ERC1155STD"
			case keyERC1155TokenReceiver:
				return "ERC1155TOKENRECEIVER"
			case keyERC1155Accepted:
				return "ERC1155ACCEPTED"
			case keyERC1155BatchAccepted:
				return "ERC1155BATCHACCEPTED"
			}
			break
		}
	}
	return "OTHER"
}

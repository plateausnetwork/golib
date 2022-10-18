//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package web3mod

type (
	Web3 interface {
		GetWalletNFTs(in GetWalletNFTsIn) GetWalletNFTsOut
	}
)

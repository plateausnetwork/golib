//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package web3

type Web3 interface {
}

type web3Impl struct {
}

func New() Web3 {
	return web3Impl{}
}

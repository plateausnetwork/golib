package moralisnft

import (
	"fmt"
	"io"
	"net/http"
)

func (m moralisNftImpl) GetList(addressWallet, addressNFT, chain, format, cursor string, limit int) ([]byte, error) {
	url := fmt.Sprintf(
		"%s%s/nft/%s?chain=%s&format=%s&cursor=%s&limit=%d",
		m.apiUrl, addressWallet, addressNFT, chain, format, cursor, limit,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (m moralisNftImpl) GetInContract(tokenAddress, tokenID, chain, format string) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/nft/%s/%s?chain=%s&format=%s",
		m.apiUrl, tokenAddress, tokenID, chain, format,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (m moralisNftImpl) GetWalletBalance(addressWallet, chain, tokenAddress string) ([]byte, error) {
	url := fmt.Sprintf("%s%s/erc20?", m.apiUrl, addressWallet)
	if chain != "" {
		url = fmt.Sprintf("%s&chain=%s", url, chain)
	}
	if tokenAddress != "" {
		url = fmt.Sprintf("%s&token_addresses=%s", url, tokenAddress)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

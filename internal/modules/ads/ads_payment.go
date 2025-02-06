package ads

import (
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SmartContractAdInserter struct {
	client   *ethclient.Client
	address  common.Address
	contract *AdsTransactor
}

func NewSmartContractAdInserter(clientURL string, contractAddress string) (*SmartContractAdInserter, error) {
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(contractAddress)
	contract, err := NewAdsTransactor(address, client)
	if err != nil {
		return nil, err
	}
	return &SmartContractAdInserter{
		client:   client,
		address:  address,
		contract: contract,
	}, nil
}

func (sai *SmartContractAdInserter) PlaceAd(adText, adImage string, payment *big.Int) error {
	auth, err := bind.NewTransactor(strings.NewReader("keyJson"), "passphrase")
	if err != nil {
		return err
	}
	tx, err := sai.contract.PlaceAd(auth, adText, adImage)
	if err != nil {
		return err
	}
	log.Printf("Ad placed with transaction hash: %s", tx.Hash().Hex())
	return nil
}

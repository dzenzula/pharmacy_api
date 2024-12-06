package blockchain

import (
	"context"
	"log"
	"math/big"
	"time"

	"pharmacy/shipment-blockchain/contracts" // Generated ABI contracts

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var blockchainClient *ethclient.Client

func SetClient(client *ethclient.Client) {
	blockchainClient = client
}

func AddShipmentToBlockchain(batch string, fromCity string, toCity string, status string, timestamp string) (string, error) {
	// Connect to contract
	contractAddress := common.HexToAddress("0x85E20C160598c1193fc3BBe9A5BF646d158A813E")
	instance, err := contracts.NewContracts(contractAddress, blockchainClient)
	if err != nil {
		return "", err
	}

	chainID, err := blockchainClient.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get Chain ID: %v", err)
	}

	// Transaction setting
	privateKey := "cfd71a355c8e7edccf0014b37de8a49e3d14958ef9ccc9d9fffd14b8e38133f0"
	// Convert string to bytes
	keyBytes := common.Hex2Bytes(privateKey)
	// Parse private key
	privKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatalf("Error creating transactor: %v", err)
	}

	// Send data to blockchain
	tx, err := instance.AddShipment(auth, batch, fromCity, toCity, status, timestamp)
	if err != nil {
		return "", err
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

// VerifyTransaction checks if a transaction exists and is mined on the blockchain
func VerifyTransaction(hashCode string) (bool, error) {
	txHash := common.HexToHash(hashCode)

	// Fetch transaction receipt
	receipt, err := blockchainClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		// If transaction is not found, it is considered invalid
		if err == ethereum.NotFound {
			return false, nil
		}
		return false, err
	}

	// Check if the transaction is successfully mined
	return receipt.Status == 1, nil
}

func UpdateShipmentOnBlockchain(batch string) (string, error) {
	// Connect to contract
	contractAddress := common.HexToAddress("0x85E20C160598c1193fc3BBe9A5BF646d158A813E")
	instance, err := contracts.NewContracts(contractAddress, blockchainClient)
	if err != nil {
		return "", err
	}

	chainID, err := blockchainClient.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get Chain ID: %v", err)
	}

	// Transaction setting
	privateKey := "cfd71a355c8e7edccf0014b37de8a49e3d14958ef9ccc9d9fffd14b8e38133f0"
	keyBytes := common.Hex2Bytes(privateKey)
	privKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatalf("Error creating transactor: %v", err)
	}

	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(20000000000)

	// Generate timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	status := "Delivered"

	log.Printf("Updating shipment: BatchID=%s, Status=%s, Timestamp=%s", batch, status, timestamp)

	// Send data to blockchain
	tx, err := instance.UpdateShipment(auth, batch, status, timestamp)
	if err != nil {
		return "", err
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

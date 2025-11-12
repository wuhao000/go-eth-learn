package utils

import (
  "log"

  "github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetClient() *ethclient.Client {
  if client != nil {
    return client
  }
  sepoliaRPC := "https://sepolia.infura.io/v3/f05f2e17cd7a4b9caf9a06d507c042a1"
  client, err := ethclient.Dial(sepoliaRPC)
  if err != nil {
    log.Fatalf("连接以太坊客户端失败: %v", err)
  }
  return client
}

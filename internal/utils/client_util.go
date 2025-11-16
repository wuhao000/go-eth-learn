package utils

import (
  "log"

  "go-eth-learn/internal/config"

  "github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetClient() *ethclient.Client {
  if client != nil {
    return client
  }

  // 从配置中获取 RPC URL
  sepoliaRPC := config.GetSepoliaRPCURL()
  client, err := ethclient.Dial(sepoliaRPC)
  if err != nil {
    log.Fatalf("连接以太坊客户端失败: %v", err)
  }
  return client
}

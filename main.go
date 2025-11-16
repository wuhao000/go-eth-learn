package main

import (
  "fmt"
  "go-eth-learn/internal/task1"
  "go-eth-learn/internal/utils"
  "log"
  "time"

  "go-eth-learn/internal"
  "go-eth-learn/internal/config"
)

func main() {
  // 加载配置
  config.LoadConfig()
  fmt.Println("=== 以太坊 Sepolia 区块头查询演示 ===\n")
  // 演示1: 获取最新区块头
  fmt.Println("1. 获取Sepolia最新区块头:")
  latestHeader, err := internal.GetSepoliaLatestHeader()
  if err != nil {
    log.Fatalf("获取最新区块头失败: %v", err)
  }
  internal.PrintHeaderInfo(latestHeader)

  // 启动区块监听（在后台运行）
  go internal.SubscribeBlock()

  // 演示2: 获取指定区块号的区块头
  fmt.Println("\n2. 获取Sepolia指定区块号的区块头:")
  blockNumber := latestHeader.Number.Int64() - 10 // 获取最新区块前10个的区块
  specificHeader, err := internal.GetSepoliaHeaderByNumber(blockNumber)
  if err != nil {
    log.Fatalf("获取指定区块头失败: %v", err)
  }
  internal.PrintHeaderInfo(specificHeader)

  // 演示3: 创建客户端并多次查询
  fmt.Println("\n3. 使用客户端进行多次查询:")
  client, err := internal.GetSepoliaClient()
  if err != nil {
    log.Fatalf("创建Sepolia客户端失败: %v", err)
  }
  defer client.Close()

  // 获取多个区块头
  for i := 0; i < 3; i++ {
    targetNumber := blockNumber - int64(i)
    fmt.Printf("\n--- 查询区块 %d ---\n", targetNumber)
    header, err := client.GetHeaderByNumber(nil) // 获取最新区块
    if err != nil {
      log.Printf("获取区块头失败: %v", err)
      continue
    }
    fmt.Printf("最新区块号: %d, 哈希: %s\n", header.Number, header.Hash().Hex()[:10]+"...")
  }

  block, err := task1.GetBlock(nil)
  if err != nil {
    log.Printf("获取区块信息失败：%v", err)
    return
  }
  task1.PrintBlock(block)
  defer utils.GetClient().Close()

  // 等待一段时间观察轮询效果
  fmt.Println("\n等待15秒观察区块轮询效果...")
  time.Sleep(15 * time.Second)

  fmt.Println("\n=== 演示完成 ===")
}

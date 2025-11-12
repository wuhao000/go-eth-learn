package task1

import (
  "context"
  "fmt"
  "go-eth-learn/internal"
  "go-eth-learn/internal/utils"
  "math/big"

  "github.com/ethereum/go-ethereum/core/types"
)

func GetBlock(blockNumber *big.Int) (*types.Block, error) {
  client := utils.GetClient()
  block, err := client.BlockByNumber(context.Background(), blockNumber)
  if err != nil {
    client.Close()
    return nil, err
  }
  return block, nil
}

func PrintBlock(block *types.Block) {
  if block == nil {
    fmt.Println("区块为空")
    return
  }

  fmt.Printf("\n=== 完整区块信息 ===\n")
  fmt.Printf("区块号: %d\n", block.Number())
  fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
  fmt.Printf("父区块哈希: %s\n", block.ParentHash().Hex())
  fmt.Printf("叔区块哈希: %v\n", block.Uncles())
  fmt.Printf("时间戳: %d\n", block.Time())
  fmt.Printf("矿工地址: %s\n", block.Coinbase().Hex())
  fmt.Printf("Gas限制: %d\n", block.GasLimit())
  fmt.Printf("已使用Gas: %d\n", block.GasUsed())
  fmt.Printf("区块大小: %d 字节\n", block.Size())
  fmt.Printf("难度: %d\n", block.Difficulty())
  fmt.Printf("随机数: %d\n", block.Nonce())
  fmt.Printf("状态根哈希: %s\n", block.Root().Hex())
  fmt.Printf("交易根哈希: %s\n", block.TxHash().Hex())
  fmt.Printf("收据根哈希: %s\n", block.ReceiptHash().Hex())
  fmt.Printf("日志布隆过滤器: %x\n", block.Bloom())
  fmt.Printf("基础费用: %s wei\n", block.BaseFee().String())

  // 显示交易信息
  fmt.Printf("交易数量: %d\n", len(block.Transactions()))
  printMax := 3
  if len(block.Transactions()) > 0 {
    fmt.Printf("前%d笔交易:\n", printMax)
    for i, tx := range block.Transactions() {
      if i >= printMax {
        break
      }

      // 获取交易发送者地址
      from := "未知"

      // 根据交易类型选择合适的signer
      var signer types.Signer
      switch tx.Type() {
      case types.LegacyTxType:
        signer = types.NewEIP155Signer(tx.ChainId())
      case types.AccessListTxType:
        signer = types.NewEIP2930Signer(tx.ChainId())
      case types.DynamicFeeTxType:
        signer = types.NewLondonSigner(tx.ChainId())
      case types.BlobTxType:
        signer = types.NewLondonSigner(tx.ChainId()) // Blob交易也使用London signer
      case types.SetCodeTxType:
        signer = types.NewLondonSigner(tx.ChainId()) // SetCode交易也使用London signer
      default:
        fmt.Printf("未知交易类型: %d\n", tx.Type())
      }

      if signer != nil {
        sender, err := types.Sender(signer, tx)
        if err == nil {
          from = sender.Hex()[:10] + "..."
        }
      }

      to := "合约创建"
      if tx.To() != nil {
        to = tx.To().Hex()[:10] + "..."
      }

      fmt.Printf("  %d. 哈希: %s, From: %s, To: %s, Value: %s wei",
        i+1,
        tx.Hash().Hex()[:20]+"...",
        from,
        to,
        tx.Value().String())

      // 获取交易收据信息
      receipt, err := utils.GetClient().TransactionReceipt(context.Background(), tx.Hash())
      if err != nil {
        fmt.Printf(", 收据获取失败: %v", err)
        fmt.Println()
        continue
      }

      // 显示交易状态
      status := "失败"
      if receipt.Status == 1 {
        status = "成功"
      }
      fmt.Printf(", 状态: %s", status)

      // 显示Gas使用情况
      fmt.Printf(", Gas: %d/%d", receipt.GasUsed, tx.Gas())

      //printLogs(receipt)

      fmt.Println()
    }
    if len(block.Transactions()) > printMax {
      fmt.Printf("  ... 还有 %d 笔交易\n", len(block.Transactions())-printMax)
    }
  }

  // 显示叔区块信息
  if len(block.Uncles()) > 0 {
    fmt.Printf("叔区块数量: %d\n", len(block.Uncles()))
    for i, uncle := range block.Uncles() {
      fmt.Printf("  %d. 哈希: %s\n", i+1, uncle.Hash().Hex())
    }
  }

  fmt.Printf("====================\n")
}

func printLogs(receipt *types.Receipt) {
  // 显示日志信息
  if len(receipt.Logs) > 0 {
    fmt.Printf(", 日志数: %d", len(receipt.Logs))
    // 打印详细日志信息
    for j, rlog := range receipt.Logs {
      fmt.Printf("\n    日志%d: 地址=%s", j+1, rlog.Address.Hex()[:10]+"...")

      // 使用增强版日志解析功能
      logInfo := internal.EnhancedFormatLogInfo(rlog)
      fmt.Printf("\n      %s", logInfo)

      // 仍然显示原始技术信息作为参考
      if len(rlog.Topics) > 0 {
        fmt.Printf("\n      原始主题数=%d", len(rlog.Topics))
        for k, topic := range rlog.Topics {
          fmt.Printf("\n        主题%d: %s", k+1, topic.Hex()[:10]+"...")
        }
      }
      if len(rlog.Data) > 0 {
        fmt.Printf("\n      数据长度=%d bytes", len(rlog.Data))
        // 如果数据不太长，显示前几个字节
        if len(rlog.Data) <= 32 {
          fmt.Printf("\n      数据=%x", rlog.Data)
        } else {
          fmt.Printf("\n      数据=%x...", rlog.Data[:32])
        }
      }
    }
  }
}

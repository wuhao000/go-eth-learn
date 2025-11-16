package task1

import (
  "context"
  "errors"
  "fmt"
  "go-eth-learn/internal/config"
  "go-eth-learn/internal/utils"
  "log"
  "os"
  "testing"

  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/rpc"
)

const (
  contractAddr = "0x57410187b45B45cA9f23E4B2265Edf2c18991D7E"
)

func TestLoad(t *testing.T) {
  client := utils.GetClient()
  auctionInstance, err := NewAuction(common.HexToAddress(contractAddr), client)
  if err != nil {
    log.Fatal(err)
  }

  // 读取只读方法（例如查看 owner）
  owner, err := auctionInstance.Admin(&bind.CallOpts{Context: context.Background()})
  if err != nil {
    log.Fatal("read failed:", err)
  }

  fmt.Println("Contract owner:", owner.Hex())
  tr, err := auctionInstance.Withdraw(&bind.TransactOpts{
    Context: context.Background(),
  })
  var rpcErr rpc.Error
  if errors.As(err, &rpcErr) && rpcErr.Error() != "execution reverted: no funds." {
    log.Fatal(rpcErr.Error())
  }
  fmt.Println(tr.Cost())
}

// TestMain 在包中的所有测试运行前执行一次，用于统一初始化
func TestMain(m *testing.M) {
  // 统一初始化配置
  config.InitConfig()

  // 运行所有测试
  code := m.Run()
  os.Exit(code)
}

package task1

import (
  "fmt"
  "math/big"
  "testing"
)

func TestTransfer(t *testing.T) {
  privateKey := ""
  targetAddress := "0x1d0ecb42d442baeb1a33af74b45e0f148012941e"
  value := big.NewInt(1000000000000000) // in wei (1 eth)
  hex := Transfer(privateKey, targetAddress, value)
  fmt.Println("tx hash:", hex)
}

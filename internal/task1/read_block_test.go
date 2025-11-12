package task1

import "testing"

func TestPrintBlock(t *testing.T) {
  block, err := GetBlock(nil)
  if err != nil {
    t.Error("读取区块信息失败")
  }
  PrintBlock(block)
}

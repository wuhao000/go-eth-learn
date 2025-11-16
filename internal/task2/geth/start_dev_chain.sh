#!/bin/bash

# 开发模式私链启动脚本
set -e

echo "=== 启动以太坊开发模式私链 ==="

# 启动geth开发节点
echo "启动geth开发节点..."
echo "节点将在以下地址启动:"
echo "HTTP: http://localhost:8545"
echo "WS: ws://localhost:8546"
echo ""
echo "开发模式特点:"
echo "- 预设账户: 0x..."
echo "- 预设余额: 大量ETH"
echo "- 自动挖矿"
echo "按 Ctrl+C 停止节点"

geth --dev --http --http.addr "0.0.0.0" --http.port 8545 --http.api "eth,net,web3,personal" --http.corsdomain "*" --ws --ws.addr "0.0.0.0" --ws.port 8546 --ws.api "eth,net,web3,personal" --ws.origins "*" --dev.period 5 console
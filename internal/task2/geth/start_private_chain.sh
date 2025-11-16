#!/bin/bash

# 私链启动脚本
set -e

echo "=== 启动以太坊私链 ==="

# 创建数据目录
DATA_DIR="./private_chain_data"
if [ ! -d "$DATA_DIR" ]; then
    echo "创建数据目录: $DATA_DIR"
    mkdir -p "$DATA_DIR"
fi

# 初始化创世区块
echo "初始化创世区块..."
geth --datadir "$DATA_DIR" init genesis.json

# 创建账户文件
ACCOUNT_FILE="$DATA_DIR/account.txt"
if [ ! -f "$ACCOUNT_FILE" ]; then
    echo "创建新账户..."
    # 创建一个密码文件
    echo "123456" > "$DATA_DIR/password.txt"

    # 创建账户
    ACCOUNT_ADDRESS=$(geth --datadir "$DATA_DIR" account new --password "$DATA_DIR/password.txt" | grep -o "0x[0-9a-fA-F]\{40\}")
    echo "$ACCOUNT_ADDRESS" > "$ACCOUNT_FILE"
    echo "账户地址: $ACCOUNT_ADDRESS"
else
    ACCOUNT_ADDRESS=$(cat "$ACCOUNT_FILE")
    echo "使用现有账户: $ACCOUNT_ADDRESS"
fi

# 启动geth节点
echo "启动geth节点..."
echo "节点将在以下地址启动:"
echo "HTTP: http://localhost:8545"
echo "WS: ws://localhost:8546"
echo ""
echo "按 Ctrl+C 停止节点"

geth --datadir "$DATA_DIR" \
    --http \
    --http.addr "0.0.0.0" \
    --http.port 8545 \
    --http.api "eth,net,web3,personal" \
    --http.corsdomain "*" \
    --ws \
    --ws.addr "0.0.0.0" \
    --ws.port 8546 \
    --ws.api "eth,net,web3,personal" \
    --ws.origins "*" \
    --nodiscover \
    --maxpeers 0 \
    --networkid 1337 \
    --mine \
    --miner.etherbase "$ACCOUNT_ADDRESS" \
    --dev \
    console
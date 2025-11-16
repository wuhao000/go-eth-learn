## 在wsl中安装go
```bash
sudo apt update
sudo apt install golang-1.21
sudo ln -s /usr/lib/go-1.21/bin/go /usr/local/bin/go
sudo ln -s /usr/lib/go-1.21/bin/gofmt /usr/local/bin/gofmt
export GOTOOLCHAIN=local
```

## 下载geth v1.11.6 （支持挖矿的最后一个版本）

解压后进入源码目录

```bash
make geth
```

初始化私链

```bash
#!/bin/bash
set -e

DATADIR="./data"
NETWORK_ID=2025

# 仅第一次执行
if [ -d "$DATADIR/geth" ]; then
    echo "数据目录已存在，跳过初始化。"
    exit 0
fi

echo "[1] 创建数据目录"
mkdir -p $DATADIR

echo "[2] 写入 genesis.json"
cat > genesis.json <<EOF
{
  "config": {
    "chainId": $NETWORK_ID,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "ethash": {}
  },
  "difficulty": "0x20000",
  "gasLimit": "0x8000000",
  "alloc": {
    "0x0000000000000000000000000000000000000001": {
      "balance": "0xffffffffffffffff"
    }
  }
}
EOF

echo "[3] 初始化创世块"
build/bin/geth --datadir $DATADIR init genesis.json

echo "[4] 创建账户（第一次执行）"
echo "minerpassword" > password.txt
MINER_ACCOUNT=$(build/bin/geth --datadir $DATADIR account new --password password.txt | grep -o '0x[a-fA-F0-9]\{40\}')
echo "$MINER_ACCOUNT" > miner.txt

echo "初始化完成！矿工账户: $MINER_ACCOUNT"

```

启动

```bash
#!/bin/bash
set -e

DATADIR="./data"
NETWORK_ID=2025

if [ ! -f miner.txt ]; then
    echo "矿工账户不存在，请先运行 init.sh 初始化链。"
    exit 1
fi

MINER_ACCOUNT=$(cat miner.txt)

echo "启动节点，矿工账户：$MINER_ACCOUNT"

nohup build/bin/geth \
  --datadir $DATADIR \
  --networkid $NETWORK_ID \
  --http \
  --http.api eth,web3,personal,net,miner \
  --http.addr 0.0.0.0 \
  --http.port 8545 \
  --http.corsdomain "*" \
  --http.vhosts="*" \
  --allow-insecure-unlock \
  --unlock $MINER_ACCOUNT \
  --password password.txt \
  --mine \
  --miner.threads=2 \
  --miner.etherbase=$MINER_ACCOUNT \
  --ipcdisable \
  > geth.log 2>&1 &

echo "Geth 已启动，日志在 geth.log"

```

通过以下命令进入控制台
```bash
build/bin/geth attach http://127.0.0.1:8545
```

使用指定的线程数开始挖矿
```bash
 miner.start(24)
```

查看当前区块高度
```bash
 eth.blockNumber
```

查询账户余额（注意地址要替换成之前日志中生成的地址）
```bash
eth.getBalance("0x000ea01e63ebcAB8879d9036B3d1771978d12084")
```

```JS
// 读取 ABI 和 bytecode
var abi = JSON.parse(require('fs').readFileSync('/mnt/i/web3/solidity-learn/contracts/task1/build/IntegerToRoman.abi', 'utf8'));
var bytecode = "0x" + require('fs').readFileSync('/mnt/i/web3/solidity-learn/contracts/task1/build/IntegerToRoman.bin', 'utf8');

// 创建合约对象
var IntegerToRoman = eth.contract(abi);

// 部署
var instance = IntegerToRoman.new({from: "0x83e61B16E254f9181EBA01f1D99d5570b136802a", data: bytecode, gas: 3000000}, function(err, res){
    if(!err){
        if(res.address){
            console.log("部署成功，合约地址:", res.address);
        } else {
            console.log("交易哈希:", res.transactionHash);
        }
    } else {
        console.error(err);
    }
});
```

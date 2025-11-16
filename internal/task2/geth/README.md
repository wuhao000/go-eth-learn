# 以太坊私链和智能合约部署

这是一个完整的以太坊私链搭建和智能合约部署项目。

## 项目结构

```
.
├── genesis.json                 # 创世区块配置文件
├── SimpleStorage.sol           # 简单存储智能合约
├── start_private_chain.sh      # 启动私链脚本
├── deploy.js                   # 合约部署脚本
├── interact.js                 # 合约交互脚本
├── package.json                # Node.js项目配置
├── README.md                   # 本说明文档
└── private_chain_data/         # 私链数据目录（运行后自动创建）
```

## 快速开始

### 1. 启动私链

```bash
./start_private_chain.sh
```

这个脚本会：
- 创建数据目录
- 初始化创世区块
- 创建一个新账户（密码：123456）
- 启动geth节点（HTTP端口：8545，WebSocket端口：8546）
- 开始挖矿

**注意：保持这个终端窗口开启，私链将持续运行。**

### 2. 部署智能合约

在新的终端窗口中运行：

```bash
npm run deploy
```

这个脚本会：
- 编译SimpleStorage.sol合约
- 部署合约到私链
- 自动测试合约功能（设置值、读取值、增加操作）
- 保存合约信息到`contract-info.json`

### 3. 与合约交互

```bash
npm run interact
```

提供交互式菜单来：
- 读取当前存储的值
- 设置新的值
- 将当前值增加1
- 查看账户余额

## 智能合约功能

SimpleStorage合约提供以下功能：
- `get()`: 读取当前存储的值
- `set(uint256 x)`: 设置新的值
- `increment()`: 将当前值增加1
- 构造函数：设置初始值（部署时设为42）

## 网络配置

- **网络ID**: 1337
- **共识机制**: Clique (PoA)
- **出块时间**: 5秒
- **初始账户余额**: 1000000 ETH
- **HTTP RPC**: http://localhost:8545
- **WebSocket**: ws://localhost:8546

## 故障排除

### 1. 端口被占用
如果8545或8546端口被占用，请修改`start_private_chain.sh`中的端口号。

### 2. 连接失败
确保私链正在运行，并且可以访问`http://localhost:8545`。

### 3. 合约编译失败
确保Solidity编译器版本正确（本项目使用0.8.0版本）。

### 4. Gas不足
脚本会自动估算Gas，如果出现Gas不足错误，可以在脚本中增加Gas限制。

## 手动操作（可选）

### 使用geth控制台
启动私链后，可以使用geth控制台直接与区块链交互：

```javascript
// 检查账户
eth.accounts

// 检查余额
eth.getBalance(eth.accounts[0])

// 发送交易
eth.sendTransaction({
  from: eth.accounts[0],
  to: "0x...",
  value: web3.toWei(1, "ether")
})
```

### 使用 Remix
1. 打开Remix IDE (https://remix.ethereum.org)
2. 导入`SimpleStorage.sol`
3. 编译合约
4. 在部署环境中选择"Web3 Provider"
5. 连接到 `http://localhost:8545`
6. 部署合约

## 清理数据

要完全重置私链，删除数据目录：

```bash
rm -rf private_chain_data
rm -f contract-info.json
```

## 许可证

MIT License
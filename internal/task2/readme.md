## P2P网络层

* 节点发现协议
* 会话与加密传输协议（RLPx）
* 子协议管理（如 eth/66、snap、les 等）

| 协议名                    | 模块位置                  | 功能描述                              |
| ---------------------- | --------------------- | --------------------------------- |
| **eth/62, 63, 65, 66** | `eth/protocols/eth/`  | 以太坊主协议：区块、交易、状态同步                 |
| **snap/1**             | `eth/protocols/snap/` | 快照同步，用于快速下载状态（PoS 后默认方式）          |
| **les/2**              | `les/`                | 轻节点协议（Light Ethereum Subprotocol） |
| **ethstats**           | `ethstats/`           | 节点状态上报                            |


### les 轻节点协议

#### 设计目标

让资源受限的客户端（手机、嵌入式设备、浏览器后端等）能以极小的存储/带宽代价访问以太坊数据（查询余额、交易、事件、调用合约读接口等）。
把“重数据与计算”放在full node / light server，客户端按需请求并通过Merkle/receipt/状态证明在本地验证返回的数据一致性。

#### 整体架构与角色

Light client（轻客户端）：只同步区块头、链的关键信息，按需向 full node 请求交易体、状态节点、存证等，并在本地用 header/stateRoot 等做验证。

Light server（LES server / full node with LES enabled）：完整保存数据并响应 light client 的请求。为了避免滥用，server 通常需要做资源限制（credit/serve quota），并非所有 full node 都自动充当 light server。
Ethereum Stack Exchange
+1

在 Geth 实现里，LES 是运行在 devp2p/RLPx 之上的一个子协议（即 les 子协议）


### 区块数据结构

#### 区块总体结构（Block Structure Overview）

在以太坊中，每个区块（Block）由三部分组成：

```ini
Block = BlockHeader + BlockBody + BlockHash
```


具体来说：

```
Block
 ├── Header       区块头（BlockHeader）
 ├── Body         区块体（BlockBody）
 │    ├── Transactions[]   交易列表
 │    └── Uncles[]         叔块列表（仅 PoW 有）
 └── Hash         通过 RLP 编码 Header 后 Keccak256 得到
```


在 Geth 源码中定义在：

> core/types/block.go

#### BlockHeader 结构详细解读

以太坊的区块头是所有共识、哈希计算、链连接的核心。

```go
type Header struct {
    ParentHash  common.Hash    // 父区块哈希
    UncleHash   common.Hash    // 叔块列表的哈希（RLP 编码后）
    Coinbase    common.Address // 矿工地址（PoW）或验证者地址（PoS）
    Root        common.Hash    // 状态树根（stateRoot）
    TxHash      common.Hash    // 交易树根（transactionsRoot）
    ReceiptHash common.Hash    // 收据树根（receiptsRoot）
    Bloom       Bloom          // 事件日志布隆过滤器
    Difficulty  *big.Int       // 挖矿难度（PoW），PoS中为零
    Number      *big.Int       // 区块高度
    GasLimit    uint64         // 当前块最大 gas 限额
    GasUsed     uint64         // 实际消耗的 gas
    Time        uint64         // 时间戳
    Extra       []byte         // 附加数据（可嵌入 32 字节字符串）
    MixDigest   common.Hash    // Ethash 随机数校验（PoW）
    Nonce       BlockNonce     // 工作量证明随机数（PoW）
    BaseFee     *big.Int       // EIP-1559 基础费用（London 之后）
}
```

各字段含义说明

| 字段                     | 说明                                                  |
| ---------------------- | --------------------------------------------------- |
| **ParentHash**         | 父区块的哈希，用于连接主链。                                      |
| **UncleHash**          | 所有叔块头的 RLP 编码哈希（仅 PoW）。                             |
| **Coinbase**           | 出块者地址（矿工奖励接收方或验证者地址）。                               |
| **Root**               | 世界状态树根（stateRoot），即所有账户状态 Merkle Patricia Trie 根哈希。 |
| **TxHash**             | 当前区块交易树根。                                           |
| **ReceiptHash**        | 所有交易收据的 Merkle 根。                                   |
| **Bloom**              | 事件日志过滤器，用于快速查找特定事件。                                 |
| **Difficulty**         | 当前出块难度（PoW）。PoS 已废弃此字段逻辑。                           |
| **Number**             | 区块号，从创世块（0）递增。                                      |
| **GasLimit / GasUsed** | Gas 限额与实际消耗值。                                       |
| **Time**               | 区块生成时间戳（秒）。                                         |
| **Extra**              | 附加字段，通常由矿池标识或共识层使用。                                 |
| **MixDigest / Nonce**  | 工作量证明验证参数（PoW 专用）。                                  |
| **BaseFee**            | EIP-1559 引入的基础 Gas 费。                               |


#### BlockBody 结构
```go
type Body struct {
  Transactions []*Transaction // 当前块中所有交易
  Uncles       []*Header      // 叔块头列表
}
```

Transactions：存储所有有效交易（包含签名、nonce、gasPrice 等）。

Uncles：在 PoW 共识中，落后 1–6 个区块但仍有效的候选块头。

作用：鼓励参与出块、提高去中心化。

奖励：叔块矿工能获得部分奖励（约 7/8）。

PoS（合并后）中没有 Uncles 概念。

#### 区块哈希计算方式

以太坊区块哈希定义如下：

```go
block.Hash() = Keccak256(RLP(BlockHeader))
```


只取 Header 做哈希；

任何区块体变化（交易内容）都会影响 TxRoot，从而改变 Header；

所以 Header Hash 就是区块唯一标识。

#### 存储关系（State / Tx / Receipt 三棵 Trie）

每个区块都存储三棵 Merkle Patricia Trie 树：

```go
StateTrie (stateRoot)
├─ Account[address].balance
├─ Account[address].nonce
├─ Account[address].storageRoot
└─ Account[address].codeHash

TransactionsTrie (txRoot)
└─ 每笔交易的 RLP 编码节点

ReceiptsTrie (receiptRoot)
└─ 每笔交易执行后的 Receipt 记录
```


这些 Trie 树最终都落在 LevelDB 中存储（chaindata 数据库）。
Geth 的数据库键值前缀包括：h（header）、b（body）、r（receipts）、s（state）。

#### PoW 与 PoS 的差异
| 项目                    | PoW（Ethash） | PoS（Beacon Chain / Merge 后）          |
| --------------------- | ----------- | ------------------------------------ |
| **共识算法**              | Ethash      | Beacon Chain（Casper FFG + LMD-GHOST） |
| **Difficulty 字段**     | 用于挖矿        | 恒为 0                                 |
| **Nonce / MixDigest** | 计算出块有效性     | 保留但无意义                               |
| **Coinbase**          | 矿工地址        | 验证者的奖励地址                             |
| **区块时间**              | 13–15s      | 12s 固定 Slot                          |
| **叔块机制**              | 有（Uncles）   | 无                                    |
| **共识验证方式**            | 重算工作量       | 验证签名与投票                              |


在 Geth 中，PoS 时代区块头结构保持兼容，只是共识引擎从 ethash → beacon。


## 区块链协议层

共识机制、交易处理、区块生成、区块验证


## 状态存储层

LevelDB、默克尔树、状态数据库


## EVM执行层

交易解释、合约执行、gas计算


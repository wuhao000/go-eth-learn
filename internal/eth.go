package internal

import (
  "context"
  "fmt"
  "go-eth-learn/internal/utils"
  "log"
  "math/big"
  "strings"
  "time"

  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/rpc"
)

// EthClient 封装以太坊客户端
type EthClient struct {
  client *ethclient.Client
}

// GetHeaderByNumber 根据区块号获取区块头
func (ec *EthClient) GetHeaderByNumber(blockNumber *big.Int) (*types.Header, error) {
  header, err := ec.client.HeaderByNumber(context.Background(), blockNumber)
  if err != nil {
    return nil, fmt.Errorf("获取区块头失败: %v", err)
  }
  return header, nil
}

// GetLatestHeader 获取最新区块头
func (ec *EthClient) GetLatestHeader() (*types.Header, error) {
  return ec.GetHeaderByNumber(nil)
}

// GetHeaderByHash 根据区块哈希获取区块头
func (ec *EthClient) GetHeaderByHash(blockHash common.Hash) (*types.Header, error) {
  header, err := ec.client.HeaderByHash(context.Background(), blockHash)
  if err != nil {
    return nil, fmt.Errorf("根据哈希获取区块头失败: %v", err)
  }
  return header, nil
}

// PrintHeaderInfo 打印区块头信息
func PrintHeaderInfo(header *types.Header) {
  fmt.Printf("=== 区块头信息 ===\n")
  fmt.Printf("区块号: %d\n", header.Number)
  fmt.Printf("区块哈希: %s\n", header.Hash().Hex())
  fmt.Printf("父区块哈希: %s\n", header.ParentHash.Hex())
  fmt.Printf("时间戳: %d\n", header.Time)
  fmt.Printf("矿工地址: %s\n", header.Coinbase.Hex())
  fmt.Printf("Gas限制: %d\n", header.GasLimit)
  fmt.Printf("已使用Gas: %d\n", header.GasUsed)
  fmt.Printf("区块大小: %s\n", header.Size().String())
  fmt.Printf("难度: %d\n", header.Difficulty)
  fmt.Printf("随机数: %d\n", header.Nonce)
  fmt.Printf("状态根哈希: %s\n", header.Root.Hex())
  fmt.Printf("交易根哈希: %s\n", header.TxHash.Hex())
  fmt.Printf("收据根哈希: %s\n", header.ReceiptHash.Hex())
  fmt.Printf("日志布隆过滤器: %x\n", header.Bloom)
  fmt.Printf("基础费用: %s\n", header.BaseFee.String())
  fmt.Printf("================\n")
}

// Close 关闭客户端连接
func (ec *EthClient) Close() {
  ec.client.Close()
}

var client *EthClient

// GetSepoliaClient 获取Sepolia测试网络客户端
func GetSepoliaClient() (*EthClient, error) {
  if client != nil {
    return client, nil
  }
  ethClient := utils.GetClient()
  client = &EthClient{
    client: ethClient,
  }
  log.Println("成功连接到Sepolia测试网络")
  return client, nil
}

// GetSepoliaHeaderByNumber 获取Sepolia网络指定区块号的区块头
func GetSepoliaHeaderByNumber(blockNumber int64) (*types.Header, error) {
  client, err := GetSepoliaClient()
  if err != nil {
    return nil, err
  }
  defer client.Close()

  return client.GetHeaderByNumber(big.NewInt(blockNumber))
}

// GetSepoliaLatestHeader 获取Sepolia网络最新区块头
func GetSepoliaLatestHeader() (*types.Header, error) {
  client, err := GetSepoliaClient()
  if err != nil {
    return nil, err
  }
  defer client.Close()

  return client.GetLatestHeader()
}

func GetPendingBalance(address common.Address) (*big.Int, *big.Float, error) {
  wei, err := client.client.PendingBalanceAt(context.Background(), address)
  ethValue := bigIntToFloat(wei)
  return wei, ethValue, err
}

// GetBalance 查询余额
func GetBalance(address common.Address) (*big.Int, *big.Float, error) {
  if client == nil {
    return nil, nil, fmt.Errorf("client not initialized")
  }

  wei, err := client.client.BalanceAt(context.Background(), address, nil)
  if err != nil {
    return nil, nil, fmt.Errorf("failed to get balance: %v", err)
  }

  ethValue := bigIntToFloat(wei)

  return wei, ethValue, nil
}

func SubscribeBlock() {
  // 使用现有的客户端连接进行轮询
  ethClient, err := GetSepoliaClient()
  if err != nil {
    log.Fatalf("Failed to connect to Ethereum node: %v", err)
  }
  defer ethClient.Close()

  fmt.Println("🚀 开始定时轮询新区块...")

  // 记录上一个区块号，用于检测新区块
  var lastBlockNumber *big.Int
  ticker := time.NewTicker(5 * time.Second) // 每5秒轮询一次
  defer ticker.Stop()

  for {
    select {
    case <-ticker.C:
      // 获取最新区块头
      header, err := ethClient.GetLatestHeader()
      if err != nil {
        log.Printf("获取区块头失败: %v", err)
        continue
      }

      // 检查是否是新区块
      if lastBlockNumber == nil || header.Number.Cmp(lastBlockNumber) > 0 {
        fmt.Printf("⛓️ 新区块 #%v, Hash: %s\n", header.Number.String(), header.Hash().Hex())
        lastBlockNumber = header.Number
      }
    }
  }
}

func bigIntToFloat(wei *big.Int) *big.Float {
  // Convert wei to eth
  fbalance := new(big.Float)
  fbalance.SetString(wei.String())
  ethValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
  return ethValue
}

// getReceipts 使用以太坊客户端获取区块收据并处理状态信息
func getReceipts() {
  receipts, err := client.client.BlockReceipts(context.Background(), rpc.BlockNumberOrHash{})
  if err != nil {
    log.Fatal(err)
  }
  for _, r := range receipts {
    // 打印收据信息
    fmt.Printf("状态 Status: %d\n", r.Status)
    fmt.Printf("累计Gas消耗 CumulativeGasUsed: %d\n", r.CumulativeGasUsed)
    fmt.Printf("布隆过滤器 Bloom: %x\n", r.Bloom)
    fmt.Printf("日志数量 Logs: %d\n", len(r.Logs))
    fmt.Printf("交易哈希 TxHash: %s\n", r.TxHash.String())
    fmt.Printf("合约地址 ContractAddress: %s\n", r.ContractAddress.String())
    fmt.Printf("Gas消耗 GasUsed: %d\n", r.GasUsed)
    fmt.Printf("区块哈希 BlockHash: %s\n", r.BlockHash.String())
    fmt.Printf("区块号 BlockNumber: %d\n", r.BlockNumber)
    fmt.Printf("交易索引 TransactionIndex: %d\n", r.TransactionIndex)
  }
}

// LogInfo 日志的可读信息
type LogInfo struct {
  EventName string
  Arguments map[string]interface{}
}

// ABICache 合约ABI缓存
type ABICache struct {
  contracts map[string]abi.ABI
}

var abiCache = &ABICache{
  contracts: make(map[string]abi.ABI),
}

// 标准ERC20 ABI
const erc20ABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "from", "type": "address"},
      {"indexed": true, "name": "to", "type": "address"},
      {"indexed": false, "name": "value", "type": "uint256"}
    ],
    "name": "Transfer",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "owner", "type": "address"},
      {"indexed": true, "name": "spender", "type": "address"},
      {"indexed": false, "name": "value", "type": "uint256"}
    ],
    "name": "Approval",
    "type": "event"
  }
]`

// 标准ERC721 ABI
const erc721ABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "from", "type": "address"},
      {"indexed": true, "name": "to", "type": "address"},
      {"indexed": true, "name": "tokenId", "type": "uint256"}
    ],
    "name": "Transfer",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "owner", "type": "address"},
      {"indexed": true, "name": "approved", "type": "address"},
      {"indexed": true, "name": "tokenId", "type": "uint256"}
    ],
    "name": "Approval",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "owner", "type": "address"},
      {"indexed": true, "name": "operator", "type": "address"},
      {"indexed": false, "name": "approved", "type": "bool"}
    ],
    "name": "ApprovalForAll",
    "type": "event"
  }
]`

// Uniswap V2 Pair ABI
const uniswapV2PairABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "sender", "type": "address"},
      {"indexed": false, "name": "amount0In", "type": "uint256"},
      {"indexed": false, "name": "amount1In", "type": "uint256"},
      {"indexed": false, "name": "amount0Out", "type": "uint256"},
      {"indexed": false, "name": "amount1Out", "type": "uint256"},
      {"indexed": true, "name": "to", "type": "address"}
    ],
    "name": "Swap",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": false, "name": "reserve0", "type": "uint256"},
      {"indexed": false, "name": "reserve1", "type": "uint256"}
    ],
    "name": "Sync",
    "type": "event"
  }
]`

// WETH ABI
const wethABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "sender", "type": "address"},
      {"indexed": false, "name": "amount", "type": "uint256"}
    ],
    "name": "Deposit",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "name": "src", "type": "address"},
      {"indexed": true, "name": "dst", "type": "address"},
      {"indexed": false, "name": "wad", "type": "uint256"}
    ],
    "name": "Withdrawal",
    "type": "event"
  }
]`

// GetStandardABI 根据合约地址获取标准ABI
func (cache *ABICache) GetStandardABI(contractAddress common.Address) (abi.ABI, error) {
  address := contractAddress.Hex()

  // 检查缓存
  if cachedABI, exists := cache.contracts[address]; exists {
    return cachedABI, nil
  }

  var abiString string
  var abiType string

  // 根据已知合约地址返回对应的ABI
  switch address {
  case "0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14": // WETH Sepolia Testnet
    abiString = wethABI
    abiType = "WETH"
  default:
    // 如果不是已知合约，尝试通过启发式方法判断
    // 在实际应用中，这里可以调用 Etherscan API 获取ABI
    // 为了演示，我们返回一个通用的ERC20 ABI
    abiString = erc20ABI
    abiType = "ERC20 (推测)"
  }

  parsedABI, err := abi.JSON(strings.NewReader(abiString))
  if err != nil {
    return abi.ABI{}, fmt.Errorf("解析%s ABI失败: %v", abiType, err)
  }

  // 缓存ABI
  cache.contracts[address] = parsedABI

  fmt.Printf("✅ 为合约 %s 加载了 %s ABI\n", address[:10]+"...", abiType)
  return parsedABI, nil
}

// ParseLogWithABI 使用ABI解析日志（更准确的方法）
func ParseLogWithABI(rlog *types.Log, contractABI abi.ABI) (*LogInfo, error) {
  if len(rlog.Topics) == 0 {
    return &LogInfo{
      EventName: "Unknown",
      Arguments: map[string]interface{}{
        "data": hexutil.Encode(rlog.Data),
      },
    }, nil
  }

  // 尝试使用ABI解析日志
  for _, event := range contractABI.Events {
    if event.ID == rlog.Topics[0] {
      // 找到匹配的事件
      parsedLog, err := contractABI.Unpack(event.Name, rlog.Data)
      if err != nil {
        return nil, fmt.Errorf("解析事件 %s 失败: %v", event.Name, err)
      }

      result := &LogInfo{
        EventName: event.Name,
        Arguments: make(map[string]interface{}),
      }

      // 将解析后的参数转换为可读格式
      for i, input := range event.Inputs {
        if i < len(parsedLog) {
          result.Arguments[input.Name] = formatValue(parsedLog[i], input.Type.String())
        }
      }

      return result, nil
    }
  }

  return nil, fmt.Errorf("ABI中未找到匹配的事件")
}

// formatValue 格式化值为可读字符串
func formatValue(value interface{}, typeStr string) interface{} {
  switch v := value.(type) {
  case *big.Int:
    if strings.Contains(typeStr, "uint256") || strings.Contains(typeStr, "uint") {
      // 对于大整数，返回字符串表示
      return v.String()
    }
    return v
  case common.Address:
    return v.Hex()
  case common.Hash:
    return v.Hex()
  case []byte:
    return hexutil.Encode(v)
  case [32]byte:
    return hexutil.Encode(v[:])
  default:
    return v
  }
}

// 基础的事件签名映射（只包含最重要的）
var basicEventSignatures = map[string]string{
  "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef":         "Transfer", // ERC20 Transfer
  "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925":         "Approval", // ERC20 Approval
  "0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526ed8d7f":         "OwnershipTransferred",
  "0x8be0079c5316591434f068f618e15c11f1e732ac65a7a11e35c7e9e566514f4a2":        "Paused",
  "0x809581d7e560f460c894b1b8f6f8b2a0f6a8a8a8a8a8a8a8a8a8a8a8a8a8a8a8a8a8a8a8": "Unpaused",
  "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822":         "Swap",       // Uniswap V3 Swap
  "0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1":         "Swap",       // Uniswap V2 Swap
  "0xe1bbbcc27279f29485ef3c967b5624118b740e6bda078679072e3a5d2a2b5d6":          "Deposit",    // WETH Deposit
  "0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65":         "Withdrawal", // WETH Withdrawal
}

// ParseLog 基础解析（使用签名映射）
func ParseLog(rlog *types.Log) *LogInfo {
  if len(rlog.Topics) == 0 {
    return &LogInfo{
      EventName: "Unknown",
      Arguments: map[string]interface{}{
        "data": hexutil.Encode(rlog.Data),
      },
    }
  }

  // 第一个主题通常是事件签名
  eventSignature := rlog.Topics[0].Hex()
  eventName := basicEventSignatures[eventSignature]

  if eventName == "" {
    eventName = "Unknown Event"
  }

  result := &LogInfo{
    EventName: eventName,
    Arguments: make(map[string]interface{}),
  }

  // 根据事件类型解析参数
  switch eventName {
  case "Transfer":
    if len(rlog.Topics) >= 3 {
      result.Arguments["from"] = rlog.Topics[1].Hex()
      result.Arguments["to"] = rlog.Topics[2].Hex()

      // 解析数据（金额）
      if len(rlog.Data) >= 32 {
        amount := new(big.Int).SetBytes(rlog.Data)
        result.Arguments["value"] = amount.String()
      }
    }

  case "Approval":
    if len(rlog.Topics) >= 3 {
      result.Arguments["owner"] = rlog.Topics[1].Hex()
      result.Arguments["spender"] = rlog.Topics[2].Hex()

      // 解析数据（金额）
      if len(rlog.Data) >= 32 {
        amount := new(big.Int).SetBytes(rlog.Data)
        result.Arguments["value"] = amount.String()
      }
    }

  case "Swap":
    // Swap事件可能有不同的签名，尝试通用解析
    if len(rlog.Topics) >= 3 {
      result.Arguments["sender"] = rlog.Topics[1].Hex()
      result.Arguments["recipient"] = rlog.Topics[2].Hex()

      // 解析数据中的金额信息
      if len(rlog.Data) >= 64 {
        amount0 := new(big.Int).SetBytes(rlog.Data[0:32])
        amount1 := new(big.Int).SetBytes(rlog.Data[32:64])
        result.Arguments["amount0"] = amount0.String()
        result.Arguments["amount1"] = amount1.String()
      }
    }

  case "Deposit":
    if len(rlog.Topics) >= 2 {
      result.Arguments["sender"] = rlog.Topics[1].Hex()
    }
    if len(rlog.Data) >= 32 {
      amount := new(big.Int).SetBytes(rlog.Data)
      result.Arguments["amount"] = amount.String()
    }

  case "Withdrawal":
    if len(rlog.Topics) >= 3 {
      result.Arguments["src"] = rlog.Topics[1].Hex()
      result.Arguments["dst"] = rlog.Topics[2].Hex()
    }
    if len(rlog.Data) >= 32 {
      amount := new(big.Int).SetBytes(rlog.Data)
      result.Arguments["amount"] = amount.String()
    }

  default:
    // 对于未知事件，显示原始主题和数据
    result.Arguments["topics"] = make([]string, len(rlog.Topics))
    for i, topic := range rlog.Topics {
      result.Arguments["topics"].([]string)[i] = topic.Hex()
    }
    result.Arguments["data"] = hexutil.Encode(rlog.Data)
  }

  return result
}

// AnalyzeEventSignature 分析事件签名，提供更严谨的解析方法
func AnalyzeEventSignature(eventSignature string, contractAddress common.Address) string {
  // 首先检查是否是常见的DeFi协议事件
  if eventSignature == "0xff48c13eda96b1cceacc6b9edeedc9e9db9d6226afbc30146b720c19d3addb1c" {
    // 这是Curve Finance的TokenExchange事件签名
    return "TokenExchange (Curve)"
  }

  // 检查其他已知的重要事件签名
  knownSignatures := map[string]string{
    // Uniswap V3相关
    "0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67": "Swap (Uniswap V3)",
    "0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1": "Swap (Uniswap V2)",
    "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822": "Swap (Uniswap V3)",

    // Curve相关
    "0xff48c13eda96b1cceacc6b9edeedc9e9db9d6226afbc30146b720c19d3addb1c": "TokenExchange (Curve)",
    "0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140": "AddLiquidity (Curve)",
    "0xdfb68d771a469df1a35ca2708fd5a4efd629199e24bf0470c8b0303ce1a0d2a9": "RemoveLiquidity (Curve)",

    // Aave相关
    "0x631042c832b07452973831137f2d73e395028b44b250dedc5abb0ee766e168ac": "Borrow (Aave)",
    "0xc41a360a802760c3c2d0a2dd5aa0c6a890ffb9cc90b11324c08f15b7bcce433":  "Repay (Aave)",
    "0x99cd89bce2a7ba3d0c6a5c2f3b0c8889c4ab70f12c3fc9e3aa1ba6661fd0c6e":  "Deposit (Aave)",
    "0x6a52787424a2ff4252b39640dca9a73267f144962fa5a8bbff6e8f3c7b5c86a":  "Withdraw (Aave/Compound)",

    // Compound相关
    "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925": "Approval",
    "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f": "Transfer (Compound)",
  }

  if eventName, exists := knownSignatures[eventSignature]; exists {
    return eventName
  }

  // 如果签名未知，尝试通过启发式方法分析
  return AnalyzeUnknownSignature(eventSignature, contractAddress)
}

// AnalyzeUnknownSignature 启发式分析未知签名
func AnalyzeUnknownSignature(eventSignature string, contractAddress common.Address) string {
  // 基于合约地址模式分析
  address := contractAddress.Hex()

  // 如果合约地址在已知的DeFi协议范围内
  if strings.HasPrefix(address, "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") {
    return "Unknown Event (Uniswap Router)"
  }

  // 基于签名的特征分析
  // 某些事件签名有特定的模式
  if strings.HasPrefix(eventSignature, "0xff48") {
    return "TokenExchange (Likely Curve Protocol)"
  }

  // 如果无法确定，返回签名前缀
  return fmt.Sprintf("Unknown Event (Signature: %s...)", eventSignature[:10])
}

// EnhancedParseLog 增强的日志解析函数（结合ABI和签名映射）
func EnhancedParseLog(rlog *types.Log) *LogInfo {
  // 首先尝试获取合约的ABI进行解析
  if contractABI, err := abiCache.GetStandardABI(rlog.Address); err == nil {
    if logInfo, err := ParseLogWithABI(rlog, contractABI); err == nil {
      // ABI解析成功，返回结果
      return logInfo
    }
  }

  // 如果ABI解析失败，使用更严谨的签名分析方法
  if len(rlog.Topics) > 0 {
    eventSignature := rlog.Topics[0].Hex()
    eventName := AnalyzeEventSignature(eventSignature, rlog.Address)

    // 尝试根据事件类型解析参数
    return ParseLogByEventType(rlog, eventName)
  }

  // 最后回退到基础解析
  return ParseLog(rlog)
}

// ParseLogByEventType 根据事件类型解析参数
func ParseLogByEventType(rlog *types.Log, eventName string) *LogInfo {
  result := &LogInfo{
    EventName: eventName,
    Arguments: make(map[string]interface{}),
  }

  // 根据事件名称进行解析
  if strings.Contains(eventName, "TokenExchange") {
    // Curve TokenExchange事件解析
    return ParseCurveTokenExchange(rlog)
  } else if strings.Contains(eventName, "Swap") {
    // 通用Swap事件解析
    return ParseSwapEvent(rlog)
  } else if strings.Contains(eventName, "Transfer") {
    // 通用Transfer事件解析
    return ParseTransferEvent(rlog)
  }

  // 默认返回基础信息
  result.Arguments["topics"] = make([]string, len(rlog.Topics))
  for i, topic := range rlog.Topics {
    result.Arguments["topics"].([]string)[i] = topic.Hex()
  }
  result.Arguments["data"] = hexutil.Encode(rlog.Data)

  return result
}

// ParseCurveTokenExchange 解析Curve TokenExchange事件
func ParseCurveTokenExchange(rlog *types.Log) *LogInfo {
  result := &LogInfo{
    EventName: "TokenExchange (Curve)",
    Arguments: make(map[string]interface{}),
  }

  // TokenExchange事件的签名是第一个topic
  if len(rlog.Topics) >= 4 {
    // Topic[0]: 事件签名
    // Topic[1]: buyer (address, indexed)
    // Topic[2]: sold_id (uint128, indexed)
    // Topic[3]: tokens_sold (uint128, indexed)

    buyer := common.BytesToAddress(rlog.Topics[1].Bytes())
    soldId := new(big.Int).SetBytes(rlog.Topics[2].Bytes())
    tokensSold := new(big.Int).SetBytes(rlog.Topics[3].Bytes())

    result.Arguments["buyer"] = buyer.Hex()
    result.Arguments["sold_id"] = soldId.String()
    result.Arguments["tokens_sold"] = tokensSold.String()

    // Data部分包含:
    // tokens_bought (uint128)
    // bought_id (uint128)
    if len(rlog.Data) >= 64 {
      tokensBought := new(big.Int).SetBytes(rlog.Data[0:32])
      boughtId := new(big.Int).SetBytes(rlog.Data[32:64])

      result.Arguments["tokens_bought"] = tokensBought.String()
      result.Arguments["bought_id"] = boughtId.String()
    }
  }

  return result
}

// ParseSwapEvent 解析通用Swap事件
func ParseSwapEvent(rlog *types.Log) *LogInfo {
  result := &LogInfo{
    EventName: "Swap",
    Arguments: make(map[string]interface{}),
  }

  if len(rlog.Topics) >= 3 {
    sender := common.BytesToAddress(rlog.Topics[1].Bytes())
    recipient := common.BytesToAddress(rlog.Topics[2].Bytes())

    result.Arguments["sender"] = sender.Hex()
    result.Arguments["recipient"] = recipient.Hex()

    // 解析data中的金额信息
    if len(rlog.Data) >= 128 {
      amount0 := new(big.Int).SetBytes(rlog.Data[0:32])
      amount1 := new(big.Int).SetBytes(rlog.Data[32:64])
      sqrtPriceX96 := new(big.Int).SetBytes(rlog.Data[64:96])
      liquidity := new(big.Int).SetBytes(rlog.Data[96:128])

      result.Arguments["amount0"] = amount0.String()
      result.Arguments["amount1"] = amount1.String()
      result.Arguments["sqrtPriceX96"] = sqrtPriceX96.String()
      result.Arguments["liquidity"] = liquidity.String()
    }
  }

  return result
}

// ParseTransferEvent 解析通用Transfer事件
func ParseTransferEvent(rlog *types.Log) *LogInfo {
  result := &LogInfo{
    EventName: "Transfer",
    Arguments: make(map[string]interface{}),
  }

  if len(rlog.Topics) >= 3 {
    from := common.BytesToAddress(rlog.Topics[1].Bytes())
    to := common.BytesToAddress(rlog.Topics[2].Bytes())

    result.Arguments["from"] = from.Hex()
    result.Arguments["to"] = to.Hex()

    // 解析data中的金额
    if len(rlog.Data) >= 32 {
      amount := new(big.Int).SetBytes(rlog.Data)
      result.Arguments["value"] = amount.String()
    }
  }

  return result
}

// EnhancedFormatLogInfo 增强的日志格式化函数
func EnhancedFormatLogInfo(rlog *types.Log) string {
  logInfo := EnhancedParseLog(rlog)

  var result strings.Builder
  result.WriteString(fmt.Sprintf("事件: %s", logInfo.EventName))

  if len(logInfo.Arguments) > 0 {
    result.WriteString(" (")
    first := true
    for key, value := range logInfo.Arguments {
      if !first {
        result.WriteString(", ")
      }
      first = false

      switch v := value.(type) {
      case string:
        if len(v) > 42 && strings.HasPrefix(v, "0x") {
          // 可能是地址，截断显示
          result.WriteString(fmt.Sprintf("%s: %s...", key, v[:10]))
        } else if len(v) > 20 {
          // 长字符串截断显示
          result.WriteString(fmt.Sprintf("%s: %s...", key, v[:17]))
        } else {
          result.WriteString(fmt.Sprintf("%s: %s", key, v))
        }
      default:
        result.WriteString(fmt.Sprintf("%s: %v", key, v))
      }
    }
    result.WriteString(")")
  }

  return result.String()
}

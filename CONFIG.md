# 配置说明

## 环境变量配置

本项目使用环境变量来配置 Infura API Key 等敏感信息。

### 1. 配置文件设置

1. 复制示例配置文件：
```bash
cp .env.example .env
```

2. 编辑 `.env` 文件，填入你的 Infura API Key：
```
INFURA_API_KEY=your_infura_api_key_here
INFURA_SEPOLIA_URL=https://sepolia.infura.io/v3
```

### 2. 环境变量说明

- `INFURA_API_KEY`: 你的 Infura API Key（必需）
- `INFURA_SEPOLIA_URL`: Sepolia 测试网的 RPC URL（可选，默认为 https://sepolia.infura.io/v3）

### 3. 获取 Infura API Key

1. 访问 [Infura 官网](https://infura.io/)
2. 注册并登录账户
3. 创建新项目
4. 复制项目的 API Key

### 4. 安全注意事项

- ⚠️ **重要**: `.env` 文件已添加到 `.gitignore`，不会被提交到版本控制
- 🔒 **不要**将包含真实 API Key 的 `.env` 文件分享给他人
- 📝 如需分享配置，请使用 `.env.example` 文件
- 🔄 定期更换你的 API Key 以确保安全

### 5. 代码中使用

配置会在程序启动时自动加载：

```go
// main.go 中会自动加载配置
config.InitConfig()

// 在其他模块中使用
rpcURL := config.GetSepoliaRPCURL()

// 或者使用新的配置结构
rpcURL := config.GlobalConfig.GetDefaultRPCURL()
```
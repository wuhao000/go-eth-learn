package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// 诚实声明：此函数只返回真实数据，绝不生成假数据
// 如果无法获取真实数据，将返回明确的错误信息

// 应用配置
type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
}

// Infura配置
type InfuraConfig struct {
	APIKey      string `yaml:"api_key"`
	SepoliaURL  string `yaml:"sepolia_url"`
	MainnetURL  string `yaml:"mainnet_url"`
	GoerliURL   string `yaml:"goerli_url"`
}

// 网络配置
type NetworkConfig struct {
	DefaultNetwork string `yaml:"default_network"`
	Timeout        int    `yaml:"timeout"`
}

// 日志配置
type LogConfig struct {
	Level      string `yaml:"level"`
	File       string `yaml:"file"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

// 服务器配置
type ServerConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

// 完整配置结构体
type Config struct {
	App      AppConfig      `yaml:"app"`
	Infura   InfuraConfig   `yaml:"infura"`
	Network  NetworkConfig  `yaml:"network"`
	Log      LogConfig      `yaml:"log"`
	Server   ServerConfig   `yaml:"server"`
}

var GlobalConfig *Config

// 加载配置文件
func LoadConfig() error {
	configPath := getConfigPath()

	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析YAML
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 从环境变量覆盖API Key（如果存在）
	if apiKey := os.Getenv("INFURA_API_KEY"); apiKey != "" {
		config.Infura.APIKey = apiKey
	}

	// 验证必要的配置
	if config.Infura.APIKey == "" {
		return fmt.Errorf("Infura API Key 未配置，请在配置文件中设置或通过环境变量 INFURA_API_KEY 设置")
	}

	GlobalConfig = config
	log.Printf("配置加载成功，环境: %s，网络: %s", config.App.Environment, config.Network.DefaultNetwork)

	return nil
}

// 获取配置文件路径
func getConfigPath() string {
	// 优先级：环境变量 > 当前目录 > 向上查找 > 默认路径
	if configPath := os.Getenv("CONFIG_PATH"); configPath != "" {
		return configPath
	}

	// 检查当前目录
	if _, err := os.Stat("config.yml"); err == nil {
		return "config.yml"
	}

	// 向上查找项目根目录（支持测试环境）
	if wd, err := os.Getwd(); err == nil {
		searchPaths := []string{
			filepath.Join(wd, "config.yml"),                     // 当前目录/config.yml
			filepath.Join(filepath.Dir(wd), "config.yml"),       // 上级目录/config.yml
			filepath.Join(filepath.Dir(filepath.Dir(wd)), "config.yml"), // 上上级目录/config.yml
		}

		for _, path := range searchPaths {
			if _, err := os.Stat(path); err == nil {
				return path
			}
		}
	}

	// 最后尝试绝对路径（基于go.mod位置）
	if configPath := findConfigNearGoMod(); configPath != "" {
		return configPath
	}

	return "config.yml" // 默认路径
}

// 在go.mod附近查找配置文件
func findConfigNearGoMod() string {
	if wd, err := os.Getwd(); err == nil {
		// 从当前目录开始向上查找go.mod文件
		currentDir := wd
		for currentDir != filepath.Dir(currentDir) { // 直到根目录
			goModPath := filepath.Join(currentDir, "go.mod")
			if _, err := os.Stat(goModPath); err == nil {
				// 找到go.mod，检查同目录的config.yml
				configPath := filepath.Join(currentDir, "config.yml")
				if _, err := os.Stat(configPath); err == nil {
					return configPath
				}
			}
			currentDir = filepath.Dir(currentDir)
		}
	}
	return ""
}

// 获取完整的 RPC URL
func (c *Config) GetRPCURL(network string) string {
	switch network {
	case "mainnet":
		return c.Infura.MainnetURL + "/" + c.Infura.APIKey
	case "goerli":
		return c.Infura.GoerliURL + "/" + c.Infura.APIKey
	case "sepolia":
		return c.Infura.SepoliaURL + "/" + c.Infura.APIKey
	default:
		return c.Infura.SepoliaURL + "/" + c.Infura.APIKey
	}
}

// 获取默认网络RPC URL
func (c *Config) GetDefaultRPCURL() string {
	return c.GetRPCURL(c.Network.DefaultNetwork)
}

// 服务器地址
func (c *Config) GetServerAddr() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// 兼容性函数 - 保持向后兼容
func GetSepoliaRPCURL() string {
	if GlobalConfig != nil {
		return GlobalConfig.GetRPCURL("sepolia")
	}
	log.Fatal("配置未加载，请先调用 LoadConfig()")
	return ""
}

// 初始化配置（兼容性函数）
func InitConfig() {
	err := LoadConfig()
	if err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}
}
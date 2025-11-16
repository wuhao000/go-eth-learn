const fs = require('fs');
const solc = require('solc');
const { Web3 } = require('web3');

async function deploy() {
    try {
        console.log("=== 部署智能合约到私链 ===");

        // 连接到私链
        const web3 = new Web3('http://localhost:8545');
        console.log("连接到私链: http://localhost:8545");

        // 检查连接
        try {
            const blockNumber = await web3.eth.getBlockNumber();
            console.log("连接成功，当前区块:", blockNumber);
        } catch (error) {
            console.error("无法连接到私链，请确保私链正在运行:", error.message);
            process.exit(1);
        }

        // 读取合约源码
        const contractSource = fs.readFileSync('SimpleStorage.sol', 'utf8');
        console.log("读取合约源码成功");

        // 编译合约
        const input = {
            language: 'Solidity',
            sources: {
                'SimpleStorage.sol': {
                    content: contractSource
                }
            },
            settings: {
                outputSelection: {
                    '*': {
                        '*': ['*']
                    }
                }
            }
        };

        const output = JSON.parse(solc.compile(JSON.stringify(input)));

        if (output.errors) {
            console.error("编译错误:");
            output.errors.forEach(err => {
                console.error(err.formattedMessage);
            });
            process.exit(1);
        }

        const contractBytecode = output.contracts['SimpleStorage.sol']['SimpleStorage'].evm.bytecode.object;
        const contractAbi = output.contracts['SimpleStorage.sol']['SimpleStorage'].abi;

        console.log("合约编译成功");
        console.log("合约字节码长度:", contractBytecode.length);

        // 获取账户
        const accounts = await web3.eth.getAccounts();
        console.log("可用账户数量:", accounts.length);

        if (accounts.length === 0) {
            console.error("没有可用账户");
            process.exit(1);
        }

        const deployer = accounts[0];
        console.log("部署账户:", deployer);

        // 检查余额
        const balance = await web3.eth.getBalance(deployer);
        console.log("账户余额:", web3.utils.fromWei(balance, 'ether'), "ETH");

        // 创建合约实例
        const contract = new web3.eth.Contract(contractAbi);

        // 部署合约，初始值为42
        console.log("开始部署合约...");
        const deployTx = contract.deploy({
            data: '0x' + contractBytecode,
            arguments: [42] // 构造函数参数，初始值为42
        });

        const estimatedGas = await deployTx.estimateGas({ from: deployer });
        console.log("预估Gas:", estimatedGas);

        const deployedContract = await deployTx.send({
            from: deployer,
            gas: Number(estimatedGas) + 100000, // 转换为Number并添加一些额外的gas
            gasPrice: await web3.eth.getGasPrice()
        });

        console.log("=== 合约部署成功 ===");
        console.log("合约地址:", deployedContract.options.address);
        console.log("交易哈希:", deployedContract.transactionHash);
        console.log("区块号:", deployedContract.blockNumber);

        // 保存合约信息
        const contractInfo = {
            address: deployedContract.options.address,
            abi: contractAbi,
            transactionHash: deployedContract.transactionHash,
            blockNumber: deployedContract.blockNumber
        };

        fs.writeFileSync('contract-info.json', JSON.stringify(contractInfo, null, 2));
        console.log("合约信息已保存到 contract-info.json");

        // 测试合约
        console.log("\n=== 测试合约功能 ===");

        // 读取初始值
        let initialValue = await deployedContract.methods.get().call();
        console.log("初始值:", initialValue);

        // 设置新值
        console.log("设置新值为 100...");
        const setTx = await deployedContract.methods.set(100).send({
            from: deployer,
            gas: 100000
        });
        console.log("设置交易哈希:", setTx.transactionHash);

        // 读取新值
        let newValue = await deployedContract.methods.get().call();
        console.log("新值:", newValue);

        // 测试increment
        console.log("调用 increment()...");
        const incTx = await deployedContract.methods.increment().send({
            from: deployer,
            gas: 100000
        });
        console.log("增量交易哈希:", incTx.transactionHash);

        let incrementedValue = await deployedContract.methods.get().call();
        console.log("增量后的值:", incrementedValue);

        console.log("\n=== 部署和测试完成 ===");

    } catch (error) {
        console.error("部署过程中出错:", error);
        process.exit(1);
    }
}

deploy();
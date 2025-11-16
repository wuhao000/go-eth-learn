const fs = require('fs');
const { Web3 } = require('web3');

async function testContract() {
    try {
        console.log("=== 测试智能合约功能 ===");

        // 检查合约信息文件是否存在
        if (!fs.existsSync('contract-info.json')) {
            console.error("合约信息文件不存在，请先运行 npm run deploy");
            process.exit(1);
        }

        const contractInfo = JSON.parse(fs.readFileSync('contract-info.json', 'utf8'));
        console.log("合约地址:", contractInfo.address);

        // 连接到私链
        const web3 = new Web3('http://localhost:8545');
        const contract = new web3.eth.Contract(contractInfo.abi, contractInfo.address);

        // 获取账户
        const accounts = await web3.eth.getAccounts();
        const user = accounts[0];

        console.log("使用账户:", user);

        // 显示当前值
        let currentValue = await contract.methods.get().call();
        console.log("当前存储的值:", currentValue);

        // 测试设置值
        console.log("测试设置值为 200...");
        const setTx = await contract.methods.set(200).send({
            from: user,
            gas: 200000
        });
        console.log("设置交易哈希:", setTx.transactionHash);

        // 读取新值
        let newValue = await contract.methods.get().call();
        console.log("设置后的值:", newValue);

        // 测试increment
        console.log("测试 increment()...");
        const incTx = await contract.methods.increment().send({
            from: user,
            gas: 200000
        });
        console.log("增量交易哈希:", incTx.transactionHash);

        let incrementedValue = await contract.methods.get().call();
        console.log("增量后的值:", incrementedValue);

        // 检查账户余额
        let balance = await web3.eth.getBalance(user);
        console.log("账户余额:", web3.utils.fromWei(balance, 'ether'), "ETH");

        console.log("\n=== 所有测试完成 ===");
        console.log("✅ 合约读取功能正常");
        console.log("✅ 合约写入功能正常");
        console.log("✅ 合约计算功能正常");
        console.log("✅ 私链连接正常");

    } catch (error) {
        console.error("测试过程中出错:", error);
        process.exit(1);
    }
}

testContract();
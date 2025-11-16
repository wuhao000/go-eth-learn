const fs = require('fs');
const { Web3 } = require('web3');

async function interact() {
    try {
        console.log("=== 与已部署的合约交互 ===");

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

        // 交互菜单
        const readline = require('readline');
        const rl = readline.createInterface({
            input: process.stdin,
            output: process.stdout
        });

        function showMenu() {
            console.log("\n=== 交互菜单 ===");
            console.log("1. 读取当前值");
            console.log("2. 设置新值");
            console.log("3. 增加1");
            console.log("4. 查看账户余额");
            console.log("5. 退出");
            rl.question("请选择操作 (1-5): ", async (choice) => {
                switch (choice) {
                    case '1':
                        let value = await contract.methods.get().call();
                        console.log("当前值:", value);
                        break;
                    case '2':
                        rl.question("请输入要设置的值: ", async (input) => {
                            try {
                                const numValue = parseInt(input);
                                if (isNaN(numValue)) {
                                    console.log("请输入有效的数字");
                                    return showMenu();
                                }
                                console.log("设置值为:", numValue);
                                const tx = await contract.methods.set(numValue).send({
                                    from: user,
                                    gas: 100000
                                });
                                console.log("交易哈希:", tx.transactionHash);
                                console.log("设置成功！");
                            } catch (error) {
                                console.error("设置失败:", error.message);
                            }
                            showMenu();
                        });
                        return;
                    case '3':
                        try {
                            console.log("增加1...");
                            const tx = await contract.methods.increment().send({
                                from: user,
                                gas: 100000
                            });
                            console.log("交易哈希:", tx.transactionHash);
                            console.log("增加成功！");
                        } catch (error) {
                            console.error("增加失败:", error.message);
                        }
                        break;
                    case '4':
                        let balance = await web3.eth.getBalance(user);
                        console.log("账户余额:", web3.utils.fromWei(balance, 'ether'), "ETH");
                        break;
                    case '5':
                        console.log("退出");
                        rl.close();
                        return;
                    default:
                        console.log("无效选择");
                }
                showMenu();
            });
        }

        showMenu();

    } catch (error) {
        console.error("交互过程中出错:", error);
        process.exit(1);
    }
}

interact();
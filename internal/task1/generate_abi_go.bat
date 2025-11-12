solcjs --abi contracts/task3/Auction.sol --base-path . --include-path node_modules -o build

abigen --abi=contracts_task3_IAuction_sol_IAuction.abi --pkg=auction --out=auction.go
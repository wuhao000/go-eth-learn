// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    uint256 private storedData;

    event DataStored(uint256 newValue, address indexed storedBy);

    constructor(uint256 initialValue) {
        storedData = initialValue;
        emit DataStored(initialValue, msg.sender);
    }

    function set(uint256 x) public {
        storedData = x;
        emit DataStored(x, msg.sender);
    }

    function get() public view returns (uint256) {
        return storedData;
    }

    function increment() public {
        storedData++;
        emit DataStored(storedData, msg.sender);
    }
}
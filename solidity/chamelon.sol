pragma solidity ^0.8.0;

import "ownable.sol";

contract Chameleon is Ownable {
    uint256 _genesis = 0x0;
    string private _bootstrap;

    constructor () {

    }

    event dataCreated(uint256 key, string random);
    event dataModified(uint256 key, string random);
    event dataRevoked(uint256 key);

    // 创建数据信息， 触发合约事件
    function create(uint256 key, string calldata random) public {
        emit dataCreated(key, random);
    }

    // 修改数据, 触发合约事件
    function modify(uint256 key, string calldata random) public {
        emit dataModified(key, random);
    }

    // 撤销数据
    function revoke(uint256 key) public {
        emit dataRevoked(key);
    }

    // 初始化合约, 设置genesis字段用于标识初始数据的索引
    function initialization(uint256 genesisKey, string calldata random, string calldata newBootstrap) public {
        require(_genesis == 0x0, "Genesis Data Initialed");
        _genesis = genesisKey;
        _bootstrap = newBootstrap;
        emit dataCreated(genesisKey, random);
    }

    // 读取启动节点信息
    function bootstrap() public view returns (string memory) {
        return _bootstrap;
    }

    function setBootstrap(string calldata new_bootstrap) public onlyOwner {
        _bootstrap = new_bootstrap;
    }

    function genesis() public view returns (uint256) {
        return _genesis;
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ODSC {
    uint256 _genesis = 0x0;
    string private _bootstrap;

    constructor () {

    }

    event dataCreated(uint256 key, bytes random);
    event dataModified(uint256 key, bytes random);
    event dataRevoked(uint256 key);

    // create data and emit dataCreated event
    function create(uint256 key, bytes calldata random) public {
        emit dataCreated(key, random);
    }

    // modify data and emit dataModified event
    function modify(uint256 key, bytes calldata random) public {
        emit dataModified(key, random);
    }

    // revoke data and emit dataRevoked event
    function revoke(uint256 key) public {
        emit dataRevoked(key);
    }

    // initilization, set genesis and bootstrap information
    function initialization(uint256 genesisKey, bytes calldata random, string calldata newBootstrap) public {
        require(_genesis == 0x0, "Genesis Data Initialed");
        _genesis = genesisKey;
        _bootstrap = newBootstrap;
        emit dataCreated(genesisKey, random);
    }

    // read bootstrap node connection information
    function bootstrap() public view returns (string memory) {
        return _bootstrap;
    }

    // set/reset bootstrap information
    function setBootstrap(string calldata new_bootstrap) public {
        _bootstrap = new_bootstrap;
    }

    // read genesis data index
    function genesis() public view returns (uint256) {
        return _genesis;
    }
}

pragma solidity ^0.8.0;

import "./ownable.sol";

contract Chameleon is Ownable {

    address owner;
    string genesis = 0x0;

    constructor () {

    }

    event dataCreated(string key, string random);
    event dataModified(string key, string random);
    event dataRevoked(string key);

    function create(string calldata key, string calldata random) public {
        emit dataCreated(key, random);
    }

    function modify(string calldata key, string calldata random) public {
        emit dataModified(key, random);
    }

    function revoke(string calldata key) public {
        emit dataRevoked(key);
    }
}

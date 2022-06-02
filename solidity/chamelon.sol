pragma solidity ^0.8.0;

contract chamelon {
    constructor () {

    }

    event dataCreated(string key, string random);
    event dataModified(string key, string random);

    function create(string calldata key, string calldata random) public {
        emit dataCreated(key, random);
    }

    function modified(string calldata key, string calldata random) public {
        emit dataModified(key, random);
    }
}

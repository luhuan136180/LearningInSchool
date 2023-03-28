pragma solidity ^0.8.4;

contract map{
    struct Student{
        uint256 id;
        uint256 score;

    }



    mapping(uint =>address) public idToAddress;
    mapping(address => address) public swapPair;
    function writeMap(uint _key,address _Value)public{
        idToAddress[_key] = _Value;
    }
}
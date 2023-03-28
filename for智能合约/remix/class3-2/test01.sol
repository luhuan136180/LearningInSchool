pragma solidity ^0.8.16

contract AddressDemo{
    address dataAddress = 0x

    function addressDemoOne(byte32 data) public pure returns (address){
        address addressConverted = address(bytes20(data));
        return addressConverted;
    }


    function  addressDemoTwo(byte32 data)public pure returns (address){
        address addressConverted = address()
    }
}
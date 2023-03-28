

pragma solidity >=0.4.22 <0.9.0;

contract helloworld {
    uint public balance;

    function sayHelloWorld() public pure returns (string memory){
        return ("Hello World! My name is Mao Haoxin");
    }
}
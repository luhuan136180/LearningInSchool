pragma solidity ^0.8.4;

contract constructors{
    address public owner;
       // 构造函数
    constructor(){
        owner = msg.sender;// 在部署合约的时候，将owner设置为部署者的地址

    }
       // 定义modifier
   modifier onlyOwner {
      require(msg.sender == owner); // 检查调用者是否为owner地址
      _; // 如果是的话，继续运行函数主体；否则报错并revert交易
   }
    function changeOwner(address _newOwner) external onlyOwner{
      owner = _newOwner; // 只有owner地址运行这个函数，并改变owner
   }

//    定义了一个changeOwner函数，运行他可以改变合约的owner，
// 但是由于onlyOwner修饰符的存在，
// 只有原先的owner可以调用，别人调用就会报错。

}
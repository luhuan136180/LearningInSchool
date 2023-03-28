pragma solidity >=0.4.16 < 0.9.0;
contract Array{
    uint[8] array1;
    bytes[5] array2;
    address[100] array3;

    uint[] array4;
    bytes[] array5;
    address[] array6;

    // uint[] memory array8 = new uint[](5);
    // bytes memory array9 = new bytes[9];


    function f() public pure{
        uint[] memory x = new uint[](3);
        x[0]=1;
        x[1]=3;
        x[2] = 4;
        g([uint(1),2,3]);
    }

    function g(uint[3] memory) public pure{

    } 
    function arrayPush() public returns (uint[] memory){
        uint[2] memory a = [uint(1),2];
        array4=a;
        array4.push(3);
        return array4;
    }

    struct Student{
        uint256 id;
        uint256 score;
    }
    Student student;//初始化一个student结构体
    function initStudent1() external{
        Student storage _student = student;
        _student.id = 11;
        _student.score = 100;
    }

 // 方法2:直接引用状态变量的struct
    function initStudent2() external{
        student.id =1;
        student.score = 80;
    }

}
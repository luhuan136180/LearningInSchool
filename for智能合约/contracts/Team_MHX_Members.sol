//源码授权协议 
//SPDX-License-Identifier: GPL-3.0

//Solidity 版本指定
pragma solidity >=0.7.0 <0.9.0;

//导入外部引用
import "hardhat/console.sol";

//合约主体
contract Team_XX {

    //成员结构体
    struct Member{
        string name;
        uint age;
    }
    
	//成员列表
    Member[]  members;

    // 添加成员
    function addMember(string memory name,uint age) public  {
        members.push(Member({name:name,age:age}));
    }

    //输出成员信息
    function printMembers() public view {
        for (uint i=0;i<members.length;i++){
            console.log("name:",members[i].name);
            console.log("age:",members[i].age);
        }
    }
    //输出成员平均年龄
    function printCalculateAvgAge() public view {
        uint ages=0;
        for(uint i=0;i<members.length;i++){
            ages+=members[i].age;
        }
        console.log("avg age:",ages/members.length);
    }
}
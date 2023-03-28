pragma solidity ^0.8.4;

contract IfElse{
    function ifElseTest(uint256 _number) public pure returns(bool){
        if(_number == 0){
            return(true);
        }else{
            return(false);
        }
    }

    function forLoopTest() public pure returns(uint256){
        uint sum  = 0;
        for(uint i = 0;i<10;i++){
            sum+=i;
        }
        return(sum);
    }
    function whileTest() public pure returns(uint256){
        uint sum = 0;
        uint i = 0;
        while(i < 10){
            sum+=i;
            i++;
        }
        return(sum);
    }
    function doWhileTest() public pure returns(uint256){
        uint sum = 0;
        uint i = 0;
        do{
        sum += i;
        i++;
        }while(i < 10);
        return(sum);
    }
    // 插入排序 正确版
    function insertionSort(uint[] memory a) public pure returns(uint[] memory) {
        // note that uint can not take negative value
        for (uint i = 1;i < a.length;i++){
            uint temp = a[i];
            uint j=i;
            while( (j >= 1) && (temp < a[j-1])){
                a[j] = a[j-1];
                j--;
            }
            a[j] = temp;
        }
        return(a);
    }

}


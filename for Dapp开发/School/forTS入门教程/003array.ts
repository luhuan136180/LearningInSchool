// 数组的类型
let fibonacci :number[] = [1,2,3,4,5];

fibonacci.push(34);

let fibonacci2 :Array<number> =[2,2,3,2];

// 用接口表示数组
interface NumberArray {
    [index: number]: number;
}
let fibonacci3: NumberArray = [1, 1, 2, 3, 5];


function sum2(){
    let args :{
        [index: number]: number;
        length: number;
        callee: Function;
    } = arguments;
}

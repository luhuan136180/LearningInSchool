interface Person {
    name:string;
    age:number;
}

let tom:Person={
    name:'Tom',
    age:25
};

//定义的变量必须与接口的属性相同

interface Person2 {
    name :string;
    age ?:number; 
}
// '?'表示该属性为可选择属性，可以不赋值
let jack:Person2={
    name : 'jack'
};


    //只读属性
    // 只能在创建的时候被赋值
interface Person3{
    readonly id :number;
    name :string;
    age ?: number;
    [propName: string]: any;
}
let alice: Person3 = {
    id: 89757,
    name: 'Tom',
    gender: 'male'
};



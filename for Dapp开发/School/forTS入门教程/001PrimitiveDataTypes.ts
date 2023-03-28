let isDone : boolean = false;
// 使用构造函数 Boolean 创造的对象不是布尔值
let createdByNewBoolean : Boolean = new Boolean(1);
let createdByBoolean : boolean = Boolean(1);
    // 数值
let decliteral : number = 6;
let hexliteral : number = 0xf00d;
// ES6 中的二进制表示法
let binaryLiteral :number = 0b1010;
// ES6 中的八进制表示法
let octalLiteral: number = 0o744;

    // 字符串
let myName : string = 'Tom';
let myAge : number = 25;

//模板字符串
let sentence : string = `Hello,my name is ${myName}.
I will be ${myAge + 1} years old next month`;
// ` 用来定义 ES6 中的模板字符串，
// ${expr} 用来在模板字符串中嵌入表达式

    // 空值

// 用 void 表示没有任何返回值的函数
function alertName():void{
    alert(`My name is Tom`)
}
// undefined 和 null 是所有类型的子类型。
// 也就是说 undefined 类型的变量，
// 可以赋值给 number 类型的变量

    // any 类型，则允许被赋值为任意类型
let myFavoriteNumber :any = 'seven';
myFavoriteNumber = 5;
// 变量如果在声明的时候，
// 未指定其类型，那么它会被识别为任意值类型：


    //类型推论：未指定雷星时，
    // TypeScript 会依照类型推论（Type Inference）
    // 的规则推断出一个类型
let myFavoriteNUmber2 = 'seven';
//此处寄予指定类型 string
// myFavoriteNUmber2 = 5;  不能将类型“number”分配给类型“string”



    //  联合类型
let myNumber3 :string|number ;
myNumber3 = 'seven';
myNumber3 = 62;


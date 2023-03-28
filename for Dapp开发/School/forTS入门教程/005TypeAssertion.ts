//语法：<类型>值
//      值 as 类型

interface Cat{
    name:string;
    run():void;
}

interface Fish{
    name:string;
    swim():void;
}

//TypeScript 不确定一个联合类型的变量到底是
// 哪个类型的时候，
// 我们只能访问此联合类型的
// 所有类型中共有的属性或方法： 
function getName(animal:Cat | Fish){
    return animal.name
}


// function isFish(animal: Cat | Fish) {
//     if (typeof animal.swim === 'function') {
//         return true;
//     }
//     return false;
// }
//以上这段代码，错误；
// 不可在还不确定类型的时候就访问其中一个类型特有的属性或方法

// 针对以上情况，我们需要将animal 断言成 Fish
function isFish(animal :Cat|Fish){
    if (typeof(animal as Fish).swim == 'function'){
        return true;
    }
    return false;
}

// 类型断言只能够「欺骗」TypeScript 编译器



//继承关系时，的断言使用
class ApiError extends Error{
    code:number = 0;
}
class HttpError extends Error{
    statusCode :number = 200;
}

function isApiError(error:Error){
    if (typeof(error as ApiError).code=='number'){
        return true;
    }
    return false;
}

//使用instanceof

function isApiError2(error:Error){
    if (error instanceof ApiError){
        return true;
    }
    return false;
}



//将任何一个类型断言为any

//将any断言成具体的类型
function getCacheData(key:string):any{
    return (window as any).cache[key];
}

interface Cat2{
    name:string;
    run():void;
}

const tom3=getCacheData('tom') as Cat2;

tom3.run()


//类型断言的限制：
// 若 A 兼容 B，那么 A 能够被断言为 B，B 也能被断言为 A

interface Animal {
    name: string;
}
interface Cat {
    name: string;
    run(): void;
}

let tom4:Cat = {
    name:'tom',
    run:()=>{console.log('run')}
};
let animal00:Animal = tom4;
function testAnimal(animal: Animal) {
    // 父类可以被断言为子类
   return (animal as Cat);
}
function testCat(cat: Cat) {
    // 子类可以被断言为父类
    return (cat as Animal);
}

// 联合类型可以被断言为其中一个类型
// 父类可以被断言为子类
// 任何类型都可以被断言为 any
// any 可以被断言为任何类型
// 要使得 A 能够被断言为 B，只需要 A 兼容 B 或 B 兼容 A 即可



//注意：类型断言会在编译结果中被删除

// 类型断言 vs 类型转换
function toBoolean(something: any): boolean {
    return something as boolean;
}

toBoolean(1);
// 返回值为 1

function toBoolean2(something: any): boolean {
    return Boolean(something);
}

toBoolean(1);
// 返回值为 true
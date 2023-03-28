//函数声明
function sum(x,y){
    return x+y;
}
//函数表达式
let mySum =function(x,y){
    return x+y;
}

function sum1(x:number,y:number):number{
    return x+y;
}

let mysum = function (x:number,y:number):number{
    return x+y;
}

interface SearchFunc{
    (sourse:string,subString:string):boolean;
}
let mySearch :SearchFunc;
mySearch = function(source: string, subString: string){
    return source.search(subString) !== -1;
}


    //带有可选参数的函数
function buildName(fitstName : string,lastName ?: string){
    if(lastName){
        return fitstName+' '+lastName
    }else{
        return fitstName;
    }
}
let tomcat = buildName('tom','cat');
let tom2 =buildName('tom');

function buildName3(firstName :string,lastName : string='Cat'){
    return firstName + ' '+ lastName;
}

let tomcat2 = buildName3('TOm','Cat');
let tom = buildName3('tom');

    // 重载允许一个函数接受不同数量或类型的参数时，作出不同的处理
function reverse(x: number): number;
function reverse(x: string): string;
function reverse(x:number|string):number|string|void{
    if (typeof x == 'number'){
        return Number(x.toString().split('').reverse().join(''));
    }else if (typeof x == 'string'){
        return x.split('').reverse().join('');
    }
}
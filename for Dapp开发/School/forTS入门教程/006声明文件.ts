// declare var 声明全局变量
// declare function 声明全局方法
// declare class 声明全局类
// declare enum 声明全局枚举类型
// declare namespace 声明（含有子属性的）全局对象
// interface 和 type 声明全局类型
// export 导出变量
// export namespace 导出（含有子属性的）对象
// export default ES6 默认导出
// export = commonjs 导出模块
// export as namespace UMD 库声明全局变量
// declare global 扩展全局变量
// declare module 扩展模块

declare var  jQuery:(selector:string) => any;
jQuery('#foo')

//一般将声明语句放到一个单独的 文件中=》声明文件
// 声明文件必须以.d.ts 为后缀


declare function jQuery2(selector:string):any;
declare function jQuery3(domReadyCallback:() => any):any;

// declare class
declare class Animal2{
    name:string;
    constructor(name:string);
    sayHi():string
}


//
declare namespace jQuery2{
    function ajax(url:string,setting?:any):void;
}
jQuery2.ajax('/api/get_something');


declare namespace jQuery3{
    function ajax(url: string, settings?: any): void;
    const version: number;
    class Event {
        blur(eventType: EventType): void
    }
    enum EventType {
        CustomClick
    }
}
jQuery3.ajax('/api/get_something');
console.log(jQuery3.version);
const e = new jQuery3.Event();

//嵌套的命名空间
declare namespace jQuery4{
    function ajax(url:string,setting?:any):void;
    namespace fn{
        function extend(object:any):void;
    }
}
jQuery4.ajax('/api/get_something');
jQuery4.fn.extend({
    check: function() {
        return this.each(function() {
            this.checked = true;
        });
    }
});


// 防止命名冲突
declare namespace jQuery4{
    interface AjaxSettings{
        method ?: 'GET'|'POST'
        data ?: any;
    }
    function ajax(url:string,settings ?: AjaxSettings):void
}
// 使用这个 interface 的时候，也应该加上 jQuery 前缀
let settings: jQuery4.AjaxSettings ={
    method :'POST',
    data:{
        name:'foo'
    }
};


//声明合并
declare function jQuery5(selector:string):any;
declare namespace jQuery5{
    function ajax(url:string,settings?:any):void;

}
jQuery5('#foo');
jQuery5.ajax('/api/get_something');


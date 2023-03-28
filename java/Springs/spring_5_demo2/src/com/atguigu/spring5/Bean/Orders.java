package com.atguigu.spring5.Bean;

public class Orders {
    private String oname;
//无参数构造
    public Orders() {
        System.out.println("第一步，通过构造器构造bean实例");

    }

    public void setOname(String oname) {
        this.oname = oname;
        System.out.println("调用set方法设置属性的值");
    }
    //创建执行初始化的方法
    public void initMethod(){
        System.out.println("执行初始化");
    }
    public  void destoryMethod(){
        System.out.println("执行销毁");
    }
}

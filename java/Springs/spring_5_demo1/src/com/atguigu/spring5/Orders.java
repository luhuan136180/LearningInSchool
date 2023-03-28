package com.atguigu.spring5;
//用于测试有参构造方法进行xml的属性注入
public class Orders {
    private String oname;
    private String address;
    private String idno;

    public Orders(String oname,String address,String idno) {
        this.oname = oname;
        this.address =address;
        this.idno=idno;
    }
    public void testOrder(){
        System.out.println(oname+"::"+address+" "+idno);
    }
}

package com.atguigu.spring5;
// 测试使用set方法在xml中配置对象并且注入属性
public class Book {
    private String bname;
    private String bauthor;
    private  String address;

    public void setBname(String bname) {
        this.bname = bname;
    }

    public void setBauthor(String bauthor) {
        this.bauthor = bauthor;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    //psvm
    public static void main(String[] args) {
        Book book = new Book();
        book.setBname("abc");
        System.out.println(book.bname);
    }

    public void testDemo(){
        System.out.println(bname+':'+bauthor+""+address);
    }
}

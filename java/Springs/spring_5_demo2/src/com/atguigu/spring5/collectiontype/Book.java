package com.atguigu.spring5.collectiontype;

import java.util.List;
//对公共集合注入进行测试，即将集合部分提取出来，供bean对象调用
public class Book {
    private List<String> list;

    public void setList(List<String> list) {
        this.list = list;
    }

    @Override
    public String toString() {
        return "Book{" +
                "list=" + list +
                '}';
    }
}

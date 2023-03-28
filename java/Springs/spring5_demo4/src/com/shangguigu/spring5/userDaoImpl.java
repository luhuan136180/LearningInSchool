package com.shangguigu.spring5;

public class userDaoImpl implements userDao{

    @Override
    public int add(int a, int b) {
        System.out.println("该方法执行了");
        return a+b;

    }

    @Override
    public String update(String id) {
        System.out.println("该方法执行了");
        return id;

    }
}

package com.shangguigu.spring5.dao;

import org.springframework.stereotype.Component;

@Component
public class UserDaoImpl1 implements UserDao{

    @Override
    public void add() {
        System.out.println("接口1号");
    }
}

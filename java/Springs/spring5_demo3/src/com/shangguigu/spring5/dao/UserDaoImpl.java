package com.shangguigu.spring5.dao;

import org.springframework.stereotype.Component;

@Component(value = "userDaoImpl0")
public class UserDaoImpl implements UserDao{

    @Override
    public void add() {
        System.out.println("接口0号");
    }
}

package com.shangguigu.spring5.Service;

import com.shangguigu.spring5.dao.UserDao;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;

@Component
public class UserService3 {
    @Resource(name = "userDaoImpl1")
    private UserDao dao;
    @Value(value = "abc")//注入普通注解
    private String name;
    public void add(){
        System.out.println(" UserService3 add....."+"::"+name);
        dao.add();

    }
}

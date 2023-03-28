package com.shangguigu.spring5.Service;

import com.shangguigu.spring5.dao.UserDao;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Component;

@Component
public class UserService2 {
    @Autowired
    @Qualifier(value = "userDaoImpl1")//该注解必须和@Autowired 连用
    private UserDao dao;
    public void add(){
        System.out.println("UserService2 add....");
        dao.add();
    }
}

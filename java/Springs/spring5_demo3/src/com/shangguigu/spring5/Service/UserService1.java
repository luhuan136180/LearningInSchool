package com.shangguigu.spring5.Service;

import com.shangguigu.spring5.dao.UserDao;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Component;

@Component
public class UserService1 {
    //定义dao类型的属性

    //不需要添加set方法
    //添加注入属性注解
    @Autowired
    @Qualifier(value = "userDaoImpl0")
    private UserDao dao;

    public void  add(){
        System.out.println("UserService1 add....");
        dao.add();

    }
}

package com.atguigu.spring5.service;

import com.atguigu.spring5.dao.UserDao;
import com.atguigu.spring5.dao.UserDaoImpl;

public class UserService {
    //创建userdao类型属性，构造set方法
    private UserDao userdao;
    public void setUserdao(UserDao userdao) {
        this.userdao = userdao;
    }
    public void add(){
        System.out.println("userservice add..........");
        userdao.update();
    }

    public void addold(){
        System.out.println("service add......");
        //多态
        UserDao userdao = new UserDaoImpl();
        userdao.update();
    }
    public static void main(String[] args) {
        UserService a = new UserService();
        a.addold();
    }
}

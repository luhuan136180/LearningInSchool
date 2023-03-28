package com.shangguigu.spring5.testdemo;

import com.shangguigu.spring5.Service.UserService1;
import com.shangguigu.spring5.Service.UserService2;
import com.shangguigu.spring5.Service.UserService3;
import com.shangguigu.spring5.config.SpringConfig;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class testService {
    @Test
    public void testService0(){
        ApplicationContext context = new ClassPathXmlApplicationContext("bean1.xml");
        UserService1 userService1 = context.getBean("userService1", UserService1.class);
        userService1.add();
    }
    @Test
    public void testService1(){
        ApplicationContext context = new ClassPathXmlApplicationContext("bean1.xml");
        UserService2 userService2 = context.getBean("userService2", UserService2.class);
        userService2.add();
    }
    @Test
    public void testService2(){
        ApplicationContext context = new ClassPathXmlApplicationContext("bean1.xml");
        UserService3 userService3 = context.getBean("userService3", UserService3.class);
        userService3.add();
    }
    @Test
    public void testService3(){
    ApplicationContext context = new AnnotationConfigApplicationContext(SpringConfig.class);
    UserService3 userService3 = context.getBean("userService3", UserService3.class);
    userService3.add();
}
}

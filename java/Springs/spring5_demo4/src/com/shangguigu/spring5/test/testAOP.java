package com.shangguigu.spring5.test;

import com.shangguigu.spring5.aopanno.user;
import com.shangguigu.spring5.config.configAop;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class testAOP {
    @Test
    public void testaop1(){
        ApplicationContext context= new ClassPathXmlApplicationContext("bean1.xml");
        user user = context.getBean("user", user.class);
        user.add();
       // user.add1();//用于测试异常通知的方法
    }
    @Test
    public void testaop2(){
        ApplicationContext context= new AnnotationConfigApplicationContext(configAop.class);//基于工厂bean
        user user = context.getBean("user", user.class);
        user.add();
        // user.add1();//用于测试异常通知的方法
    }
}

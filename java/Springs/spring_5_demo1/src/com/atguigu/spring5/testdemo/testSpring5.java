package com.atguigu.spring5.testdemo;

import com.atguigu.spring5.Book;
import com.atguigu.spring5.Orders;
import com.atguigu.spring5.User;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;
import org.springframework.core.annotation.Order;

public class testSpring5 {
    @Test
    public void testAdd() {
        //加载spring的配置文件
        ApplicationContext context=
                new ClassPathXmlApplicationContext("bean1.xml");
        //获取配置创建的对象

        User user = context.getBean("user", User.class);
        System.out.println("user");
        user.add();
    }
    @Test
    public void testBook1(){//对用set方法进行注入属性的测试
        ApplicationContext context1 = new ClassPathXmlApplicationContext("bean1.xml");
        Book book = context1.getBean("book",Book.class);//s:跟的字符需要跟需要的bean中创建的对象的ID一致
        Book book2= context1.getBean("book02",Book.class);
        Book book1= new Book();
        book1.setBname("ssss");
        book1.testDemo();
        System.out.println();
        System.out.println("book");
        book.testDemo();
        System.out.println("book02");
        book2.testDemo();
        //p 的简化测试
        Book book03 = context1.getBean("book03",Book.class);
        System.out.println("book03");
        book03.testDemo();
    }
    @Test
    public void testOrder() {
        ApplicationContext context = new ClassPathXmlApplicationContext("bean1.xml");
        Orders order01 = context.getBean("order01",Orders.class);
        System.out.println("order01");
        order01.testOrder();
    }
}

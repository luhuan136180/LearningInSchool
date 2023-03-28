package com.atguigu.spring5.testDemo;

import com.atguigu.spring5.collectiontype.Book;
import com.atguigu.spring5.collectiontype.Stu;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;


public class testdemo01 {
    @Test
    public void testCollection(){
        ApplicationContext context=new ClassPathXmlApplicationContext("bean1.xml");
        Stu stu=context.getBean("stu01",Stu.class);
        System.out.println(stu.toString());
        System.out.println();
        stu.add();
    }

    @Test
    public void testCollection1(){//对集合的属性注入
        ApplicationContext context=new ClassPathXmlApplicationContext("bean2.xml");
        Book book = context.getBean("book01",Book.class);
        System.out.println(book.toString());

    }
}

package com.atguigu.spring5.testDemo;

import com.atguigu.spring5.Bean.Orders;
import com.atguigu.spring5.autowire.emp;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class teatDemo3 {
    @Test
    public void testBean01(){
        System.out.println(".");
        ApplicationContext context =new ClassPathXmlApplicationContext("bean4.xml");//xml配置中，bean的创建在此处
        System.out.println("..");
        Orders order=context.getBean("order",Orders.class);//获取创建的bean实例对象
        System.out.println();

        System.out.println(order);
    }
    @Test
    public void  testAutWire(){//测试自动装配
        ApplicationContext context =new ClassPathXmlApplicationContext("bean5.xml");
        emp emp01=context.getBean("emp01", emp.class);
        System.out.println(emp01.toString());
        System.out.println("1");
        emp01.test();
        System.out.println("2");
        emp01.toString();
        System.out.println("3");

    }

}

package com.atguigu.spring5.testDemo;

import com.atguigu.spring5.collectiontype.Course;
import com.atguigu.spring5.factoryBean.myBean1;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class testDemo2 {
    @Test
    public void testFactoryBean1(){//普通bean的xml


        ApplicationContext context = new ClassPathXmlApplicationContext("bean3.xml");
        myBean1 mybean01 =context.getBean("mybean01",myBean1.class);
        System.out.println(mybean01);
    }
    @Test
    public void testFactoryBean2(){//工厂bean测试


        ApplicationContext context = new ClassPathXmlApplicationContext("bean3.xml");
        Course course =context.getBean("mybean02",Course.class);
        System.out.println(course);
    }
}

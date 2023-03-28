package com.shangguigu.spring5.Test;

import com.shangguigu.spring5.service.TestService;
import com.shangguigu.spring5.service.TestServiceImpl;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class TestJDBC {
    @Test
    public void main() {
        ApplicationContext context= new ClassPathXmlApplicationContext("bean1.xml");
        TestServiceImpl testServiceImpl = context.getBean("testServiceImpl", TestServiceImpl.class);
        testServiceImpl.testJDBC();

    }
}

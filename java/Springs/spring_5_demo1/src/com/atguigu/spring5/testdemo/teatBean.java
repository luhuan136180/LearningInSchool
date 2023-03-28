package com.atguigu.spring5.testdemo;

import com.atguigu.spring5.neibuBean.Emp;
import com.atguigu.spring5.service.UserService;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class teatBean {
    @Test
    public  void testAdd(){//对外部bean写法进行测试
        ApplicationContext context = new ClassPathXmlApplicationContext("bean2.xml");
        UserService userservice01= context.getBean("userservice01",UserService.class);
        userservice01.add();
    }

    @Test
    public void testAdd02(){//对内部bean写法进行测试
        ApplicationContext context = new ClassPathXmlApplicationContext("bean3.xml");
        Emp emp01= context.getBean("emp01",Emp.class);
        emp01.add();
    }

    @Test
    public void testAdd03(){//对联级写法进行测试
        ApplicationContext context = new ClassPathXmlApplicationContext("bean3.xml");
        Emp emp02= context.getBean("emp02",Emp.class);
        emp02.add();
    }
}

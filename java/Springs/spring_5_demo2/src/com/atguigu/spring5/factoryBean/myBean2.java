package com.atguigu.spring5.factoryBean;

import com.atguigu.spring5.collectiontype.Course;
import org.springframework.beans.factory.FactoryBean;


public class myBean2 implements FactoryBean<Course> {

    //定义返回的bean,修改源代码，返回course类型
    @Override
    public Course getObject() throws Exception {
        Course course= new Course();
        course.setCname("王棵得到恋爱手册");
        return course;
    }

    @Override
    public Class<?> getObjectType() {
        return null;
    }

    @Override
    public boolean isSingleton() {
        return FactoryBean.super.isSingleton();
    }
}

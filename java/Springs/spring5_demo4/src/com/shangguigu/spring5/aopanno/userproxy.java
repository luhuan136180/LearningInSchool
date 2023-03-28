package com.shangguigu.spring5.aopanno;

import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.*;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;

import java.io.PrintStream;
@Component
@Aspect     //生成代理对象
@Order(1)
public class userproxy {
    //抽取相同的切入点
    @Pointcut(value ="execution(* com.shangguigu.spring5.aopanno.user.add())" )//抽取相同的切入点
    public  void pointdemo(){//抽取相同的切入点

    }
    //在增强类的里面，在作为通知方法上面添加通知类型注解，使用切入点表达式
    //@Before(value = "execution(* com.shangguigu.spring5.aopanno.user.add())")//前置通知的注解符号
    @Before(value = "pointdemo()")
    public void before(){
        System.out.println("before.....");
    }

    @After(value = "execution(* com.shangguigu.spring5.aopanno.user.add())")//前置通知的注解符号
    public void after(){//最终通知
        System.out.println("after.....");
    }

    @AfterReturning(value = "execution(* com.shangguigu.spring5.aopanno.user.add())")//前置通知的注解符号
    public void afterreturning(){//后置通知
        System.out.println("afterreturning.....");//后置，且程序正常执行时才会触发该通知的使用
    }

    @AfterThrowing(value = "execution(* com.shangguigu.spring5.aopanno.user.add1())")//前置通知的注解符号
    public void afterThrowing(){
        System.out.println("AfterThrowing.....");
    }

    @Around(value = "execution(* com.shangguigu.spring5.aopanno.user.add())")//前置通知的注解符号
    public void around(ProceedingJoinPoint proceedingJoinPoint) throws Throwable{//环绕通知
        System.out.println("环绕之前.....");
        //执行被增强的函数
        proceedingJoinPoint.proceed();
        System.out.println("环绕之后.....");
    }
}

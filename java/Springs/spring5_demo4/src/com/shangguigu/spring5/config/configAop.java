package com.shangguigu.spring5.config;

import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.EnableAspectJAutoProxy;

@Configuration//
@ComponentScan(basePackages = {"com.shangguigu.spring5"})//扫描
@EnableAspectJAutoProxy(proxyTargetClass = true)//AOP
public class configAop {
}

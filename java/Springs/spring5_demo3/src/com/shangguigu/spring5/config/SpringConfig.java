package com.shangguigu.spring5.config;

import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

@Configuration//注解，替代项目了文件，作为配置类
@ComponentScan(basePackages = {"com.shangguigu"})//自动扫描
public class SpringConfig {

}

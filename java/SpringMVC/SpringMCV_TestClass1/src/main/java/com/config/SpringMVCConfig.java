package com.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.multipart.MultipartResolver;
import org.springframework.web.multipart.commons.CommonsMultipartResolver;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;
import org.springframework.web.servlet.view.InternalResourceViewResolver;


@Configuration
@EnableWebMvc
@ComponentScan(basePackages = {"controller"})
public class SpringMVCConfig implements WebMvcConfigurer {
    /**
     * config view resolver
     */

    @Bean
    public InternalResourceViewResolver getViewResolver() {
        InternalResourceViewResolver viewResolver = new InternalResourceViewResolver();
        viewResolver.setPrefix("/WEB-INF/jsp/");
        viewResolver.setSuffix(".jsp");
        return viewResolver;
    }


    /**
     * config static resource
     */
    @Override
    public void addResourceHandlers(ResourceHandlerRegistry registry) {
        registry.addResourceHandler("/html/**").addResourceLocations("/html/");
    }

    /**
     * MultipartResolver配置
     */
    @Bean
    public MultipartResolver multipartResolver() {

        CommonsMultipartResolver multipartResolver = new CommonsMultipartResolver();
        //设置上传文件的最大值，单位为字节
        multipartResolver.setMaxUploadSize(5400000);
        //设置请求的编码格式，默认为iso-8859-1
        multipartResolver.setDefaultEncoding("UTF-8");
        return multipartResolver;
    }
}
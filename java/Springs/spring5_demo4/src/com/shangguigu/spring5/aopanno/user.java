package com.shangguigu.spring5.aopanno;

import org.springframework.stereotype.Component;

import java.io.PrintStream;
@Component
public class user {
    public  void add(){
        System.out.println("add......");
    }
    public  void add1(){
        System.out.println("add......");
        System.out.println(10/0);
    }
}

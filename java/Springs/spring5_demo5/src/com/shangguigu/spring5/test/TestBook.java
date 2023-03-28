package com.shangguigu.spring5.test;

import com.shangguigu.spring5.Service.BookService;
import com.shangguigu.spring5.entity.Book;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import java.util.ArrayList;
import java.util.List;

public class TestBook {
    @Test
    public void testJdbcTemplate(){
        ApplicationContext context=new ClassPathXmlApplicationContext("bean1.xml");
        BookService bookService = context.getBean("bookService", BookService.class);
        //添加
        /*Book book=new Book();
        book.setUserID("1");
        book.setUserName("java");
        book.setUstatus("a");
        bookService.addBook(book);*/
        //修改
        /*Book book =new Book();
        book.setUserID("1");
        book.setUserName("javase");
        book.setUstatus("c");
        bookService.updateBook(book);*/
        
       // int count = bookService.findCount();
        //System.out.println(count);
        
       // Book bookInfo = bookService.findOne("1");
        //System.out.println(bookInfo);
        
       // List<Book> all=bookService.findAll();
        //System.out.println(all);
// 批量添加
//        ArrayList<Object[]> batchArgs = new ArrayList<>();
//        Object[] o1={"5","c++","a"};
//        Object[] o2={"6","c+","g"};
//        Object[] o3={"7","c#","h"};
//        batchArgs.add(o1);
//        batchArgs.add(o2);
//        batchArgs.add(o3);
//        bookService.batchAddBook(batchArgs)

        //批量修改
//        ArrayList<Object[]> batchArgs = new ArrayList<>();
//        Object[] o1={"sss","a","5"};
//        Object[] o2={"ssss","a1","6"};
//        Object[] o3={"sssss","a2","7"};
//        batchArgs.add(o1);
//        batchArgs.add(o2);
//        batchArgs.add(o3);
//        bookService.batchUpdateBook(batchArgs);

        //删除批量
        ArrayList<Object[]> batchArgs = new ArrayList<>();
        Object[] o1={"5"};
        Object[] o2={"6"};
        Object[] o3={"7"};
        batchArgs.add(o1);
        batchArgs.add(o2);
        batchArgs.add(o3);
        bookService.batchDelateBook(batchArgs);
    }

}

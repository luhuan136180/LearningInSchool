package com.shangguigu.spring5.Dao;

import com.shangguigu.spring5.entity.Book;
import org.springframework.stereotype.Service;

import java.util.List;


public interface BookDao {


    //添加的方法
    void add(Book book);
    //更新
    void update(Book book);
    //删除
    void delete(Book book);

    int selectCount();//查询表里有多少条记录，返回某个值

    Book findBookInfo(String id);


    List<Book> findBookAll();

    void batchAddBook(List<Object[]> batchArgs);

    void batchUpdateBook(List<Object[]> batchArgs);

    void batchDelateBook(List<Object[]> batchArgs);
}

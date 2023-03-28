package com.shangguigu.spring5.Service;

import com.shangguigu.spring5.Dao.BookDao;
import com.shangguigu.spring5.entity.Book;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class BookService {
    //注入dao属性
    @Autowired
    private BookDao bookDao;
    //添加
    public void addBook(Book book){
        bookDao.add(book);
    }
    //修改
    public void updateBook(Book book){
        bookDao.update(book);
    }
    //删除
    public void deleteBook(Book book){
        bookDao.delete(book);
    }
    public int findCount(){//查询表中有多少条记录
        return bookDao.selectCount();
    }
    //查询返回对象
    public Book findOne(String id){
        return bookDao.findBookInfo(id);
    }
    //返回集合的查询(所有)
    public List<Book> findAll(){
        return bookDao.findBookAll();
    }
    //批量添加
    public void batchAddBook(List<Object[]> batchArgs){
        bookDao.batchAddBook(batchArgs);
    }
    //批量修改
    public void batchUpdateBook(List<Object[]> batchArgs){
        bookDao.batchUpdateBook(batchArgs);
    }
    public void batchDelateBook(List<Object[]> batchArgs){
        bookDao.batchDelateBook(batchArgs);
    }
}

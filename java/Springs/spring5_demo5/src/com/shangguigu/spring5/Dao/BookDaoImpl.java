package com.shangguigu.spring5.Dao;

import com.shangguigu.spring5.entity.Book;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.BeanPropertyRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

@Repository
public class BookDaoImpl implements BookDao{
    //注入jdbcTemplate
    @Autowired
    private JdbcTemplate jdbcTemplate;



    //添加的方法
    @Override
    public void add(Book book) {
        //1.创建sql语句
        String sql="insert into t_book values(?,?,?)";
        //2.调用方法实现
        Object[] args={ book.getUserID(), book.getUserName(), book.getUstatus()};
        int update = jdbcTemplate.update(sql, args);
        System.out.println(update);
    }

    @Override
    public void update(Book book) {
        String sql="update t_book set userName =?,ustatus =? where userID=?";
        Object[] args={ book.getUserName(), book.getUstatus(), book.getUserID()};
        int update = jdbcTemplate.update(sql, args);
        System.out.println(update);
    }

    @Override
    public void delete(Book book) {
        String sql="delete from t_book where userID=?";
        Object[] args={ book.getUserID()};
        int update = jdbcTemplate.update(sql, args);
        System.out.println(update);
    }

    @Override
    public int selectCount() {//查询表的记录数
        String sql="Select count(*) from t_book";
        //第二参数shi 返回类型的class
        //queryForObjectsql(String sql ,class<T> requiredType)用于返回某个值，第一个参数是sql语句，第二参数是返回类型的class
        Integer count = jdbcTemplate.queryForObject(sql, Integer.class);
        return count;
    }

    @Override
    /*queryForObject:三个参数，,第一个：sql语句，第二个：RowMapper；第三个：sql的语句值
    * RowMapper:是一个接口，返回不同类型数据，
    * queryForObject(sql, new BeanPropertyRowMapper<Book>(Book.class), id);
    * */
    public Book findBookInfo(String id) {
        String sql="select *from t_book where userID=?";
        Book book = jdbcTemplate.queryForObject(sql, new BeanPropertyRowMapper<Book>(Book.class), id);
        return book;
    }

    @Override
    /*query:三个参数，,第一个：sql语句，第二个：RowMapper；第三个：sql的语句值(返回所有时可以不写)
     * RowMapper:是一个接口，返回不同类型数据，
     * query(sql, new BeanPropertyRowMapper<Book>(Book.class), id);
     *
     * */
    public List<Book> findBookAll() {
        String sql="select *from t_book";
        List<Book> bookList = jdbcTemplate.query(sql, new BeanPropertyRowMapper<Book>(Book.class));
        return bookList;
    }

    @Override
    public void batchAddBook(List<Object[]> batchArgs) {
        String sql="insert into t_book values(?,?,?)";
        int[] ints = jdbcTemplate.batchUpdate(sql, batchArgs);
        System.out.println(Arrays.toString(ints));
    }

    @Override
    public void batchUpdateBook(List<Object[]> batchArgs) {
        String sql="update t_book set username=?,ustatus =? where userID=?";
        int[] ints = jdbcTemplate.batchUpdate(sql, batchArgs);
        System.out.println(Arrays.toString(ints));
    }

    @Override
    public void batchDelateBook(List<Object[]> batchArgs) {
        String sql="delete from t_book where userID=?";
        int[] ints = jdbcTemplate.batchUpdate(sql, batchArgs);
        System.out.println(Arrays.toString(ints));
    }
}

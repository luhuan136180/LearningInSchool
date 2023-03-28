package com.shangguigu.spring5.Dao;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

@Repository
public class UserDaoImpl implements UserDao{
    @Autowired
    private JdbcTemplate jdbcTemplate;

    @Override

    //多钱
    public void addMoney() {
        String sql="update t_account set money=money+? where username=?";
        jdbcTemplate.update(sql,100,"marry");
    }

    @Override
    //少钱
    public void reduceMoney() {
        String sql="update t_account set money=money-? where username=?";
        jdbcTemplate.update(sql,100,"kuccy");
    }
}

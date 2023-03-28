package com.shangguigu.spring5.Dao;

import com.shangguigu.spring5.entiy.MyUser;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.BeanPropertyRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.stereotype.Repository;


import java.util.List;
@Repository
public class TestDaoImpl implements TestDao{
    @Autowired
    private JdbcTemplate jdbcTemplate;
    @Override
    public int update(String sql, Object[] param) {
        return jdbcTemplate.update(sql,param);
    }

    @Override
    public List<MyUser> query(String sql, Object[] param) {
        RowMapper<MyUser> rowMapper=new BeanPropertyRowMapper<MyUser>(MyUser.class);
        return jdbcTemplate.query(sql,rowMapper,param);
    }
}

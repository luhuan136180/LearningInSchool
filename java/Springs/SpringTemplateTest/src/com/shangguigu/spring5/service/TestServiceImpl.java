package com.shangguigu.spring5.service;

import com.shangguigu.spring5.Dao.TestDao;
import com.shangguigu.spring5.entiy.MyUser;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.PrintStream;
import java.util.ArrayList;
import java.util.List;

@Service
public class TestServiceImpl {
    @Autowired
    private TestDao testDao;


    public void testJDBC() {
        String insertSql = "insert into user values(?,?,?)";
        String selectSql = "select * from user";
//        Object[] param1 = {"1","qukuailian1","男"};
//        Object[] param2 = {"2","qukuailian2","女"};
//        Object[] param3 = {"3","qukuailian3","男"};
//        Object[] param4 = {"4","qukuailian4","女"};
//        ArrayList<Object> args = new ArrayList<>();
//        args.add(param1);
//        args.add(param2);
//        args.add(param3);
//        args.add(param4);
//        testDao.update(insertSql,param1);
//        testDao.update(insertSql,param2);
//        testDao.update(insertSql,param3);
//        testDao.update(insertSql,param4);
        List<MyUser> list=testDao.query(selectSql,null);
       // System.out.println(list);
        for(MyUser m :list){
            System.out.println(m);
        }
    }
}

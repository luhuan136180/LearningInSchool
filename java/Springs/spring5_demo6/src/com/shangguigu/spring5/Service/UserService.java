package com.shangguigu.spring5.Service;

import com.shangguigu.spring5.Dao.UserDao;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

@Service
//该注解既可以加在类上，也可以加到方法上面
@Transactional(propagation = Propagation.REQUIRED,isolation = Isolation.SERIALIZABLE,timeout = -1,readOnly = true)
public class UserService {
    @Autowired
    private UserDao userdao;

    public void accountMoney(){

//        userdao.reduceMoney();
//        userdao.addMoney();
        //以上写法没有达成原子性，通过异常会只执行一般，违背原子性
        userdao.reduceMoney();
        userdao.addMoney();

    }
}

package com.shangguigu.spring5;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
import java.util.Arrays;

public class JDKProxy {
    public static void main(String[] args) {
        //创建接口实现类代理对象
       Class[] interfaces={userDao.class};
       // Proxy.newProxyInstance(JDKProxy.class.getClassLoader(), interfaces, new InvocationHandler() {
           // @Override
           // public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
           //     return null;
           // }
        //});
        //三个参数，第一个：类加载器
        //第二个：接口地址
        //第三个：增强代码
        userDaoImpl userDao=new userDaoImpl();
        userDao dao = (userDao) Proxy.newProxyInstance(JDKProxy.class.getClassLoader(), interfaces, new userDaoProxy(userDao));//做一个强转
        int result=dao.add(1,2);
        System.out.println(result);
        String a=dao.update("ada");
        System.out.println(a);
    }
}
//创建代理对象  代码
class userDaoProxy implements InvocationHandler{
    //1.创建的是谁的代理对象，就把谁给传进来
    //通过有参数构造来传递
   // public userDaoProxy(userDaoImpl userdao){}
    //此处用通用写法
    private Object obj;
    public userDaoProxy(Object obj){
        this.obj=obj;
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {//该方法用于写增强的方法部分
        //方法之前
        System.out.println("方法之前执行"+method.getName()+"::被传递的参数"+ Arrays.toString(args));
        //被增强的方法执行
        Object res=method.invoke(obj,args);

        //方法之后
        System.out.println("执行之后"+ obj);
        return res;
    }
}
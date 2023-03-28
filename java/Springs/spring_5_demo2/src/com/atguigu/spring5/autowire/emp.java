package com.atguigu.spring5.autowire;

public class emp {
    private dept dept;
    private String ename;

    public void setDept(com.atguigu.spring5.autowire.dept dept) {
        this.dept = dept;
    }

    public void setEname(String ename) {
        this.ename = ename;
    }

    @Override
    public String toString() {
        return "emp{" +
                "dept=" + dept +
                ", ename='" + ename + '\'' +
                '}';
    }

    public  void test(){

        System.out.println(dept);
    }
}

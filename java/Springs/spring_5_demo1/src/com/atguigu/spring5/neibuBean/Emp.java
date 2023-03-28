package com.atguigu.spring5.neibuBean;

public class Emp {
    private String sname;
    private String gender;
    private Dept dept;
    public void setSname(String sname) {
        this.sname = sname;
    }
    public void setGender(String gender) {
        this.gender = gender;
    }
    public void setDept(Dept dept) {
        this.dept = dept;
    }
    public void add(){
        System.out.println(sname +"::"+gender+"::"+dept);
    }
}

package com.shangguigu.spring5.entiy;

public class MyUser {
    private Integer uid;
    private  String uname;
    private  String usex;

    public Integer getUid() {
        return uid;
    }
    public String getUname() {
        return uname;
    }
    public String getUsex() {
        return usex;
    }

    public void setUid(Integer uid) {
        this.uid = uid;
    }
    public void setUname(String uname) {
        this.uname = uname;
    }
    public void setUsex(String usex) {
        this.usex = usex;
    }

    @Override
    public String toString() {
        return "MyUser{" +
                "uid=" + uid +
                ", uname='" + uname + '\'' +
                ", usex='" + usex + '\'' +
                '}';
    }
}

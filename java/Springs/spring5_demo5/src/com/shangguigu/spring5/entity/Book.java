package com.shangguigu.spring5.entity;
//数据库中表的实体化
public class Book {
    private String userID;
    private String userName;
    private String ustatus;

    @Override
    public String toString() {
        return "Book{" +
                "userID='" + userID + '\'' +
                ", userName='" + userName + '\'' +
                ", ustatus='" + ustatus + '\'' +
                '}';
    }

    public String getUserID() {
        return userID;
    }
    public String getUserName() {
        return userName;
    }
    public String getUstatus() {
        return ustatus;
    }

    public void setUserID(String userID) {
        this.userID = userID;
    }
    public void setUserName(String userName) {
        this.userName = userName;
    }
    public void setUstatus(String ustatus) {
        this.ustatus = ustatus;
    }
}

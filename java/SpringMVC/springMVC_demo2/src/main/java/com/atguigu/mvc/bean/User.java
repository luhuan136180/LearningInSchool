package com.atguigu.mvc.bean;

public class User {
    private String username;
    private String password;
    private String age;
    private String id;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }
    public String getPassword() {
        return password;
    }
    public String getAge() {
        return age;
    }

    public void setUsername(String username) {
        this.username = username;
    }
    public void setPassword(String password) {
        this.password = password;
    }
    public void setAge(String age) {
        this.age = age;
    }

    public User(String username, String password, String age, String id) {
        this.username = username;
        this.password = password;
        this.age = age;
        this.id = id;
    }

    @Override
    public String toString() {
        return "User{" +
                "username='" + username + '\'' +
                ", password='" + password + '\'' +
                ", age='" + age + '\'' +
                ", id='" + id + '\'' +
                '}';
    }
}

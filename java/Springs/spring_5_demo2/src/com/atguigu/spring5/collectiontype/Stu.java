package com.atguigu.spring5.collectiontype;

import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.Set;

public class Stu {
    //用于测试xml中通过bean操作，配置集合bean对象
    private String[] courses;//数组
    private List<String> list;//链
    private Map<String,String> maps;//图
    private Set<String> set;

    private List<Course> ListCourses;//该属性用于测试bean中集合中注入外部类型值

    public void setSet(Set<String> set) {
        this.set = set;
    }

    public void setCourses(String[] courses) {
        this.courses = courses;
    }

    public void setList(List<String> list) {
        this.list = list;
    }

    public void setMaps(Map<String, String> maps) {
        this.maps = maps;
    }

    public void setListCourses(List<Course> listCourses) {
        ListCourses = listCourses;
    }

    @Override
    public String toString() {
        return "Stu{" +
                "courses=" + Arrays.toString(courses) +
                ", list=" + list +
                ", maps=" + maps +
                ", set=" + set +
                '}';
    }
    public void add(){
        System.out.println(Arrays.toString(courses));
        System.out.println(ListCourses);
    }
}

package com.atguigu.mvc.controller;

import com.atguigu.mvc.bean.User;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

import javax.servlet.http.HttpServletRequest;
import java.util.Arrays;

@Controller
public class paramController {

    @RequestMapping(value = "/testServletAPI")
    //形参位置的request表示当前请求
    public String testServletAPI(HttpServletRequest request){
        String username=request.getParameter("username");
        String password=request.getParameter("password");
        System.out.println("username:"+username+",password:"+password);
        return "success";
    }
    @RequestMapping("/testParam")
    public String testParam(
            @RequestParam(value = "user_name") String username,
            String password,
            String[] hobby) {
        System.out.println("username:"+username+",password:"+password+"hobby:"+ Arrays.toString(hobby));
        return "success";

    }
    @RequestMapping("/testBean")
    public String testBean(User user){
        System.out.println(user);
        return "success";
    }

}

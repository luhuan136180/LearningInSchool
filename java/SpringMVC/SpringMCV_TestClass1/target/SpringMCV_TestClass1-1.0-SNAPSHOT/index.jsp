<%--
  Created by IntelliJ IDEA.
  User: 半世琉璃
  Date: 2021/12/5
  Time: 20:10
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" pageEncoding="utf-8" %>
<html>
<head>
    <meta charset="utf-8">
    <title></title>
</head>
<body>
<form action="${pageContext.request.contextPath}/multifile" method="post" enctype="multipart/form-data">

    选择文件1：<input type="file" name="myfile"><br>
    文件描述1：<input type="text" name="description"><br>
    选择文件2：<input type="file" name="myfile"><br>
    文件描述2：<input type="text" name="description"><br>
    选择文件3：<input type="file" name="myfile"><br>
    文件描述3：<input type="text" name="description"><br>
    <input type="submit" value="提交">
</form>
</body>
</html>

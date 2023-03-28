<%--
  Created by IntelliJ IDEA.
  User: 半世琉璃
  Date: 2021/12/8
  Time: 10:15
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" pageEncoding="UTF-8" %>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<html>
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<table>
    <tr>
        <td>详情</td>
        <td>文件名</td>
    </tr>
    <%--同时取两个数组的元素--%>
    <c:forEach items="${multiFileDomain.description}" var="description" varStatus="loop">
        <tr>
            <td>${description}</td>
            <td>${multiFileDomain.myfile[loop.count-1].originalFilename}</td>
        </tr>
    </c:forEach>
</table>
</body>
</html>
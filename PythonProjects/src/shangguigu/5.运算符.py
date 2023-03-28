a=3
b=2
print(a+b,"加法")
print(a-b,'jian')
print(a*b,'')
print(a/b)
print(a//b,'取整除')
print(a%b,'')
print(a**b,'指数')
#如果是两个字符串做加法运算，会直接把这两个字符串拼接成一个字符串。
str1 = 'hello '
str = 'world'
print(str+str1)

#如果是数字和字符串做乘法运算，会将这个字符串重复多次。
str2 = 'hello'
c =5
print(str2*c)
#赋值运算符
num=10
a=b=4
print(a)
print(b)
print(num)
num1,c,d=100,3.24,"hello"
print(c)
print(d)
print(num1)
#逻辑运算符
print("逻辑运算符")
#and :并
#or : 或
#not:非
a = 34
a > 10 and print('hello world')
a < 10 and print('hello world')
a >10 or print('你好世界1')
a <10 or print('你好世界2')


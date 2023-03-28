# .创建一个列表
a=['aaaaa','ssssss','dddddd','qqqqqq','zzzzzz']
# 4.分别通过下标和切片两种方法获取列表中的第4个元素
i=0
j=4
while i<len(a):
    if i==j-1:
        print(a[i])
    i+=1

print(a[j-1])
# 5.分别通过下标和切片两种方法修改列表中的第4个元素
i=0
while i<len(a):
    if i==j-1:
        a[i]='被修改'
        print(a[i])
    i+=1

a[j-1]='再次被修改'
print(a[j-1])
# 6.分别通过如下四种方法（append、insert、extend、切片）添加列表元素
temp='zxczcz'
print(a)
a.append(temp)
a.insert(1,'lalaala')
a.insert(2,'lalal2')
print(a)
b=['a','c','d','e']
a.extend(b)
print(a)
a[len(a):]=['9']
print(a)
# 7.分别通过如下5种方法（del、pop、remove、clear、切片）删除列表元素
del a[0]
print("删除后",a)
a.pop()
print("shanchuhou",a)
a.remove('再次被修改')
print('删除后',a)
del a[:2]
print('删除后',a)
# 8.分别把index方法、count方法、reserve方法、sort方法用于不同列表进行处理
a.append('a')
print('a出现的次数',a.count('a'))
print('zxczcz所在索引位置',a.index('zxczcz'))
a.reverse()
print(a)
a.sort(key=None,reverse=False)
print(a)
# 9.分别把常用内置函数如max()、min()、sum()、len()、sorted()、reversed()用于不同列表进行处理
dight=[34,12,33,76,12,32,9,17,27,43]
print('列表中的最大数：',max(dight))
print('列表中的最小数：',min(dight))
print('列表中的数求和：',sum(dight))
print('列表中的元素个数：',len(dight))
print('列表中的元素个数：',len(dight))
sorted(dight,reverse=False)
print(dight)
reversed(dight)
print(dight)
# 10.连接两个列表
str=['aaa','bbb','ccc']
str2=['ddd','eee','fff']
str.extend(str2)
print(str)
# 11.使一个列表重复6次输出一个新列表
i=1
for i in range(6):
    str.extend(str)
# 12.编写一个代码使用到列表推导式
list = [i for i in range(10)]
print(list)
# 13.创建一个元组然后删除此元组
yuanzu=(2,3,'aa')
print(yuanzu)
del yuanzu
# 14.分别通过下标和切片两种方法获取元组中的第4个元素
tuple1=(1,2,5,6,3,2,4,5)
print(tuple1[3])
print(tuple1[3:4])
# 15.分别把index方法、count方法用于不同元组进行处理
print(tuple1.index(1))
print(tuple1.count(2))
#16.分别把常用内置函数如max()、min()、sum()、len()、sorted()、reversed()用于不同元组进行处理
print(
    max(tuple1)
)
print(min(tuple1))
print(sum(tuple1))
print(len(tuple1))
sorted(tuple1,reverse=False)
print(tuple1)
reversed(tuple1)
print(tuple1)
# 17.编写一个代码使用到生成器表达式，输出形式为元组
# def add(n, i):
#     return n + i
# def test():
#     for i in range(4):
#         yield i
# g = test()  # 初始化生成器对象
# for n in [1, 10]:
#     g = (add(n, i) for i in g)
# res = list(g)
# print(res)
# 18.使用三种方式分别创建不同的字典
a1={'name':'xiaoming','age':12}
a2=dict(name='小红',age=18)
k=["name","age"]
v=["小刚",15]
a3=dict(zip(k,v))
print(a1)
print(a2)
print(a3)

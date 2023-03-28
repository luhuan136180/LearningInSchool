#append
a=['aaaaa','ssssss','dddddd']
print("‐‐‐‐‐添加之前，列表A的数据‐‐‐‐‐A=%s" % a)
temp='zxczcz'
a.append(temp)
print("‐‐‐‐‐添加之后，列表A的数据‐‐‐‐‐A=%s" % a)
#insert
strs=['a','c','d','e']
print(strs)
strs.insert(1,'b')
print(strs)
#extend
b=['aa','ss','zzz']
a.extend(b)
print(a)
#修改：
a[0]='china'
print(a)
#查找元素 : in ;in not
name=['tom','smith','marry','jack']
findname='marry'
if findname in name:
    print("找到")
else:
    print("meizhoadao ")
#删除元素
'''
del：根据下标进行删除
pop：删除最后一个元素
remove：根据元素的值进行删除
'''
#name=['tom','smith','marry','jack']
del name[0]
print("删除后",name)
name.pop()
print("shanchuhou",name)
name.remove('marry')
print('删除后',name)
i=0

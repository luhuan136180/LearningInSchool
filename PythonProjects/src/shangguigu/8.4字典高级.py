#查看
info = {'name':'班长','age':18}
print(info['name'])
print(info)

print(info.get('age'))
#修改
print('修改前',info)
info['age']=20
print('修改后',info)
'''
果在使用 变量名['键'] = 数据 时，这个“键”在字典中，
不存在，那么就会新增这个元素
'''
info['id']=123713298731
print(info)
#删除
#del :
#删除部分：
del info['name']
print(info)
#整个删除
'''del info
print(info)'''
#clear:
info.clear()
print(info)
#遍历：
print("开始练习遍历")
person={'name':'laoma','age':18,'sex':'男'}
#遍历字典的key（键）
for key in person.keys():
    print(key)
#遍历字典的value（值）
for value in person.values():
    print(value)
#遍历字典的key-value（键值对）
for key,value in person.items():
    print('key=%s,value=%s'%(key,value))
#遍历字典的项（元素）
for item in person.items():
    print(item)
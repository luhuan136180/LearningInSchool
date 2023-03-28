import json
import json
file = open('names.txt', 'w')
names = ['zhangsan', 'lisi', 'wangwu', 'jerry', 'henry', 'merry', 'chris']
# file.write(names) 出错，不能直接将列表写入到文件里
result=json.dumps(names)
print(type(result))
file.write(result)
file.close()

import json
file = open('names.txt', 'w')
names = ['zhangsan', 'lisi', 'wangwu', 'jerry', 'henry', 'merry', 'chris']
# dump方法可以接收一个文件参数，在将对象转换成为字符串的同时写入到文件里
json.dump(names, file)
file.close()
'''
反序列
'''
# 调用loads方法，传入一个字符串，可以将这个字符串加载成为Python对象
result = json.loads('["zhangsan", "lisi", "wangwu", "jerry", "henry", "merry", "chris"]')
print(type(result)) # <class 'list'>
# 以可读方式打开一个文件
file = open('names.txt', 'r')
# 调用load方法，将文件里的内容加载成为一个Python对象
result = json.load(file)
print(result)
file.close()
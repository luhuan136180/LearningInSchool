#open(文件路径，访问模式）
'''fp=open('test2.txt','w')
fp.write('add ....\n')

fp.write('hello world, i am here!\n' * 5)
fp.close()
'''
'''
w:覆盖
a：追加
r:只读

'''
fp=open('test2.txt','r')
content=fp.read()
print(content)

f = open('test2.txt', 'r')
content = f.read(5) # 最多读取5个数据
print(content)
print("‐"*30) # 分割线，用来测试
content = f.read() # 从上次读取的位置继续读取剩下的所有的数据
print(content)
f.close() # 关闭文件，这个可是个好习惯哦
#readline
print('realine~')
f = open('test2.txt', 'r')
content = f.readline()
print("1:%s" % content)
content = f.readline()
print("2:%s" % content)
f.close()
#readlines
print('readlines~')
f = open('test2.txt', 'r')
content = f.readlines()
print(type(content))
for temp in content:
    print(temp)
f.close()

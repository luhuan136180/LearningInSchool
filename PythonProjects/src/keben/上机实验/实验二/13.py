str='hello,the world'
list=list(str)
set1=set(list)
print(set1)
for i in set1:
    count=list.count(i)
    print(print(i,'出现的次数：',count))

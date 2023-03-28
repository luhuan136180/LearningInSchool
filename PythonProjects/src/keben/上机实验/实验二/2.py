# 2
i=input('给定的车牌号：')
list1=list(i)
list1.reverse()
if int(list1[0])%2==0:
    print('不能上路')
else:
    print('能上路')
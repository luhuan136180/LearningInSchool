
# 有一个正方形，四个角的坐标（x,y)分别是（1，-1），（1，1），（-1，-1），（-1，1），x是横轴，y是纵轴。写一个程序，判断一个给定的点是否在这个正方形内（包括正方形边界）。
a=[1,-1]
b=[1,1]
c=[-1,-1]
d=[-1,1]
print('----')
i=input().split()
print(i)
if float(i[0])>float(-1) and float(i[0])<float(1):
    if float(i[1])>float(-1) and float(i[1])<float(1):
        print('在正方形内')
    else:
        print('不在正方形内')
else:
    print('不在正方形内')
print('完成')
x=input('请输入x的值')
x=int(x)
if x>=0 and x<5:
    print('y=',-x+2.5)
elif x<10:
    print('y=',2-1.5*(x-3)*(x-3))
elif x<20:
    print('y=',x/2-1.5)

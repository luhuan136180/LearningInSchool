
#一只大象口渴了，要喝20升水才能解渴，但现在只有一个深h厘米，底面半径为r厘米的小圆桶(h和r都是整数)。问大象至少要喝多少桶水才会解渴。（键盘输入h、r）
a=input('输入h和r').split()
s=3.14*int(a[1])*int(a[1])
v=int(a[0])*s
print(v)
n=int(20.0/v)+1
print('需要喝',n,'桶水')
# 对于多项式f(x) = ax3+bx2+cx+d和给定的a, b, c, d, x，计算f(x)的值。
s = input().split()
x = float(s[0])
a = float(s[1])
b = float(s[2])
c = float(s[3])
d = float(s[4])
final = (a*x*x*x)+ (b*x*x)+ c*x + d
print("%.7f" % final)

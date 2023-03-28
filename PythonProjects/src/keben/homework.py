num1=input("num1=")
num2=input("num2=")
if num1>num2:
    num1,num2=num2,num1
print("num1=",num1,";num2=",num2)
#
num3=input("num3=")
if int(num3)%2 ==0:
    print(num3,'是oushu')
else:
    print(num3,'不是偶数')

score=input("请输入成绩：")
if int(score)>=60:
    print("成绩合格")
else:
    print('成绩不及格')

a=23
i=1
while i>=0:
    num4=input("请输入一个数字")
    num=int(num4)
    if num<a:
        print('小了')
    elif num>a:
        print('大了')
    else:
        print('猜对了')
        break
score2=input("请输入成绩：")
w=int(score2)
if w>90:
    print("优秀")
elif w>80:
    print("lianghao")
elif w>70:
    print('中等')
elif w>60:
    print('及格')
else:
    print('不及格')

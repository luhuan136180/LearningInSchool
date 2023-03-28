age=30
if age>=10:
    print("我已经成年l")
if age<40:
    print("i am not a older")
#age1=eval(input("qingshuru你的年龄："))
age1=14
if age1>18:
    print("can play computer")
else:
    print("未成年，不允许")

'''
if xxx1:
事情1
elif xxx2:
事情2
elif xxx3:
事情3
'''
score = 77
if score>=90:
    print("本次考试有")
elif score>80:
    print("kaoshidengjiwei:8")
elif score>70:
    print("bencikaoshi :7")
elif score>=60:
    print('本次考试，等级为D')
elif score<60:
    print('本次考试，等级为E')
#for循环
for s in "hello":
    print(s)
for i in range(5):
    print(i)
sum=0
for i in range(100):
    sum+=(i+1)
print(sum)
#range
for x in range(2,10,3):
    print(x)
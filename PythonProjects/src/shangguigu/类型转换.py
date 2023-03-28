#一下为强制转换
print(int("123"))
str01 = '123'
int01 = int(str01)
print(type(str01),str01)
print(type(int01),int01)
flo01 = 1.23
int02 = int(flo01)
print(type(int02),int02)
print(type(flo01),flo01)
bool01 = True
int03 = int(bool01)
print(
    type(bool01),bool01
)
print(type(int03),int03)
print("-------------")
print("#其他类型转浮点型")
#其他类型转浮点型
f1 = float("12.34")
print(type(f1),f1)
print(type("123"))
f2 = float(23)
print(23)
print(f2)
print(type(f2))
print("-------------")
print("#装换成字符串      str()")
str02 = str(45)
str03 = str(23.45)
str04 = str(True)
print(type(str02),str02,type(str03),type(str04))

print("-------------")
print("#装换成buerxing     bool()")
print("对非0的整数转换，都换为true")
print(bool(''))
print(bool(""))
print(bool(0))
print(bool({}))
print(bool([]))
print(bool(()))
print(bool(2))
print(bool(-1))
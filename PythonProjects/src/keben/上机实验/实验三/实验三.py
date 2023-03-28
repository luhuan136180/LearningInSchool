str=input('请输入密文')
list=list(str)
# list[0]=chr(ord(list[0])+1)
print(list)
i=0
for i in range(len(list)):
    if list[i] != ' ':
        list[i]=chr((ord('a')+13-96)%26+96).lower()
print(list)

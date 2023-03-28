
# 根据邮件的重量和用户是否选择加急计算邮费。计算规则：重量在1000克以内(包括1000克), 基本费8元。超过1000克的部分，每500克加收超重费4元，不足500克部分按500克计算；如果用户选择加急，多收5元。
i=input('输入重量')
i=int(i)
consum=8
if i<=1000:
    print()
else:
    j=int((i-1000)/500)+1
    print(j)
    consum=consum+j*4
j=input('是否加急，输入y/n')
if j=='y':
    consum+=5
print('consum=',consum)
i=0
while i<5:
    print('welcome to learn Python')
    i=i+1
i=0
sum=0
while i<=100:
    if i%2==0:
        sum=sum+i
    i=i+1
print('sum=',sum)


str='beijing'
for i in str:

    print(i,end=',')
print()
for i in str:
    print(i,end='')
print()
for i in range(11):
    print(i,end=',')
print()
i=1
sum=0
for i in range(1,101,2):
    sum+=i
print(sum)

j=1
for j in range(1,10):
    for k in range(1,j+1):
        print(k,"*",j,'=',j*k,end='\t')
    print()
j=1
while j<10:
    k=1
    for k in range(1,j+1):
        print(k,"*",j,'=',j*k,end='\t')
    j+=1
    print()


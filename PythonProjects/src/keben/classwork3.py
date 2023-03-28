n=10
res=0

for i in range(n):
    sum=1
    for j in range(i+1):
        sum=sum*(j+1)
    res+=sum

print(res)
for i in range(4):
    print(i)
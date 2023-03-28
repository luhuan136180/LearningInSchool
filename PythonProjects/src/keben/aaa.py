str='hello the world,thanks'
l=list(str)
print(l)
ste1=set(l)
l2=[]
for i in ste1:
    count=l.count(i)
    l2.append(count)
print(l2)
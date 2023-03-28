score=['89','98','99','56','76','67','48','83','77','92','92']
print('成绩最高的为：',max(score))
score.sort()
print('成绩最高的为：',score[-1])
print('成绩最低的为：',score[0])
sum=0
i=0
while i<len(score):
    sum=int(score[i])+sum
    i=i+1
print("所有学生的平均成绩为：",sum/len(score))
sum=0
for i in score:
    if i<'60':
        sum+=1
print('不及格的人数为：',sum)
sum=0
for i in score:
    if int(i)>90:
        sum+=1
print("高于90分的人数：",sum)
sum=0
for i in score:
    if int(i)>=60 and int(i)<=80:
        sum+=1
        print(i)
print("60-80分之间的人数：",sum)
i=0
while i<len(score):
    j=i
    if j+1<len(score) and score[j+1]==score[i]:
        print(score[i],"有相同成绩")
    i+=1

score.reverse()
print('从高到低排序',score)
print('第二名的成绩为：',score[1])
m=input('请输入成绩')
print('成绩为：',m,'的排名为：',score.index(m)+1)
print('棋盘装米问题：')
print(sum([2**n for n in range(64)]))
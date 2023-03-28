s='Chain'
print(len(s))
#find:查找指定内容在字符串中是否存在，如果存在就返回该内容在字符串中第一次
#出现的开始位置索引值，如果不存在，则返回-1.
s1='chain'
print(s1.find('c'))
# startswitch判断字符串是不是以谁谁谁开头/结尾
s2='chain'
print(s2.startswith('h'))
print(s2.endswith('n'))
#count:返回 str在start和end之间 在 mystr里面出现的次数
s3='aaabbbb'
print(s3.count('a'))
#replace:替换字符串中指定的内容，如果指定次数count，则替换不会超过count次
s4='ccddddd'
print(s4)
s4.replace('c','d')
print(s4)
#split:通过参数的内容切割字符串
s5='2#s#f#4#g'
print(s5.split('#'))
#upper,lower
s6='china'
print(s6.upper())
#strip去空格

s8='   a   b    '
print(len(s8))
print(len(s8.strip()))
#join
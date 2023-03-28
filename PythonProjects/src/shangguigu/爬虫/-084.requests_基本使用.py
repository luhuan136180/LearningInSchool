import requests

url = 'http://www.baidu.com'
# 一个类型和六个属性
'''
类型 ：models.Response

r.text : 获取网站源码
r.encoding ：访问或定制编码方式
r.url ：获取请求的url
r.content ：响应的字节类型
r.status_code ：响应的状态码
r.headers ：响应的头信息'''
response = requests.get(url=url)
print(type(response))
# 设置响应的编码格式
response.encoding = 'utf-8'

# 以字符串的形式来返回了网页的源码
print(response.text)

# 返回一个url地址
print(response.url)
# 返回的是二进制的数据
print(response.content)

# 返回响应的状态码
print(response.status_code)
# 返回的是响应头
print(response.headers)

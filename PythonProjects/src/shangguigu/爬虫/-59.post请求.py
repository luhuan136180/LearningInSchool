# post请求

import urllib.request
import urllib.parse
url='https://fanyi.baidu.com/sug'
headers={
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39'
}
data={
    'kw':'spider'
}
# post请求的参数 必须要进行编码
data2=urllib.parse.urlencode(data)
print(data2)
data=urllib.parse.urlencode(data).encode('utf-8')
print(data)
# post的请求的参数 是不会拼接在url的后面的  而是需要放在请求对象定制的参数中
# post请求的参数 必须要进行编码
request=urllib.request.Request(url=url,data=data,headers=headers)
# 模拟浏览器向服务器发送请求
response=urllib.request.urlopen(request)
# 获取响应的数据
content=response.read().decode('utf-8')
print(content)

# 字符串--》json对象
import json

obj = json.loads(content)
print(obj)

# 编码之后 必须调用encode方法 data = urllib.parse.urlencode(data).encode('utf-8')
# 参数是放在请求对象定制的方法中  request = urllib.request.Request(url=url,data=data,headers=headers)



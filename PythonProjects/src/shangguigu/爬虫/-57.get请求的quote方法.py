# https://www.baidu.com/s?wd=%E5%91%A8%E6%9D%B0%E4%BC%A6


# 需求 获取 https://www.baidu.com/s?wd=周杰伦        的网页源码
import urllib.request
import urllib.parse

url = 'https://www.baidu.com/s?wd='
headers={
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32'
}

# 将周杰伦三个字变成unicode编码的格式
# 我们需要依赖于urllib.parse
name=urllib.parse.quote('周杰伦')
urllib = url + name
print(url)
# 请求对象的定制
request = urllib.request.Request(url=url,headers=headers)
# 模拟浏览器向服务器发送请求
response = urllib.request.urlopen(request)
# 获取响应的内容
content = response.read().decode('utf-8')
# 打印数据
print(content)



url = 'http://www.baidu.com/s?'
data = {
'name':'小刚',
'sex':'男',
}
data = urllib.parse.urlencode(data)
url = url + data
print(url)
'''headers={
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32'
}
request = urllib.request.Request(url=url,headers=headers)
response=urllib.request.urlopen(request)
print(response.read().decode('utf-8'))'''
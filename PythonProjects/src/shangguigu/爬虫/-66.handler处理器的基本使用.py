import urllib.request

url = 'http://www.baidu.com'

headers ={
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32'
}
request = urllib.request.Request(url = url,headers = headers)

# handler   build_opener  open

# （1）获取hanlder对象
hander=urllib.request.HTTPHandler()

# （2）获取opener对象
opener=urllib.request.build_opener(hander)

# (3) 调用open方法
response=opener.open(request)

content=response.read().decode('utf-8')
print(content)
import urllib.request

url='http://www.baidu.com/s?wd=ip'
headers={
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32'
}

request=urllib.request.Request(url=url,headers=headers)
# 模拟浏览器访问服务器
#response=urllib.request.urlopen(request)


proxies={
    'http':'120.42.46.226:6666'
}
#重点代码 # handler  build_opener  open
handler=urllib.request.ProxyHandler(proxies=proxies)

opener=urllib.request.build_opener(handler)

response=opener.open(request)

# 获取响应的信息
content = response.read().decode('utf-8')

# 保存
with open('daili.html','w',encoding='utf-8')as fp:
    fp.write(content)
import urllib.request
rul="https://www.baidu.com/"
# rul='https://cd.58.com/danche/?utm_source=sem-baidu-pc&spm=209092640546.44538031140&PGTID=0d100000-0006-6e8e-bfbd-dc4b72d4244a&ClickID=2'
headers={
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36'
}
request=urllib.request.Request(headers=headers,url=rul)
response=urllib.request.urlopen(request)
content=response.read().decode('utf-8')

# 解析网页源码 来获取我们想要的数据
from lxml import etree

# 解析服务器响应的文件
tree=etree.HTML(content)

# 获取想要的数据  xpath的返回值是一个列表类型的数据
# result=tree.xpath('//input[@id="su"]/@value')[0]
result = tree.xpath('//ul[@id="hotsearch-content-wrapper"]//a/span[@class="title-content-title"]/text()')
#//input[@id="su"]/@value

print(result)
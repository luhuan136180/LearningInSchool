import requests
url = 'http://www.baidu.com/s?'


headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36',
}

data = {
    'wd':'ip'
}

response = requests.get(url=url,params=data,headers=headers)
content=response.text
with open('daili.html','w',encoding='utf-8')as fp:
    fp.write(content)
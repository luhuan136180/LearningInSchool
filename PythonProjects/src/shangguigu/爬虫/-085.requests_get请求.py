# urllib
# (1) 一个类型以及六个方法
# （2）get请求
# （3）post请求   百度翻译
# （4）ajax的get请求
# （5）ajax的post请求
# （6）cookie登陆 微博
# （7）代理


# requests
# (1)一个类型以及六个属性
# （2）get请求
# （3）post请求
# （4）代理
# （5）cookie  验证码

import requests
url ='https://www.baidu.com/s'

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32'
}

data ={
    'wd' : '北京'
}

# url  请求资源路径
# params 参数
# kwargs 字典

reponse = requests.get(url=url,params=data,headers=headers)

content = reponse.text

print(content)
# 适用的场景：数据采集的时候 需要绕过登陆 然后进入到某个页面
# 个人信息页面是utf-8  但是还报错了编码错误  因为并没有进入到个人信息页面 而是跳转到了登陆页面
# 那么登陆页面不是utf-8  所以报错

# 什么情况下访问不成功？
# 因为请求头的信息不够  所以访问不成功

import urllib.request

url = 'https://weibo.cn/7758768475/info'
headers = {
    # ':authority':' weibo.cn',
    # ':method':' GET',
    # ':path':' /7758768475/info',
    # ':scheme':' https',
    'accept':' text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9',
    # 'accept-encoding':' gzip, deflate, br',
    'accept-language':' zh-CN,zh;q=0.9',
    'cache-control':' max-age=0',
    #     cookie中携带着你的登陆信息   如果有登陆之后的cookie  那么我们就可以携带着cookie进入到任何页面
    'cookie':' _T_WM=071b53b7e5b86e205d9f09e6554e30ac; WEIBOCN_WM=3349; H5_wentry=H5; backURL=https%3A%2F%2Fweibo.cn; SUB=_2A25Pc_GMDeRhGeFJ7loW9ibIzDmIHXVsn5_ErDV6PUJbktANLRejkW1Nf5zkC12zye1LzHX_630RjMk-9iR8KqVo; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WF_OIECiro5OKmsqfo0WNT35NHD95QNS0-RS0qRShMfWs4DqcjMi--NiK.Xi-2Ri--ciKnRi-zNeK.41hMR1KM0eBtt; SSOLoginState=1651999196',
    'referer': 'http://weibo.cn/',
    'sec-ch-ua':' " Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"',
    'sec-ch-ua-mobile':' ?0',
    'sec-ch-ua-platform':' "Windows"',
    'sec-fetch-dest':' document',
    'sec-fetch-mode':' navigate',
    'sec-fetch-site':' same-origin',
    'sec-fetch-user':' ?1',
    'upgrade-insecure-requests':' 1',
    'user-agent':' Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36',
}

request=urllib.request.Request(url=url,headers=headers)
response=urllib.request.urlopen(request)
content=response.read().decode('utf-8')
print(content)
# 将数据保存到本地
with open('weibo.html','w',encoding='utf-8')as fp:
    fp.write(content)
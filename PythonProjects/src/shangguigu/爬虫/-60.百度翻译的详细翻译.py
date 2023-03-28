import urllib.request
import urllib.parse
url='https://fanyi.baidu.com/v2transapi?from=en&to=zh'
#定义需要爬去的数据
headers={
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39',
    'Cookie': 'BIDUPSID=5026A8DB126738F57F5A023334613537; PSTM=1640231838; __yjs_duid=1_24aa82c8bbd78ea3288b62456b3f990f1640245127294; BAIDUID=415FC4DA9F7D99B63FC30579FF9FB8C4:FG=1; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; SOUND_SPD_SWITCH=1; HISTORY_SWITCH=1; SOUND_PREFER_SWITCH=1; APPGUIDE_10_0_2=1; BDUSS=pzbjVqY2ozNjZMSjJseWlJeHV3anR3YWxUUi1ZTWtZekRoSVBKa2JORnFYWU5pRVFBQUFBJCQAAAAAAAAAAAEAAAC2EGp~RGdnc3RrdmMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGrQW2Jq0FtiM2; BDUSS_BFESS=pzbjVqY2ozNjZMSjJseWlJeHV3anR3YWxUUi1ZTWtZekRoSVBKa2JORnFYWU5pRVFBQUFBJCQAAAAAAAAAAAEAAAC2EGp~RGdnc3RrdmMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGrQW2Jq0FtiM2; RT="z=1&dm=baidu.com&si=6iljoawwo8f&ss=l2pkrwfd&sl=k&tt=pfc&bcn=https%3A%2F%2Ffclog.baidu.com%2Flog%2Fweirwood%3Ftype%3Dperf&ld=2zkss&ul=3kkdo&hd=3kkel"; BDORZ=FFFB88E999055A3F8A630C64834BD6D0; ariaDefaultTheme=undefined; BAIDUID_BFESS=848DBD23899C450955664C581BB8F88F:FG=1; BA_HECTOR=0g2h008k8la0ak24gr1h7e6jb0q; BDRCVFR[7FEYkXni5q3]=mk3SLVN4HKm; BDRCVFR[r_74CZIkx53]=aeXf-1x8UdYcs; delPer=0; PSINO=1; H_PS_PSSID=36309_31253_34812_34584_35978_36281_36234_26350_36311_36061; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1651809104,1651975540; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1651978656; ab_sr=1.0.1_YmI3MWI2NjRkY2ZkMDVhOTNmZDE4Njg2YzVkNzM3YmEwZDBjMjMyYTA5OWNmN2YyOTI4Y2MzNmNjZGJmNDg4NTk3Y2EzNzdlN2RiZmZjNzIyYmU5N2EwNDVjZmFhOWVlNzZhMWJmYTY2M2E4ZWNmNzU3MDM4NmU5YWViMWFkMzZlNjE3NGUzZWY0N2M1MjY0NzZhNTA1Zjg4MjE4NmY2NGE2NWQ5NDI5NTI0MjMwZTQyNWE4ZmQ0Yzk5YmEyYTlk',
}
data={
    'from':' en',
    'to':' zh',
    'query':' spider',
    'transtype':' realtime',
    'simple_means_flag':' 3',
    'sign':' 63766.268839',
    'token':' 19836812e18b0c6a1c84a7f31e610e1c',
    'domain':' common',
}

# post请求的参数  必须进行编码 并且要调用encode方法
data = urllib.parse.urlencode(data).encode('utf-8')
# 请求对象的定制
request=urllib.request.Request(url=url,data=data,headers=headers)
# 模拟浏览器向服务器发送请求
response=urllib.request.urlopen(request)

# 获取响应的数据
content=response.read().decode('utf-8')
print(content)
import json
obj=json.loads(content)
print(obj)

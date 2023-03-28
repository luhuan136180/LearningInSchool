#https://movie.douban.com/j/chart/top_list?type=5&interval_id=100%3A90&action=&start=0&limit=20

#https://movie.douban.com/j/chart/top_list?type=5&interval_id=100%3A90&action=&start=20&limit=20

#https://movie.douban.com/j/chart/top_list?type=5&interval_id=100%3A90&action=&start=40&limit=20

# page    1  2   3   4
# start   0  20  40  60

# start （page - 1）*20

# 下载豆瓣电影前10页的数据
# （1） 请求对象的定制
# （2） 获取响应的数据
# （3） 下载数据
import urllib.request
import urllib.parse
def creat_request(page):
    base_url='https://movie.douban.com/j/chart/top_list?type=5&interval_id=100%3A90&action=&'
    headers={
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36'
    }
    data={
        'start':(page-1)*20,
        'limit':'20'
    }
    url=base_url+urllib.parse.urlencode(data)
    request=urllib.request.Request(url=url,headers=headers)
    return request

def  get_content(request):
    response=urllib.request.urlopen(request)
    content=response.read().decode('utf-8')
    return content

def down_load(page,content):
    with open('douban_' + str(page) + '.json','w',encoding='utf-8')as fp:
        fp.write(content)

if  __name__ =='__main__':
    start_page=int(input('请输入其实的页码'))
    end_paga=int(input('请输入结束的页面'))
    for page in range(start_page,end_paga+1):
        request=creat_request(page)
        #获取响应的数据
        content= get_content(request)
        #下载
        down_load(page,content)
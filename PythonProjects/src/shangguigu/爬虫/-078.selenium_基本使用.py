from selenium import webdriver
from selenium.webdriver.chrome.service import Service

# 尝试传参
# (2) 创建浏览器操作对象
s = Service("chromedriver.exe")

browser = webdriver.Chrome(service=s)
# url = 'https://www.baidu.com'
#
# driver.get(url)


url='https://www.jd.com/'

browser.get(url)
# page_source获取网页源码
content=browser.page_source
print(content)



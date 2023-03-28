from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By

path=Service('chromedriver.exe')
browser=webdriver.Chrome(service=path)
url='https://www.baidu.com'
browser.get(url)

# 元素定位

# 根据id来找到对象
button=browser.find_element(By.ID,'su')
print(button)

# 根据标签属性的属性值来获取对象的
button=browser.find_element(By.NAME,'wd')
print(button)
# 根据xpath语句来获取对象
button=browser.find_element(By.XPATH,'//input[@id="su"]')
print(button)

# 根据标签的名字来获取对象
button=browser.find_element(By.TAG_NAME,'input')
print(button)

# 使用的bs4的语法来获取对象
button=browser.find_element(By.CSS_SELECTOR,'#su')

#链接文本
#链接文本
button=browser.find_element(By.LINK_TEXT,'新闻')

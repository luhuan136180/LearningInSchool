from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
def share_browser():
    chrome_options = Options()
    chrome_options.add_argument('--headless')
    chrome_options.add_argument('--disable-gpu')
    service = Service("chromedriver.exe")
    browser = webdriver.Chrome(options=chrome_options,service=service)
    return browser

browser = share_browser()
url = 'https://www.baidu.com'
browser.get(url)
input =browser.find_element(By.ID,'su')

# 获取标签的属性
print(input.get_attribute('class'))
# 获取标签的名字
print(input.tag_name)

# 获取元素文本
button=browser.find_element(By.LINK_TEXT,'新闻')
print(button.text)


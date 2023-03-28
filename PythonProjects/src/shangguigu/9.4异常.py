try:
    f = open('test.txt', 'r')
    print(f.read())
except FileNotFoundError:
    print('文件没有找到,请检查文件名称是否正确')
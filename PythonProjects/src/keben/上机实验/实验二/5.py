# 给定一个整数（可以是负的），请将该数各个位上数字反转得到一个新数。新数不得有多余的前导0。i=in
def reverse(x):
    if int(x) == 0:
        return 0

    x = x.replace("-","")  # abs() 也可以用
    re = x[::-1]  # 反转
    re = [-1,1][int(x) > 0] * int(re)  # 正负 if - else 判断
    if not -2**31 <= re <= 2**31-1:
        return 0
    return re

if __name__ == "__main__":
    x = input('qingshuru一个数')
    print(reverse(x))

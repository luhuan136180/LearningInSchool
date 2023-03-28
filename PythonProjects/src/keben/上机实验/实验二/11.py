n = input('请输入的整数：')
lst = []
for i in range(int(n)):
    z = eval(input('请输入总病例数：'))
    lst.append(z)
    y = eval(input('请输入有效病历数：'))
    lst.append(y)

j = lst[1] / lst[0]

num = 1
for o in range(2, len(lst), 2):
    p = lst[o+1]/lst[o]
    if p-j > 0.05:
        print(f"第{num}组：better")
    elif j-p > 0.05:
        print(f"第{num}组：worse")
    else:
        print(f"第{num}组：same")
    num += 1
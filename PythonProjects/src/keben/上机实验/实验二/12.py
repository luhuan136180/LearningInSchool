list1=['jack','marry','tom','smith','lion']
list2=['jack','tom','lion']
list3=['marry','tom','smith']
program1={"name":"爱泉护泉",'名单':list1}
program2={"name":"绿色出行",'名单':list2}
program3={"name":"勤俭节约",'名单':list3}
list_a = [a for a in list1 if a in list2]
list_b = [a for a in list3 if a in list2]
list_c = [a for a in list_b if a in list_a]
print('参见了三个活动的人有：',list_c)
list_a.extend(list_b)
set1=set(list_a)
print('参见了两个活动的人有：',set1)

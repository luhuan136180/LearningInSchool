package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pushV int整型一维数组
 * @param popV int整型一维数组
 * @return bool布尔型
 */
/*
思路：创建切片模拟一个临时栈
        遍历push数组，将元素按顺序输入{模拟栈的压入}切片中
        每输入一次，就将temp栈顶和输出数组的指针指向元素比较，若栈不为空且栈顶元素等于指向的元素
        栈压出，指针向后移
        若最后temp栈为空，及全部按照popV数组输出，true
*/
func IsPopOrder(pushV []int, popV []int) bool {
	// write code here

	temp := make([]int, 0) //创建一个临时切片（）
	i := 0
	for _, v := range pushV {
		temp = append(temp, v)
		for len(temp) != 0 && temp[len(temp)-1] == popV[i] {
			temp = temp[:len(temp)-1] //删除最后一个元素
			i++
		}

	}

	return len(temp) == 0
}

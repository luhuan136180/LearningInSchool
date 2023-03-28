package com.suanfafenxi.test2;

import java.util.Arrays;

public class insertSort1 {
    public static void main(String[] args) {
        int[] arr = {6, 3, 8, 2, 9, 1, 19, 29};
        insertionSort(arr);
        System.out.println("2020131128" + "毛浩昕");
        System.out.println(Arrays.toString(arr));
        int k1=19;
        int k2=7;
        System.out.println("查询数字"+k1);
        System.out.println(binarySearch(arr, 0, arr.length-1, k1));

        System.out.println("查询数字"+k2);
        System.out.println(binarySearch(arr, 0, arr.length-1, k2));
    }

    public static void insertionSort(int[] arr) {
        int count=0;

        for (int i = 1; i < arr.length; i++) {
            int insertVal = arr[i];
            int j = i - 1;//给insertval找插入的位置

            while (j >= 0) {
                if (insertVal < arr[j]) {
                    arr[j + 1] = arr[j];
                   count = count+1;
                } else {
                    count = count + 1;
                    break;
                }
                j--;
            }

            if (j + 1 != i) {
                arr[j + 1] = insertVal;
            }
        }
        System.out.println(count);
    }

    public static int binarySearch(int[] arr, int left, int right, int k) {
        if (left > right) {
            return -1;
        }
        int mid = (left + right) / 2;

        if (k == arr[mid]) {
            return mid;
        }else if (k<arr[mid]){
            return binarySearch(arr,left,mid-1,k);
        }else if(k>arr[mid]){
            return binarySearch(arr,mid+1,right,k);
        }else{
            return -1;
        }

    }
}


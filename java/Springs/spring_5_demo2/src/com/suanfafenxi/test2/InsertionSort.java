package com.suanfafenxi.test2;

import java.util.Arrays;

public class InsertionSort {
    public static void main(String[] args) {
        int[] arr = {6, 3, 8, 2, 9, 1, 19, 29};
        insertionSort(arr);
        System.out.println("2020131128"+"毛浩昕");
        System.out.println(Arrays.toString(arr));

    }
    public static void insertionSort(int[] arr){
        for(int i=1;i<arr.length;i++) {
            int insertVal=arr[i];
            int j=i-1;//给insertval找插入的位置
            while(j>=0&&insertVal<arr[j]) {
                arr[j+1]=arr[j];
                j--;
            }
            if(j+1!=i) {
                arr[j+1]=insertVal;
            }
        }

    }
}

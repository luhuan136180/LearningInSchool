package com.suanfafenxi.test1;

import java.lang.reflect.Array;
import java.util.Arrays;

public class SelectkSort {
    public static void main(String[] args) {
        int[] arr = {6, 3, 8, 2, 9, 1, 19, 29};
        System.out.println("2020131135"+','+"毛浩昕");
        SelectSort(arr);

        System.out.println(Arrays.toString(arr));

    }

    public static void SelectSort(int[] arr){
        int temp=0;//设置
        int count=0;
        for(int i=0;i<arr.length-1;i++){
            int min=arr[i];
            int minIndex=i;
            for(int j=i+1;j<arr.length;j++){
                if(min>arr[j]) {
                    min=arr[j];
                    minIndex=j;
                }
                count=count+1;
                }
            if(minIndex!=i){
                arr[minIndex]=arr[i];
                arr[i]=min;
            }
        }
        System.out.println("一共比较了："+count);
    }
}

package com.suanfafenxi.test1;
import java.util.*;
public class BubbleSort {
    public static void main(String []args){
        int[] arr = {6, 3, 8, 2, 9, 1, 19, 29};
        System.out.println("2020131135"+','+"唐浩航");


        System.out.println("::"+Arrays.toString(arr));
    }
    public static void bubbleSort(int[] arr) {
        int temp=0;
        int count=0;
        boolean flag=false;//创造标识，表明是否进行交换

        for(int i=0;i<arr.length-1;i++) {
            for(int j=0;j<arr.length-1-i;j++) {
                if(arr[j]>arr[j+1]) {
                    flag=true;
                    temp=arr[j];
                    arr[j]=arr[j+1];
                    arr[j+1]=temp;

                }
                count=count+1;
            }
            if(flag) {
                flag=false;
            }else {
                break;
            }

        }
        System.out.println("一共比较了："+count);
    }
}



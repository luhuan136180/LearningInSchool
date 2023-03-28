package com.suanfafenxi.test3;

import java.util.Arrays;

public class quickSort {
    public static void main(String[] args) {
        int[] arr = {6, 3, 8, 2, 9, 1, 19, 29};
        quickSort(arr,0,arr.length);
        System.out.println("2020131128" + "毛浩昕");

        System.out.println(Arrays.toString(arr));

    }

       public static void quickSort(int[] a,int lo, int hi){
                 if(lo<hi){
                         int p=partition(a, lo, hi);
                         quickSort(a,lo,p-1);
                         quickSort(a, p+1, hi);
                     }
             }
    public static int partition(int[] a,int lo, int hi){
                 int i,j,privot;
                 privot=a[hi];
                 i=lo;
                 for(j=lo;j<=hi-1;j++){
                         if(a[j]<privot){
                                 swap(a, i, j);
                                 i++;
                             }
                     }
                 swap(a,i,hi);
        return i;
             }
     public static void swap(int[] a,int lo,int hi){
                 int s=a[lo];
                 a[lo]=a[hi];
                 a[hi]=s;
             }

}

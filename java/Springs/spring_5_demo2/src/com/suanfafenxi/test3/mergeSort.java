package com.suanfafenxi.test3;

import java.util.Arrays;

public class mergeSort {
    public  static int count=0;
    public static void main(String[] args) {
        int[] arr = {6, 3, 8, 2, 9, 1, 19, 29};

        mergeSort(arr);
        System.out.println("2020131128" + "毛浩昕");
        System.out.println(count);
        System.out.println(Arrays.toString(arr));
    }
    public  static void mergeSort(int[] n){
        if(n.length>1){
            int mid =n.length/2;
            int[] b=new int[mid];
            int[] c=new int[mid];
            System.arraycopy(n,0,b,0,mid);
            System.arraycopy(n,mid,c,0,n.length-mid);

            meger(n,b,c,count);


        }
    }
    public static void meger(int[] a,int[] b ,int[] c,int count){
        int i=0;
        int j=0;
        int k=0;
        while(i<b.length&&j<c.length){
            if(b[i]<=c[j]){
                a[k]=b[i];
                i++;
                count=count+1;
            }else {
                a[k] = c[j];
                j++;
                count=count+1;
            }
            k++;
            }

        if (i==b.length){
            System.arraycopy(c,j,a,k,c.length-j);
        }else{
            System.arraycopy(b,i,a,k,b.length-i);
        }
    }
}

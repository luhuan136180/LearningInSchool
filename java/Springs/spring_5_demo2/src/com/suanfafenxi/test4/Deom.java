package com.suanfafenxi.test4;

import java.util.Arrays;

public class Deom {
    public static int[][] DPKnapsack(int[] weight, int[] value, int maxweight) {
        int n=weight.length;
        int[][] maxvalue=new int[n+1][maxweight+1];

        for (int i = 0; i < maxweight+1; i++) {
            maxvalue[0][i]=0;
            System.out.print(maxvalue[0][i]);
        }

        for (int i = 0; i < n+1; i++) {
            maxvalue[i][0]=0;
            System.out.print(maxvalue[i][0]);
        }

        for (int i = 1; i <=n; i++) {
            for (int j = 1; j <=maxweight; j++) {
                maxvalue[i][j]=maxvalue[i-1][j];

                if (weight[i-1]<=j) {
                    if (maxvalue[i-1][j-weight[i-1]]+value[i-1]>maxvalue[i-1][j]){
                        maxvalue[i][j] = maxvalue[i-1][j-weight[i-1]]+value[i-1];
                    }
//                    System.out.println(maxvalue[i][j]);
                }
            }
        }
//        for (int[] date : maxvalue) {
//            System.out.println(Arrays.deepToString(maxvalue));
//        }
        return maxvalue;
    }

    public static void main(String[] args) {
        int weight[]= {0,7,3,4,5,6,2,8};
        int value[]= {0,42,1,40,25,30,19,50};
        int maxweight = 21;

//        System.out.println(DPKnapsack(weight, value, maxweight));
        int[][] ints = DPKnapsack(weight, value, maxweight);
        System.out.println(Arrays.deepToString(ints));
        // for (int i = 0; i < ints[0].length; i++) {
        //   for (int j = 0; j < ints.length; j++) {

        // }
        // }

    }
    public static int[] optimalKnapSack(int[][] V,int[] w,int[] v,int W){
        int[] L = new int[w.length];
        int k=0;
        int j = W;
        for (int i=w.length; i>1; i--){
            if (V[i][j] > V[i-1][j]){
                k++;
                L[k] = i-1;  // include item i;
                j = j-w[i-1];
            }
        }
        return L;
    }
}

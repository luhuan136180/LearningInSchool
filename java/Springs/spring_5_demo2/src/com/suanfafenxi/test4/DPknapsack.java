package com.suanfafenxi.test4;

import java.util.Arrays;

public class DPknapsack {
    public static void main(String[] args) {
        int[] weight = {0,7,3,4,5,6,2,8};
        int[] value = {0,42,12,40,25,30,19,50};
        int W=21;
        int[][] d=dpKnapsack(weight,value,W);
        for (int i = 1; i <= weight.length; i++) {
            for (int j = 1; j <= W; j++) {
                System.out.print(d[i][j] + "\t");
                if (j == W) {
                    System.out.println();
                }
            }
        }
        int[] max= printResult(d,weight);
        System.out.println(Arrays.toString(max));

        System.out.println();
    }
    public static int[][] dpKnapsack(int[] w, int[] v, int C){
        int length = w.length;

        int[][] dp = new int[length][C + 1];
        //初始化第一行
        //仅考虑容量为C的背包放第0个物品的情况
        for (int i = 0; i <= C; i++) {
            dp[0][i] = w[0] <= i ? v[0] : 0;
        }
        //填充其他行和列
        for (int i = 1; i < length; i++) {
            for (int j = 0; j <= C; j++) {
                dp[i][j] = dp[i - 1][j];
                if (w[i] <= j) {
                    dp[i][j] = Math.max(dp[i][j], v[i] + dp[i - 1][j - w[i]]);
                }
            }
        }
        return dp;
    }
    public static int[] optimalKnapSack(int[][] V,int[] w,int[] v,int W){
        int[] L = new int[w.length+1];
        int k=0;
        int j = W;
        for (int i=w.length-1; i>=0; i--){
            if (V[i][j] > V[i-1][j]){
                k=k+1;
                L[k] = i-1;  // include item i;
                j = j-w[i-1];
            }
        }
        return L;
    }
    public static int[] optimalKnapSack1(int[][] V,int[] w,int[] v,int W){
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
    public static int[] printResult(int[][] F, int[] weight) {
        int v = 21;
        int n = 7;
        boolean[] isAdd = new boolean[n + 1];
        for(int i = n; i >= 1; i--) {
            if(F[i][v] == F[i-1][v])
                isAdd[i] = false;
            else {
                isAdd[i] = true;
                v -= weight[i];
            }
        }

        int[] re = new int[n + 1];
        for (int i = 0; i < isAdd.length; i++) {
            if(isAdd[i]){
                re[i] = 1;
            }else {
                re[i] = 0;
            }
        }

        return re;
    }
}

package com.suanfafenxi.test4;

public class napsack {

    //初始化二维数组

    public static int N;//物品个数
    public static int V;//背包载重
    public static int[] weight = {0,7,3,4,5,6,2,8};

    public static int[] value = {0,42,12,40,25,30,19,50};
    public static void main(String[] args) {
        N = weight.length - 1;
        V = 21;

        F1 = new int[N + 1][V + 1];
        int[][] res = DPknapsack(weight, value);

        for(int i = 0; i < res.length; i++) {
            for (int j = 0; j < res[i].length; j++) {
                System.out.print(res[i][j] + " ");
            }
            System.out.println();
        }
        int[] ints = printResult(res, weight);

        for (int i = 0; i < ints.length; i++) {
            System.out.print(ints[i] + " " ) ;
        }

        //贪心算法
        initF1();



        System.out.println("贪心");
        for (int i = 0; i < F1.length; i++) {
            for (int j = 0; j < F1[i].length; j++) {
                System.out.print(F1[i][j] + " ");
            }
            System.out.println();
        }
    }

    public static int[][] DPknapsack(int[] w, int[] v){
        int[][] F = new int[N+1][V+1];
        for(int i = 1; i <= N; i++) {
            for(int j = 0; j <= V; j++) {
                //如果容量为j的背包放得下第i个物体
                if(j >= w[i]) {
                    F[i][j] = Math.max(F[i - 1][j - w[i]] + v[i], F[i - 1][j]);
                }else {
                    //放不下，只能选择不放第i个物体
                    F[i][j] = F[i - 1][j];
                }
            }
        }
        return F;
    }

    /**
     *
     * @param F 二维数组
     * @param weight weight数组
     */
    public static int[] printResult(int[][] F, int[] weight) {
        int v = V;
        int n = N;
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
    public static int F1[][];

    public static void initF1(){
        for (int i = 0; i < F1.length; i++) {
            for (int j = 0; j < F1[i].length; j++) {
                F1[i][j] = -1;
            }
        }
    }

}
package com.suanfafenxi.test4;
public class Untitled {
    public static void main(String[] args) {
        int[] weight = {0, 7, 3, 4, 5, 6, 2, 8};
        int[] value = {0, 42, 12, 40, 25, 30, 19, 50};
        int bagSize = 21;

        int[][] dp;
        dp = testWeightBagProblem(weight, value, bagSize);

        //打印dp数组
        for (int i = 0; i <= weight.length; i++){
            for (int j = 0; j <= bagSize; j++){
                System.out.print(dp[i][j] + " ");
            }
            System.out.print("\n");
        }

        int[] item = optimalKnapsack(dp, weight, value, bagSize);

        for(int i=0; i<item.length; i++) {
            System.out.print(item[i] + " ");
        }

    }

    public static int[][] testWeightBagProblem(int[] weight, int[] value, int bagSize){
        int wLen = weight.length, value0 = 0;
        //定义dp数组：dp[i][j]表示背包容量为j时，前i个物品能获得的最大价值
        int[][] dp = new int[wLen + 1][bagSize + 1];
        //初始化：背包容量为0时，能获得的价值都为0
        for (int i = 0; i <= wLen; i++){
            dp[i][0] = value0;
        }
        //遍历顺序：先遍历物品，再遍历背包容量
        for (int i = 1; i <= wLen; i++){
            for (int j = 1; j <= bagSize; j++){
                if (j < weight[i - 1]){
                    dp[i][j] = dp[i - 1][j];
                }else{
                    dp[i][j] = Math.max(dp[i - 1][j], dp[i - 1][j - weight[i - 1]] + value[i - 1]);
                }
            }
        }

        return dp;
    }

    public static int[] optimalKnapsack(int[][] V, int[] w, int v[], int W){
        int[] item = new int[w.length];

        findWhat(w.length-1, W, V, item, w, v);

        return item;
    }

    static void findWhat(int i, int j, int[][] dp, int[] item, int[] w, int[] v) {				//最优解情况
        if (i >= 0) {
            if (dp[i][j] == dp[i - 1][j]) {
                item[i] = 0;
                findWhat(i - 1, j, dp, item, w, v);
            }
            else if (j - w[i] >= 0 && dp[i][j] == dp[i - 1][j - w[i]] + v[i]) {
                item[i] = 1;
                findWhat(i - 1, j - w[i], dp, item, w, v);
            }
        }
    }

}

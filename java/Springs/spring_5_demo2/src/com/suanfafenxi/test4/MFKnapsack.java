package com.suanfafenxi.test4;



public class MFKnapsack {
    public static void main(String[] args) {
        int[] weight = {0,7,3,4,5,6,2,8}; //物品重量 // [6, 5] [2,3][3,4]
        int[] val = {0,42,12,40,25,30,19,50}; //物品价值
        int m = 21; //背包容量
        int n = val.length; //物品个数

        int[][] f = new int[n+1][m+1]; //f[i][j]表示前i个物品能装入容量为j的背包中的最大价值
        int[][] path = new int[n+1][m+1];
        //初始化第一列和第一行
        for(int i=0;i<f.length;i++){
            f[i][0] = 0;
        }
        for(int i=0;i<f[0].length;i++){
            f[0][i] = 0;
        }
        //通过公式迭代计算
        for(int i=1;i<f.length;i++){
            for(int j=1;j<f[0].length;j++){
                if(weight[i-1]>j)
                    f[i][j] = f[i-1][j];
                else{
                    if(f[i-1][j]<f[i-1][j-weight[i-1]]+val[i-1]){
                        f[i][j] = f[i-1][j-weight[i-1]]+val[i-1];
                        path[i][j] = 1;
                    }else{
                        f[i][j] = f[i-1][j];
                    }
                    //f[i][j] = Math.max(f[i-1][j], f[i-1][j-weight[i-1]]+val[i-1]);
                }
            }
        }
        for(int i=0;i<f.length;i++){
            for(int j=0;j<f[0].length;j++){
                System.out.print(f[i][j]+" ");
            }
            System.out.println();
        }

        int i=f.length-1;
        int j=f[0].length-1;
        while(i>0&&j>0){
            if(path[i][j] == 1){
                System.out.print("第"+i+"个物品装入 ");
                j -= weight[i-1];
            }
            i--;
        }
    }
}

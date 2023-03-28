package com.suanfafenxi.test1;
import java.util.ArrayList;
import java.util.List;

/**
 * 整数规划中的 0-1背包问题
 */
public class KnapsackProblem {

    public static void main(String[] args) {
        // 启动
        exhaustAlgorithm();
    }

    /**
     * 1.穷举法
     *
     */
    public static void  exhaustAlgorithm(){

        // 1.算例定义：背包限制重量100  有4个物品 每个物品的重量和价值
        double knapsackLimit = 15;
        int goodsNum = 4;
        int[] weight = {7,3,4,5,6};
        int[] value = {42,12,40,25,30};

        // 2.根据商品数量生成解空间(递归+回溯)
        List<List<Integer>> solutionSpace = getSolutionSpace(goodsNum);

        // 3.评价每个解
        double maxVal = 0;
        List<Integer> solve = null;
        for(int i=0; i<solutionSpace.size();i++){
            List<Integer> curr = solutionSpace.get(i);
            double w = 0;
            double v = 0;
            for(int j=0; j<curr.size(); j++){
                if(curr.get(j) == 1){
                    w += weight[j];
                    v += value[j];
                }
            }

            // 满足约束条件
            if(w<=knapsackLimit){
                if(v>=maxVal){
                    maxVal = v;
                    solve = curr;
                }
            }
        }

        // 4.输出最优解
        System.out.println("最大价值："+ maxVal + "  最优方案："+solve);
    }

    /**
     * 穷举解空间中的所有解
     */
    public static List<List<Integer>> getSolutionSpace(int goodsNum){

        List<List<Integer>> result = new ArrayList<>();
        List<Integer> path = new ArrayList<>();

        // 递归的层数 相当于每个物品就是一层for
        Integer index = 1;
        Integer length = goodsNum;

        // 递归
        dfs(result,path,length,index);

        System.out.println("解空间元素个数: "+result.size());
        System.out.println("解空间详情：");
        result.forEach(System.out::println);

        return result;
    }

    public static void dfs(List<List<Integer>> result,List<Integer> path, int length, int index){

        // 终止条件
        if(index>length){
            // 获取一个解
            result.add(new ArrayList<>(path));
            return;
        }

        // 每个物品有两种选择方式 1：选   0：不选
        for(int i=0; i<2; i++){
            path.add(i);
            dfs(result,path,length,index+1);
            // 回溯
            path.remove(index-1);
        }
    }
}


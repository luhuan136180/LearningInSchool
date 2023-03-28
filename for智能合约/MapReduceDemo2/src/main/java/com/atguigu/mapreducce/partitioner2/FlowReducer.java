package com.atguigu.mapreducce.partitioner2;

import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Reducer;

import java.io.IOException;


public class FlowReducer extends Reducer<Text, FlowBean,Text, FlowBean> {
    private FlowBean outV = new FlowBean();

    @Override
    protected void reduce(Text key, Iterable<FlowBean> values, Reducer<Text, FlowBean, Text, FlowBean>.Context context) throws IOException, InterruptedException {

        //遍历集合累加值
        long totalUp =0;
        long totaldown = 0;
        for (FlowBean value : values) {
            totalUp += value.getUpFlow();
            totaldown += value.getDownFlow();
        }
        //2.封装outK，
        outV.setUpFlow(totalUp);
        outV.setDownFlow(totaldown);
        outV.setSumFlow();

        //写出
        context.write(key,outV);
    }
}

package com.atguigu.mapreducce.writableComparable;


import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

//mapper四个参数：行偏移量，一行的数据，数据（手机号） ，flowbean（自行设置的输出数据，包含上行流量，下行流量...）

public class FlowMapper extends Mapper<LongWritable,Text, FlowBean, Text> {

    private FlowBean outK = new FlowBean();
    private Text outV = new Text();


    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, FlowBean, Text>.Context context) throws IOException, InterruptedException {


        //获取一行
        String line = value.toString();

        //切割
        String[] split =line.split("\t");

        //封装
        outV.set(split[0]);
        outK.setUpFlow(Long.parseLong(split[1]));
        outK.setDownFlow(Long.parseLong(split[2]));
        outK.setSumFlow();
        //
        context.write(outK,outV);
    }
}

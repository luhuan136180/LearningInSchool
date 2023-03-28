package com.atguigu.mapreducce.partitioner2;

import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

//mapper四个参数：行偏移量，一行的数据，数据（手机号） ，flowbean（自行设置的输出数据，包含上行流量，下行流量...）

public class FlowMapper extends Mapper<LongWritable,Text, Text, FlowBean> {

    private Text outK =new Text();
    private FlowBean outV = new FlowBean();


    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, Text, FlowBean>.Context context) throws IOException, InterruptedException {
        //获取一行
        //1	13736230513	192.196.100.1	www.atguigu.com	2481	24681	200
        String line = value.toString();

        //切割
        String[] split = line.split("\t");

        //3.抓取想要的数据
        //手机号
        //上行流量和下行流量：。。。。。
        String phone =split[1];
        String up =split[split.length -3];
        String down =split[split.length -2];


        //封装
        outK.set(phone);
        outV.setUpFlow(Long.parseLong(up));
        outV.setDownFlow(Long.parseLong(down));
        outV.setSumFlow();


        //写出
        context.write(outK,outV);
    }
}

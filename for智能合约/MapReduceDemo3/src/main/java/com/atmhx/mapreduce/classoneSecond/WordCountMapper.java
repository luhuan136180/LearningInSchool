package com.atmhx.mapreduce.classoneSecond;

import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

//
public class WordCountMapper extends Mapper<LongWritable, Text, IntWritable,Text> {

    private Text outV = new Text();
    private IntWritable outK = new IntWritable();

    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, IntWritable, Text>.Context context) throws IOException, InterruptedException {

        //获取一行数据
        String line = value.toString();

        //切片
        //[时间,liuliang]
        String[] split = line.split("\t");

        //
        outK.set(Integer.parseInt(split[1]));
        outV.set(split[0]);

        context.write(outK,outV);
    }
}

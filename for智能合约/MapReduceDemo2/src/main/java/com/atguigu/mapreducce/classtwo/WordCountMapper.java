package com.atguigu.mapreducce.classtwo;

import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

//
public class WordCountMapper extends Mapper<LongWritable, Text,Text, IntWritable> {

    private Text outk = new Text();
    private IntWritable outV = new IntWritable();

    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, Text, IntWritable>.Context context) throws IOException, InterruptedException {
        //super.map(key, value, context);

        //1.获取一行
        //zhangsan	88
        String line = value.toString();

        //2.切割
        // [zhangsan,88]
        String[] words = line.split("\t");

        //循环写出
        outk.set(words[0]);
        outV.set(Integer.parseInt(words[1]));
        context.write(outk,outV);
    }
}

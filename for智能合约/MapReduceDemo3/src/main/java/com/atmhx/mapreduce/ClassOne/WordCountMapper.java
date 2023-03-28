package com.atmhx.mapreduce.ClassOne;

import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

//
public class WordCountMapper extends Mapper<LongWritable, Text,Text, IntWritable> {

    private Text outk = new Text();
    private IntWritable outV = new IntWritable(1);

    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, Text, IntWritable>.Context context) throws IOException, InterruptedException {
        //super.map(key, value, context);

        //1.获取一行
        //Nehru,2016-01-01
        String line = value.toString();

        //2.切割
        // [Nehru,2016-01-01]
        String[] words = line.split(",");

        //循环写出
        outk.set(words[1]);
        context.write(outk,outV);
    }
}

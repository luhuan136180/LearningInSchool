package com.atmhx.mapreduce.classthree;

import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

//
public class WordCountMapper extends Mapper<LongWritable, Text,Text, Text> {

    private Text outk = new Text();
    private Text outV = new Text();

    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, Text, Text>.Context context) throws IOException, InterruptedException {
        //1.获取一行
        //13718855152 112
        String line = value.toString();

        //2.切割
        //[13718855152,112]
        String[] words = line.split("\t");

        //循环写出
        outk.set(words[1]);
        outV.set(words[0]);
        context.write(outk,outV);
    }
}

package com.atmhx.mapreduce.classtwo;


import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Reducer;

import java.io.IOException;

public class WordCountReducer extends Reducer<Text, IntWritable,Text,IntWritable> {

    int sum;
    IntWritable outV = new IntWritable();

    @Override
    protected void reduce(Text key, Iterable<IntWritable> values, Reducer<Text, IntWritable, Text, IntWritable>.Context context) throws IOException, InterruptedException {
        int sum=0;
        int num=0;
        for (IntWritable value : values){
            sum+=value.get();
            num+=1;
        }

        outV.set(sum/num);

        //写出
        context.write(key,outV);
    }
}

package com.atmhx.mapreduce.classoneSecond;


import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Reducer;

import java.io.IOException;

public class WordCountReducer extends Reducer<IntWritable,Text,Text,IntWritable> {

    int sum;
    IntWritable outV = new IntWritable();

    @Override
    protected void reduce(IntWritable key, Iterable<Text> values, Reducer<IntWritable, Text, Text, IntWritable>.Context context) throws IOException, InterruptedException {
        //
        for (Text value:values){
            context.write(value,key);
        }
    }
}

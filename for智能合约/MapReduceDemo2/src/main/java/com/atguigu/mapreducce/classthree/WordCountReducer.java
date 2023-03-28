package com.atguigu.mapreducce.classthree;


import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Reducer;

import java.io.IOException;

public class WordCountReducer extends Reducer<Text, Text,Text,Text> {

    Text outV = new Text();

    @Override
    protected void reduce(Text key, Iterable<Text> values, Reducer<Text, Text, Text, Text>.Context context) throws IOException, InterruptedException {
        //
        String words="";
        for (Text value : values){
            String line = value.toString();
            words += line+"\t";
        }

        outV.set(words);
        context.write(key,outV);
    }
}

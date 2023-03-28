package com.atguigu.mapreducce.ClassOne;


import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Partitioner;

//kv对，需要和mapper的输出以及reducer的输入KV 对保持相同
public class TimePartitioner extends Partitioner<Text, IntWritable> {


    @Override
    public int getPartition(Text text, IntWritable intWritable, int numPartitions) {
        //2016-01-01
        String Time = text.toString();
        String timepre = Time.substring(5,7);

        //设置分区号
        int partition;

        if ("01".equals(timepre)){
            partition =  0;
        }else if ("02".equals(timepre)){
            partition = 1;
        }else {
            partition = 2;
        }

        return partition;
    }
}

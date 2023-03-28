package com.atguigu.mapreducce.partitioner2;

import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Partitioner;

//kv对，需要和mapper的输出以及reducer的输入KV 对保持相同
public class ProvincePartitioner extends Partitioner<Text, FlowBean> {
    @Override
    public int getPartition(Text text, FlowBean flowBean, int numPartitions) {

        //text是手机号
        String phone = text.toString();
        String prephone = phone.substring(0,3);

        //设置分区号
        int partition;

        if ("136".equals(prephone)){
            partition = 0;//分区0
        }else if ("137".equals(prephone)){
            partition = 1;
        }else if ("138".equals(prephone)){
            partition = 2;
        }else if ("139".equals(prephone)){
            partition = 3;
        }else {
            partition = 4;;
        }


        return partition;
    }
}

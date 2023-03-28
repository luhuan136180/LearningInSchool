package com.atguigu.mapreducce.writableComparable;

import org.apache.hadoop.io.Writable;
import org.apache.hadoop.io.WritableComparable;

import java.io.DataInput;
import java.io.DataOutput;
import java.io.IOException;

/*
* 1.定义类实现writable接口
* 2.重写序列化和反序列化
* 3.重写空参构造
* 4.toString方法重写
*
* */
public class FlowBean implements WritableComparable<FlowBean> {

    private long upFlow;
    private long downFlow;
    private long sumFlow;

    //重写空参构造
    public FlowBean(){

    }



    public long getUpFlow() {
        return upFlow;
    }

    public void setUpFlow(long upFlow) {
        this.upFlow = upFlow;
    }

    public long getDownFlow() {
        return downFlow;
    }

    public void setDownFlow(long downFlow) {
        this.downFlow = downFlow;
    }

    public long getSumFlow() {
        return sumFlow;
    }

    public void setSumFlow(long sumFlow) {
        this.sumFlow = sumFlow;
    }
    public void setSumFlow() {
        this.sumFlow = this.downFlow + this.upFlow;
    }

    @Override
    public void write(DataOutput dataOutput) throws IOException {

        dataOutput.writeLong(upFlow);
        dataOutput.writeLong(downFlow);
        dataOutput.writeLong(sumFlow);
    }

    @Override
    public void readFields(DataInput dataInput) throws IOException {

        this.upFlow= dataInput.readLong();
        this.downFlow = dataInput.readLong();
        this.sumFlow = dataInput.readLong();

    }

    @Override
    public String toString() {
        return upFlow + "\t" + downFlow + "\t" + sumFlow;
    }

    @Override
    public int compareTo(FlowBean o) {

        //总流量的倒叙排序 ：-1倒叙；
        //返回正数表示原值大于形参，负数表示原值小于形参，0就是俩相同
        if (this.sumFlow > o.sumFlow){
            return -1;
        }else if (this.sumFlow < o.sumFlow){
            return 1;
        }else {
            return 0;
        }

    }
}

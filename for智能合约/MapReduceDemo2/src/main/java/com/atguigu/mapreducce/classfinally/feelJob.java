package com.atguigu.mapreducce.classfinally;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.FileSystem;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;

public class feelJob {
	public static void main(String[] args) throws Exception {
		Configuration conf = new Configuration();
		conf.set("fs.defaultFS", "hdfs://192.168.83.133:9000");
		System.setProperty("HADOOP_USER_NAME","root");
		FileSystem fileSystem = FileSystem.get(conf);
		Job job = Job.getInstance(conf, "feelmodel");
		job.setJarByClass(feelJob.class);
		job.setMapperClass(feelMapper.class);
		job.setCombinerClass(feelReducer.class);
		job.setReducerClass(feelReducer.class);
		job.setOutputKeyClass(Text.class);
		job.setOutputValueClass(IntWritable.class);
		FileInputFormat
				.setInputPaths(job, new Path("hdfs://192.168.83.133:9000/input_2020131132/2020131132_上传文件.txt"));
		FileOutputFormat.setOutputPath(job, new Path(
				"hdfs://192.168.83.133:9000/output_2020131132/2020131132_模型"));
		
		int isSuccessed = job.waitForCompletion(true) ? 0 : 1;
		if (isSuccessed == 0) {
			System.out.println("执行成功");
	        fileSystem.copyToLocalFile(new Path(
					"hdfs://192.168.83.133:9000/output_2020131132/2020131132_模型/part-r-00000"), new Path("e:/2020131132_模型.txt"));
	        System.exit(isSuccessed);
		}
		
	}
}

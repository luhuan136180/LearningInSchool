package com.atguigu.hdfs;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.*;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.Arrays;

public class HdfsClient {

    private FileSystem fs;
    @Before
    public void init() throws IOException, URISyntaxException, InterruptedException{
        //连接的集群nn地址
        URI uri =new URI("hdfs://master:8020");
        //创建一个配置文件
        Configuration configuration =new Configuration();

        //用户
        String user = "mhx";

        fs = FileSystem.get(uri,configuration,user);
    }

    @After
    public void close() throws IOException {
        fs.close();
    }
    //创建文件夹
    @Test
    public void testmkdir() throws IOException, URISyntaxException, InterruptedException {
        //连接的集群nn地址

        //创建一个文件夹
        fs.mkdirs(new Path("/xiyou/huaguoshan546"));

        //关闭资源
        fs.close();
    }
    //上传
    /*
    * 参数优先级
    * hdfs-default.xml => hdfs-site.xml => 在项目资源目录下的配置文件 => 代码里面的设置
    *
    * */
    @Test
    public void testPut() throws IOException{
        fs.copyFromLocalFile(false,true,new Path("D:\\software\\data\\sunwukong.txt"),new Path("/xiyou/huaguoshan1"));
    }

    @Test
    public void testGet() throws IOException {
        //最后一个参数：false 表示开启校验
        fs.copyToLocalFile(false,new Path("hdfs://master:8020/xiyou/huaguoshan1"),new Path("D:\\software\\receive\\"),false);
    }


    @Test
    public void testRm() throws IOException {
        //参数：第一个参数：要删除的路径；第二个参数：是否递归删除

        //删除文件
        fs.delete(new Path("/mysql-community-client-8.0.25-1.el7.x86_64.rpm"),false);

        //删除空目录
//        fs.delete(new Path("/xiyou"),false);

        //删除非空目录
        fs.delete(new Path("/xiyou"),true);
    }

    //文件的更名和移动
    @Test
    public void testmv() throws IOException {
//        //第一个参数：目标原地址；第二个参数：目标文件路径
//        fs.rename(new Path("/input/word.txt"),new Path("/input/ss.txt"));
//
//
//        //移动
//        fs.rename(new Path("/input/ss.txt"),new Path("/ss.txt"));

        //目录更名
        fs.rename(new Path("/input"),new Path("/input-changeName"));
    }

    //获取文件详细信息
    @Test
    public void fileDetail() throws IOException {
        RemoteIterator<LocatedFileStatus> listFiles = fs.listFiles(new Path("/"),true);

        //遍历文件
        while(listFiles.hasNext()){
            LocatedFileStatus fileStatus = listFiles.next();

            System.out.println("========" + fileStatus.getPath() + "=========");
            System.out.println(fileStatus.getPermission());
            System.out.println(fileStatus.getOwner());
            System.out.println(fileStatus.getGroup());
            System.out.println(fileStatus.getLen());
            System.out.println(fileStatus.getModificationTime());
            System.out.println(fileStatus.getReplication());
            System.out.println(fileStatus.getBlockSize());
            System.out.println(fileStatus.getPath().getName());

            // 获取块信息
            BlockLocation[] blockLocations = fileStatus.getBlockLocations();
            System.out.println(Arrays.toString(blockLocations));
        }

    }

    @Test
    public void classwork() throws IOException{
        //创建一个文件夹
        fs.mkdirs(new Path("/input_2020131128"));
        //
        fs.copyFromLocalFile(false,true,new Path("C:\\Users\\半世琉璃\\Desktop\\作业\\hadoop\\课程考核2022\\test.txt"),new Path("/input_2020131128"));

    }


}

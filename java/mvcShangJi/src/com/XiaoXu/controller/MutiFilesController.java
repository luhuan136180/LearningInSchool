package com.XiaoXu.controller;

/**
 * @author 小栩
 * @version 1.0
 * @date 2021/12/1 16:56
 */
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.multipart.MultipartFile;
import com.XiaoXu.pojo.MultiFileDomain;
import javax.servlet.http.HttpServletRequest;
import java.io.File;
import java.util.List;

@Controller
public class MutiFilesController {

    private static  final Log logger = LogFactory.getLog(MutiFilesController.class);

    /**
     *多文件上传
     */
    @RequestMapping("/multifile")
    public String multiFileUpload(@ModelAttribute MultiFileDomain multiFileDomain,HttpServletRequest request){

        String realpath= request.getServletContext().getRealPath("uploadfiles");
        System.out.println(realpath);
        File targetDir= new File(realpath);

        if(!targetDir.exists()){
            targetDir.mkdirs();
        }
        List<MultipartFile> files = multiFileDomain.getMyfile();

        for(int i=0;i<files.size();i++){
            MultipartFile file = files.get(i);
            String fileName = file.getOriginalFilename();
            File targetFile = new File(realpath,fileName);

            //上传
            try{
                file.transferTo(targetFile);
            }catch (Exception e){
                e.printStackTrace();
            }
        }
        logger.info("成功");
        return "showMulti";
    }
}


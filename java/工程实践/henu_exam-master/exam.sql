/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50731
Source Host           : localhost:3306
Source Database       : exam

Target Server Type    : MYSQL
Target Server Version : 50731
File Encoding         : 65001

Date: 2020-12-25 19:38:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for exam
-- ----------------------------
DROP TABLE IF EXISTS `exam`;
CREATE TABLE `exam` (
  `e_id` int(11) NOT NULL AUTO_INCREMENT,
  `t_id` varchar(20) DEFAULT NULL COMMENT '˭?????',
  `e_name` varchar(20) DEFAULT NULL COMMENT '???Կ?Ŀ',
  `e_starttime` varchar(40) DEFAULT NULL COMMENT '??ʼʱ?',
  `e_stoptime` varchar(40) DEFAULT NULL,
  `e_autostart` tinyint(2) DEFAULT NULL,
  `e_status` tinyint(2) DEFAULT NULL COMMENT 'nothing\r\n\r\n',
  `q_path` varchar(255) DEFAULT NULL,
  `q_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`e_id`),
  KEY `FK_t_e_fk` (`t_id`),
  CONSTRAINT `FK_t_e_fk` FOREIGN KEY (`t_id`) REFERENCES `teacher` (`t_id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8 COMMENT='????ʵ?';

-- ----------------------------
-- Records of exam
-- ----------------------------
INSERT INTO `exam` VALUES ('1', '001', 'c++', '2020-12-13T10:30', '2020-12-13T12:00', '1', '0', '/file/1/0427628CB9184E3BBE6445B706941C60.docx', 'C++期末考试题.docx');
INSERT INTO `exam` VALUES ('6', '001', 'JavaWeb', '2020-12-20T08:00', '2020-12-23T22:18', '1', '1', '/file/6/3081661683A94B8FAD8BD05AF1372C8A.docx', 'Java Web期末考试题.docx');
INSERT INTO `exam` VALUES ('7', '001', '18-3算法设计与分析', '2020-12-12T22:23', '2021-01-01T02:19', '1', '2', '/file/7/0F27CA79CE61494185B085A5094FEDBC', '');
INSERT INTO `exam` VALUES ('8', '001', '机器学习与数据挖掘', '2020-12-19T22:21', '2020-12-24T22:23', '0', '3', '/file/8/0930E36AC340423FB943D81C16E28CCE.docx', '机器学习与数据挖掘期末考试.docx');
INSERT INTO `exam` VALUES ('9', '001', '计算机网络', '2020-12-25T08:30', '2020-12-25T10:00', '0', '4', null, null);
INSERT INTO `exam` VALUES ('13', '001', '算法竞赛课', '2020-12-28T17:09', '2020-12-31T17:12', '1', '0', null, null);
INSERT INTO `exam` VALUES ('15', '001', '大学物理', '2020-12-27T08:00', '2020-12-27T10:00', '1', '0', '/file/15/56A8E61B3BF5466C8049628ADC28B790.docx', '大学物理期末考试.docx');
INSERT INTO `exam` VALUES ('16', '001', '逻辑设计', '2021-01-02T10:20', '2021-01-21T12:00', '1', '0', '/file/16/19A2089136724BEF99D74647482EE998.docx', '逻辑设计期末考试卷.docx');
INSERT INTO `exam` VALUES ('17', '001', '形势与政策', '2020-12-19T15:00', '2020-12-19T17:30', '1', '0', '/file/17/0755C16FBE0E4D88A773D756B16DEA23.docx', '形势与政策期末试卷.docx');
INSERT INTO `exam` VALUES ('18', '001', '数据结构', '2020-12-20T08:00', '2020-12-21T10:00', '1', '0', '/file/18/DA969CC60E69407D9745CE5BDC065D06.docx', '数据结构期末考试.docx');
INSERT INTO `exam` VALUES ('19', '009', '高数A', '2020-12-20T14:00', '2020-12-21T07:30', '1', '0', '/file/19/F5642D26A8184D05B4BE451C82BCDE6A.docx', '高数A期末模拟考试.docx');
INSERT INTO `exam` VALUES ('22', '011', '创新实践', '2020-12-23 12:50:00', '2020-12-23 19:30:00', '1', '0', '/file/22/9A20ADC2B5A54E86ABFF63140AF76580.docx', '试卷.docx');
INSERT INTO `exam` VALUES ('29', '011', '创新实践2', '2020-12-22 812:50:00', '2020-12-22 13:30:00', '1', '0', null, null);
INSERT INTO `exam` VALUES ('30', '011', '高数A', '2020-12-22 12:50:00', '2020-12-22 13:50:00', '1', '0', null, null);

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
  `f_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '???',
  `s_id` varchar(40) DEFAULT NULL,
  `f_name` varchar(255) DEFAULT NULL,
  `f_path` varchar(1024) DEFAULT NULL,
  `f_size` int(11) DEFAULT NULL,
  `f_time` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`f_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8 COMMENT='?ļ??????еģ????????';

-- ----------------------------
-- Records of file
-- ----------------------------
INSERT INTO `file` VALUES ('12', '1812050089', '《Java Web程序设计》实验指导书SSM.pdf', '/file/1812050089/5F4392E583C94883B89CDA6C2F7560D1.pdf', null, '2020/12/20 17:38');
INSERT INTO `file` VALUES ('13', '1812050089', '大作业参考题目.docx', '/file/1812050089/9FC6576F1F2A4AEE9F74F2F66787F861.docx', null, '2020/12/20 17:38');
INSERT INTO `file` VALUES ('14', '1812050089', 'henu.jpg', '/file/1812050089/45BB8334A4234D049BC23C9100C4EBE2.jpg', null, '2020/12/20 17:38');
INSERT INTO `file` VALUES ('17', '1812050001', '试卷.docx', '/file/1812050001/4DAFFFC3F6C54B5E99938CB6438CD283.docx', null, '2020/12/23 18:49');
INSERT INTO `file` VALUES ('18', '1812050089', 'student.xls', '/file/1812050089/F33457BDDD894C4290F73532614F32B1.xls', null, '2020/12/24 06:58');
INSERT INTO `file` VALUES ('19', '1812050089', '李福连选课.xlsx', '/file/1812050089/76A7CD638A044BC18A463CB02B9CCB5A.xlsx', null, '2020/12/24 06:59');
INSERT INTO `file` VALUES ('20', '1812050089', '期末提交材料组织方式.zip', '/file/1812050089/14D8EC3C368341F396F87420DA9650EC.zip', null, '2020/12/24 06:59');
INSERT INTO `file` VALUES ('21', '1812050089', 'images.jpg', '/file/1812050089/64218425CB0041BF8AFFA8BC1E777F21.jpg', null, '2020/12/24 07:01');

-- ----------------------------
-- Table structure for global
-- ----------------------------
DROP TABLE IF EXISTS `global`;
CREATE TABLE `global` (
  `g_id` int(11) NOT NULL,
  `task_interval` int(11) DEFAULT NULL,
  `record_number` int(11) DEFAULT NULL,
  `max_advance_time` int(11) DEFAULT NULL,
  `min_file_size` int(11) DEFAULT NULL,
  `max_file_size` int(11) DEFAULT NULL,
  `teacher_can_clean` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`g_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ϵͳʵ???洢ȫ????????Ϣ';

-- ----------------------------
-- Records of global
-- ----------------------------
INSERT INTO `global` VALUES ('1', '22', '13', '11', '4', '200', '1');

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice` (
  `n_id` int(11) NOT NULL AUTO_INCREMENT,
  `e_id` int(11) DEFAULT NULL,
  `n_text` varchar(1024) DEFAULT NULL,
  `n_time` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`n_id`),
  KEY `FK_e_n_fk` (`e_id`),
  CONSTRAINT `FK_e_n_fk` FOREIGN KEY (`e_id`) REFERENCES `exam` (`e_id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 COMMENT='ʵʱ??Ϣʵ?';

-- ----------------------------
-- Records of notice
-- ----------------------------
INSERT INTO `notice` VALUES ('8', '6', 'test', '2020-12-15 15:00');
INSERT INTO `notice` VALUES ('9', '8', 'test', '2020-12-20 14:55');
INSERT INTO `notice` VALUES ('10', '6', '考试即将开始，请同学们做好准备', '2020-12-20 21:05');
INSERT INTO `notice` VALUES ('11', '18', '同学们注意考试时间，合理安排做题进度！', '2020-12-20 21:15');
INSERT INTO `notice` VALUES ('12', '19', '考试明天开始，请勿错过时间', '2020-12-20 21:28');
INSERT INTO `notice` VALUES ('13', '1', '考试即将考试，请同学们提前15分钟进入系统考试。如有问题，请及时联系老师。', '2020-12-20 21:28');
INSERT INTO `notice` VALUES ('14', '16', '考试即将考试，请同学们提前15分钟进入系统考试。如有问题，请及时联系老师。', '2020-12-20 23:11');
INSERT INTO `notice` VALUES ('15', '15', '考试即将考试，请同学们提前15分钟进入系统考试。如有问题，请及时联系老师。', '2020-12-20 23:12');
INSERT INTO `notice` VALUES ('19', '7', '记得明天演示哈~', '2020-12-24 06:43');

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student` (
  `s_i` int(11) NOT NULL AUTO_INCREMENT,
  `e_id` int(11) DEFAULT NULL,
  `s_id` varchar(20) DEFAULT NULL,
  `s_name` varchar(20) DEFAULT NULL,
  `s_class` varchar(20) DEFAULT NULL,
  `s_ipaddress` varchar(40) DEFAULT NULL,
  `s_online` tinyint(2) DEFAULT NULL,
  `s_score` int(5) DEFAULT NULL,
  `s_comment` varchar(1024) DEFAULT NULL,
  `s_fpath` varchar(255) DEFAULT NULL,
  `s_fname` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`s_i`),
  KEY `FK_s_f_fk2` (`e_id`),
  CONSTRAINT `FK_s_f_fk2` FOREIGN KEY (`e_id`) REFERENCES `exam` (`e_id`)
) ENGINE=InnoDB AUTO_INCREMENT=305 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of student
-- ----------------------------
INSERT INTO `student` VALUES ('1', '7', '1812050089', '李福连', '18-3', null, null, '100', 'good', '/file/7/0303DC26CE17444097AC372672FE1A55.docx', '项目特色.docx');
INSERT INTO `student` VALUES ('4', '7', '235', '儿童', '18-3', null, null, '83', '不错', null, null);
INSERT INTO `student` VALUES ('7', '8', '1812050089', '李福连', '18-3', null, null, '99', '超级棒', '/file/8/002AC7FB0D1247F6AA8F6806908A524F.html', 'txt.html');
INSERT INTO `student` VALUES ('8', '9', '1812050089', '李福连', '18-3', null, null, '98', '超级nice', '/file/9/3B8FCA726ACF4F08934B332E4DFF7D02.pdf', '《Java Web程序设计》实验指导书SSM.pdf');
INSERT INTO `student` VALUES ('9', '6', '1812050089', '李福连', '18-3', null, null, '100', '很好很好', '/file/6/BBF180242C534BB7A0092DE48E533F63.docx', '1812050089-李福连-Java Web答案.docx');
INSERT INTO `student` VALUES ('11', '7', '1', '张三', '1', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('12', '7', '2', '赵四', '1', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('13', '7', '1', '张三', '1', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('14', '7', '2', '赵四', '1', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('15', '7', '1', '张三', '1', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('16', '7', '2', '赵四', '1', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('17', '7', '1', '张三', '1', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('18', '7', '2', '赵四', '1', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('27', '18', '211205008', '李一花', '18-3班', null, null, '65', null, '/file/18/5C823F1465BD424B8D8D07B140791D9D.docx', '211205008.docx');
INSERT INTO `student` VALUES ('28', '18', '211205009', '成宝拉', '18-3班', null, null, '100', null, '/file/18/CD6F52B6B7864ED4BD2263A8F0CC2AEC.docx', '211205009.docx');
INSERT INTO `student` VALUES ('29', '18', '211205010', '金正峰', '18-3班', null, null, '45', null, '/file/18/27D5F48C59124E25A88FC730098833B1.docx', '211205010.docx');
INSERT INTO `student` VALUES ('30', '18', '211205011', '成余晖', '18-4班', null, null, '86', null, '/file/18/2C5029ACBC194BD18B983DF8A9A7539B.docx', '211205011.docx');
INSERT INTO `student` VALUES ('31', '18', '211205012', '张美玉', '18-4班', null, null, '88', null, '/file/18/ECBC9CB9343E4C6AA07949C0BBF2B93C.docx', '211205012.docx');
INSERT INTO `student` VALUES ('32', '17', '2112050030', '张艺馨', '21-1班', null, null, '65', null, '/file/17/893B45086AA249D7BADD46289763BCD9.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('33', '17', '2112050031', '甘绿蕊', '21-1班', null, null, '95', null, '/file/17/9CFDD74E7E524355ACB91D474E47858D.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('35', '18', '1812050089', '李福连', '18-3', null, null, '96', null, '/file/18/CA20353F0E2748719911436BA0F5B830.docx', '1812050089-李福连.docx');
INSERT INTO `student` VALUES ('36', '17', '2112050032', '张韩复', '21-1班', null, null, '45', null, '/file/17/B33235A6B5F54B4892C53EC51A1D05DA.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('37', '17', '2112050033', '杨杨', '21-1班', null, null, '88', null, '/file/17/0C53AE9C0B484E3296FCD20140D8824F.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('38', '17', '2112050034', '孙程琦', '21-1班', null, null, '74', null, '/file/17/2A763024A0064441A9D7A032795CB015.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('39', '17', '2112050035', '林俊雄', '21-1班', null, null, '53', null, '/file/17/E0D17C21E8704C17929896D159A5A3F5.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('40', '17', '2112050036', '简玲玲', '21-1班', null, null, '92', null, '/file/17/CF89311ABA284B1F84191A7830537FAE.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('41', '17', '2112050037', '袁博', '21-1班', null, null, '84', null, '/file/17/A017F53FB48F4655836B9BB6AAD23A62.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('42', '17', '2112050038', '孟姣娇', '21-1班', null, null, '71', null, '/file/17/EF60E4DEC8FE480CA8289BE473A8C6FC.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('43', '17', '2112050039', '崔流苏', '21-1班', null, null, '69', null, '/file/17/92F469D0A3C6402A9EA1B3008A582876.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('44', '17', '2112050040', '文从安', '21-1班', null, null, '60', null, '/file/17/6EA2832AA8EA4E28807484EFA9E760DA.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('45', '15', '18110001', '张三说', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('46', '15', '18110002', '赵四', '18-4', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('47', '15', '18110003', '张思思', '18-5', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('48', '15', '18110004', '刘福对', '18-6', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('49', '15', '18110005', '郝红梅', '18-7', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('50', '15', '18110006', '蔡司分', '18-8', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('51', '15', '18110007', '欧喷', '18-9', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('52', '15', '18110008', '王多余', '18-10', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('53', '15', '18110009', '彩虹姐', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('54', '15', '18110010', '刘一段', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('55', '15', '18110011', '张昊天', '18-4', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('56', '15', '18110012', '艾豪华', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('57', '17', '2112050041', '许岚玉', '21-1班', null, null, '81', null, '/file/17/DCB10B05A005481CB4B57C088B222946.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('58', '17', '2112050042', '周晓晨', '21-1班', null, null, '84', null, '/file/17/F842B61B64A741308017DED1F9339622.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('59', '17', '2112050043', '唐楚楠', '21-1班', null, null, '50', null, '/file/17/2CD855E0B0BC4DE8A3055DFD170259CB.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('60', '17', '2112050044', '颜澜起', '21-1班', null, null, '63', null, '/file/17/BC43B7763CE84AD2B9B7587A2A32C5FA.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('61', '17', '2112050045', '周舟', '21-1班', null, null, '78', null, '/file/17/3CBA3B40E01646209D92CBA36E9C72B9.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('62', '18', '18110001', '张三说', '18-3', null, null, '92', null, '/file/18/73BAD8467967474DBCD1880470512150.docx', '18110001.docx');
INSERT INTO `student` VALUES ('63', '18', '18110002', '赵四', '18-4', null, null, '83', null, '/file/18/C3E9BAFEB7CF4CE7A1E82DECC6DDFC8A.docx', '18110002.docx');
INSERT INTO `student` VALUES ('118', '16', '18110001', '张三说', '18-3', null, null, '54', null, null, null);
INSERT INTO `student` VALUES ('119', '16', '18110002', '赵四', '18-4', null, null, '78', null, null, null);
INSERT INTO `student` VALUES ('120', '16', '18110003', '张思思', '18-5', null, null, '89', null, null, null);
INSERT INTO `student` VALUES ('121', '16', '18110004', '刘福对', '18-6', null, null, '79', null, null, null);
INSERT INTO `student` VALUES ('122', '16', '18110005', '郝红梅', '18-7', null, null, '67', null, null, null);
INSERT INTO `student` VALUES ('123', '16', '18110006', '蔡司分', '18-8', null, null, '98', null, null, null);
INSERT INTO `student` VALUES ('124', '16', '18110007', '欧喷', '18-9', null, null, '79', null, null, null);
INSERT INTO `student` VALUES ('125', '16', '18110008', '王多余', '18-10', null, null, '76', null, null, null);
INSERT INTO `student` VALUES ('126', '16', '18110009', '彩虹姐', '18-3', null, null, '89', null, null, null);
INSERT INTO `student` VALUES ('127', '16', '18110010', '刘一段', '18-3', null, null, '59', null, null, null);
INSERT INTO `student` VALUES ('128', '16', '18110011', '张昊天', '18-4', null, null, '99', null, null, null);
INSERT INTO `student` VALUES ('129', '16', '18110012', '艾豪华', '18-3', null, null, '100', null, null, null);
INSERT INTO `student` VALUES ('160', '6', '181205001', '周学成', '18-3', null, null, '59', '再接再励', '/file/6/5D991337074246C6BEAF78948E3BEBDA.docx', '181205001-周学成-Java Web答案 .docx');
INSERT INTO `student` VALUES ('161', '6', '181205002', '方小菲', '18-3', null, null, '32', null, '/file/6/3EA859955674403CB0E0633007F8B123.docx', '181205002-方小菲-Java Web答案.docx');
INSERT INTO `student` VALUES ('162', '6', '181205003', '宇泽', '18-3', null, null, '77', null, '/file/6/0E9ACA5C423840B190BAFA381BC52C4E.docx', '181205003-宇泽-Java Web答案.docx');
INSERT INTO `student` VALUES ('163', '6', '181205004', '冯珊珊', '18-3', null, null, '60', null, '/file/6/06B242BCDF8949BB91B469F09165A559.docx', '181205004-冯珊珊-Java Web答案.docx');
INSERT INTO `student` VALUES ('164', '6', '181205005', '楮墨', '18-3', null, null, '78', null, '/file/6/564BD84EC6C5442989E94EE90948955E.docx', '181205005-楮墨-Java Web答案.docx');
INSERT INTO `student` VALUES ('165', '6', '181205006', '华域', '18-3', null, null, '98', null, '/file/6/340C06A3B74B4FD2B4A7C04DBB896908.docx', '181205006-华域-Java Web答案  .docx');
INSERT INTO `student` VALUES ('166', '6', '181205007', '上官杰', '18-3', null, null, '88', null, '/file/6/FF61F7E1702D47EE94E40B1E248257F4.docx', '181205007-上官杰-Java Web答案.docx');
INSERT INTO `student` VALUES ('167', '6', '181205008', '赵元凯', '18-3', null, null, '77', null, '/file/6/549D50DCC81E49F69DDC90C7D989C625.docx', 'Javaweb答卷.docx');
INSERT INTO `student` VALUES ('168', '6', '181205009', '吴怡', '18-3', null, null, '67', null, '/file/6/8750139B3F1D4C0396B93C69DD509A6F.docx', 'Javaweb答卷.docx');
INSERT INTO `student` VALUES ('169', '6', '181205010', '郑佳', '18-3', null, null, '66', null, '/file/6/6B630BB604F14D678F9E7FCCBD3086C5.docx', 'Javaweb答卷.docx');
INSERT INTO `student` VALUES ('170', '6', '181205011', '黄璞祎', '18-3', null, null, '100', null, null, null);
INSERT INTO `student` VALUES ('171', '6', '181205012', '余兵', '18-3', null, null, '93', null, '/file/6/49906E8D76FF4D5CB943B7489ECAA29A.docx', 'Javaweb答卷.docx');
INSERT INTO `student` VALUES ('172', '6', '181205013', '赵佳吉', '18-3', null, null, '98', null, null, null);
INSERT INTO `student` VALUES ('173', '6', '181205014', '岳庆元', '18-3', null, null, '95', null, null, null);
INSERT INTO `student` VALUES ('174', '6', '181205015', '肖龙', '18-3', null, null, '80', null, null, null);
INSERT INTO `student` VALUES ('175', '13', '18110001', '张三说', '18-3', null, null, '77', null, null, null);
INSERT INTO `student` VALUES ('176', '13', '18110002', '赵四', '18-4', null, null, '66', null, null, null);
INSERT INTO `student` VALUES ('177', '13', '18110003', '张思思', '18-5', null, null, '98', null, null, null);
INSERT INTO `student` VALUES ('178', '13', '18110004', '刘福对', '18-6', null, null, '86', null, null, null);
INSERT INTO `student` VALUES ('179', '13', '18110005', '郝红梅', '18-7', null, null, '54', null, null, null);
INSERT INTO `student` VALUES ('180', '13', '18110006', '蔡司分', '18-8', null, null, '76', null, null, null);
INSERT INTO `student` VALUES ('181', '13', '18110007', '欧喷', '18-9', null, null, '50', null, null, null);
INSERT INTO `student` VALUES ('182', '13', '18110008', '王多余', '18-10', null, null, '95', null, null, null);
INSERT INTO `student` VALUES ('183', '13', '18110009', '彩虹姐', '18-3', null, null, '86', null, null, null);
INSERT INTO `student` VALUES ('184', '13', '18110010', '刘一段', '18-3', null, null, '82', null, null, null);
INSERT INTO `student` VALUES ('185', '13', '18110011', '张昊天', '18-4', null, null, '66', null, null, null);
INSERT INTO `student` VALUES ('186', '13', '18110012', '艾豪华', '18-3', null, null, '91', null, null, null);
INSERT INTO `student` VALUES ('187', '19', '1812050001', '张三', '1', null, null, '60', null, '/file/19/EF2703B01C9346389EFC15C0D773C680.txt', '答案.txt');
INSERT INTO `student` VALUES ('188', '19', '1812050002', '赵四', '1', null, null, '59', null, null, null);
INSERT INTO `student` VALUES ('189', '19', '1812050003', '李弗瑞', '1', null, null, '90', null, null, null);
INSERT INTO `student` VALUES ('190', '19', '1812050004', '张党', '1', null, null, '43', null, null, null);
INSERT INTO `student` VALUES ('191', '19', '1812050005', '刘大山', '1', null, null, '87', null, null, null);
INSERT INTO `student` VALUES ('192', '19', '1812050006', '王五', '1', null, null, '34', null, null, null);
INSERT INTO `student` VALUES ('193', '19', '1812050007', '七峰', '1', null, null, '66', null, null, null);
INSERT INTO `student` VALUES ('194', '19', '1812050008', '赵莹莹', '1', null, null, '75', null, null, null);
INSERT INTO `student` VALUES ('195', '19', '1812050009', '任我行', '1', null, null, '13', null, null, null);
INSERT INTO `student` VALUES ('196', '19', '1812050010', '周智若', '2', null, null, '88', null, null, null);
INSERT INTO `student` VALUES ('197', '19', '1812050011', '李舒', '2', null, null, '93', null, null, null);
INSERT INTO `student` VALUES ('198', '19', '1812050012', '张三丰', '2', null, null, '8', null, null, null);
INSERT INTO `student` VALUES ('199', '19', '1812050013', '张无忌', '2', null, null, '79', null, null, null);
INSERT INTO `student` VALUES ('200', '19', '1812050014', '宋小宝', '2', null, null, '85', null, null, null);
INSERT INTO `student` VALUES ('201', '19', '1812050015', '陈赫', '2', null, null, '80', null, null, null);
INSERT INTO `student` VALUES ('202', '19', '1812050016', '陈颖', '2', null, null, '56', null, null, null);
INSERT INTO `student` VALUES ('203', '19', '1812050017', '黄晓明', '2', null, null, '83', null, null, null);
INSERT INTO `student` VALUES ('204', '19', '1812050018', '韩红', '2', null, null, '63', null, null, null);
INSERT INTO `student` VALUES ('205', '19', '1812050019', '周杰伦', '2', null, null, '76', null, null, null);
INSERT INTO `student` VALUES ('206', '19', '1812050020', '林俊杰', '2', null, null, '98', null, null, null);
INSERT INTO `student` VALUES ('207', '17', '18110001', '张三说', '18-3', null, null, '83', null, '/file/17/82A50212014D45C2AA7DDD3DF0D61142.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('208', '17', '18110002', '赵四', '18-4', null, null, '89', null, '/file/17/E326CF5F47AB4B1EAC640E918C8153DB.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('209', '17', '18110003', '张思思', '18-5', null, null, '91', null, '/file/17/F9F980E471C84C5EACCD5F0257067D92.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('211', '17', '18110004', '刘福对', '18-6', null, null, '88', null, '/file/17/0B513E33C8034501A11A619A3F80FE97.docx', '形势与政策答卷.docx');
INSERT INTO `student` VALUES ('213', '9', '18110002', '赵四', '18-4', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('214', '9', '18110003', '张思思', '18-5', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('215', '9', '18110004', '刘福对', '18-6', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('216', '9', '18110005', '郝红梅', '18-7', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('217', '9', '18110006', '蔡司分', '18-8', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('218', '9', '18110007', '欧喷', '18-9', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('219', '9', '18110008', '王多余', '18-10', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('220', '9', '18110009', '彩虹姐', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('221', '9', '18110010', '刘一段', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('222', '9', '18110011', '张昊天', '18-4', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('223', '9', '18110012', '艾豪华', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('224', '18', '211205001', '乌黎樊', '18-3班', null, null, null, null, '/file/18/343DCADE998149FDA29CFF0806CA6173.docx', '211205001-乌黎樊.docx');
INSERT INTO `student` VALUES ('225', '6', '181205016', '乌黎樊', '18-3', null, null, '93', null, '/file/6/209D8AB1470A4BF981D7159651C71001.docx', '181205016-乌黎樊.docx');
INSERT INTO `student` VALUES ('226', '6', '181205017', '金正八', '18-3', null, null, '66', null, '/file/6/D6129D051C4D46DA8926FFE67613261D.docx', '181205017-金正八.docx');
INSERT INTO `student` VALUES ('227', '1', '181205001', '周学成', '18-3', null, null, '84', null, null, null);
INSERT INTO `student` VALUES ('228', '1', '181205002', '方小菲', '18-3', null, null, '93', null, null, null);
INSERT INTO `student` VALUES ('229', '1', '181205003', '宇泽', '18-3', null, null, '81', null, null, null);
INSERT INTO `student` VALUES ('231', '1', '181205004', '冯珊珊', '18-3', null, null, '74', null, null, null);
INSERT INTO `student` VALUES ('233', '1', '181205005', '楮墨', '18-3', null, null, '60', null, null, null);
INSERT INTO `student` VALUES ('235', '1', '181205006', '华域', '18-3', null, null, '88', null, null, null);
INSERT INTO `student` VALUES ('237', '1', '181205007', '上官杰', '18-3', null, null, '89', null, null, null);
INSERT INTO `student` VALUES ('239', '1', '181205008', '赵元凯', '18-3', null, null, '97', null, null, null);
INSERT INTO `student` VALUES ('241', '1', '181205009', '吴怡', '18-3', null, null, '69', null, null, null);
INSERT INTO `student` VALUES ('243', '1', '181205010', '郑佳', '18-3', null, null, '77', null, null, null);
INSERT INTO `student` VALUES ('247', '1', '181205012', '余兵', '18-3', null, null, '89', null, null, null);
INSERT INTO `student` VALUES ('251', '1', '181205014', '岳庆元', '18-3', null, null, '67', null, null, null);
INSERT INTO `student` VALUES ('285', '22', '1812050001', '刘一', '18-3', null, null, '0', null, '/file/22/67D67E1AAA5744CEA3B35D01E259C590.docx', '试卷.docx');
INSERT INTO `student` VALUES ('286', '22', '1812050002', '陈二', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('287', '22', '1812050003', '张三', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('288', '22', '1812050004', '李四', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('289', '22', '1812050005', '王五', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('295', '29', '1812050001', '刘一', '18-3', null, null, '3000', null, null, null);
INSERT INTO `student` VALUES ('296', '29', '1812050002', '陈二', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('297', '29', '1812050003', '张三', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('298', '29', '1812050004', '李四', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('299', '29', '1812050005', '王五', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('300', '30', '1812050001', '刘一', '18-3', null, null, '0', null, null, null);
INSERT INTO `student` VALUES ('301', '30', '1812050002', '陈二', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('302', '30', '1812050003', '张三', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('303', '30', '1812050004', '李四', '18-3', null, null, null, null, null, null);
INSERT INTO `student` VALUES ('304', '30', '1812050005', '王五', '18-3', null, null, null, null, null, null);

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher` (
  `t_id` varchar(20) NOT NULL COMMENT '˭?????',
  `t_pwd` varchar(40) NOT NULL COMMENT '??ʦ???',
  `t_name` varchar(20) NOT NULL COMMENT '??ʦ??ʵ????',
  `t_isadmin` tinyint(2) NOT NULL COMMENT '?Ƿ?Ϊ????Ա',
  PRIMARY KEY (`t_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='?û?????ģ????ʦ??????Աʵ?';

-- ----------------------------
-- Records of teacher
-- ----------------------------
INSERT INTO `teacher` VALUES ('001', '0adc3949ba59abbe56e057f20f883e', '李福连', '1');
INSERT INTO `teacher` VALUES ('003', '0adc3949ba59abbe56e057f20f883e', '何小明', '1');
INSERT INTO `teacher` VALUES ('006', '0adc3949ba59abbe56e057f20f883e', '张无忌', '0');
INSERT INTO `teacher` VALUES ('007', '0adc3949ba59abbe56e057f20f883e', '刘丽萍', '1');
INSERT INTO `teacher` VALUES ('008', '0adc3949ba59abbe56e057f20f883e', '冯婧', '1');
INSERT INTO `teacher` VALUES ('009', '0adc3949ba59abbe56e057f20f883e', '李帅博', '1');
INSERT INTO `teacher` VALUES ('010', '0adc3949ba59abbe56e057f20f883e', '李虹霞', '1');
INSERT INTO `teacher` VALUES ('011', '0adc3949ba59abbe56e057f20f883e', 'nmy', '1');

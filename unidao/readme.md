1. 创建数据库 CREATE DATABASE xmetau;
2. 创建如下数据库表
// 登陆账号表
CREATE TABLE IF NOT EXISTS `T_LoginAccount`(
   `FId` BIGINT UNSIGNED AUTO_INCREMENT COMMENT '主键ID',
   `FUserId` BIGINT NOT NULL COMMENT '用户ID',
   `FLoginType` INT NOT NULL DEFAULT '0' COMMENT '登录类型',
   `FAppId` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '业务ID',
   `FOpenId` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '业务唯一键',
   `FUnionId` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '业务统一键',
   `FCreateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `FModifyTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`FId`),
    UNIQUE INDEX IndexUserId(`FUserId`),
    UNIQUE INDEX IndexLoginOpenId(`FOpenId`,`FAppId`),
    UNIQUE INDEX IndexUnionId(`FUnionId`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

// 用户属性表
CREATE TABLE IF NOT EXISTS `T_UserAttribute`(
  `FId` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `FUserId` bigint(20) NOT NULL COMMENT '用户ID',
  `FNickName` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `FAvatarUrl` varchar(2048) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '头像',
  `FNoviceTraining` int(11) NOT NULL DEFAULT '0' COMMENT '是否完成新手训练',
  `FPlayLevel` int(11) NOT NULL DEFAULT '0' COMMENT '棋力等级',
  `FGoldCoin` int(11) NOT NULL DEFAULT '0' COMMENT '金币',
  `FWinNum` int(11) NOT NULL DEFAULT '0' COMMENT '赢棋次数',
  `FLoseNum` int(11) NOT NULL DEFAULT '0' COMMENT '输棋次数',
  `FTieNum` int(11) NOT NULL DEFAULT '0' COMMENT '平局次数',
  `FGameName` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'douzi' COMMENT '游戏名称',
  `FCreateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `FModifyTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`FId`),
  UNIQUE KEY `IndexUserId` (`FUserId`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

// 游戏设置表
CREATE TABLE IF NOT EXISTS `T_GameSetting` (
  `FId` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `FSetType` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '设置类型',
  `FSetId` bigint(20) NOT NULL COMMENT '设置ID',
  `FSetKey` varchar(128) NOT NULL DEFAULT '' COMMENT '设置Key',
  `FSetValue` varchar(4096) NOT NULL DEFAULT '' COMMENT '设置值',
  `FCreateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `FModifyTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`FId`),
  UNIQUE KEY `IndexTypeUserKey` (`FSetId`,`FSetKey`,`FSetType`),
  KEY `IndexUserId` (`FSetId`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

3. 修改代码 tables/config.go 修改如下访问数据库的参数，使用自己的密码，IP和端口
	  DataBasePassword = "数据库密码"
	  DataBaseIP       = "数据库IP"
	  DataBasePort     = 数据库端口

4. 编译代码，go build, 生成可执行文件unidao

5. 启动unidao服务
  nohup ./unidao > /data/log/unidao.log 2>&1 &

6. 编译tools 目录 go build test_login.go;   运行 ./test_login，是否成功

7. 查看数据库 T_LoginAccount 和 T_UserAttribute 都会插入一条用户记录

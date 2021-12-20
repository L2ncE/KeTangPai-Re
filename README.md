# 课堂派项目-CSAGolang后端考核

### 功能列表

*粗体为基础功能要求*

- [x] **账号注册登录,更改密码**
- [x] **创建课堂,删除课堂**
- [x] **布置作业,删除作业**
- [x] **上传-下载课件资料**
- [x] **发布话题-话题讨论**
- [ ] **上课签到**
- [ ] **课中提问(学生or老师)和回答（抢答or抽答）**
- [ ] **成绩管理**
- [x] 项目说明文档
- [x] 对入参进行检验，如用户名长度，密码长度
- [x] 找回密码，通过密保实现
- [x] 查看全部话题概略
- [x] 查看一条话题详细信息和其下属评论
- [x] 更改话题与删除话题-话题讨论
- [x] 匿名进行话题讨论功能
- [x] 能够对话题-话题讨论点赞
- [x] 随时可以打开或关闭课堂
- [ ] 对学生、老师、学校等等进行权限管理，可参考RABC
- [ ] 单点登录，可参考JWT
- [ ] 对题目自动改错给分，选填给出标准答案自动改错即可
- [ ] 创建题库，可随机生成试卷
- [ ] 直播，可参考使用RTMP协议或自定义协议，也可通过拆解组合视频的方式完成
- [ ] 弹幕，可参考WS
- [ ] 上课发言热榜，建议使用redis
- [ ] 资料上传时实现断点续传和秒传
- [ ] 部署，建议使用docker
- [ ] RPC和MQ
- [x] ~~高并发~~(不确定)
- [ ] 微服务

### 功能介绍与实现思路

### 数据库相关代码

##### 创建数据库"text" (这个数据库的名字很草率,主要是代码里面不好改)

```mysql
CREATE DATABASE test
```

##### mysql建表

*用户表*

```mysql
CREATE TABLE `User`
(
    `Id`          BIGINT(20)                 NOT NULL AUTO_INCREMENT,
    `Name`        VARCHAR(20)  DEFAULT '',
    `Password`    VARCHAR(20)  DEFAULT '123456',
    `Question`    VARCHAR(255) DEFAULT NULL,
    `Answer`      VARCHAR(255) DEFAULT NULL,
    `ClassroomId` BIGINT(20)   DEFAULT NULL,
    `Sign`        BOOL         DEFAULT FALSE NULL,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
```

*话题表*

```mysql
CREATE TABLE `topic`
(
    `Id`         BIGINT(20)             NOT NULL AUTO_INCREMENT,
    `Name`       VARCHAR(20)  DEFAULT '',
    `Context`    VARCHAR(255) DEFAULT NULL,
    `PostTime`   DATETIME     DEFAULT NULL,
    `UpdateTime` DATETIME     DEFAULT NULL,
    `Likes`      BIGINT(20)   default 0 null,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
```

*评论表*

```mysql
CREATE TABLE `Comment`
(
    `Id`          BIGINT(20)             NOT NULL AUTO_INCREMENT,
    `TopicId`     BIGINT(20)             NOT NULL,
    `Name`        VARCHAR(20)  DEFAULT '',
    `Context`     VARCHAR(255) DEFAULT NULL,
    `CommentTime` DATETIME     DEFAULT NULL,
    `Likes`       BIGINT(20)   default 0 null,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
```

*教室表*

```mysql
CREATE TABLE `ClassRoom`
(
    `Id`           BIGINT(20)                NOT NULL AUTO_INCREMENT,
    `ClassName`    VARCHAR(20) DEFAULT '',
    `CreatorName`  VARCHAR(20) DEFAULT '',
    `CreateTime`   DATETIME    DEFAULT NULL,
    `LastOpenTime` DATETIME    DEFAULT NULL,
    `Status`       BOOL        DEFAULT FALSE NULL,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
```

*作业表*

```mysql
CREATE TABLE `Homework`
(
    `Id`          BIGINT(20) NOT NULL AUTO_INCREMENT,
    `ClassRoomId` BIGINT(20) NOT NULL,
    `Name`        VARCHAR(20)  DEFAULT '',
    `Context`     VARCHAR(255) DEFAULT NULL,
    `PostTime`    DATETIME     DEFAULT NULL,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
```
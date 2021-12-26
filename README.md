# 课堂派项目-CSAGolang后端考核

####前言:
花了不少时间搞这个考核项目,没有所有功能都完美实现,很多东西都不会也只有现学 ,万幸最后还是实现了不少功能,
github上的commit记录也超过了百条,也算是对得起自己。

### 功能列表

*粗体为基础功能要求*

- [x] **账号注册登录,更改密码**
- [x] **创建课堂,删除课堂**
- [x] **布置作业,删除作业**
- [x] **上传-下载课件资料**
- [x] **发布话题-话题讨论**
- [x] **上课签到**
- [x] **课中提问和回答**
- [x] **成绩管理**
- [x] 项目说明文档
- [x] JWT登录
- [x] 对入参进行检验，如用户名长度，密码长度
- [x] 找回密码，通过密保实现
- [x] 查看全部话题概略
- [x] 查看一条话题详细信息和其下属评论
- [x] 更改话题与删除话题-话题讨论
- [x] 匿名进行话题讨论功能
- [x] 能够对话题-话题讨论点赞
- [x] 管理员系统,能够直接删除话题与话题讨论以及教室
- [x] 随时可以打开或关闭课堂
- [x] 对学生、老师进行权限管理，实现多对一关联关系
- [ ] 对题目自动改错给分，选填给出标准答案自动改错即可
- [ ] 创建题库，可随机生成试卷
- [ ] 弹幕，可参考WS
- [ ] 上课发言热榜，建议使用redis
- [ ] 资料上传时实现断点续传和秒传
- [ ] 部署，建议使用docker
- [ ] RPC和MQ
- [x] ~~高并发~~(不确定)

### 功能介绍与实现思路
####1.账号注册登录,更改密码
①注册就将输入的数据,用户名密码等post到数据库中,并将输入的用户名与已有的用户名进行比对,不能重复,同时限制用户名与密码的长度。  
②登录时输入用户名与密码,进入数据库中进行比对,若不同则输入密保,密保正确则可以更换密码,若错误则登陆失败。  
③更换密码需要在登录后重新输入老密码,并输入一串新的密码,在通过密码验证与长度限制验证后就将数据库中的密码所更换。
####2.创建课堂,删除课堂
简单的增删改查,教室表在下方给出,值得一提的是每个教室都有开关状态,可以创建者可以自行打开或关闭
####3.布置作业,删除作业
依旧是简单的增删改查,每个作业都有对应的教室

### 数据库相关代码.

##### 创建数据库"text" (这个数据库的名字很草率,主要是代码里面不好改)

```mysql
CREATE DATABASE test
```

##### mysql建表

*用户表*

```mysql
CREATE TABLE `User`
(
    `Id`              BIGINT(20) NOT NULL AUTO_INCREMENT,
    `Name`            VARCHAR(20)  DEFAULT '',
    `Password`        VARCHAR(20)  DEFAULT '123456',
    `Question`        VARCHAR(255) DEFAULT NULL,
    `Answer`          VARCHAR(255) DEFAULT NULL,
    `ClassroomIdSign` BIGINT(20)   DEFAULT NULL,
    `Status`          VARCHAR(20)  DEFAULT NULL,
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

*成绩表*

```mysql
CREATE TABLE `Grade`
(
    `Id`         BIGINT(20) NOT NULL AUTO_INCREMENT,
    `Name`       VARCHAR(20) DEFAULT '',
    `Subject`    VARCHAR(20) DEFAULT NULL,
    `Grade`      BIGINT(20)  DEFAULT NULL,
    `Poster`     VARCHAR(20) DEFAULT '',
    `PostTime`   DATETIME    DEFAULT NULL,
    `UpdateTime` DATETIME    DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
```
# 课堂派项目-CSAGolang后端考核

### 功能列表

- [x] 账号注册登录注销
- [ ] 创建课堂,布置作业
- [ ] 上传-下载课件资料
- [ ] 发布话题-话题讨论
- [ ] 上课签到
- [ ] 课中提问(学生or老师)和回答（抢答or抽答）
- [ ] 成绩管理
- [ ] 代码格式规范、注释充分、层次清晰
- [ ] 项目说明文档
- [ ] 对学生、老师、学校等等进行权限管理，可参考RABC
- [ ] 单点登录，可参考JWT
- [ ] 找回密码，可使用邮件、短信，或可选择使用生物信息进行验证
- [ ] 对题目自动改错给分，选填给出标准答案自动改错即可
- [ ] 创建题库，可随机生成试卷
- [ ] 直播，可参考使用RTMP协议或自定义协议，也可通过拆解组合视频的方式完成
- [ ] 弹幕，可参考WS
- [ ] 上课发言热榜，建议使用redis
- [ ] 资料上传时实现断点续传和秒传
- [ ] 部署，建议使用docker
- [ ] RPC和MQ
- [ ] 高并发
- [ ] 微服务

### 功能介绍与实现思路

### 数据库相关代码

##### 创建数据库"text" (这个数据库的名字很草率,主要是代码里面不好改)

```bash
CREATE DATABASE test
```

##### mysql建表

*用户表*

```mysql
CREATE TABLE `User`
(
    `Id`       BIGINT(20) NOT NULL AUTO_INCREMENT,
    `Name`     VARCHAR(20)  DEFAULT '',
    `Password` VARCHAR(20)  DEFAULT '123456',
    `Question` VARCHAR(255) DEFAULT NULL,
    `Answer`   VARCHAR(255) DEFAULT NULL,
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
    `PostTime`   datetime     DEFAULT NULL,
    `UpdateTime` datetime     DEFAULT NULL,
    `Likes`      bigint       default 0 null,
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
    `CommentTime` datetime     DEFAULT NULL,
    `Likes`       bigint       default 0 null,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
```

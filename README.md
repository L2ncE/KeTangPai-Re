# 课堂派项目-CSAGolang后端考核

### 前言😎

花了不少时间搞这个考核项目,但是临近期末,不能将所有空闲时间都投入到里面,很多功能都不是它的最优解,例如排行榜用redis在人数更多时效率会高不少,但是官方文档和源码都没怎么看明白。很多东西都不会也只有现学 ,万幸最后还是实现了不少功能,
学习到了不少新知识,GitHub上的commit记录也超过了150条,算是收获满满了。

### 功能列表🤣

*粗体为基础功能要求*

1. **账号注册登录,更改密码**
2. **创建课堂,删除课堂**
3. **布置作业,删除作业**
4. **上传-下载课件资料**
5. **发布话题-话题讨论**
6. **上课签到**
7. **课中提问和回答**
8. **成绩管理**
9. 项目说明文档
10. JWT登录
11. 对入参进行检验，如用户名长度，密码长度
12. 找回密码，通过密保实现
13. 查看全部话题概略
14. 查看一条话题详细信息和其下属评论
15. 更改话题与删除话题-话题讨论
16. 匿名进行话题讨论功能
17. 能够对话题-话题讨论点赞
18. 管理员系统,能够直接删除话题与话题讨论
19. 随时可以打开或关闭课堂
20. 对学生、老师进行权限管理，实现多对一关联关系
21. 上课发言热榜
22. ~~高并发~~(不确定)

### 功能介绍与实现思路😍

#### 1.账号注册登录,更改密码

①注册就将输入的数据,用户名密码等post到数据库中,并将输入的用户名与已有的用户名进行比对,不能重复,同时限制用户名与密码的长度。
②登录时输入用户名与密码,进入数据库中进行比对,若不同则输入密保,密保正确则可以更换密码,若错误则登陆失败。 ③更换密码需要在登录后重新输入老密码,并输入一串新的密码,在通过密码验证与长度限制验证后就将数据库中的密码所更换。

#### 2.创建课堂,删除课堂

简单的增删改查,教室表在下方给出,值得一提的是每个教室都有开关状态,可以创建者可以自行打开或关闭

#### 3.布置作业,删除作业

依旧是简单的增删改查,每个作业都有对应的教室

#### 4.上传-下载课件资料

①上传文件直接使用FormFile上传,并使用SaveUploadedFile保存,在Postman上可以简单实现,其中限制了上传的最大尺寸

Headers

| Key          | Value               |
| ------------ | ------------------- |
| Content-Type | multipart/form-data |

body

选择form-data

然后选择文件file

②下载文件

没有实现得很完美,只能下载本地文件

传路径与文件名就可以实现

#### 5.发布话题-话题讨论

①发布话题即简单的sql语句操作

②发布话题讨论同上,其中话题讨论是捆绑的话题,每一个讨论都有其对应的话题id

#### 6.上课签到

用户表中有签到教室id,通过更新用户的签到教室id进行签到

#### 7.课中提问和回答

即时地进行提问与回答,也就只能提问回答了,也不会存到库中,~~真正实现可能还需要前端的帮助?~~
进行问答后,发言数量会加一,以便进行热榜排名

#### 8.成绩管理

简单的增删改查,成绩只能由老师修改,可不能自己偷偷改成满分

#### 9.项目说明文档

这就是哦

#### 10.JWT登录

非常简陋的JWT登录,不过也能用,每次登录会使用HS256加密,将其传入header中间件会进行解密,成功就登陆成功可以进行其他操作了

#### 11.对入参进行检验，如用户名长度，密码长度

在注册时进行简单的检验

#### 12.找回密码，通过密保实现

注册时会让你输入一个问题和答案,即密保,如果登录时密码错误就会让你输入密保,正确了就会让你重新输入密码,错了就直接return

#### 13.查看全部话题概略

通过访问话题模型来找到所有话题相关数据,并使用切片进行储存,再封装到服务中,最后在api中进行调用

#### 14.查看一条话题详细信息和其下属评论

通过传入话题的id来访问对应的详细信息以及评论信息,在api层将其合并,成功调用时可以顺利输出

#### 15.更改话题与删除话题-话题讨论

普普通通的增删改查

#### 16.匿名进行话题讨论功能

匿名评论不需要登录,在数据库中Name也是Anonymity,其他的和正规评论区别不大,也是简单的数据库操作(当然也不能评论不存在的话题)

#### 17.能够对话题-话题讨论点赞

每次调用点赞时会使数据库中点赞数加一,也是简单的数据库操作(需要登录后才能点赞)

#### 18.管理员系统,能够直接删除话题与话题讨论

需要用户的身份为管理员,此时删除时不会再去判断你是否为话题/话题讨论的发布者,可以无视这些直接删除,删除操作和之前的没什么不同

#### 19.随时可以打开或关闭课堂

看教室表就会发现,每个教室都会有自己的"状态",创建者可以自己打开或关闭,每次调用打开或关闭时数据库就会进行相应的更新

#### 20.对学生、老师进行权限管理，实现多对一关联关系

每个用户都有自己的身份,普通用户只能注册为学生或老师,特殊用户为管理员,不同身份的人能够进行的操作也会有不同。例如老师可以更改学生的成绩,但是学生自己并没有权限进行修改,管理员可以随时删除话题,以防出现违规内容

#### 21.上课发言热榜

可以查看排行榜,每个用户都有自己的发言次数,进行提问或回答都会将发言次数加一,查看排行榜时会将用户和发言次数按发言次数的多少从高到低排序(数据库操作不是最优解)

### 数据库相关代码😉

##### 创建数据库"text" (这个数据库的名字很草率,主要是代码里面不好改)

```mysql
CREATE DATABASE test;
```

##### mysql建表

*用户表*

```mysql
CREATE TABLE `User`
(
    `Id`              BIGINT(20)             NOT NULL AUTO_INCREMENT,
    `Name`            VARCHAR(20)  DEFAULT '',
    `Password`        VARCHAR(20)  DEFAULT '123456',
    `Question`        VARCHAR(255) DEFAULT NULL,
    `Answer`          VARCHAR(255) DEFAULT NULL,
    `ClassroomIdSign` BIGINT(20)   DEFAULT NULL,
    `Status`          VARCHAR(20)  DEFAULT NULL,
    `SpeechNum`       BIGINT(20)   DEFAULT 0 NULL,
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
    `Likes`      BIGINT(20)   DEFAULT 0 NULL,
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
    `Likes`       BIGINT(20)   DEFAULT 0 NULL,
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
    `PosterName`  VARCHAR(20)  DEFAULT NULL,
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
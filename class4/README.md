# MySQL



![](https://s1.ax1x.com/2022/11/17/zmnUVs.png)

### 什么是MySQL

#### 概念

> MySQL是一个**关系型数据库**，一种开源关系数据库管理系统（RDBMS），它使用最常用的数据库管理语言-结构化查询语言（SQL）进行数据库管理

##### **数据库**

顾名思义，是可以存储、管理数据的一类软件

> 举个例子，我们进出校的时候需要刷脸，机器一下就把你识别出来了，那它是怎么"知道"你长什么样的呢？
>
> 没错，就是通过数据库知道的
>
> 我们的人脸信息都是被编码后放在了数据库中，这样机器每次把刷到的人脸信息和数据库的一对比，自然就知道了

##### **关系型数据库**

是建立在关系模型基础上的数据库

> - 关系型数据库是由多张能互相联接的二维行列表格组成的数据库。
> - 关系模型由关系数据结构、关系操作集合、关系完整性约束三部分组成。
> - 当前主流的关系型数据库有Oracle、DB2、PostgreSQL、Microsoft SQL Server、Microsoft Access、MySQL、浪潮K-DB、MariaDB、SqLite等

简单来说，关系型数据库就是数据以**数据表**的形式进行存储与管理的数据库

那么，表又是啥？

表是一种二维行列表格，关系型数据库中的数据通常以表的形式存储，如图所示

![](https://s1.ax1x.com/2022/11/19/zK9FQs.png)

![](https://s1.ax1x.com/2022/11/19/zK99Jg.png)



- `列（column）` - 表中的一个字段 ，如图2中`patient_id`  。所有表都是由一个或多个列组成的
- `行（row）` - 表中的一个记录，比如第一行，这条记录，记录下了小红的信息(病人号、姓名、年龄、身高体重等等)
- `主键（primary key）` - 一列（或一组列），其值能够唯一标识表中每一行。

> 比如图2中每个patient_id，都标识了一个**独一无二**的病人
>
> 又比如b站的uid，也标识了一个**独一无二**的用户



##### DBMS

![](https://s1.ax1x.com/2022/11/19/zK9T00.md.png)

> 是(Database Management System)数据库管理系统的简称
>
> 是指对数据进行管理的大型系统软件

*DBMS又分为三类：*

1. *关系数据库系统(RDBMS)*
2. *面向对象数据库系统(OODBMS)*
3. *对象关系数据库系统(ORDBMS)*



##### SQL

> SQL 是一种**操作数据库**的语言，包括创建数据库、删除数据库、查询记录、修改记录、添加字段等。
>
> SQL 是关系型数据库的标准语言，所有的关系型数据库管理系统（RDBMS）都将 SQL 作为其标准处理语言。

也就是说，SQL也是一种语言，不过这种语言是用来操作关系型数据库的

SQL有以下用途：

- 允许用户访问关系型数据库系统中的数据；(**查数据**)
- 允许用户描述数据；(**自定义要查哪些数据**)
- 允许用户定义数据库中的数据，并处理该数据；(**建表**)
- 允许将 SQL 模块、库或者预处理器嵌入到其它编程语言中；(**后文在golang中使用SQL**)
- 允许用户创建和删除数据库、表、数据项（记录）；
- *允许用户在数据库中创建视图、存储过程、函数；*
- *允许用户设置对表、存储过程和视图的权限；*

下图是SQL的体系结构：

<img src="https://s1.ax1x.com/2022/11/19/zK9omq.md.png" style="zoom:67%;" />

**SQL命令分为三类**

> 具体语法后文会讲,记住关键词即可

1.DDL - Data Definition Language，数据定义语言

> 对数据的结构和形式进行定义，一般用于数据库和表的创建、删除、修改等。

| 命令   | 说明                                               |
| ------ | -------------------------------------------------- |
| CREATE | 用于在数据库中创建一个新表、一个视图或者其它对象。 |
| ALTER  | 用于修改现有的数据库，比如表、记录。               |
| DROP   | 用于删除整个表、视图或者数据库中的其它对象         |

2.DML - Data Manipulation Language，数据处理语言

> 对数据库中的数据进行处理，一般用于数据项（记录）的插入、删除、修改和查询。

| 命令   | 说明                                               |
| ------ | -------------------------------------------------- |
| CREATE | 用于在数据库中创建一个新表、一个视图或者其它对象。 |
| ALTER  | 用于修改现有的数据库，比如表、记录。               |
| DROP   | 用于删除整个表、视图或者数据库中的其它对象         |

3.DCL - Data Control Language，数据控制语言

> 对数据库中的数据进行处理，一般用于数据项（记录）的插入、删除、修改和查询。

| 命令   | 说明                                 |
| ------ | ------------------------------------ |
| SELECT | 用于从一个或者多个表中检索某些记录。 |
| INSERT | 插入一条记录。                       |
| UPDATE | 修改记录。                           |
| DELETE | 删除记录。                           |



#### 原理

> 当前阶段，这部分自行了解即可，看不懂也没关系，知识储备足够的时候自然会懂

[MySQL底层原理]:https://juejin.cn/post/6892914758006079496
[MySQL官方文档]:https://dev.mysql.com/doc/



### 如何使用MySQL

要想使用MySQL，我们首先得学会SQL语句

> MySQL其实是有许多友好的图形化客户端的，可以很轻松地生成SQL语句，但是对于开发人员来说，最好还是得会SQL



#### SQL语法

SQL语法其实很简单，就是几个**关键词**(如上)与基础的几段结构

让我们先来了解一下SQL的基本规则

1. **SQL 语句要以分号`;`结尾**，在 RDBMS （关系型数据库）当中，SQL 语句是逐条执行的，一条 SQL 语句代表着数据库的一个操作
2. **SQL 语句不区分大小写**，例如，不管写成 SELECT 还是 select，解释都是一样的。表名和列名也是如此，但是，**插入到表中的数据是区分大小写的**
3. **SQL 语句的单词之间必须使用半角空格（英文空格）或换行符来进行分隔**



**开始**

------



> 前提是已经安装好了 并且正确配置环境变量，如果不会，在群里提问、自行百度或者向学长提问

##### 1. 登录你的mysql

首先，打开你的cmd或者windows powershell(如果是mac就打开你的终端)

然后输入

```shell
mysql -u root -p
```

root是你的用户名

然后根据提示输入密码来启动MySQL

除此之外，我们也可以在goland中连接数据库，这样可以直观地管理数据，并且还可以一键生成SQL

##### 2.创建一个数据库

###### **创建数据库**

```sql
creat database student;
```

```sql
creat database 
student;
```

这二者是等效的

让我们看看结果：

```sql
mysql> create database school;
Query OK, 1 row affected (0.01 sec)
```



###### **删除数据库**

```sql
drop database school;
```



###### 查看所有数据库

```sql
show databases;
```

当然，我们也可以使用LIKE从句来指定要查看的数据库:

1.完全匹配(只有名为`school`的数据库)

```sql
show databases like 'school';
```

2.查看名字中包含`sch`的数据库(**%表示任意个字符**)

```sql
show databases like '%sch%';
```

3.查看以`sch`开头的数据库

```sql
show databases like 'sch%';
```

4.查看以`ool`结尾的数据库

```sql
show databases like '%ool';
```

结果如图所示

![](https://s1.ax1x.com/2022/11/19/zK9CWQ.png)

[tips]:mysql 有四个自带的数据库，分别是`mysql`,`information_schema`,`performance_schema`,`sys`;

###### 使用数据库

这样我们就成功创建了一个数据库了

那我们怎么使用呢？只需要输入:

```sql
student school;
```

接下来就可以使用这个数据库了



##### 3.创建一张数据表

###### 基本格式

```sql
CREATE TABLE <表名> ([表定义选项])[表选项][分区选项];
```

###### 表的结构

一张表通常由多个**字段**组成，每个字段都有其独特的名称

表的结构如下所示:

![](https://s1.ax1x.com/2022/11/19/zK9Pzj.png)

[Field]:字段名
[Type]:字段类型
[NULL]:是允许为空
[Key]:是否主键
[extra]:拓展



下面是Mysql的数据类型，大致分为**数值**、**日期/时间**和**字符串**(字符)类型

[数值类型]：

![](https://s1.ax1x.com/2022/11/19/zK9ll9.png)

[时间/日期]：

![](https://s1.ax1x.com/2022/11/19/zK9QSJ.png)

[字符串/字符类型]：

![](https://s1.ax1x.com/2022/11/19/zK9KW4.png)

详情见:

[菜鸟教程]:https://www.runoob.com/mysql/mysql-data-types.html



###### 创建表

其对应的SQL语句如下:

```sql
CREATE TABLE `student` (
		`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
		`name` VARCHAR(20) DEFAULT '0',
		`sex` VARCHAR(8) DEFAULT '',
		`age` INT(11) ,
		PRIMARY KEY(`id`)
	)ENGINE=InnoDB AUTO_INCREMENT=1 CHARSET=utf8mb4;
```

对应结构

```sql
CREATE TABLE <表名> (
	[表定义选项]
	)[表选项][分区选项];
```

`ENGINE=InnoDB`:指定 MySQL引擎为InnoDB

`AUTO_INCREMENT=1`:自增的字段每次自增1

`CHARSET=utf8mb4`：指定编码为utf8base64以支持中文



######  以SQL语句形式查看表结构

```sql
show create table student;
```



###### 删除表

```sql
drop table student;
```



##### 4.增删查改

###### 添加一条记录 Insert

> 由于id是自增，所以添加的时候无需插入id

```sql
insert into 表名(列1，列2，列3，列4，...)  values(值，值，值)
```

```sql
insert into student(name,sex,age) values('小明','男',18);
```

当字段允许空时或者设置了默认值时，也可以不插入这个字段，字段会默认为空或者默认值

```sql
insert into student(name,sex) values('小红','女');
```



###### 查询记录 Select

```sql
select 列1,列2 from 表名 where ...;
```

查询所有学生

```sql
select *from student;
```

查询所有学生的姓名与年龄

```sql
select name,age from student;
```

查询男学生

```sql
select name,age from student where sex='男';
```

如果想改变输出时的字段名 可以使用**AS**

```sql
select name as '姓名' from student;
```



###### **子查询**

> 子查询是嵌套在另一个语句，如：select update delete insert中的查询

```sql
select name,age from student where name in(
	select name from student where sex='男'
)order by id desc;
```

我们将这条语句拆解一下

最外层：

```sql
select name,age from student where name in(..)order by id desc;
```

里层:

```sql
select name from student where sex='男';
```

外层在里层结果的基础上再进行查询

最后的:

```sql
order by id desc;
```

意思是结果按id排序，降序(desc)

升序是`asc`



###### 更新记录 Update

```sql
update <表名> set 列1=值1，列2=值2，.... where...
```

```sql
update student set age=19,sex='女' where id=1;
```

这里用`where`来限定要更新的记录,因为update会**更新所有**符合限定条件的记录，如果不限定，会更新所有记录

条件之间可以用`and`、`or`连接

```sql
update student set age=19,sex='女' where id>=1 and name='小明';
```



###### 删除记录 Delete

```sql
delete from student where id=1;
```



##### 5.修改数据表

###### 修改字段

```sql
alter table <表名> change <旧字段名> <新字段名> <新数据类型>；
```

```sql
alter table student change name 姓名 varchar(32);
```

###### 删除字段

```sql
alter table student drop 姓名;
```

###### 添加字段

> 需要注意的是，添加的字段会默认在表最后一列

```sql
alter table <表名> add <新字段名><数据类型>[约束条件];
```

```sql
alter table student add name varchar(12) default '李华' not null;
```

###### 在开头添加字段

> 末尾加上first即可

```sql
alter table student add name varchar(12) default '李华' not null first;
```

###### 在中间添加字段

> 末尾加after接字段名即可

```sql
alter table student add name varchar(12) default '李华' not null after id;
```





#### 用go操作MySQL

> 我们已经学会了如何直接用SQL语句操作MySQL,那如何用go操作MySQL呢？
>
> 使用标准库database/sql库即可

##### 连接数据库

> 以下代码简单实现了MySQL的连接

```go
package main

import (
    _ "github.com/go-sql-driver/mysql" //我们使用的mysql，需要导入相应驱动包，否则会报错
	"database/sql"					   //标准库
	"log"
)

// 定义一个全局对象db
var db *sql.DB

func initDB() {
	var err error 
    // 设置一下dns charset:编码方式 parseTime:是否解析time类型 loc:时区
	dsn := "studentname:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开mysql驱动
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}

func main() {
	//初始化连接
	initDB()
}
```

##### CRUD

###### 单行查询

单行查询`db.QueryRow()`执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）

```go
func (db *DB) QueryRow(query string, args ...interface{}) *Row
```

具体示例代码：

```go
// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from student where id=?"
	var u student
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}
```

###### 多行查询

多行查询`db.Query()`执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。

```go
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
```

具体示例代码：

```go
// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from student where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u student
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}
```

###### 插入数据

插入、更新和删除操作都使用`Exec`方法。

```go
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
```

Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。

具体插入数据示例代码如下：

```go
// InsertStudent 插入数据
func InsertStudent(st model.Student) {
	sqlStr := "insert into student(name,age) values (?,?)"
	_, err := db.Exec(sqlStr, st.Name, st.Age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}
```

###### 更新数据

具体更新数据示例代码如下：

```go
// UpdateStudentAgeById 更新数据
func UpdateStudentAgeById(st model.Student) {
	sqlStr := "update student set age=? where id=1"
	_, err := db.Exec(sqlStr, st.Age, st.Id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	log.Println("update success")
}
```

###### 删除数据

具体删除数据的示例代码如下：

```go
// DeleteStudentById 删除数据
func DeleteStudentById(id int) {
	sqlStr := "delete from student where id=?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	log.Println("delete success")
}
```



#### GORM

> 这部分暂时不是重点，会基础CRUD即可

[gorm官方文档]:https://gorm.io/zh_CN/docs/index.html

##### 什么是gorm

> The fantastic ORM library for Golang aims to be developer friendly.

就是一个go语言ORM(Object Relational Mapping)框架

通过ORM，我们可以把对象映射到关系型数据库中，实现无需编写原生sql就能操作关系型数据库

我们直接上手操作就行

##### gorm基础

###### 安装

```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

###### 模型定义

```go
type Student struct {
    Id		   int		`gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` //column指定在数据库中对应的字段
  Name         string 	`gorm:"column:name" json:"name"`
  Sex          string	`gorm:"column:sex" json:"sex"`
  Age          string	`gorm:"column:age" json:"age"`
}
```

> GORM 倾向于**约定优于配置:**
>
> - 默认情况下，GORM 使用 `ID` 作为主键，使用结构体名的 `蛇形复数` 作为表名，字段名的 `蛇形` 作为列名，并使用 `CreatedAt`、`UpdatedAt` 字段追踪创建、更新时间
>
> - gorm提供了一个gorm.Model结构体，其包括字段 `ID`、`CreatedAt`、`UpdatedAt`、`DeletedAt`，我们可以直接把它嵌入我们的结构体
> - 我们可以使用tag来控制字段，详见官方文档
> - 需要注意的是，在gorm中，表名默认是结构体名的复数，列名默认是字段名的蛇形小写。比如上述结构体，gorm会认为表名是studnets，那如何解决这个问题呢？只需要在init的时候加一句`db.SingularTable(true)`或者在查询时临时指定表名

###### 连接到数据库

> GORM 官方支持的数据库类型有： MySQL, PostgreSQL, SQlite, SQL Server

连接到MySQL:

```go
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
  dsn := "student:pass@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```

GORM 允许通过一个现有的数据库连接来初始化 `*gorm.DB`，例如：

```go
import (
  "database/sql"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	dsn := "student:pass@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := dsn.DSN 
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
    db.SingularTable(true)
	log.Println("connect success:")
}
```



###### CRUD

> 这部分多看官方文档，很详细

**插入:**

```go
func insert(st model.Student) {
	res := db.Table("student").Create(&st)
	if res.Error != nil {
		log.Println("insert err:", res.Error)
	}
	log.Println("insert success")
}
```

用指定的字段创建记录

> 创建记录并更新给出的字段。

```go
db.Select("Name", "Age", "CreatedAt").Create(&st)
// INSERT INTO `students` (`name`,`age`,`created_at`) VALUES ("老王", 18, "2020-07-04 11:05:21.775")
```

> 创建一个记录且一同忽略传递给略去的字段值。

```go
db.Omit("Name", "Age", "CreatedAt").Create(&st)
// INSERT INTO `students` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
```

批量插入

> 要有效地插入大量记录，请将一个 `slice` 传递给 `Create` 方法。 GORM 将生成单独一条SQL语句来插入所有数据，并回填主键的值，钩子方法也会被调用。

```go
var sts = []model.student{{Name: "老王1"}, {Name: "老王2"}, {Name: "老王3"}}
db.Create(&st)

for _, st := range sts {
   st.Id // 1,2,3
}
```

使用 `CreateInBatches` 分批创建时，你可以指定每批的数量，例如：

```go
var students = []student{{name: "老王_1"}, ...., {Name: "老王_10000"}}

// 数量为 100
db.CreateInBatches(students, 100)
```



**查询：**

```go
	st := model.Student{}
	//查询第一条记录
	db.First(&st)
	log.Println("success find:", st)

	//如果主键是数字类型，可以使用内联条件检索
	db.First(&st, 10)
	// SELECT * FROM student WHERE id = 10;

	db.First(&st, "10")
	// SELECT * FROM student WHERE id = 10;

	db.Find(&st, []int{1, 2, 3})
	// SELECT * FROM student WHERE id IN (1,2,3);

	//如果主键是字符串（例如像 uuid），查询将被写成这样：
	//db.First(&st, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM student WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	//当目标对象有一个主要值时，将使用主键构建条件，例如：
	st.Id = 1
	db.First(&st)
	st = model.Student{}
	// SELECT * FROM student WHERE id = 10;

	var result model.Student
	db.Model(model.Student{Id: 10}).First(&result)
	// SELECT * FROM student WHERE id = 10;

	//检索全部记录
	var sts []model.Student
	db.Find(&sts)

	//随机获取一条记录
	db.Take(&st)
	log.Println("success take:", st)

	//获取最后一条记录
	db.Last(&st)
	log.Println("success last:", st)

    //条件查询:
    // 获取第一条匹配的记录
    db.Where("name = ?", "老王").First(&st)
    // SELECT * FROM student WHERE name = '老王' ORDER BY id LIMIT 1;

    // 获取所有匹配的记录
    db.Where("name <> ?", "老王").Find(&st)
    // SELECT * FROM student WHERE name <> '老王';

    // IN
    db.Where("name IN ?", []string{"老王", "老王 2"}).Find(&st)
    // SELECT * FROM student WHERE name IN ('老王','老王 2');

    // LIKE
    db.Where("name LIKE ?", "%jin%").Find(&st)
    // SELECT * FROM student WHERE name LIKE '%jin%';

    // AND
    db.Where("name = ? AND age >= ?", "老王", "22").Find(&st)
    // SELECT * FROM student WHERE name = '老王' AND age >= 22;

    // Time
    db.Where("updated_at > ?", lastWeek).Find(&st)
    // SELECT * FROM student WHERE updated_at > '2000-01-01 00:00:00';

    // BETWEEN
    db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&st)
    // SELECT * FROM student WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	//当然，也可以使用原生sql进行扫描
	var st model.student
	db.Raw("SELECT id, name, ag,sex FROM student WHERE name = ?", 3).Scan(&st)
```



**更新：**

```go
//----`Save` 会保存所有的字段，即使字段是零值
db.First(&st)

student.Name = "yxh"
student.Age = 12
db.Save(&st)
// UPDATE students SET name='老王 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;

//----更新单个列
// 条件更新
db.Model(&model.Student{}).Where("active = ?", true).Update("name", "hello")
// UPDATE students SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// student 的 ID 是 `111`
db.Model(&model.Student{}).Update("name", "hello")
// UPDATE students SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// 根据条件和 model 的值进行更新
db.Model(&model.Student{}).Where("active = ?", true).Update("name", "hello")
// UPDATE students SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true


//----更新多个列
// 根据 `struct` 更新属性，只会更新非零值的字段
db.Model(&st).Updates(model.student{Name: "hello", Age: 18, Active: false})
// UPDATE students SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// 根据 `map` 更新属性
db.Model(&st).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE students SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

//----更新选定字段
// 使用 Map 进行 Select
// student's ID is `111`:
db.Model(&st).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE students SET name='hello' WHERE id=111;

db.Model(&st).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE students SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// 使用 Struct 进行 Select（会 select 零值的字段）
db.Model(&st).Select("Name", "Age").Updates(model.student{Name: "new_name", Age: 0})
// UPDATE students SET name='new_name', age=0 WHERE id=111;

// Select 所有字段（查询包括零值字段的所有字段）
db.Model(&st).Select("*").Update(model.student{Name: "jinzhu", Role: "admin", Age: 0})

// Select 除 Role 外的所有字段（包括零值字段的所有字段）
db.Model(&st).Select("*").Omit("Role").Update(model.student{Name: "jinzhu", Role: "admin", Age: 0})
```



**删除:**

```go
st:=model.Student{Id:10}
db.Delete(&st)
// DELETE from student where id = 10;

// 带额外条件的删除
db.Where("name = ?", "老王").Delete(&st)
// DELETE from student where id = 10 AND name = "老王";
```



# 作业

发送到邮箱fengxiangrui@lanshana.email

提交格式：第四次作业-2011111188-wx-LvX

**截止时间**：下一次上课之前

## lv0

在MySQL中创建一个database，在里面创建一张数据表student，然后使用go语言操作MySQL，向里面插入十条记录，然后全部读出并打印

## lv1

在上次作业的基础上，使用MySQL进行数据持久化(database/sql和gorm任选)

## lv2

请你过一遍MySQL与SQL理论知识，包括但不限于以下内容:

- 事务
- 规范化理论
- MySQL架构
- 触发器、函数、视图、存储过程
- 联表查询

无需提交

## lv3

如果你学有余力，使用gin框架+MySQL实现QQ的好友功能

- 登录与创建账号
- 加好友
- 删好友
- 查看所有好友
- 好友分组
- 好友搜索

你可能需要知道的知识点:

- 模糊查询
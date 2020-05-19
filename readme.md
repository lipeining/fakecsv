
## fakecsv

### 格式
```js

```

### todo

### 理论基础
#### load data
```
代码直接生成 csv 文件，按行分配，字段使用 , 分割，使用 load data 导入
```
```
LOAD DATA [LOW_PRIORITY | CONCURRENT] [LOCAL] INFILE 'file_name'
[REPLACE | IGNORE]
INTO TABLE tbl_name
[PARTITION (partition_name,...)]
[CHARACTER SET charset_name]
[{FIELDS | COLUMNS}
[TERMINATED BY 'string']
[[OPTIONALLY] ENCLOSED BY 'char']
[ESCAPED BY 'char']
]
[LINES
[STARTING BY 'string']
[TERMINATED BY 'string']
]
[IGNORE number {LINES | ROWS}]
[(col_name_or_user_var,...)]
[SET col_name = expr,...]

（1） fields关键字指定了文件记段的分割格式，如果用到这个关键字，MySQL剖析器希望看到至少有下面的一个选项： 
terminatedby分隔符：意思是以什么字符作为分隔符
enclosed by字段括起字符
escaped by转义字符

terminated by描述字段的分隔符，默认情况下是tab字符（\t） 
enclosed by描述的是字段的括起字符。
escaped by描述的转义字符。默认的是反斜杠（backslash：\）  

例如：load data infile "/home/xxx/xxx txt" replace into table Orders fields terminated by',' enclosed by '"';

（2）lines 关键字指定了每条记录的分隔符默认为'\n'即为换行符

如果两个字段都指定了那fields必须在lines之前。如果不指定fields关键字缺省值与如果你这样写的相同： fields terminated by'\t'enclosed by ’ '' ‘ escaped by'\\'

如果你不指定一个lines子句，缺省值与如果你这样写的相同： lines terminated by '\n'

例如：load data infile "/xxx/load.txt" replace into tabletest fields terminated by ',' lines terminated by '/n';

load data infile '/xxx/xxx.txt' into table t0 character set gbk fieldsterminated by ',' enclosed by '"' lines terminated by '\n' (`name`,`age`,`description`) set update_time=current_timestamp;
```

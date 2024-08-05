# word

单词格式转换


## 打包
```bash
go build -o word.exe
```


## 使用
```bash
.\word.exe word -h

# 转大写
.\word.exe word -u -s hello

# 转小写
.\word.exe word -u -s HELLO

# 驼峰转下划线
.\word.exe word -d -s HelloWorld

# 下划线转驼峰
.\word.exe word -c -s hello_world
```
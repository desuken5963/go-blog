# go_blog

## db の接続

app コンテナからアプリケーションの db に接続する。

```
root@hogehoge:/app# mysql -u workuser -pPassw0rd -h db go_blog_db
```

ユーザーの作成

```
MySQL [go_blog_db]> CREATE USER 'workuser'@'%' IDENTIFIED BY 'Passw0rd!';
```

ユーザーの確認

```
MySQL [go_blog_db]> SELECT host, user FROM mysql.user;
```

ユーザーの削除

```
MySQL [go_blog_db]> DROP USER 'workuser'@'%';
```

## goose の扱い

マイグレーションの状態確認

```
root@hogehoge:/app# goose status
```

マイグレーションファイルの作成

```
root@hogehoge:/app# goose create create_hogehoges sql
```

マイグレーションの実行

```
root@hogehoge:/app# goose up
```

スキーマの後方への更新

```
root@hogehoge:/app# goose down
```

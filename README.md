# go_blog

## db の接続

app コンテナからアプリケーションの db に接続する。

```
root@hogehoge:/app# mysql -u root -ppassword -h db go_blog_db
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

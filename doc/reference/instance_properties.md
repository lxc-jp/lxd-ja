(instance-properties)=
# インスタンスプロパティ

インスタンスプロパティはインスタンスが作成されたときに設定されます。
これらは{ref}`プロファイル <profiles>`の一部にはできません。

以下のインスタンスプロパティが利用可能です。

```{list-table}
   :header-rows: 1
   :widths: 2 1 4

* - プロパティ
  - 読み取り専用
  - 説明
* - `name`
  - yes
  - インスタンス名 ({ref}`instance-name-requirements`参照)
* - `architecture`
  - no
  - インスタンスアーキテクチャ
```

(instance-name-requirements)=
## インスタンス名の要件

インスタンス名は`lxc rename`コマンドでインスタンスをリネームすることでのみ変更できます。

有効なインスタンス名は次の要件を満たさなければなりません。

- 名前は1～63文字である必要があります。
- 名前はASCIIテーブルの文字、数字、ダッシュのみを含む必要があります。
- 名前は数字またはダッシュで始まってはいけません。
- 名前はダッシュで終わってはいけません。

これらの要件は、インスタンス名がDNSレコードとして、ファイルシステム上で、色々なセキュリティプロファイル、そしてインスタンス自身のホスト名として使えるように定められています。

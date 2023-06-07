---
orphan: true
---

# 設定オプション

いくつかのインスタンスオプション:

```{config:option} agent.nic_config instance
:shortdesc: インスタンスデバイスと同じ名前とMTUを設定する
:default: "`false`"
:type: bool
:liveupdate: "`no`"
:condition: 仮想マシン

デフォルトのネットワークインターフェイスの名前とMTUをインスタンスデバイスと同じに設定するかどうかを制御します（コンテナでは自動的に行われます）
```

```{config:option} migration.incremental.memory.iterations instance
:shortdesc: 最大転送操作回数
:condition: コンテナ
:default: 10
:type: integer
:liveupdate: "yes"

インスタンスを停止する前に通過させる転送操作の最大回数
```

```{config:option} cluster.evacuate instance
:shortdesc: インスタンスの退避時の操作方法
:default: "`auto`"
:type: string
:liveupdate: "no"

インスタンスの退避時に行う操作を制御します（`auto`、`migrate`、`live-migrate`、または`stop`）
```

これらは、第二引数として `instance` スコープを指定する必要があります。
デフォルトのスコープは `server` なので、この引数は必須ではありません。

いくつかのサーバーオプション:

```{config:option} backups.compression_algorithm server
:shortdesc: イメージの圧縮アルゴリズム
:type: string
:scope: global
:default: "`gzip`"

新しいイメージに使用する圧縮アルゴリズム（`bzip2`、`gzip`、`lzma`、`xz`、または`none`）
```

```{config:option} instances.nic.host_name
:shortdesc: ホスト名の生成方法
:type: string
:scope: global
:default: "`random`"

`random`に設定されている場合、ランダムなホストインターフェイス名をホスト名として使用します。`mac`に設定されている場合、`lxd<mac_address>`（最初の2桁を省略したMAC）の形式でホスト名を生成します。
```

```{config:option} instances.placement.scriptlet
:shortdesc: カスタムで設定するインスタンスの自動配置ロジック
:type: string
:scope: global

カスタムで設定するインスタンスの自動配置ロジックの {ref}`clustering-instance-placement-scriptlet` を格納します
```

```{config:option} maas.api.key
:shortdesc: MAASを管理するためのAPIキー
:type: string
:scope: global

MAASを管理するためのAPIキー
```

他のスコープも可能です。
このスコープは、主に短い説明や説明で、利用可能なオプションでフォーマットを使用できることを示しています。

```{config:option} test1 something
:shortdesc: テスト

テスト。
```

```{config:option} test2 something
:shortdesc: こんにちは！ **太字** と `コード`

これが実際のテキストです。

2つの段落で構成されています。

そしてリスト:

- 項目
- 項目
- 項目

そして表:

キー                                 | タイプ      | スコープ     | デフォルト                                          | 説明
:--                                 | :---      | :----     | :------                                          | :----------
`acme.agree_tos`                    | bool      | global    | `false`                                          | ACME利用規約に同意する
`acme.ca_url`                       | string    | global    | `https://acme-v02.api.letsencrypt.org/directory` | ACMEサービスのディレクトリリソースへのURL
`acme.domain`                       | string    | global    | -                                                | 証明書が発行されるドメイン
`acme.email`                        | string    | global    | -                                                | アカウント登録に使用されるメールアドレス
```

```{config:option} test3 something
:shortdesc: テスト
:default: "`false`"
:type: タイプ
:liveupdate: Pythonはオプションを解析するため、"no"は"False"に変換されます - これを防ぐためにテキストの周りに引用符を付けてください（"no"または"`no`"）
:condition: "yes"
:readonly: "`maybe` - オプションがコードで始まる場合も引用符を追加してください"
:resource: リソース,
:managed: 管理された
:required: 必須
:scope: （これは「global」や「local」のようなもので、オプションのスコープ（`server`、`instance`など）では**ありません）

内容
```

オプションを参照するには、{config:option}を使用してください。
リンクテキストを上書きすることはできません。
サーバーオプション（デフォルト）を除いて、スコープを指定する必要があります。

{config:option}instance:migration.incremental.memory.iterations

{config:option}something:test1

{config:option}maas.api.key

索引はこちらです：
{ref}config-options

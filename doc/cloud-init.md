---
discourse: 12559
relatedlinks: https://cloudinit.readthedocs.org/
---

(cloud-init)=
# `cloud-init`を使用するには

```{youtube} https://www.youtube.com/watch?v=8OCG15TAldI
```

[`cloud-init`](https://cloud-init.io/)はLinuxディストリビューションのインスタンスの自動的な初期化とカスタマイズのためのツールです。

インスタンスに`cloud-init`設定を追加することで、インスタンスの最初の起動時に`cloud-init`に特定のアクションを実行させることができます。
可能なアクションには、例えば以下のようなものがあります：

* パッケージの更新とインストール
* 特定の設定の適用
* ユーザーの追加
* サービスの有効化
* コマンドやスクリプトの実行
* VMのファイルシステムをディスクのサイズに自動的に拡張する

詳細な情報は{ref}`cloud-init:index`を参照してください。

```{note}
`cloud-init`アクションはインスタンスの最初の起動時に一度だけ実行されます。
インスタンスの再起動ではアクションは再実行されません。
```

## イメージ内の`cloud-init`サポート

`cloud-init`を使用するには、`cloud-init`がインストールされたイメージをベースにインスタンスを作る必要があります。

* `ubuntu`および`ubuntu-daily` {ref}`イメージサーバ <remote-image-servers>`からのすべてのイメージは`cloud-init`をサポートしています。
* [`images`リモート](https://images.linuxcontainers.org/)からのイメージには`cloud-init`が有効化されたバリアントがあり、通常デフォルトバリアントよりもサイズが大きくなります。
クラウドバリアントは`/cloud`接尾辞を使用します。例えば、`images:ubuntu/22.04/cloud`。

## 設定オプション

LXDは、`cloud-init`の設定に対して`cloud-init.*`と`user.*`の2つの異なる設定オプションセットをサポートしています。
どちらのセットを使用する必要があるかは、使用するイメージの`cloud-init`サポートによって異なります。
一般的には、新しいイメージは`cloud-init.*`設定オプションをサポートし、古いイメージは`user.*`をサポートしていますが、例外も存在する可能性があります。

以下の設定オプションがサポートされています。

* `cloud-init.vendor-data`または`user.vendor-data` ({ref}`cloud-init:vendordata`を参照)
* `cloud-init.user-data`または`user.user-data` ({ref}`cloud-init:user_data_formats`を参照)
* `cloud-init.network-config`または`user.network-config` ({ref}`cloud-init:network_config`を参照)

設定オプションの詳細については、[`cloud-init`インスタンスオプション](instance-options-cloud-init)と、`cloud-init`ドキュメント内の{ref}`LXDデータソース <cloud-init:datasource_lxd>`を参照してください。

### ベンダーデータとユーザーデータ

`vendor-data`と`user-data`の両方が、`cloud-init`に{ref}`クラウド構成データ <explanation/format:cloud config data>`を提供するために使用されます。

主な考え方は、`vendor-data`は一般的なデフォルト構成に使用され、`user-data`はインスタンス固有の構成に使用されることです。
これは、プロファイルで`vendor-data`を指定し、インスタンス構成で`user-data`を指定する必要があることを意味します。
LXDはこの方法を強制しませんが、プロファイルとインスタンス構成の両方で`vendor-data`と`user-data`を使用することができます。

インスタンスに対して`vendor-data`と`user-data`の両方が提供される場合、`cloud-init`は2つの構成をマージします。
しかし、両方の設定で同じキーを使った場合、マージは不可能になるかもしれません。
この場合、指定されたデータをどのようにマージするべきかを`clout-init`に指定してください。
{ref}`cloud-init:merging_user_data`を参照して手順を確認してください。

## `cloud-init`の設定方法

インスタンスの`cloud-init`を設定するには、対応する設定オプションをインスタンスが使用する{ref}`プロファイル <profiles>`または{ref}`インスタンス構成 <instances-configure>`に直接追加します。

インスタンスに直接`cloud-init`を設定する場合、`cloud-init`はインスタンスの最初の起動時にのみ実行されることに注意してください。
つまり、インスタンスを起動する前に`cloud-init`を設定する必要があります。
これを行うには、`lxc launch`の代わりに`lxc init`でインスタンスを作成し、設定が完了した後に起動します。

### `cloud-init`設定のYAMLフォーマット

`cloud-init`のオプションでは、YAMLの[literalスタイルフォーマット](https://yaml.org/spec/1.2.2/#812-literal-style)が必要です。
パイプ記号(`|`)を使用して、パイプの後にインデントされたテキスト全体を、改行とインデントを保持したまま`cloud-init`に単一の文字列として渡すことを示します。

`vendor-data`および`user-data`のオプションは通常、`#cloud-config`で始まります。

例：

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    package_upgrade: true
    packages:
      - package1
      - package2
```

```{tip}
構文が正しいかどうかを確認する方法については、{ref}`cloud-init:reference/faq:how can i debug my user data?`を参照してください。
```

## `cloud-init`のステータスを確認する方法

`cloud-init`はインスタンスの最初の起動時に自動的に実行されます。
設定されたアクションによっては、完了するまでに時間がかかる場合があります。

`cloud-init`のステータスを確認するには、インスタンスにログインして以下のコマンドを入力します。

    cloud-init status

結果が`status: running`の場合、`cloud-init`はまだ実行中です。結果が`status: done`の場合、完了しています。

また、`--wait`フラグを使用して、`cloud-init`が完了したときにのみ通知を受け取ることができます：

```{terminal}
:input: cloud-init status --wait
:user: root
:host: instance

.....................................
status: done
```

## ユーザーデータやベンダーデータを指定する方法

`user-data`と`vendor-data`の設定は、例えば、パッケージのアップグレードやインストール、ユーザーの追加、コマンドの実行などに使用することができます。

提供される値は、最初の行で`cloud-init`に渡される{ref}`ユーザーデータ形式 <cloud-init:user_data_formats>`のタイプを示す必要があります。
パッケージのアップグレードやユーザーの設定などのアクティビティには、`#cloud-config`が使用するデータ形式です。

構成データは、インスタンスのルートファイルシステム内の以下のファイルに保存されます：

* `/var/lib/cloud/instance/cloud-config.txt`
* `/var/lib/cloud/instance/user-data.txt`

### 例

以下のセクションでは、さまざまな例のユースケースに対するユーザーデータ（またはベンダーデータ）の設定を参照してください。

より高度な{ref}`例 <cloud-init:yaml_examples>`は、`cloud-init`ドキュメントで見つけることができます。

#### パッケージのアップグレード

インスタンスが作成された直後に、インスタンスのリポジトリからパッケージをアップグレードするためには、`package_upgrade`キーを使用します：

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    package_upgrade: true
```

#### パッケージのインストール

インスタンスのセットアップ時に特定のパッケージをインストールするには、`packages`キーを使用し、パッケージ名をリストとして指定します：

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    packages:
      - git
      - openssh-server
```

#### タイムゾーンの設定

インスタンス作成時にインスタンスのタイムゾーンを設定するには、`timezone`キーを使用します：

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    timezone: Europe/Rome
```

#### コマンドの実行

コマンド（マーカーファイルの書き込みなど）を実行するには、`runcmd`キーを使用し、コマンドをリストとして指定します：

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    runcmd:
      - [touch, /run/cloud.init.ran]
```

#### ユーザーアカウントの追加

ユーザーアカウントを追加するには、`user`キーを使用します。
デフォルトユーザーやサポートされているキーに関する詳細は、`cloud-init`ドキュメント内の{ref}`cloud-init:reference/examples:including users and groups`の例を参照してください。

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    user:
      - name: documentation_example
```

## ネットワーク構成データを指定する方法

デフォルトでは、`cloud-init`はインスタンスの`eth0`インターフェイスにDHCPクライアントを設定します。
デフォルトの構成を上書きするために、`network-config`オプションを使用して独自のネットワーク構成を定義することができます（これはテンプレートの構造によるものです）。

その後、`cloud-init`はUbuntuリリースに応じて`ifupdown`か`netplan`を使用して、システム上の関連するネットワーク構成をレンダリングします。

構成データは、インスタンスのルートファイルシステム内の以下のファイルに保存されます：

* `/var/lib/cloud/seed/nocloud-net/network-config`
* `/etc/network/interfaces.d/50-cloud-init.cfg` (`ifupdown`を使用している場合)
* `/etc/netplan/50-cloud-init.yaml` (`netplan`を使用している場合)

### 例

特定のネットワークインターフェースに静的なIPv4アドレスを設定し、カスタム名前サーバーを使用するための次の設定を使用します：

```yaml
config:
  cloud-init.network-config: |
    version: 1
    config:
      - type: physical
        name: eth1
        subnets:
          - type: static
            ipv4: true
            address: 10.10.101.20
            netmask: 255.255.255.0
            gateway: 10.10.101.1
            control: auto
      - type: nameserver
        address: 10.10.10.254
```

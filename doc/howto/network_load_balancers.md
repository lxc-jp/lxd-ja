---
discourse: 14317
---

(network-load-balancers)=
# ネットワークロードバランサーを設定するには

```{note}
ネットワークロードバランサーは現状では {ref}`network-ovn` でのみ利用できます.
```

ネットワークロードバランサーは、外部 IP アドレス上の特定のポートを、ロードバランサーが属するネットワークの内部 IP アドレス上の特定のポートにフォワードできるという点で、ネットワークフォワードに似ています。ネットワークロードバランサーとネットワークフォワードの違いは、ロードバランサーは内向きのトラフィックを複数の内部のバックエンドアドレスで共有するのに使えることです。

この機能は外部 IP アドレスの数に限りがあったり、複数のインスタンスで単一の外部アドレスとそのアドレス上のポートを共有したい場合に有用です。

ロードバランサーは以下の要素で構成されます。

- 単一の外部リッスン IP アドレス。
- 内部 IP アドレスとオプショナルなポートレンジからなる単一あるいは複数の名前付きバックエンド。
- 単一または複数の名前付きバックエンドにフォワードされるように設定された単一または複数のリッスンポートレンジ。

## ネットワークロードバランサーを作成する

ネットワークロードバランサーを作成するには以下のコマンドを使用します。

```bash
lxc network load-balancer create <network_name> <listen_address> [configuration_options...]
```

それぞれのロードバランサーはネットワークに割り当てられます。
ロードバランサーには単一の外部リッスンアドレスが必要です (どのアドレスがロードバランス可能かについてのさらなる情報は {ref}`network-load-balancers-listen-addresses` 参照)。

### ロードバランサーのプロパティ

ネットワークロードバランサーは以下のプロパティを持ちます。

プロパティ       | 型           | 必須 | 説明
:--              | :--          | :--  | :--
`listen_address` | string       | yes  | リッスンする IP アドレス
`description`    | string       | no   | ネットワークロードバランサーの説明
`config`         | string set   | no   | キー/バリュー形式の設定オプション (`user.*` カスタムキーのみがサポートされます)
`backends`       | backend list | no   | {ref}`バックエンド仕様 <network-load-balancers-backend-specifications>` のリスト
`ports`          | port list    | no   | {ref}`ポート仕様 <network-load-balancers-port-specifications>` のリスト

(network-load-balancers-listen-addresses)=
### リッスンアドレスの要件

有効なリッスンアドレスは以下の要件を満たす必要があります。

- 許可されるリッスンアドレスはアップリンクのネットワークの `ipv{n}.routes` 設定またはプロジェクトの `restricted.networks.subnets` 設定 (設定されている場合) に定義されている必要があります。
- リッスンアドレスは他のネットワークやネットワーク内のエンティティで使用されているサブネットと重なってはいけません。

(network-load-balancers-backend-specifications)=
## バックエンドを設定する

ターゲットのアドレス (と省略可能なポート) をネットワークロードバランサーに定義するためにバックエンド仕様を追加できます。
バックエンドのターゲットアドレスはロードバランサーが関連付けられているネットワークと同じサブネット内である必要があります。

バックエンド仕様を追加するには以下のコマンドを使用します。

```bash
lxc network load-balancer backend add <network_name> <listen_address> <backend_name> <listen_ports> <target_address> [<target_ports>]
```

ターゲットポートは省略可能です。
省略した場合、ロードバランサーはバックエンドのリッスンポートをバックエンドのターゲットポートとして使用します。

トラフィックを別のポートにフォワードしたい場合、2つの選択肢があります。

- 単一のターゲットポートを指定し、全てのリッスンポートへのトラフィックをこのターゲットポートにフォワードする。
- 一組のターゲットポートをリッスンポートと同じ数のポートで指定し、リッスンポートの最初のポートをターゲットポートの最初のポート、リッスンポートの2番目のポートをターゲットポートの2番目のポート、というようにトラフィックをフォワードする。

### バックエンドのプロパティ

ネットワークロードバランサーのバックエンドは以下のプロパティを持ちます。

プロパティ       | 型     | 必須 | 説明
:--              | :--    | :--  | :--
`name`           | string | yes  | バックエンドの名前
`target_address` | string | yes  | フォワード先の IP アドレス
`target_port`    | string | no   | ターゲットポート (例 `70,80-90` や `90`)、空の場合 {ref}`ポート <network-load-balancers-port-specifications>` の `listen_port` と同じ
`description`    | string | no   | バックエンドの説明

(network-load-balancers-port-specifications)=
## ポートを設定する

ネットワークロードバランサーにポート指定を追加し、リッスンアドレスの特定のポートから、単一または複数のバックエンド上の特定のポートにトラフィックを転送できます。

ポート仕様を追加するには以下のコマンドを使用します。

```bash
lxc network load-balancer port add <network_name> <listen_address> <protocol> <listen_ports> <backend_name>[,<backend_name>...]
```

単一のリッスンポートまたは一組のポートを指定できます。
指定されたバックエンドはポートのリッスンポート設定と互換性があるターゲットポート設定を持たなければなりません。

### ポートのプロパティ

ネットワークロードバランサーのポートは以下のプロパティを持ちます。

プロパティ       | 型           | 必須 | 説明
:--              | :--          | :--  | :--
`protocol`       | string       | yes  | ポートのプロトコル (`tcp` または `udp`)
`listen_port`    | string       | yes  | リッスンポート (例 `80,90-100`)
`target_backend` | backend list | yes  | フォワード先のバックエンドの名前
`description`    | string       | no   | ポートの説明

## ネットワークロードバランサーを編集する

ネットワークロードバランサーを編集するには以下のコマンドを使用します。

```bash
lxc network load-balancer edit <network_name> <listen_address>
```

このコマンドは YAML 形式のネットワークロードバランサーの設定を編集用に開きます。
一般の設定、バックエンド、ポート仕様を編集できます。

## ネットワークロードバランサーを削除する

ネットワークロードバランサーを削除するには以下のコマンドを使用します。

```bash
lxc network load-balancer delete <network_name> <listen_address>
```

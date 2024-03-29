# ネットワークを作成し設定するには

マネージドネットワークを作成し設定するには、 `lxc network` コマンドとそのサブコマンドを使用します。
どのコマンドでも `--help` を追加すると使用方法と利用可能なフラグについてより詳細な情報を表示できます。

(network-types)=
## ネットワークタイプ

以下のネットワークタイプが利用できます。

```{list-table}
   :header-rows: 1

* - ネットワークタイプ
  - ドキュメント
  - 設定オプション
* - `bridge`
  - {ref}`network-bridge`
  - {ref}`network-bridge-options`
* - `ovn`
  - {ref}`network-ovn`
  - {ref}`network-ovn-options`
* - `macvlan`
  - {ref}`network-macvlan`
  - {ref}`network-macvlan-options`
* - `sriov`
  - {ref}`network-sriov`
  - {ref}`network-sriov-options`
* - `physical`
  - {ref}`network-physical`
  - {ref}`network-physical-options`

```

## ネットワークを作成する

ネットワークを作成するには以下のコマンドを実行します。

```bash
lxc network create <name> --type=<network_type> [configuration_options...]
```

利用可能なネットワークタイプ一覧と設定オプションへのリンクは {ref}`network-types` を参照してください。

`--type` 引数を指定しない場合、デフォルトのタイプ `bridge` が使用されます。

(network-create-cluster)=
### クラスタ内にネットワークを作成する

LXD クラスタを実行していてネットワークを作成したい場合、各クラスタメンバーに別々にネットワークを作成する必要があります。
この理由はネットワーク設定は、例えば親ネットワークインタフェースの名前のように、クラスタメンバー間で異なるかもしれないからです。

このため、まず `--target=<cluster_member>` フラグとメンバー用の適切な設定を指定して保留中のネットワークを作成する必要があります。
全てのメンバーで同じネットワーク名を使うようにしてください。
次に実際にセットアップするために `--target` フラグなしでネットワークを作成してください。

例えば、以下の一連のコマンドで 3 つのクラスタメンバー上に `UPLINK` という名前の物理ネットワークをセットアップします。

```{terminal}
:input: lxc network create UPLINK --type=physical parent=br0 --target=vm01

Network UPLINK pending on member vm01
:input: lxc network create UPLINK --type=physical parent=br0 --target=vm02
Network UPLINK pending on member vm02
:input: lxc network create UPLINK --type=physical parent=br0 --target=vm03
Network UPLINK pending on member vm03
:input: lxc network create UPLINK --type=physical
Network UPLINK created
```

{ref}`cluster-config-networks`も参照してください。

(network-attach)=
## インスタンスにネットワークをアタッチする

マネージドネットワークを作成後、それをインスタンスに{ref}`NICデバイス <devices-nic>`としてアタッチできます。

そのためには、以下のコマンドを使います。

    lxc network attach <network_name> <instance_name> [<device_name>] [<interface_name>]

デバイス名とインタフェース名は省略可能ですが、少なくともデバイス名は指定することをお勧めします。
指定しない場合、LXDはネットワーク名をデバイス名として使用しますが、紛らわしく問題を起こすかもしれません。
例えば、LXDイメージは`eth0`インタフェースにIP自動設定を行いますが、インタフェースの名前が違うと機能しません。

例えば、`my-network`というネットワークを`my-instance`というインタンスに`eth0`デバイスとしてアタッチするには、以下のコマンドを入力します。

    lxc network attach my-network my-instance eth0

### NICデバイスを追加する

`lxc network attach`コマンドはインスタンスにNICデバイスを追加するショートカットです。
別の方法として、通常通りネットワーク設定でNICデバイスを追加できます。

    lxc config device add <instance_name> <device_name> nic network=<network_name>

この方法を使う場合、必要に応じてネットワークのデフォルト設定をオーバーライドするように追加の設定をコマンドに追加できます。
全ての利用可能なデバイスオプションについては{ref}`NICデバイス <devices-nic>`を参照してください。

## ネットワークを設定する

既存のネットワークを設定するには、 `lxc network set` と `lxc network unset` コマンド (単一の設定項目を設定する場合) または `lxc network edit` コマンド (設定全体を編集する場合) のどちらかを使います。
特定のクラスタメンバーの設定を変更するには、 `--target` フラグを追加してください。

例えば、以下のコマンドは物理ネットワークの DNS サーバーを設定します。

```bash
lxc network set UPLINK dns.nameservers=8.8.8.8
```

利用可能な設定オプションはネットワークタイプによって異なります。
各ネットワークタイプの設定オプションへのリンクは {ref}`network-types` を参照してください。

高度なネットワーク機能を設定するためには別のコマンドがあります。
以下のドキュメントを参照してください。

- {doc}`/howto/network_acls`
- {doc}`/howto/network_forwards`
- {doc}`/howto/network_load_balancers`
- {doc}`/howto/network_zones`
- {doc}`/howto/network_ovn_peers` (OVN のみ)

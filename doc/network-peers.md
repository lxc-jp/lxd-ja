# ネットワークピアの設定 <!-- Network Peers configuration -->

ネットワークピアは 2 つの OVN ネットワーク間でルーティングの関係を作成できます。
これにより 2 つのネットワーク間での通信がアップリンクのネットワーク経由ではなく OVN サブシステム内で完結できます。
<!--
Network peers allow the creation of routing relationships between two OVN networks.
This allows for traffic between those two networks to stay within the OVN subsystem rather than having to transit
via the uplink network.
-->

ピアリングが双方向であることを確実にするため、ピアリング内の両方のネットワークがセットアップの行程を完了する必要があります。
<!--
Both networks in the peering are required to complete a setup step to ensure that the peering is mutual.
-->

例。
<!--
E.g.
-->

```
lxc network peer create <local_network> foo <target_project/target_network> --project=local_network
lxc network peer create <target_network> foo <local_project/local_network> --project=target_project
```

ピアのセットアップの行程でプロジェクトかネットワーク名の指定が正しくない場合、対応するプロジェクトや
ネットワークが存在しないというエラーを上記のコマンドは表示しません。
これは他のプロジェクトの（訳注：悪意の）ユーザーがプロジェクトやネットワークが存在するかを確認できないようにするためです。
<!--
If either the project or network name specified in the peer setup step is incorrect, the user will not get an error
from the command explaining that the respective project/network does not exist. This is to prevent a user in a
different project from being able to discover whether a project and network exists.
-->

## プロパティー <!-- Properties -->
ネットワークピアの設定は以下の通りです。
<!--
The following are network peer properties:
-->

プロパティー <!-- Property --> | 型 <!-- Type --> | 必須 <!-- Required --> | 説明 <!-- Description -->
:--              | :--        | :--      | :--
name             | string     | yes      | ローカルネットワーク上のネットワークピアの名前 <!-- Name of the Network Peer on the local network -->
description      | string     | no       | ネットワークピアの説明 <!-- Description of Network Peer -->
config           | string set | no       | 設定のキーバリューペアー (`user.*` のカスタムキーのみサポート) <!-- Config key/value pairs (Only `user.*` custom keys supported) -->
ports            | port list  | no       | ネットワークフォワードのポートリスト <!-- Network forward port list -->
target_project   | string     | yes      | 対象のネットワークがどのプロジェクト内に存在するか (作成時に必須) <!-- Which project the target network exists in (required at create time). -->
target_network   | string     | yes      | どのネットワークとピアを作成するか (作成時に必須) <!-- Which network to create a peer with (required at create time). -->
status           | string     | --       | 作成中か作成完了 (対象のネットワークと相互にピアリングした状態) かを示すステータス <!-- Status indicates if pending or created (mutual peering exists with the target network). -->

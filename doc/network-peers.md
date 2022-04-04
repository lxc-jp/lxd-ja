---
discourse: 12165
---

# ネットワークピアの設定

ネットワークピアは 2 つの OVN ネットワーク間でルーティングの関係を作成できます。
これにより 2 つのネットワーク間での通信がアップリンクのネットワーク経由ではなく OVN サブシステム内で完結できます。

ピアリングが双方向であることを確実にするため、ピアリング内の両方のネットワークがセットアップの行程を完了する必要があります。

例。

```
lxc network peer create <local_network> foo <target_project/target_network> --project=local_network
lxc network peer create <target_network> foo <local_project/local_network> --project=target_project
```

ピアのセットアップの行程でプロジェクトかネットワーク名の指定が正しくない場合、対応するプロジェクトや
ネットワークが存在しないというエラーを上記のコマンドは表示しません。
これは他のプロジェクトの（訳注：悪意の）ユーザーがプロジェクトやネットワークが存在するかを確認できないようにするためです。

## プロパティ
ネットワークピアの設定は以下の通りです。

プロパティ | 型 | 必須 | 説明
:--              | :--        | :--      | :--
name             | string     | yes      | ローカルネットワーク上のネットワークピアの名前
description      | string     | no       | ネットワークピアの説明
config           | string set | no       | 設定のキーバリューペアー (`user.*` のカスタムキーのみサポート)
ports            | port list  | no       | ネットワークフォワードのポートリスト
target_project   | string     | yes      | 対象のネットワークがどのプロジェクト内に存在するか (作成時に必須)
target_network   | string     | yes      | どのネットワークとピアを作成するか (作成時に必須)
status           | string     | --       | 作成中か作成完了 (対象のネットワークと相互にピアリングした状態) かを示すステータス

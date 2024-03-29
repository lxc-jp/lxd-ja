(cluster-recover)=
# クラスタを復旧するには

クラスタの 1 つまたは複数のメンバーがオフラインまたは到達不能になるかもしれません。
この場合、このメンバー上での操作と全てのメンバーにまたがる状態変更が必要な操作は不可能になります。
詳細は{ref}`clustering-offline-members`と{ref}`cluster-automatic-evacuation`を参照してください。

オフラインのクラスタメンバーを復旧させるかクラスタから削除すると、通常通り操作が可能となります。
これができない場合、故障の原因となったケースに応じて、クラスタを復旧させるいくつかの方法があります。
詳細は以下のセクションを参照してください。

```{note}
復旧が必要な状態のクラスタにいる場合、LXD クライアントが LXD デーモンに接続できないため、ほとんどの `lxc` コマンドは動きません。

このため、クラスタを復旧するコマンドは LXD デーモン (`lxd`) が直接提供します。
全ての利用可能なコマンドの概要を見るには `lxd cluster --help` を実行してください。
```

## 過半数割れからの復旧

各 LXD クラスタは([`cluster.max_voters`](server-options-cluster)で設定した)特定のメンバー数を持ち、これが分散データベースのvoterメンバーの数を決定します。
クラスタメンバーの過半数を恒久的に失った場合 (例えば、 3 つのメンバーのクラスタで 2 つのメンバーを消失した場合)、クラスタは過半数を失い利用不可能となります。
しかし、最低 1 つのデータベースメンバーが生き残っていれば、クラスタを復旧できます。

このためには、以下のステップを実行してください。

1. クラスタ内の生き残っているどれかのメンバーにログオンし以下のコマンドを実行します。

       sudo lxd cluster list-database

   このコマンドはデータベースロールの 1 つを持つクラスタメンバーを表示します。
1. 一覧表示されたデータベースメンバーの 1 つを新しいリーダーとして選択します。
   そのマシンにログオンします (すでにログオンしたマシンと異なる場合)。
1. そのマシンで LXD デーモンが実行中でないことを確認します。
   例えば、snap を使用している場合は以下のようにします。

       sudo snap stop lxd

1. まだオンラインである他の全てのクラスタメンバーにログオンし LXD デーモンを停止します。
1. 新しいリーダーとして選択したサーバーで以下のコマンドを実行します。

       sudo lxd cluster recover-from-quorum-loss

1. 全てのマシンで再び LXD デーモンを開始し、新しいリーダーを始動させます。
   例えば、snap を使用している場合は以下のようにします。

       sudo snap start lxd

これでデータベースはオンラインに戻るはずです。
データベースから情報が削除されることはありません。
失ったクラスタメンバーについての情報は、そのメンバーのインスタンスについてのメタデータも含めて、全て残っています。
失ったインスタンスを再び作成する必要がある場合に、この情報がさらに復旧を進める上で役に立ちます。

失ったクラスタメンバーを恒久的に削除するには、強制削除します。
{ref}`cluster-manage-delete-members` を参照してください。

## アドレス変更からクラスタメンバーを復旧する

クラスタのいくつかのメンバーがもう到達不能な場合、あるいはクラスタ自体が IP アドレスまたはリッスニングポート番号の変更のために到達不能な場合、クラスタを再設定できます。

そうするためには、クラスタの各メンバーでクラスタ設定を編集し、IP アドレスとリッスニンググポート番号を必要に応じて変更します。
この操作中はメンバーは削除できません。
クラスタ設定は完全なクラスタの記述を含む必要がありますので、全てのクラスタメンバー上で全てのクラスタメンバーを変更する必要があります。

異なるメンバーの {ref}`clustering-member-roles` を編集できますが、以下の制限があります。

- `database*` ロールを持たないクラスタメンバーは、グローバルデータベースがないため、 voter になれません。
- 少なくとも 2 つのメンバー (2 つのメンバーからなるクラスタの場合を除く、この場合は 1 つで十分) が voter にとどまる必要があります。そうでなければ過半数が成り立ちません。

各クラスタメンバーにログオンして以下のステップを実行します。

1. LXD デーモンを停止します。
   例えば、snap を使用していれば以下のようにします。

       sudo snap stop lxd

1. 以下のコマンドを実行します。

       sudo lxd cluster edit

1. クラスタメンバーがクラスタの他のメンバーについて持っている情報の YAML 表現を編集します。

   ```yaml
   # Latest dqlite segment ID: 1234

   members:
     - id: 1             # メンバーの内部 ID (読み取り専用)
       name: server1     # クラスタメンバー名 (読み取り専用)
       address: 192.0.2.10:8443 # メンバーの最終の既知のアドレス (変更可能)
       role: voter              # メンバーの最終の既知のロール (変更可能)
     - id: 2             # メンバーの内部 ID (読み取り専用)
       name: server2     # クラスタメンバー名 (読み取り専用)
       address: 192.0.2.11:8443 # メンバーの最終の既知のアドレス (変更可能)
       role: stand-by           # メンバーの最終の既知のロール (変更可能)
     - id: 3             # メンバーの内部 ID (読み取り専用)
       name: server3     # クラスタメンバー名 (読み取り専用)
       address: 192.0.2.12:8443 # メンバーの最終の既知のアドレス (変更可能)
       role: spare              # メンバーの最終の既知のロール (変更可能)
   ```

   アドレスとロールを編集できます。

全てのクラスタメンバー上でこの変更をした後、全てのメンバーで LXD デーモンを再び開始します。
例えば、snap を使用していれば以下のようにします。

    sudo snap start lxd

クラスタは全てのメンバーが入った状態で再び完全に利用可能になるはずです。
データベースから情報が削除されることはありません。
クラスタメンバーとそれらのインスタンスについての情報は全て残っています。

## Raft のメンバーシップを手動で変更する

場合によっては、なんらかの予期せぬ挙動のために Raft のメンバーシップ設定を手動で変更する必要があるかもしれません。

例えば、クラスタメンバーをクリーンでない状態で削除した場合、 `lxc cluster list` では表示されないが Raft 設定の一部として残るという状態になるかも知れません。
Raft の設定を見るには以下のコマンドを使用します。

    lxd sql local "SELECT * FROM raft_nodes"

この場合、残ったノードを削除するには以下のコマンドを使用します。

    lxd cluster remove-raft-node <address>

---
discourse: 15871
---

(cluster-form)=
# クラスタを形成するには

LXD クラスタを形成するときはブートストラップサーバーから始めます。
このブートストラップサーバーは既存の LXD サーバーでもよいですし新しくインストールしたものでもよいです。

ブートストラップサーバーを初期化した後、クラスタに追加のサーバーをジョインできます。
詳細は {ref}`clustering-members` を参照してください。

LXD クラスタを形成するために初期化プロセス中に設定をインタラクティブに指定することもできますし、完全な設定を含むプリシードファイルを使うこともできます。

素早く自動的にベーシックなLXDクラスタをセットアップするにはMicroCloudが使えます。
ただし、このプロジェクトはまだ初期段階なことに注意してください。

## クラスタをインタラクティブに設定する

クラスタを形成するには、まずブートストラップサーバー上で `lxd init` を実行する必要があります。その後クラスタにジョインさせたい他のサーバー上でもそのコマンドを実行します。

クラスタをインタラクティブに形成する際、クラスタを設定するために `lxd init` のプロンプトの質問に回答します。

### ブートストラップサーバーを初期化する

ブートストラップサーバーを初期化するには、 `lxd init` を実行して希望の設定に応じて質問に回答します。

ほとんどの質問はデフォルト値を受け入れることができますが、以下の質問には適切に答えるようにしてください。

- `Would you like to use LXD clustering?`

  **yes** を選択。
- `What IP address or DNS name should be used to reach this server?`

  他のサーバーがアクセスできる IP または DNS のアドレスを確実に使用してください。
- `Are you joining an existing cluster?`

  **no** を選択。
- `Setup password authentication on the cluster?`

  {ref}`認証トークン <authentication-token>` (推奨) を使う場合 **no** を、{ref}`トラストパスワード <authentication-trust-pw>` を使う場合 **yes** を選択。

<details>
<summary>ブートストラップ上での <code>lxd init</code> の完全な例を見るには展開してください</summary>

```{terminal}
:input: lxd init

Would you like to use LXD clustering? (yes/no) [default=no]: yes
What IP address or DNS name should be used to reach this server? [default=192.0.2.101]:
Are you joining an existing cluster? (yes/no) [default=no]: no
What member name should be used to identify this server in the cluster? [default=server1]:
Setup password authentication on the cluster? (yes/no) [default=no]: no
Do you want to configure a new local storage pool? (yes/no) [default=yes]:
Name of the storage backend to use (btrfs, dir, lvm, zfs) [default=zfs]:
Create a new ZFS pool? (yes/no) [default=yes]:
Would you like to use an existing empty block device (e.g. a disk or partition)? (yes/no) [default=no]:
Size in GiB of the new loop device (1GiB minimum) [default=9GiB]:
Do you want to configure a new remote storage pool? (yes/no) [default=no]:
Would you like to connect to a MAAS server? (yes/no) [default=no]:
Would you like to configure LXD to use an existing bridge or host interface? (yes/no) [default=no]:
Would you like to create a new Fan overlay network? (yes/no) [default=yes]:
What subnet should be used as the Fan underlay? [default=auto]:
Would you like stale cached images to be updated automatically? (yes/no) [default=yes]:
Would you like a YAML "lxd init" preseed to be printed? (yes/no) [default=no]:
```

</details>

初期化プロセスが終了したら、最初のクラスタメンバーが起動してネットワーク上で利用可能になるはずです。
これは `lxc cluster list` で確認できます。

### 追加のサーバーをジョインさせる

これでクラスタに追加のサーバーをジョインできるようになりました。

```{note}
追加するサーバーは新規にインストールした LXD サーバーにしたほうがよいです。
既存のサーバーを使う場合、既存のデータは消失するので、ジョインする前にデータを確実にクリアしてください。
```

クラスタにサーバーをジョインさせるには、クラスタ上で `lxd init` を実行します。
既存のクラスタにジョインするには root 権限が必要ですので、コマンドを root で実行するか `sudo` をつけて実行するのを忘れないでください。

基本的に、初期化プロセスは以下のステップからなります。

1. 既存のクラスタにジョインをリクエストする。

   `lxd init` の最初の質問に適切に回答します。

   - `Would you like to use LXD clustering?`

     **yes** を選択。
   - `What IP address or DNS name should be used to reach this server?`

     他のサーバーがアクセスできる IP または DNS のアドレスを確実に使用してください。
   - `Are you joining an existing cluster?`

     **yes** を選択。
   - `Do you have a join token?`

     ブートストラップサーバーを {ref}`認証トークン <authentication-token>` (推奨) を使うように設定した場合 **yes** を、{ref}`トラストパスワード <authentication-trust-pw>` を使うように設定した場合 **no** を選択。
1. クラスタで認証する。

   ブートスラップサーバーを設定する際に選んだ認証方法に応じて2 つの方法があります。

   `````{tabs}

   ````{group-tab} 認証トークン (推奨)
   {ref}`認証トークン <authentication-token>` を使うようにクラスタを設定した場合、新メンバーごとにジョイントークンを生成する必要があります。
   そのためには、既存のクラスタメンバー (例えばブートストラップサーバー) で以下のコマンドを実行します。

       lxc cluster add <new_member_name>

   このコマンドは設定時に有効な([`cluster.join_token_expiry`](server-options-cluster)参照)一回限りのジョイントークンを返します。
   `lxd init` のプロンプトでジョイントークンを求められたときにこのトークンを入力してください。

   ジョイントークンは既存のオンラインメンバーのアドレス、一回限りのシークレットとクラスタ証明書のフィンガープリントを含みます。
   ジョイントークンがこれらの質問に自動で回答できるので、 `lxd init` 中に回答が必要な質問の量を減らすことができます。
   ````
   ````{group-tab} トラストパスワード
   {ref}`トラストパスワード <authentication-trust-pw>` を使うようにクラスタを設定した場合、認証プロセスを開始できるまでに `lxd init` はより多くの情報を必要とします。

   1. 新しいクラスタメンバーの名前を指定します。
   1. 既存のクラスタメンバーのアドレスを提供します (ブートストラップサーバーまたはすでに追加済みの他のサーバー)。
   1. クラスタのフィンガープリントを検証します。
   1. フィンガープリントが正しければ、クラスタで認証するトラストパスワードを入力します。
   ````

   `````

1. クラスタにジョインする際サーバーの全てのローカルデータが消失することを確認します。
1. サーバー固有の設定を行います (詳細は {ref}`clustering-member-config` を参照)。

   デフォルト値を受け入れることもできますし、各サーバーにカスタム値を指定することもできます。

<details>
<summary>追加のサーバー上で <code>lxd init</code> を実行する完全な例を見るには展開してください</summary>

`````{tabs}

````{group-tab} 認証トークン (推奨)

```{terminal}
:input: sudo lxd init

Would you like to use LXD clustering? (yes/no) [default=no]: yes
What IP address or DNS name should be used to reach this server? [default=192.0.2.102]:
Are you joining an existing cluster? (yes/no) [default=no]: yes
Do you have a join token? (yes/no/[token]) [default=no]: yes
Please provide join token: eyJzZXJ2ZXJfbmFtZSI6InJwaTAxIiwiZmluZ2VycHJpbnQiOiIyNjZjZmExZDk0ZDZiMjk2Nzk0YjU0YzJlYzdjOTMwNDA5ZjIzNjdmNmM1YjRhZWVjOGM0YjAxYTc2NjU0MjgxIiwiYWRkcmVzc2VzIjpbIjE3Mi4xNy4zMC4xODM6ODQ0MyJdLCJzZWNyZXQiOiJmZGI1OTgyNjgxNTQ2ZGQyNGE2ZGE0Mzg5MTUyOGM1ZGUxNWNmYmQ5M2M3OTU3ODNkNGI5OGU4MTQ4MWMzNmUwIn0=
All existing data is lost when joining a cluster, continue? (yes/no) [default=no] yes
Choose "size" property for storage pool "local":
Choose "source" property for storage pool "local":
Choose "zfs.pool_name" property for storage pool "local":
Would you like a YAML "lxd init" preseed to be printed? (yes/no) [default=no]:
```

````
````{group-tab} トラストパスワード

```{terminal}
:input: sudo lxd init

Would you like to use LXD clustering? (yes/no) [default=no]: yes
What IP address or DNS name should be used to reach this server? [default=192.0.2.102]:
Are you joining an existing cluster? (yes/no) [default=no]: yes
Do you have a join token? (yes/no/[token]) [default=no]: no
What member name should be used to identify this server in the cluster? [default=server2]:
IP address or FQDN of an existing cluster member (may include port): 192.0.2.101:8443
Cluster fingerprint: 2915dafdf5c159681a9086f732644fb70680533b0fb9005b8c6e9bca51533113
You can validate this fingerprint by running "lxc info" locally on an existing cluster member.
Is this the correct fingerprint? (yes/no/[fingerprint]) [default=no]: yes
Cluster trust password:
All existing data is lost when joining a cluster, continue? (yes/no) [default=no] yes
Choose "size" property for storage pool "local":
Choose "source" property for storage pool "local":
Choose "zfs.pool_name" property for storage pool "local":
Would you like a YAML "lxd init" preseed to be printed? (yes/no) [default=no]:
```

````
`````

</details>

初期化プロセスが終わった後、サーバーが新しいクラスタメンバーとして追加されます。
これは `lxc cluster list` で確認できます。

## クラスタをプリシードファイルで設定する

クラスタを形成するには、まずブートストラップサーバー上で `lxd init` を実行します。
その後、クラスタにジョインさせたい他のサーバーでもこのコマンドを実行します。

`lxd init` の質問にインタラクティブに回答する代わりに、プリシードファイルを使って必要な情報を提供できます。
以下のコマンドを使って `lxd init` にファイルをフィードできます。

    cat <preseed-file> | lxd init --preseed

サーバーごとに異なるプリシードファイルが必要です。

### ブートストラップサーバーを初期化する

プリシードファイルの必要な中身は認証に {ref}`認証トークン <authentication-token>` (推奨) を使うか {ref}`トラストパスワード <authentication-trust-pw>` を使うかに応じて異なります。

`````{tabs}

````{group-tab} 認証トークン (推奨)
クラスタリングを有効にするには、ブートストラップサーバー用のプリシードファイルは以下のフィールドを含む必要があります。

```yaml
config:
  core.https_address: <IP_address_and_port>
cluster:
  server_name: <server_name>
  enabled: true
```

ブートストラップサーバー用のプリシードファイルの例を以下に示します。

```yaml
config:
  core.https_address: 192.0.2.101:8443
  images.auto_update_interval: 15
storage_pools:
- name: default
  driver: dir
networks:
- name: lxdbr0
  type: bridge
profiles:
- name: default
  devices:
    root:
      path: /
      pool: default
      type: disk
    eth0:
      name: eth0
      nictype: bridged
      parent: lxdbr0
      type: nic
cluster:
  server_name: server1
  enabled: true
```

````
````{group-tab} トラストパスワード
クラスタリングを有効にするには、ブートストラップサーバー用のプリシードファイルは以下のフィールドを含む必要があります。

```yaml
config:
  core.https_address: <IP_address_and_port>
  core.trust_password: <trust_password>
cluster:
  server_name: <server_name>
  enabled: true
```

ブートストラップサーバー用のプリシードファイルの例を以下に示します。

```yaml
config:
  core.trust_password: the_password
  core.https_address: 192.0.2.101:8443
  images.auto_update_interval: 15
storage_pools:
- name: default
  driver: dir
networks:
- name: lxdbr0
  type: bridge
profiles:
- name: default
  devices:
    root:
      path: /
      pool: default
      type: disk
    eth0:
      name: eth0
      nictype: bridged
      parent: lxdbr0
      type: nic
cluster:
  server_name: server1
  enabled: true
```

````
`````

### 追加のサーバーをジョインさせる

プリシードファイルの必要な中身は認証に {ref}`認証トークン <authentication-token>` (推奨) を使うか {ref}`トラストパスワード <authentication-trust-pw>` を使うかに応じて異なります。

新しいクラスタメンバー用のプリシードファイルは参加するサーバーに固有のデータと設定値を含む `cluster` セクションのみが必要です。

`````{tabs}

````{group-tab} 認証トークン (推奨)
追加のサーバーのプリシードファイルは以下の項目を含む必要があります。

```yaml
cluster:
  enabled: true
  server_address: <IP_address_of_server>
  cluster_token: <join_token>
```

新しいクラスタメンバー用のプリシードファイルの例を以下に示します。

```yaml
cluster:
  enabled: true
  server_address: 192.0.2.102:8443
  cluster_token: eyJzZXJ2ZXJfbmFtZSI6Im5vZGUyIiwiZmluZ2VycHJpbnQiOiJjZjlmNmVhMWIzYjhiNjgxNzQ1YTY1NTY2YjM3ZGUwOTUzNjRmM2MxMDAwMGNjZWQyOTk5NDU5YzY2MGIxNWQ4IiwiYWRkcmVzc2VzIjpbIjE3Mi4xNy4zMC4xODM6ODQ0MyJdLCJzZWNyZXQiOiIxNGJmY2EzMDhkOTNhY2E3MGJmYThkMzE0NWM4NWY3YmE0ZmU1YmYyNmJiNDhmMmUwNzhhOGZhMDczZDc0YTFiIn0=
  member_config:
  - entity: storage-pool
    name: default
    key: source
    value: ""
```

````
````{group-tab} トラストパスワード
追加のサーバーのプリシードファイルは以下の項目を含む必要があります。

```yaml
cluster:
  server_name: <server_name>
  enabled: true
  cluster_address: <IP_address_of_bootstrap_server>
  server_address: <IP_address_of_server>
  cluster_password: <trust_password>
  cluster_certificate: <certificate> # これか cluster_certificate_path を使用します
  cluster_certificate_path: <path_to-certificate_file> # これか cluster_certificate を使用します
```

  YAML 互換の `cluster_certificate` キーを作成するにはブートストラップサーバー上で以下のどちらかのコマンドを実行してください。

   - snap を使用している場合: `sed ':a;N;$!ba;s/\n/\n\n/g' /var/snap/lxd/common/lxd/cluster.crt`
   - そうでない場合: `sed ':a;N;$!ba;s/\n/\n\n/g' /var/lib/lxd/cluster.crt`

  あるいは、`cluster.crt` ファイルをブートストラップサーバーからジョインさせたいサーバーにコピーして `cluster_certificate_path` キーにそのパスを指定します。

新しいクラスタメンバー用のプリシードファイルの例を以下に示します。

```yaml
cluster:
  server_name: server2
  enabled: true
  server_address: 192.0.2.102:8443
  cluster_address: 192.0.2.101:8443
  cluster_certificate: "-----BEGIN CERTIFICATE-----

opyQ1VRpAg2sV2C4W8irbNqeUsTeZZxhLqp4vNOXXBBrSqUCdPu1JXADV0kavg1l

2sXYoMobyV3K+RaJgsr1OiHjacGiGCQT3YyNGGY/n5zgT/8xI0Dquvja0bNkaf6f

...

-----END CERTIFICATE-----
"
  cluster_password: the_password
  member_config:
  - entity: storage-pool
    name: default
    key: source
    value: ""
```

````
`````

## MicroCloudを使う

```{youtube} https://www.youtube.com/watch?v=ZSZoLnp-Ip0
```

LXDクラスタを手動でセットアップする代わりに、[MicroCloud](https://snapcraft.io/microcloud)を使ってすぐに使えるLXDクラスタとCephストレージの環境を作ることができます。

これに必要なsnapパッケージをインストールするには、以下のコマンドを実行します。

    snap install lxd microceph microcloud

次に以下のコマンドでブートストラッププロセスを開始します。

    microcloud init

初期化の行程中に、MicroCloudは他のサーバーを検出、クラスタをセットアップし、Cephに追加するディスクを尋ねるプロンプトを表示します。

初期化が完了したら、CephとLXDクラスタの両方が作られ、LXD自体はネットワークとクラスタ内で使用するのに適したストレージが設定された状態になります。

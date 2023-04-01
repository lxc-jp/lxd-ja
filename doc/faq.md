# よく聞かれる質問 (FAQ)

## 一般的な問題

### LXD サーバをリモートからアクセス可能にするには？

デフォルトでは LXD サーバはローカルの Unix ソケットのみをリッスンしているためネットワークからはアクセス可能ではありません。
リッスンする追加のアドレスを指定することでネットワークから LXD を利用可能にできます。
これは `core.https_address` 設定で実現できます。

現状のサーバ設定を確認するには、以下のコマンドを実行します。

```bash
lxc config show
```

リッスンするアドレスを設定するには、まず利用可能なアドレスを調べた上で、次にサーバで `config set` コマンドを実行します。

```bash
ip addr
lxc config set core.https_address 192.0.2.1
```

{ref}`security_remote_access` も参照してください。

### HTTPS 経由で `lxc remote add` を実行したらパスワードを聞かれたがどうすればよいか？

デフォルトではセキュリティ上の理由から LXD はパスワードを設定していないため、 `lxc remote add` でリモートは追加できません。
パスワードを設定するには LXD が実行中のホスト上で以下のコマンドを実行します。

```bash
lxc config set core.trust_password SECRET
```

これでリモートパスワードが設定されるので、 `lxc remote add` 実行時にこのパスワードを使用できます。

あるいはクライアント証明書を `.config/lxc/client.crt` (`~/.config/lxc/client.crt` またはSnapユーザーの場合は `~/snap/lxd/common/config/client.crt`) からサーバにコピーして以下のコマンドで追加すれば、パスワードを設定しなくてもサーバにアクセスできます。

```bash
lxc config trust add client.crt
```

詳細は {doc}`authentication` を参照してください。

### 自分のホームディレクトリをコンテナ内にバインドマウントできますか？

はい。ディスクデバイスを使って以下のようにすればできます。

```bash
lxc config device add container-name home disk source=/home/${USER} path=/home/ubuntu
```

非特権コンテナの場合は、以下のいずれかも必要です。

- `lxc config device add` の実行に `shift=true` を指定する。これは `shiftfs` がサポートされているかに依存します（`lxc info` 参照）。
- `raw.idmap` エントリーを使用する（[ユーザー名前空間 (user namespace) 用の ID のマッピング](userns-idmap.md) 参照）。
- マウントしたいホームディレクトリに再帰的な POSIX ACL を設定する。

上記のいずれかを実行すればコンテナ内のユーザーは read/write パーミッションに沿ってアクセス可能です。
上記のいずれも設定しない場合、アクセスしようとすると UID/GID (65536:65536) のオーバーフローが発生し、全ユーザーで読み取り可能 (world readable) 以外のファイルへのアクセスは失敗します。

特権コンテナではコンテナ内の UID/GID が外部と同じなためこの問題はありません。
しかしこれは特権コンテナのセキュリティの問題のほとんどの原因でもあります。

### LXD コンテナ内で Docker を動かすには？

LXD のコンテナ内で Docker を動かすにはコンテナの `security.nesting` プロパティを `true` にする必要があります。

```bash
lxc config set <container> security.nesting true
```

LXD コンテナはカーネルモジュールをロードすることはできないので、お使いの Docker の設定によっては、ホストで追加のカーネルモジュールをロードする必要があるかもしれないことに注意してください。

コンテナが必要とするカーネルモジュールのカンマ区切りリストを以下のコマンドで指定すればホストでそれらのモジュールをロードできます。

```bash
lxc config set <container> linux.kernel_modules <modules>
```

コンテナ内に `/.dockerenv` ファイルを作成するとネストした環境内で実行しているために発生するエラーを Docker が無視するようにできるという報告もあります。

### `lxc`はどこに設定を保存していますか？

`lxc`コマンドは、設定を`~/.config/lxc`に保存します。Snapユーザーの場合は、`~/snap/lxd/common/config`に保存されます。

そのディレクトリにはさまざまな設定ファイルが格納されており、その中には以下のものがあります：

- `client.crt`：クライアント証明書（オンデマンドで生成）
- `client.key`：クライアントキー（オンデマンドで生成）
- `config.yml`：設定ファイル（`remotes`、`aliases`などの情報）
- `servercerts/`：`remotes`に属するサーバー証明書が格納されているディレクトリ

## ネットワークの問題

大規模な[プロダクション環境](performance-tuning)では、複数の VLAN を持ち、LXD クライアントを直接それらの VLAN に接続するのが一般的です。
`netplan` と `systemd-networkd` を使っている場合、いくつかの最悪の問題を引き起こす可能性があるバグに遭遇するでしょう。

### VLAN ベースのブリッジでは `netplan` で `systemd-networkd` が使えない

執筆時点（2019-03-05）では、`netplan` は VLAN にアタッチされたブリッジにランダムな MAC アドレスを割り当てられません。
常に同じ MAC アドレスを選択するため、同じネットワークセグメントに複数のマシンが存在する場合、レイヤー 2 の問題が発生します。
複数のブリッジを作成することも困難です。代わりに `network-manager` を使ってください。
設定例は次のようになります。管理アドレスが 10.61.0.25 で、VLAN102 をクライアントのトラフィックに使います。

    network:
      version: 2
      renderer: NetworkManager
      ethernets:
        eth0:
          dhcp4: no
          accept-ra: no
          # This is the 'Management Address'
          addresses: [ 10.61.0.25/24 ]
          gateway4: 10.61.0.1
          nameservers:
            addresses: [ 1.1.1.1, 8.8.8.8 ]
        eth1:
          dhcp4: no
          accept-ra: no
          # A bogus IP address is required to ensure the link state is up
          addresses: [ 10.254.254.25/32 ]

      vlans:
        vlan102:
          accept-ra: no
          dhcp4: no
          id: 102
          link: eth1

      bridges:
        br102:
          accept-ra: no
          dhcp4: no
          interfaces: [ "vlan102" ]
          # A bogus IP address is required to ensure the link state is up
          addresses: [ 10.254.102.25/32 ]
          parameters:
            stp: false

#### 注意事項

- `eth0` はデフォルトゲートウェイの指定がある管理インタフェースです
- `vlan102` は `eth1` を使います
- `br102` は `vlan102` を使います。そして bogus な /32 の IP アドレスが割り当てられています。

他に重要なこととして、`stp: false` を設定することがあります。そうしなければ、ブリッジは最大で 10 秒間 `learning` 状態となります。これはほとんどの DHCP リクエストが投げられる期間よりも長いです。
クロスコネクトされてループを引き起こす可能性はありませんので、このように設定しても安全です。

### port security に気をつける

スイッチは MAC アドレスの変更を許さず、不正な MAC アドレスのトラフィックをドロップするか、ポートを完全に無効にするものが多いです。
ホストから LXD インスタンスに ping できたとしても、異なったホストから ping できない場合は、これが原因の可能性があります。
この原因を突き止める方法は、アップリンク（この場合は `eth1`）で `tcpdump` を実行することです。
すると、応答は送るが ACK を取得できない "ARP Who has `xx.xx.xx.xx` tell `yy.yy.yy.yy`"、もしくは ICMP パケットが行き来しているものの、決して他のホストで受け取られないのが見えるでしょう。

### 不必要に特権コンテナを実行しない

特権コンテナはホスト全体に影響する処理を行うことができます。例えば、ネットワークカードをリセットするために、`/sys` 内のものを使えます。
これは **ホスト全体** に対してリセットを行い、ネットワークの切断を引き起こします。
ほぼすべてのことが非特権コンテナで実行できます。コンテナ内から NFS マウントしたいというような、通常とは異なる特権が必要なケースでは、バインドマウントを使う必要があるかもしれません。

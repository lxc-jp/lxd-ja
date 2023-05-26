(server-expose)=
# LXDをネットワークに公開するには

デフォルトでは、LXDはUnixソケットを介してローカルユーザーからのみ使用でき、ネットワーク経由でアクセスすることはできません。

LXDをネットワークに公開するには、ローカルUnixソケット以外のアドレスをリッスンするように設定する必要があります。これを行うには、[`core.https_address`](server) サーバー設定オプションを設定します。

例えば、LXDサーバをポート`8443`でアクセスできるようにするには、以下のコマンドを入力します。

    lxc config set core.https_address :8443

特定のIPアドレスからのアクセスを許可するには、`ip addr`を使用して利用可能なアドレスを見つけ、それを設定します。例えば：

```{terminal}
:input: ip addr

1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: enp5s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether 00:16:3e:e3:f3:3f brd ff:ff:ff:ff:ff:ff
    inet 10.68.216.12/24 metric 100 brd 10.68.216.255 scope global dynamic enp5s0
       valid_lft 3028sec preferred_lft 3028sec
    inet6 fd42:e819:7a51:5a7b:216:3eff:fee3:f33f/64 scope global mngtmpaddr noprefixroute
       valid_lft forever preferred_lft forever
    inet6 fe80::216:3eff:fee3:f33f/64 scope link
       valid_lft forever preferred_lft forever
3: lxdbr0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 00:16:3e:8d:f3:72 brd ff:ff:ff:ff:ff:ff
    inet 10.64.82.1/24 scope global lxdbr0
       valid_lft forever preferred_lft forever
    inet6 fd42:f4ab:4399:e6eb::1/64 scope global
       valid_lft forever preferred_lft forever
:input: lxc config set core.https_address 10.68.216.12
```

全てのリモートクライアントはLXDに接続して公開利用とマークされた任意のイメージにアクセスできます。

(server-authenticate)=
## LXDサーバでの認証

リモートAPIにアクセスできるようにするには、クライアントはLXDサーバに認証しなければなりません。
いくつかの認証方法があります。詳細は{ref}`authentication`を参照してください。

お勧めの方法はクライアントのTLS証明書をトラストトークンを使ってサーバのトラストストアに追加することです。
トラストトークンを使ってクライアントを認証するには、以下の手順を実行します。

1. サーバで、以下のコマンドを入力します。

       lxc config trust add

   追加したいクライアントの名前を入力します。
   クライアント証明書を追加するのに使用できるトークンをコマンドが生成し表示します。
1. クライアントで、以下のコマンドでサーバを追加します。

       lxc remote add <remote_name> <token>

   % Include content from [../authentication.md](../authentication.md)
```{include} ../authentication.md
    :start-after: <!-- Include start NAT authentication -->
    :end-before: <!-- Include end NAT authentication -->
```

詳細や他の認証方法については{ref}`authentication`を参照してください。

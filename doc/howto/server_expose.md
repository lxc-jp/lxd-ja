(server-expose)=
# LXDをネットワークに公開するには

デフォルトでは、LXDはUnixソケットを使ってローカルユーザだけが使えます。

LXDをネットワークに公開するには[`core.https_address`](server)サーバ設定をセットします。
例えば、LXDサーバをポート`8443`でアクセスできるようにするには、以下のコマンドを入力します。

    lxc config set core.https_address :8443

全てのリモートクライアントはLXDに接続して公開利用とマークされた任意のイメージにアクセスできます。

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

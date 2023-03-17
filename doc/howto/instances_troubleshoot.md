---
discourse: 7362
---

(instances-troubleshoot)=
# インスタンスの起動に失敗する問題のトラブルシューティング方法

インスタンスの起動に失敗し、エラー状態になる場合、これは通常、インスタンスの作成に使用したイメージまたはサーバー設定に関連する大きな問題を示しています。

問題のトラブルシューティングを行うには、以下の手順を完了してください：

1. 関連するログファイルとデバッグ情報を保存します：

   インスタンスログ
   : インスタンスログを表示するには、次のコマンドを入力します：

         lxc info <instance_name> --show-log

   コンソールログ
   : コンソールログを表示するには、次のコマンドを入力します：

         lxc console <instance_name> --show-log

   詳細なサーバ情報
   : LXDのsnapパッケージには、デバッグ用の関連サーバー情報を収集するツールが含まれています。
   それを実行するには、次のコマンドを入力します：

         sudo lxd.buginfo

1. LXDサーバーを実行しているマシンを再起動します。
1. インスタンスをもう一度起動してみてください。
   エラーが再発した場合は、ログを比較して同じエラーかどうか確認します。

   同じエラーであり、ログ情報からエラーの原因を特定できない場合は、[フォーラム](https://discuss.linuxcontainers.org)で質問を投稿してください。
   収集したログファイルを含めるようにしてください。

## トラブルシューティングの例

この例では、systemdが起動できないRHEL 7システムを調査しましょう。

```{terminal}
:input: lxc console --show-log systemd

Console log:

Failed to insert module 'autofs4'
Failed to insert module 'unix'
Failed to mount sysfs at /sys: Operation not permitted
Failed to mount proc at /proc: Operation not permitted
[!!!!!!] Failed to mount API filesystems, freezing.
```

ここでのエラーは、 /sys と /proc がマウントできないと言っています - これは、非特権コンテナでは正しいです。
しかし、LXDは可能であればこれらのファイルシステムを自動的にマウントします。

{doc}`コンテナの要件 <../container-environment>`では、すべてのコンテナには空の `/dev`、`/proc`、`/sys` ディレクトリが必要であり、`/sbin/init` が存在していなければなりません。
これらのディレクトリが存在しない場合、LXDはそれらをマウントできず、`systemd`がその後それを試みます。
これは非特権コンテナなので、`systemd`はこれを行う能力がなく、コンテナはフリーズします。

したがって、何も変更される前の環境を確認でき、`raw.lxc` 設定パラメータを使用してコンテナ内の`init`システムを明示的に変更できます。
これは、Linuxカーネルのコマンドラインで `init=/bin/bash` を設定するのと同等です。

    lxc config set systemd raw.lxc 'lxc.init.cmd = /bin/bash'

これがどのように見えるかを示します：

```{terminal}
:input: lxc config set systemd raw.lxc 'lxc.init.cmd = /bin/bash'

:input: lxc start systemd
:input: lxc console --show-log systemd

Console log:

[root@systemd /]#
```

これでコンテナが起動したので、確認して期待通りにうまく動作していないことがわかります：

```{terminal}
:input: lxc exec systemd bash

[root@systemd ~]# ls
[root@systemd ~]# mount
mount: failed to read mtab: No such file or directory
[root@systemd ~]# cd /
[root@systemd /]# ls /proc/
sys
[root@systemd /]# exit
```

LXDは自動的に回復しようとするため、起動時にいくつかのディレクトリが作成されました。
コンテナをシャットダウンして再起動すると問題が解決しますが、元の原因はまだ残っています - テンプレートには必要なファイルが含まれていません。

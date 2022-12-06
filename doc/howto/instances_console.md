---
discourse: 9223
---

(instances-console)=
# コンソールにアクセスするには

インスタンスのコンソールにアタッチするには `lxc console` コマンドを使います。
コンソールは起動時に既に利用可能になり、必要なら、ブートメッセージを見て、コンテナや仮想マシンの起動時の問題をデバッグするのに使えます。

インタラクティブなコンソールに接続するには、以下のコマンドを入力します。

    lxc console <instance_name>

ログ出力を見るには `--show-log` フラグを渡します。

    lxc console <instance_name> --show-log

インスタンスが起動したらすぐにコンソールにアタッチできます。

    lxc start <instance_name> --console
    lxc start <instance_name> --console=vga

## グラフィカルなコンソールにアタッチする (仮想マシンの場合)

```{youtube} https://www.youtube.com/watch?v=pEUsTMiq4B4
```

仮想マシンでは、コンソールにログオンしてグラフィカルな出力を見ることができます。
コンソールを使えば、例えば、グラフィカルなインタフェースを使ってオペレーティングシステムをインストールしたりデスクトップ環境を実行できます。

さらなる利点は `lxd-agent` プロセスが実行していなくても、コンソールは利用可能です。
これは `lxd-agent` が起動する前や `lxd-agent` が全く利用可能でない場合にもコンソール経由で仮想マシンにアクセスできることを意味します。

仮想マシンにグラフィカルなアウトプットを持つ VGA コンソールを開始するには、 SPICE クライアント (例えば、`virt-viewer` または `spice-gtk-client`) をインストールする必要があります。
次に以下のコマンドを入力します。

    lxc console <vm_name> --type vga

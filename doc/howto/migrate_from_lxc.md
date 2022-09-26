(migrate-from-lxc)=
# LXC から LXD にコンテナをマイグレートするには

LXD は LXC のコンテナを LXD サーバにインポートするためのツール (`lxc-to-lxd`) を提供しています。
LXD コンテナは LXD サーバと同じマシン上に存在する必要があります。

このツールは LXC コンテナを分析し、データと設定の両方を新しい LXD コンテナにマイグレートします。

```{note}
あるいは LXC コンテナ内で `lxd-migrate` ツールを使用して LXD にマイグレートすることもできます ({ref}`import-machines-to-instances` 参照)。
しかし、このツールは LXC コンテナの設定は一切マイグレートしません。
```

## ツールを取得する

snap をお使いの場合、`lxc-to-lxd` は自動でインストールされます。
`lxd.lxc-to-lxd` で利用できます。

そうでない場合、 `go` (バージョン 1.18 以降) がインストールされていることを確認の上、以下のコマンドでツールをインストールしてください。

    go install github.com/lxc/lxd/lxc-to-lxd@latest

## LXC コンテナを用意する

一度に1つのコンテナをマイグレートすることもできますし、同時にあなたの全ての LXC コンテナをマイグレートすることもできます。

```{note}
マイグレートされたコンテナは元のコンテナと同じ名前を使用します。
LXD にインスタンス名としてすでに存在する名前を持つコンテナをマイグレートすることはできません。

このため、マイグレーションプロセスを開始する前に名前の衝突を引き起こす可能性のある LXC コンテナはリネームしてください。
```

マイグレーションプロセスを開始する前に、マイグレートしたいコンテナを停止してください。

## マイグレーションプロセスを開始する

コンテナをマイグレートするには `sudo lxd.lxc-to-lxd [flags]` と実行してください。
(このコマンドはあなたが snap を使用していると想定しています。そうでない場合 `lxd.lxc-to-lxd` を `lxc-to-lxd` と読み替えてください。以下の例でも同様です)

例えば、全てのコンテナをマイグレートするには

    sudo lxd.lxc-to-lxd --all

`lxc1` コンテナだけをマイグレートするには

    sudo lxd.lxc-to-lxd --containers lxc1

2 つのコンテナ (`lxc1` と `lxc2`) をマイグレートし LXD 内の `my-storage` ストレージプールを使用するには 

    sudo lxd.lxc-to-lxd --containers lxc1,lxc2 --storage my-storage

実際に実行せずに全てのコンテナのマイグレートをテストするには

    sudo lxd.lxc-to-lxd --all --dry-run

全てのコンテナをマイグレートするが、`rsync` の帯域幅を 5000 KB/s に限定するには

    sudo lxd.lxc-to-lxd --all --rsync-args --bwlimit=5000

全ての利用可能なフラグを確認するには `sudo lxd.lxc-to-lxd --help` と実行してください。

```{note}
`linux64` アーキテクチャがサポートされない (`linux64` architecture isn't supported) というエラーが出る場合、ツールを最新版にアップデートするか LXC コンテナ内のアーキテクチャを `linux64` から `amd64` か `x86_64` に変更してください。
```

## 設定を確認する

このツールは LXC の設定と (1つまたは複数の) コンテナの設定を分析し、可能な限りの範囲で設定をマイグレートします。
以下のような実行結果が出力されます。

```bash
Parsing LXC configuration
Checking for unsupported LXC configuration keys
Checking for existing containers
Checking whether container has already been migrated
Validating whether incomplete AppArmor support is enabled
Validating whether mounting a minimal /dev is enabled
Validating container rootfs
Processing network configuration
Processing storage configuration
Processing environment configuration
Processing container boot configuration
Processing container apparmor configuration
Processing container seccomp configuration
Processing container SELinux configuration
Processing container capabilities configuration
Processing container architecture configuration
Creating container
Transferring container: lxc1: ...
Container 'lxc1' successfully created
```

マイグレーションプロセスが完了したら、設定を確認し、必要に応じて、マイグレートした LXD コンテナを起動する前に LXD 内の設定を更新してください。

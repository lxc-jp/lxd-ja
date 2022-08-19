(import-machines-to-instances)=
# 物理または仮想マシンを LXD インスタンスにインポートするには

LXD は既存のディスクやイメージに基づく LXD インスタンスを作成するツール (`lxd-migrate`) を提供しています。

このツールは Linux マシン上で実行できます。
まずLXD サーバに接続して空のインスタンスを作成します。このインスタンスはマイグレーション中またはマイグレーション後に設定を変更できます。
次にこのツールはあなたが用意したディスクまたはイメージからインスタンスにデータをコピーします。

このツールはコンテナと仮想マシンの療法を作成できます。
* コンテナを作成する際は、コンテナのルートファイルシステムを含むディスクまたはパーティションを用意する必要があります。
  例えば、これはあなたがツールを実行しているマシンまたはコンテナの `/` ルートディスクかもしれません。
* 仮想マシンを作成する際は、起動可能なディスク、パーティション、またはイメージを用意する必要があります。
  これは単にファイルシステムを用意するだけでは不十分であり、実行中のコンテナから仮想マシンを作成することはできないことを意味します。
  また使用中の物理マシンから仮想マシンを作成することもできません。これはマイグレーションツールがコピーしようとするディスクを使用中になるからです。
  変わりに、起動可能なディスク、起動可能なパーティション、または現在使用中でないディスクを用意してください。

既存のマシンを LXD インスタンスにマイグレートするには以下の手順を実行してください。

1. 最新の [LXD release](https://github.com/lxc/lxd/releases) の **Assets** セクションから `bin.linux.lxd-migrate` ツールをダウンロードしてください。
1. ツールをインスタンスを作成したいマシン上に配置して
   (通常 `chmod u+x bin.linux.lxd-migrate` を実行して) 実行可能にしてください。
1. マシンに `rsync` がインストールされているか確認してください。
   インストールされていない場合は (例えば `sudo apt install rsync` で) インストールしてください。
1. 以下のようにツールを実行します。

       ./bin.linux.lxd-migrate

   ツールはマイグレーソンに必要な情報を入力するようプロンプトを出します。

   ```{tip}
   ツールをインタラクティブに実行する代わりの方法として、設定をパラメータでコマンドに指定することもできます。
   詳細は `./bin.linux.lxd-migrate --help` を参照してください。
   ```

   1. LXD サーバの URL を、 IP アドレスまたは DNS 名で指定してください。
   1. 証明書のフィンガープリントを確認してください。
   1. 認証の方法を選択してください ({ref}`authentication` 参照)。

      例えば、証明書トークンを選ぶ場合、 LXD サーバにログオンしてマイグレーションツールを実行中のマシン用のトークンを `lxc config trust add` で作成してください。
      次に生成されたトークンをツールを認証するのに使用してください。
   1. コンテナと仮想マシンのどちらを作成するか選択してください。
   1. 作成するインスタンスの名前を指定してください。
   1. ルートファイルシステム (コンテナの場合)、起動可能なディスク、パーティションまたはイメージファイル (仮想マシンの場合) のパスを指定します。
   1. コンテナの場合、任意で追加でファイルシステムのマウントを指定できます。
   1. 仮想マシンの場合、セキュアブートがサポートさているかを指定します。
   1. 任意で、新しいインスタンスを設定します。
      プロファイルを指定するか、オプションやストレージを変更したりネットワークを設定する設定オプションを直接指定できます。

      あるいは、マイグレーション後に新しいインスタンスを設定することもできます。
   1. マイグレーションの設定が完了したら、マイグレーションプロセスを開始します。

   <details>
   <summary>展開して出力の例を見る</summary>

   ```
   Please provide LXD server URL: https://192.0.2.7:8443
   Certificate fingerprint: xxxxxxxxxxxxxxxxx
   ok (y/n)? y

   1) Use a certificate token
   2) Use an existing TLS authentication certificate
   3) Generate a temporary TLS authentication certificate
   Please pick an authentication mechanism above: 1
   Please provide the certificate token: xxxxxxxxxxxxxxxx

   Remote LXD server:
     Hostname: bar
     Version: 5.4

   Would you like to create a container (1) or virtual-machine (2)?: 1
   Name of the new instance: foo
   Please provide the path to a root filesystem: /
   Do you want to add additional filesystem mounts? [default=no]:

   Instance to be created:
     Name: foo
     Project: default
     Type: container
     Source: /

   Additional overrides can be applied at this stage:
   1) Begin the migration with the above configuration
   2) Override profile list
   3) Set additional configuration options
   4) Change instance storage pool or volume size
   5) Change instance network

   Please pick one of the options above [default=1]: 3
   Please specify config keys and values (key=value ...): limits.cpu=2

   Instance to be created:
     Name: foo
     Project: default
     Type: container
     Source: /
     Config:
       limits.cpu: "2"

   Additional overrides can be applied at this stage:
   1) Begin the migration with the above configuration
   2) Override profile list
   3) Set additional configuration options
   4) Change instance storage pool or volume size
   5) Change instance network

   Please pick one of the options above [default=1]: 4
   Please provide the storage pool to use: default
   Do you want to change the storage size? [default=no]: yes
   Please specify the storage size: 20GiB

   Instance to be created:
     Name: foo
     Project: default
     Type: container
     Source: /
     Storage pool: default
     Storage pool size: 20GiB
     Config:
       limits.cpu: "2"

   Additional overrides can be applied at this stage:
   1) Begin the migration with the above configuration
   2) Override profile list
   3) Set additional configuration options
   4) Change instance storage pool or volume size
   5) Change instance network

   Please pick one of the options above [default=1]: 5
   Please specify the network to use for the instance: lxdbr0

   Instance to be created:
     Name: foo
     Project: default
     Type: container
     Source: /
     Storage pool: default
     Storage pool size: 20GiB
     Network name: lxdbr0
     Config:
       limits.cpu: "2"

   Additional overrides can be applied at this stage:
   1) Begin the migration with the above configuration
   2) Override profile list
   3) Set additional configuration options
   4) Change instance storage pool or volume size
   5) Change instance network

   Please pick one of the options above [default=1]: 1
   Instance foo successfully created
   ```
   </details>

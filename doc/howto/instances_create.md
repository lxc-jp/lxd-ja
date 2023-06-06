(instances-create)=
# インスタンスを作成するには

インスタンスを作成するには、`lxc init` か `lxc launch` コマンドを使用できます。
`lxc init` コマンドはインスタンスを作成だけしますが、`lxc launch` コマンドは作成して起動します。

## 使い方

コンテナを作成するには以下のコマンドを入力します。

    lxc launch|init <image_server>:<image_name> <instance_name> [flags]

イメージ
: イメージは必要最小限のオペレーティングシステム (例えば、Linux ディストリビューション) と LXD 関連の情報を含みます。
  さまざまなオペレーティングシステムのイメージがビルトインのリモートイメージサーバで利用できます。
  詳細は {ref}`images` を参照してください。

  イメージがローカルにない場合、イメージサーバとイメージの名前を指定 (例えば、LXD のビルトインのイメージサーバ上の 22.04 Ubuntu イメージなら `images:ubuntu/22.04`) する必要があります。

インスタンス名
: インスタンス名は LXD の運用環境 (そしてクラスタ内) でユニークである必要があります。
  追加の要件については {ref}`instance-properties` を参照してください。

フラグ
: フラグの完全なリストについては `lxc launch --help` か `lxc init --help` を参照してください。
  よく使うフラグは以下のとおりです。

  - `--config` は新しいインスタンスの設定オプションを指定します
  - `--device` はプロファイルを通して提供されるデバイスの {ref}`デバイスオプション <devices>` を上書きします
  - `--profile` は新しいインスタンスに使用する {ref}`プロファイル <profiles>` を指定します
  - `--network` や `--storage` は新しいインスタンスに指定のネットワークやストレージプールを使用させます
  - `--target` は指定のクラスタメンバー上にインスタンスを作成します
  - `--vm` はコンテナではなく仮想マシンを作成します

## 設定ファイルを渡す

インスタンス設定をフラグとして指定する代わりに、YAML ファイルでコマンドに渡すことができます。

例えば、`config.yaml` の設定でコンテナを起動するには、以下のコマンドを入力します。

    lxc launch images:ubuntu/22.04 ubuntu-config < config.yaml

```{tip}
YAML ファイルの必要な文法を見るには既存のインスタンス設定の中身を確認 (`lxc config show <instance_name> -e`) してください。
```

## 例

以下の例では `lxc launch` を使用しますが、同じように `lxc init` も使用できます。

### コンテナを起動する

`images` サーバの Ubuntu 22.04 のイメージで `ubuntu-container` というインスタンス名でコンテナを起動するには、以下のコマンドを入力します。

    lxc launch images:ubuntu/22.04 ubuntu-container

### 仮想マシンを起動する

`images` サーバの Ubuntu 22.04 のイメージで `ubuntu-vm` というインスタンス名で仮想マシンを起動するには、以下のコマンドを入力します。

    lxc launch images:ubuntu/22.04 ubuntu-vm --vm

### コンテナを指定の設定で起動する

コンテナを起動しリソースを 1 つの vCPU と 192MiB の RAM に限定するには、以下のコマンドを入力します。

    lxc launch images:ubuntu/22.04 ubuntu-limited --config limits.cpu=1 --config limits.memory=192MiB

### 指定のクラスタメンバー上で仮想マシンを起動する

クラスタメンバー `server2` 上で仮想マシンを起動するには、以下のコマンドを入力します。

    lxc launch images:ubuntu/22.04 ubuntu-container --vm --target server2

### 指定のインスタンスタイプでコンテナを起動する

LXD ではクラウドのシンプルなインスタンスタイプが使えます。これは、インスタンスの作成時に指定できる文字列で表されます。

3 つの指定方法があります:

- `<instance type>`
- `<cloud>:<instance type>`
- `c<CPU>-m<RAM in GB>`

例えば、次の 3 つは同じです:

- `t2.micro`
- `aws:t2.micro`
- `c1-m1`

コマンドラインでは、インスタンスタイプは次のように指定します:

```bash
lxc launch ubuntu:22.04 my-instance -t t2.micro
```

使えるクラウドとインスタンスタイプのリストは [`https://github.com/dustinkirkland/instance-type`](https://github.com/dustinkirkland/instance-type) で確認できます。

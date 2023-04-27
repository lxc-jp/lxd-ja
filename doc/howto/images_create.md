(images-create)=
# イメージを作成するには

独自のイメージを作成し共有したい場合、既存のインスタンスやスナップショットをベースにすることもできますし、一から独自のイメージを作ることもできます。

(images-create-publish)=
## インスタンスやスナップショットからイメージを発行する

インスタンスやインスタンススナップショットを新しいインスタンスのベースとして使いたい場合、それらからイメージを作成し発行するのが良いです。

インスタンスからイメージを発行するには、インスタンスが停止されていることを確認してください。
次に以下のコマンドを入力します。

    lxc publish <instance_name> [<remote>:]

スナップショットからイメージを発行するには、以下のコマンドを入力します。

    lxc publish <instance_name>/<snapshot_name> [<remote>:]

どちらの場合も`--alias`フラグで新しいイメージにエイリアスを設定し、`--expire`で有効期限を設定し、`--public`でイメージを公開状態にすることができます。
利用可能な全てのフラグ一覧は`lxc publish --help`を参照してください。

発行のプロセスはインスタンスやスナップショットからtarballを生成した後圧縮するため、かなりの時間がかかるかもしれません。特にI/OとCPUの負荷が高いため、発行の操作はLXDで直列化(訳注:1つずつ順に実行)されます。

### 発行用にインスタンスを準備する

インスタンスからイメージを発行する前に、イメージに含めるべきでない全てのデータをクリーンアップしてください。
通常、これは以下のデータを含みます。

- インスタンスメタデータ(編集には`lxc config metadata`を使ってください)
- ファイルテンプレート(編集には`lxc config template`を使ってください)
- インスタンス自身の内部のインスタンスに特有なデータ(例えば、ホストのSSH鍵と`dbus/systemd machine-id`)

(images-create-build)=
## イメージをビルドする

独自イメージをビルドするには、[`distrobuilder`](https://github.com/lxc/distrobuilder)が使用できます。

インストール手順とツールの使い方は[`distrobuilder`のドキュメント](https://distrobuilder.readthedocs.io/en/latest/)を参照してください。
(instances_manage)=
# インスタンスを管理するには

全てのインスタンスを一覧表示するには以下のコマンドを入力します。

    lxc list

表示するインスタンスをフィルターできます。例えば、インスタンスタイプ、状態、またはインスタンスが配置されているクラスタメンバーでフィルターできます。

    lxc list type=container
    lxc list status=running
    lxc list location=server1

インスタンス名でフィルターもできます。
複数のインスタンスを一覧表示するには、名前の正規表現を使います。
例えば以下のようにします。

    lxc list ubuntu.*

全てのフィルターオプションを見るには `lxc list --help` と入力します。

## インスタンスの情報を表示する

インスタンスについての詳細情報を表示するには以下のコマンドを入力します。

    lxc info <instance_name>

インスタンスの最新のログの行を表示するにはコマンドに `--show-log` を追加します。

    lxc info <instance_name> --show-log

## インスタンスを起動する

インスタンスを起動するには以下のコマンドを入力します。

    lxc start <instance_name>

インスタンスが存在しないか既に稼働中の場合はエラーになります。

起動する際にコンソールにすぐにアタッチするには `--console` フラグを渡します。
例えば以下のようにします。

    lxc start <instance_name> --console

詳細は {ref}`instances-console` を参照してください。

(instances-manage-stop)=
## インスタンスを停止する

インスタンスを停止するには以下のコマンドを入力します。

    lxc stop <instance_name>

インスタンスが存在しないか稼働中ではない場合はエラーになります。

## インスタンスを削除する

インスタンスがもう不要な場合、削除できます。
削除する前にインスタンスを停止する必要があります。

インスタンスを削除するには以下のコマンドを入力します。

    lxc delete <instance_name>

```{caution}
このコマンドはインスタンスとそのスナップショットを永久的に削除します。
```

### 間違ってインスタンスを削除するのを防ぐ

間違ってインスタンスを削除するのを防ぐには 2 つの方法があります。

- `lxc delete` コマンドを使うたびに承認のプロンプトを表示するには、エイリアスを作成します。

       lxc alias add delete "delete -i"

- 特定のインスタンスが削除されることを防ぐためには、そのインスタンスの [`security.protection.delete`](instance-options-security) を `true` に設定します。
  手順は {ref}`instances-configure` を参照してください。

## インスタンスを再構築する

インスタンスのrootディスクを一掃して再初期化したいがインスタンスの設定は維持したい場合、インスタンスを再構築できます。

再構築はスナップショットが1つも存在しないインスタンスでのみ可能です。

再構築の前にインスタンスを停止します。
そして、以下のコマンドのいずれかを入力します。

- 別のイメージでインスタンスを再構築する。

        lxc rebuild <image_name> <instance_name>

- 空のルートディスクでインスタンスを再構築する。

        lxc rebuild <instance_name> --empty

`rebuild`コマンドについてのより詳細な情報は`lxc rebuild --help`を参照してください。

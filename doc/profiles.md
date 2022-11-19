(profiles)=
# プロファイルを使用するには

プロファイルは一組の設定オプションを保持します。
プロファイルにはインスタンスオプション、デバイスとデバイスオプションを含められます。

1つのインスタンスには任意の数のプロファイルを適用できます。
プロファイルは指定された順番に適用され、その結果最後に指定したプロファイルが特定のキーを上書きします。
どのような場合でも、インスタンス固有の設定はプロファイル由来のものを上書きします。

```{note}
プロファイルはコンテナと仮想マシンに適用できます。
ですので、どちらのタイプに有効なオプションとデバイスを含めることができます。

インスタンスタイプに適用できない設定を含むプロファイルを適用すると、この設定は無視されエラーにはなりません。
```

新しいインスタンスを起動する際にプロファイルを指定しない場合は、自動的には`default`プロファイルが適用されます。
このプロファイルはネットワークインタフェースとルートディスクを定義します。
`default` プロファイルはリネームや削除はできません。

## プロファイルを表示する

全ての利用可能なプロファイルを一覧表示するには以下のコマンドを入力します。

    lxc profile list

プロファイルの内容を表示するには以下のコマンドを入力します。

    lxc profile show <profile_name>

## 空のプロファイルを作成する

空のプロファイルを作成するには以下のコマンドを入力します。

    lxc profile create <profile_name>

## プロファイルを編集する

プロファイルの特定の設定オプションを設定するか、あるいはYAML形式でプロファイル全体を編集できます。

### プロファイルの特定の設定オプションを設定する

プロファイルのインスタンスオプションを設定するには、`lxc profile set`コマンドを使います。
プロファイル名とインスタンスオプションのキーとバリューを指定します。

    lxc profile set <profile_name> <option_key>=<option_value> <option_key>=<option_value> ...

プロファイルのインスタンスデバイスを追加と変更するには、`lxc profile device add`コマンドを使います。
プロファイル名、デバイス名、デバイスタイプと({ref}`デバイスタイプ <device-types>`ごとの)必要に応じてデバイスオプションを指定します。

    lxc profile device add <instance_name> <device_name> <device_type> <device_option_key>=<device_option_value> <device_option_key>=<device_option_value> ...

以前にプロファイルに追加したデバイスのインスタンスデバイスオプションを設定するには、`lxc profile device set`コマンドを使います。

    lxc profile device set <instance_name> <device_name> <device_option_key>=<device_option_value> <device_option_key>=<device_option_value> ...

### プロファイル全体を編集する

っ子の設定オプションを別々に設定する代わりに、YAML形式で一度にすべてのオプションを提供できます。

既存のプロファイルまたはインスタンス設定の中身で必要なマークアップをチェックします。
例えば、`default`プロファイルは以下のようになっているかもしれません。

    config: {}
    description: Default LXD profile
    devices:
      eth0:
        name: eth0
        network: lxdbr0
        type: nic
      root:
        path: /
        pool: default
        type: disk
    name: default
    used_by:

インスタンスオプションは`config`の下の配列として提供されます。
インスタンスデバイスとインスタンスデバイスオプションは`devices`の下に提供されます。

ターミナルの標準エディタを使ってプロファイルを編集するには、以下のコマンドを入力します。

    lxc profile edit <profile_name>

別の方法として、設定を含むYAMLファイル( 例えば、`profile.yaml`)を作成して、以下のコマンドで設定をプロファイルに書き込めます。

    lxc profile edit <profile_name> < profile.yaml

## インスタンスにプロファイルを適用する

インスタンスにプロファイルを適用するには以下のコマンドを入力します。

    lxc profile add <instance_name> <profile_name>

```{tip}
プロファイル追加後に`lxc config show <instance_name>`を実行して設定を確認します。

プロファイルが`profiles`の下に一覧表示されます。
しかし、プロファイルからの設定オプションは`config`の下には表示されません(`--expanded`フラグを追加しない限り)。
この挙動の理由はこれらの設定はプロファイルからは取得されインスタンス設定から取得されるわけではないからです。

これはプロファイルを編集する場合、変更はプロファイルを使用している全てのインスタンスに自動的に適用されることを意味します。
```

インスタンスの起動時に`--profile`フラグを追加してプロファイルを指定することもできます。

    lxc launch <image> <instance_name> --profile <profile> --profile <profile> ...

## インスタンスからプロファイルを削除する

インスタンスからプロファイルを削除するには以下のコマンドを入力します。

    lxc profile remove <instance_name> <profile_name>

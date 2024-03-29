(instances-configure)=
# インスタンスを設定するには

{ref}`instance-options` を設定するか {ref}`devices` を設定することでインスタンスを設定できます。

設定方法は以下の項を参照してください。

```{note}
異なるインスタンス設定を保管し再利用するには、{ref}`プロファイル <profiles>` を使用してください。
```

(instances-configure-options)=
## インスタンスオプションを設定する

{ref}`インスタンスを作成する <instances-create>` 際にインスタンスオプションを指定できます。

インスタンスが作成された後にインスタンスオプションを更新するには、`lxc config set` コマンドを使います。
インスタンス名とインスタンスオプションのキーとバリューを指定します。

    lxc config set <instance_name> <option_key>=<option_value> <option_key>=<option_value> ...

利用可能なオプションの一覧とどのオプションがどのインスタンスタイプで利用可能かの情報は {ref}`instance-options` を参照してください。

例えば、コンテナのメモリーリミットを変更するには、以下のコマンドを入力します。

    lxc config set my-container limits.memory=128MiB

```{note}
一部のインスタンスオプションはインスタンスが稼働中に即座に更新されます。
他のインスタンスオプションはインスタンスの再起動後に更新されます。

どのオプションがインスタンス稼働中に即座に反映されるかの情報は {ref}`instance-options` の "ライブアップデート" 列を参照してください。
```

(instances-configure-devices)=
## デバイスを設定する

インスタンスにインスタンスデバイスを追加や設定するには、`lxc config device add` コマンドを使います。
一般的に、デバイスはコンテナの稼働中に追加または削除できます。
VMはいくつかのデバイスタイプではホットプラグをサポートしますが、全てではありません。

インスタンス名、デバイス名、デバイスタイプと ({ref}`デバイスタイプ <devices>` ごとに) 必要に応じてデバイスオプションを指定します。

    lxc config device add <instance_name> <device_name> <device_type> <device_option_key>=<device_option_value> <device_option_key>=<device_option_value> ...

利用可能なデバイスタイプとそのオプションについては {ref}`devices` を参照してください。

```{note}
各デバイスのエントリはインスタンスごとにユニークな名前により識別します。

プロファイルに定義されたデバイスは、プロファイルがインスタンスに割り当てられる順番にインスタンスに適用されます。
インスタンス設定内に直接定義されたデバイスは最後に適用されます。
各ステージで、より以前のステージに同じ名前のデバイスがある場合は、デバイスエントリ全体が最後の定義により上書きされます。

デバイス名は最大64文字です。
```

例えば、ホストシステムの `/share/c1` 上のストレージをインスタンスのパス `/opt` に追加するには、以下のコマンドを入力します。

    lxc config device add my-container disk-storage-device disk source=/share/c1 path=/opt

以前追加したデバイスのインスタンスデバイスオプションを設定するには、以下のコマンドを入力します。

    lxc config device set <instance_name> <device_name> <device_option_key>=<device_option_value> <device_option_key>=<device_option_value> ...

```{note}
デバイスオプションは {ref}`インスタンスの作成 <instances-create>` 時に `--device` フラグを使って指定することもできます。
これは {ref}`プロファイル <profiles>` を通して提供されるデバイスのデバイスオプションを上書きしたい場合に有用です。
```

デバイスを除去するには、`lxc config device remove` コマンドを使います。
利用可能なコマンドの完全なリストは `lxc config device --help` を参照してください。

## インスタンス設定を表示する

書き込み可能なインスタンスプロパティ、インスタンスオプション、デバイスとデバイスオプションを含むインスタンスの現在の設定を表示するには、以下のコマンドを入力します。

    lxc config show <instance_name> --expanded

(instances-configure-edit)=
## インスタンス設定全体を編集する

書き込み可能なインスタンスプロパティ、インスタンスオプション、デバイスとデバイスオプションを含むインスタンス設定全体を編集するには、以下のコマンドを入力します。

    lxc config edit <instance_name>

```{note}
利便性のため、`lxc config edit` コマンドは読み取り専用のインスタンスプロパティを含む設定全体を表示します。
しかし、これらのプロパティは変更できません。
変更しても無視されます。
```

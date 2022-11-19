(instance-config)=
# インスタンスの設定

インスタンス設定は以下の異なるカテゴリから構成されます。

インスタンスプロパティ
: インスタンスプロパティはインスタンスが作成されるときに設定されます。
  これには、例えば、インスタンス名やアーキテクチャが含まれます。
  これらのプロパティはインスタンス生成時に指定されます。
  いくつかのプロパティは読み取り専用で作成後は変更できませんが、他のプロパティは {ref}`インスタンス設定全体を編集する <instances-configure-edit>` 際に更新できます。

  YAML 設定内では、プロパティはトップレベルにあります。

  利用可能なインスタンスプロパティのリファレンスは {ref}`instance-properties` を参照してください。

インスタンスオプション
: インスタンスオプションはインスタンスに直接関連する設定オプションです。
  これには、例えば、起動時のオプション、セキュリティ設定、ハードウェアのリミット、カーネルモジュール、スナップショット、そしてユーザの鍵を含みます。
  これらのオプションはインスタンスの作成時に (`--config key=value` フラグを使って) キー/バリューペアで指定できます。
  作成後は `lxc config set` や `lxc config unset` コマンドで変更できます。

  YAML 設定内では、オプションは `config` エントリの下に配置されます。

  利用可能なインスタンスオプションのリファレンスは {ref}`instance-options` を参照してください。

```{toctree}
:maxdepth: 1
:hidden:

../reference/instance_properties.md
../reference/instance_options.md
QEMU 設定をオーバーライド <../howto/instance_qemu_config.md>
../reference/instance_units.md
```

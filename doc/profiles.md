# プロファイル
<!-- Profiles -->
<!--
Profiles can store any configuration that an instance can (key/value or devices)
and any number of profiles can be applied to an instance.
-->
プロファイルはインスタンスが保持できる（キー・バリューやデバイスなどの）あらゆる
設定を保持することができ、プロファイルをいくつでもインスタンスに適用することが
できます。

<!--
Profiles are applied in the order they are specified so the last profile to
specify a specific key wins.
-->
プロファイルは指定された順番に適用され、その結果最後に指定したプロファイルが
特定のキーを上書きします。

<!--
In any case, instance-specific configuration always overrides that coming from
the profiles.
-->
どのような場合でも、インスタンス固有の設定はプロファイル由来のものを上書きします。

## デフォルトのプロファイル <!-- Default profile -->
<!--
If not present, LXD will create a `default` profile.
-->
まだ存在していない場合は、LXDは `default` プロファイルを作成します。

<!--
The `default` profile cannot be renamed or removed.
The `default` profile is set for any new instance created which doesn't
specify a different profiles list.
-->
`default` プロファイルはリネームや削除はできません。
`default` プロファイルは異なるプロファイルリストを指定せずに作られたあらゆる
新規のインスタンスに設定されます。

## 設定 <!-- Configuration -->
<!--
As profiles aren't specific to containers or virtual machines, they may
contain configuration and devices that are valid for either type.
-->
プロファイルはコンテナや仮想マシンに固有なものではないため、どちらのインスタンスタイプでも有効な設定やデバイスを含めることができます。

<!--
This differs from the behavior when applying those config/devices
directly to an instance where its type is then taken into consideration
and keys that aren't allowed result in an error.
-->
これはこれらの設定やデバイスをインスタンスに直接適用するときの挙動とは異なります。
その場合はインスタンスタイプが考慮され、許可されないキーはエラーになります。

<!--
See [instance configuration](instances.md) for valid configuration options.
-->
有効な設定のオプションについては [インスタンス設定](instances.md) を
参照してください。

# プロファイル <!-- Profiles -->
<!--
Profiles can store any configuration that a container can (key/value or
devices) and any number of profiles can be applied to a container.
-->
プロファイルはコンテナが保持できる（キー・バリューやデバイスなどの）あらゆる
設定を保持することができ、プロファイルをいくつでもコンテナに適用することが
できます。

<!--
Profiles are applied in the order they are specified so the last profile to
specify a specific key wins.
-->
プロファイルは指定された順番に適用され、その結果最後に指定したプロファイルが
特定のキーを上書きします。

<!--
In any case, resource-specific configuration always overrides that coming from
the profiles.
-->
どのような場合でも、リソース固有の設定はプロファイル由来のものを上書きします。

<!--
If not present, LXD will create a `default` profile.
-->
まだ存在していない場合は、LXDは `default` プロファイルを作成します。

<!--
The `default` profile cannot be renamed or removed.
-->
`default` プロファイルはリネームや削除はできません。

<!--
The `default` profile is set for any new container created which doesn't
specify a different profiles list.
-->
`default` プロファイルは異なるプロファイルリストを指定せずに作られたあらゆる
新規のコンテナに設定されます。

<!--
See [container configuration](containers.md) for valid configuration options.
-->
有効な設定のオプションについては [コンテナ設定](containers.md) を
参照してください。

# プロジェクト設定
<!-- Project configuration -->
LXD はあなたの LXD サーバを分割する方法としてプロジェクトをサポートしています。
それぞれのプロジェクトはプロジェクトに固有なインスタンスのセットを持ち、また
プロジェクトに固有なイメージやプロファイルを持つこともできます。
<!--
LXD supports projects as a way to split your LXD server.
Each project holds its own set of instances and may also have its own images and profiles.
-->

プロジェクトに何を含めるかは `features` 設定キーによって決められます。
機能が無効の場合はプロジェクトは `default` プロジェクトから継承します。
<!--
What a project contains is defined through the `features` configuration keys.
When a feature is disabled, the project inherits from the `default` project.
-->

デフォルトでは全ての新規プロジェクトは全体のフィーチャーセットを取得し、
アップグレード時には既存のプロジェクトは新規のフィーチャーが有効には
なりません。
<!--
By default all new projects get the entire feature set, on upgrade,
existing projects do not get new features enabled.
-->

key/value 設定は現在サポートされている以下のネームスペースによって
名前空間が分けられています。
<!--
The key/value configuration is namespaced with the following namespaces
currently supported:
-->

 - `features` プロジェクトのフィーチャーセットのどの部分が使用中か <!-- (What part of the project featureset is in use) -->
 - `limits` プロジェクトに属するコンテナーと VM に適用されるリソース制限 <!-- (Resource limits applied on containers and VMs belonging to the project) -->
 - `user` ユーザメタデータに対する自由形式の key/value <!-- (free form key/value for user metadata) -->

キー <!-- Key -->               | 型 <!-- Type --> | 条件 <!-- Condition -->             | デフォルト値 <!-- Default -->                   | 説明 <!-- Description -->
:--                             | :--       | :--                   | :--                       | :--
features.images                 | boolean   | -                     | true                      | プロジェクト用のイメージとイメージエイリアスのセットを分離する <!-- Separate set of images and image aliases for the project -->
features.profiles               | boolean   | -                     | true                      | プロジェクト用のプロファイルを分離する <!-- Separate set of profiles for the project -->
limits.containers               | integer   | -                     | -                         | プロジェクト内に作成可能なコンテナーの最大数 <!-- Maximum number of containers that can be created in the project -->
limits.virtual-machines         | integer   | -                     | -                         | プロジェクト内に作成可能な VM の最大数 <!-- Maximum number of VMs that can be created in the project -->
limits.cpu                      | integer   | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.cpu" 設定の合計の最大値 <!-- Maximum value for the sum of individual "limits.cpu" configs set on the instances of the project -->
limits.memory                   | integer   | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.memory" 設定の合計の最大値 <!-- Maximum value for the sum of individual "limits.memory" configs set on the instances of the project -->
limits.processes                | integer   | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.processes" 設定の合計の最大値 <!-- Maximum value for the sum of individual "limits.processes" configs set on the instances of the project -->

これらのキーは lxc ツールを使って以下のように設定できます。
<!--
Those keys can be set using the lxc tool with:
-->

```bash
lxc project set <project> <key> <value>
```
## プロジェクトの制限 <!-- Project limits -->

注意: `limits.*` 設定キーの 1 つを設定する際はプロジェクト内の **全ての** インスタンスに直接あるいはプロファイル経由で同じ設定キーを設定 **する必要があります**。
<!--
Note that to be able to set one of the `limits.*` config keys, **all** instances
in the project **must** have that same config key defined, either directly or
via a profile.
-->

それに加えて
<!--
In addition to that:
-->

- `limits.cpu` 設定キーを使うにはさらに CPU ピンニングが使用されて **いない** 必要があります。 <!-- The `limits.cpu` config key also requires that CPU pinning is **not** used. -->
- `limits.memory` 設定キーはパーセント **ではなく** 絶対値で設定する必要があります。 <!-- The `limits.memory` config key must be set to an absolute value, **not** a percentage. -->

プロジェクトに設定された `limits.*` 設定キーは直接あるいはプロファイル経由でプロジェクト内のインスタンスに設定した個々の `limits.*` 設定キーの値の **合計値** に対しての hard な上限として振る舞います。
<!--
The `limits.*` config keys defined on a project act as a hard upper bound for
the **aggregate** value of the individual `limits.*` config keys defined on the
project's instances, either directly or via profiles.
-->

例えば、プロジェクトの `limits.memory` 設定キーを `50GB` に設定すると、プロジェクト内のインスタンスに設定された全ての `limits.memory` 設定キーの個々の値の合計が `50GB` 以下に維持されることを意味します。
インスタンスの作成あるいは変更時に `limits.memory` の値を全体の合計が `50GB` を超えるように設定しようとするとエラーになります。
<!--
For example, setting the project's `limits.memory` config key to `50GB` means
that the sum of the individual values of all `limits.memory` config keys defined
on the project's instances will be kept under `50GB`. Trying to create or modify
an instance assigning it a `limits.memory` value that would make the total sum
exceed `50GB`, will result in an error.
-->

同様にプロジェクトの `limits.cpu` 設定キーを `100` に設定すると、個々の `limits.cpu` の値の **合計** が `100` 以下に維持されることを意味します。
<!--
Similarly, setting the project's `limits.cpu` config key to `100`, means that
the **sum** of individual `limits.cpu` values will be kept below `100`.
-->

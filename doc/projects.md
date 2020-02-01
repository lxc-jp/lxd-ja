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
 - `user` ユーザメタデータに対する自由形式の key/value <!-- (free form key/value for user metadata) -->

キー <!-- Key -->               | 型 <!-- Type --> | 条件 <!-- Condition -->             | デフォルト値 <!-- Default -->                   | 説明 <!-- Description -->
:--                             | :--       | :--                   | :--                       | :--
features.images                 | boolean   | -                     | true                      | プロジェクト用のイメージとイメージエイリアスのセットを分離する <!-- Separate set of images and image aliases for the project -->
features.profiles               | boolean   | -                     | true                      | プロジェクト用のプロファイルを分離する <!-- Separate set of profiles for the project -->


これらのキーは lxc ツールを使って以下のように設定できます。
<!--
Those keys can be set using the lxc tool with:
-->

```bash
lxc project set <project> <key> <value>
```

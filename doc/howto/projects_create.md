(projects-create)=
# プロジェクトの作成と設定方法

プロジェクトは作成時または後で設定することができます。
ただし、プロジェクトにインスタンスが含まれている場合、有効になっている機能を変更することはできません。

## プロジェクトを作成する

プロジェクトを作成するには、`lxc project create`コマンドを使用します。

`--config`フラグを使用して設定オプションを指定できます。
利用可能な設定オプションについては、{ref}`ref-projects`を参照してください。

例えば、インスタンスを分離し、デフォルトプロジェクトのイメージとプロファイルにアクセスを許可する`my-project`というプロジェクトを作成するには、次のコマンドを入力します：

    lxc project create my-project --config features.images=false --config features.profiles=false

セキュリティに関する機能（例えば、コンテナのネスト）へのアクセスをブロックし、バックアップを許可する`my-restricted-project`というプロジェクトを作成するには、次のコマンドを入力します：

    lxc project create my-restricted-project --config restricted=true --config restricted.backups=allow

(projects-configure)=
## プロジェクトの設定
プロジェクトを設定するには、特定の設定オプションを設定するか、プロジェクト全体を編集できます。

いくつかの設定オプションは、インスタンスが含まれていないプロジェクトに対してのみ設定できます。

## 特定の設定オプションを設定する

特定の設定オプションを設定するには、`lxc project set`コマンドを使用します。

例えば、`my-project`で作成できるコンテナの数を5つに制限するには、次のコマンドを入力します：

    lxc project set my-project limits.containers=5

特定の設定オプションを解除するには、`lxc project unset`コマンドを使用します。

```{note}
設定オプションを解除すると、デフォルト値に設定されます。
このデフォルト値は、プロジェクトが作成されたときに設定される初期値と異なる場合があります。
```

### プロジェクトを編集する

プロジェクトの設定全体を編集するには、`lxc project edit`コマンドを使用します。

例：

    lxc project edit my-project

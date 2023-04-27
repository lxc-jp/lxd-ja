(projects-work)=
# 異なるプロジェクトで作業する方法

`default`プロファイル以外にもプロジェクトがある場合、LXDで作業する際に正しいプロジェクトを使用するか、対象となるプロジェクトを確認する必要があります。

```{note}
{ref}`特定のユーザーに制限されたプロジェクト <projects-confined>`がある場合、すべてのプロジェクトを表示できるのは、LXDへのフルアクセス権を持つユーザーのみです。

フルアクセス権を持たないユーザーは、アクセス権があるプロジェクトの情報のみを表示できます。
```

## プロジェクトの一覧表示

すべてのプロジェクト（閲覧許可があるもの）を一覧表示するには、次のコマンドを入力します：

    lxc project list

デフォルトでは、出力はリスト形式で表示されます：

```{terminal}
:input: lxc project list
:scroll:

+----------------------+--------+----------+-----------------+-----------------+----------+---------------+---------------------+---------+
|      NAME            | IMAGES | PROFILES | STORAGE VOLUMES | STORAGE BUCKETS | NETWORKS | NETWORK ZONES |     DESCRIPTION     | USED BY |
+----------------------+--------+----------+-----------------+-----------------+----------+---------------+---------------------+---------+
| default              | YES    | YES      | YES             | YES             | YES      | YES           | Default LXD project | 19      |
+----------------------+--------+----------+-----------------+-----------------+----------+---------------+---------------------+---------+
| my-project (current) | YES    | NO       | NO              | NO              | YES      | YES           |                     | 0       |
+----------------------+--------+----------+-----------------+-----------------+----------+---------------+---------------------+---------+
```

異なる出力形式を要求するには、`--format`フラグを追加します。
詳細については、`lxc project list --help`を参照してください。

## プロジェクトの切り替え

デフォルトでは、LXDで実行するすべてのコマンドは、現在使用しているプロジェクトに影響します。
どのプロジェクトを使用しているかを確認するには、`lxc project list`コマンドを使用します。

別のプロジェクトに切り替えるには、次のコマンドを入力します：

    lxc project switch <project_name>

## プロジェクトをターゲットにする

別のプロジェクトに切り替える代わりに、コマンドを実行する際に特定のプロジェクトをターゲットにすることができます。
多くのLXDコマンドは、`--project`フラグをサポートしており、異なるプロジェクトでアクションを実行できます。

```{note}
許可があるプロジェクトだけをターゲットにすることができます。
```

以下のセクションでは、プロジェクトを切り替える代わりにターゲットにする典型的な例をいくつか紹介します。

### 特定のプロジェクト内のインスタンスをリストする

特定のプロジェクト内のインスタンスをリストするには、`lxc list`コマンドに`--project`フラグを追加します。

例：

    lxc list --project my-project

### インスタンスを別のプロジェクトに移動する

インスタンスを1つのプロジェクトから別のプロジェクトに移動するには、次のコマンドを入力します：

    lxc move <instance_name> <new_instance_name> --project <source_project> --target-project <target_project>

ターゲットプロジェクトにその名前のインスタンスが存在しない場合、同じインスタンス名を維持できます。

例えば、インスタンス`my-instance`を`default`プロジェクトから`my-project`に移動し、インスタンス名を維持するには、次のコマンドを入力します：

    lxc move my-instance my-instance --project default --target-project my-project

### プロファイルを別のプロジェクトにコピーする

デフォルトの設定でプロジェクトを作成すると、プロファイルはプロジェクト内で隔離されます（[`features.profiles`](project-features) が `true` に設定されています）。
そのため、プロジェクトはデフォルトのプロファイル（`default`プロジェクトの一部）にアクセスできず、インスタンスを作成しようとすると次のようなエラーが表示されます。

```{terminal}
:input: lxc launch images:ubuntu/22.04 my-instance

Creating my-instance
Error: Failed instance creation: Failed creating instance record: Failed initialising instance: Failed getting root disk: No root device could be found
```

これを修正するには、`default`プロジェクトのデフォルトプロファイルの内容を現在のプロジェクトのデフォルトプロファイルにコピーします。
そのためには、次のコマンドを入力してください：

    lxc profile show default --project default | lxc profile edit default

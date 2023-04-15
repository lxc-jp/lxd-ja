(projects-confine)=
# 特定のユーザーにプロジェクトを制限する方法

プロジェクトを使用して、異なるユーザーまたはクライアントの活動を制限できます。
詳細については、{ref}`projects-confined`を参照してください。

特定のユーザーにプロジェクトを制限する方法は、選択した認証方法によって異なります。

## 特定のTLSクライアントにプロジェクトを制限する

```{youtube} https://www.youtube.com/watch?v=4iNpiL-lrXU&t=525s
```

LXDサーバーへの接続に使用されるTLSクライアント証明書を制限することで、特定のプロジェクトへのアクセスを制限できます。
詳細については、{ref}`authentication-tls-certs`を参照してください。

クライアント証明書が追加された時点からアクセスを制限するには、トークン認証を使用するか、クライアント証明書をサーバーに直接追加する必要があります。
パスワード認証を使用する場合、クライアント証明書が追加された後にのみ制限できます。

制限されたクライアント証明書を追加するには、次のコマンドを使用します:

````{tabs}

```{group-tab} トークン認証

    lxc config trust add --projects <project_name> --restricted

```

```{group-tab} クライアント証明書を追加

    lxc config trust add <certificate_file> --projects <project_name> --restricted
```

````

クライアントは、通常の方法でサーバーをリモートに追加できます（`lxc remote add <server_name> <token>` または `lxc remote add <server_name> <server_address>`）。そして、指定されたプロジェクトのみにアクセスできます。

既存の証明書のアクセスを制限するには（アクセス制限が変更されるか、証明書が信頼パスワードで追加されたため）、次のコマンドを使用します：

    lxc config trust edit <fingerprint>

`restricted`が`true`に設定されていることを確認し、`projects`の下に証明書がアクセスを許可するプロジェクトを指定してください。

```{note}
リモートを追加するときに`--project`フラグを指定できます。
この設定では、指定されたプロジェクトが事前に選択されます。
ただし、これによってクライアントがこのプロジェクトに制限されるわけではありません。
```

## 特定のRBACロールにプロジェクトを制限する

```{youtube} https://www.youtube.com/watch?v=VE60AbJHT6E
```

Canonical RBACサービスを使用している場合、RBACロールはそのロールを持つユーザーが実行できる操作を定義します。
詳細については、{ref}`authentication-rbac`を参照してください。

RBACを使用してプロジェクトを制限するには、RBACインターフェイスで対象のプロジェクトに移動し、必要に応じて異なるユーザーやグループにRBACロールを割り当てます。

## 特定のLXDユーザーにプロジェクトを制限する

```{youtube} https://www.youtube.com/watch?v=6O0q3rSWr8A
```

[LXD snap](https://snapcraft.io/lxd)を使用する場合、snapに含まれるマルチユーザーLXDデーモンを設定して、特定のユーザーグループ内のすべてのユーザーのために動的にプロジェクトを作成できます。

そのためには、`daemon.user.group`設定オプションを対応するユーザーグループに設定します：

    sudo snap set lxd daemon.user.group=<user_group>

LXDを使用できるようにしたいすべてのユーザーアカウントがこのグループのメンバーであることを確認してください。

グループのメンバーがLXDコマンドを発行すると、LXDはこのユーザーのために制限されたプロジェクトを作成し、このプロジェクトに切り替えます。
この時点でLXDが{ref}`初期化 <initialize>`されていない場合、自動的に初期化されます（デフォルト設定で）。

プロジェクトの設定をカスタマイズしたい場合（例えば、制限や制約を課すために）、プロジェクトが作成された後で行うことができます。
プロジェクト設定を変更するには、LXDへの完全なアクセスが必要であり、つまり`lxd`グループの一部であり、設定したLXDユーザーグループの一部であるだけではありません。

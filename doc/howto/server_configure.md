(server-configure)=
# LXDサーバーを設定するには

LXDサーバーで利用可能な全て設定オプションについては{ref}`server`を参照してください。

LXDサーバーがクラスタの一部の場合、一部のオプションはクラスタに適用され、また別のオプションはローカルサーバー、つまりクラスタメンバーにのみ適用されます。
{ref}`server`オプションの表で、クラスタに適用されるオプションは`global`スコープと表記され、ローカルサーバーのみに適用されるオプションは`local`スコープと表記されます。

## サーバーオプションを設定する

以下のコマンドでサーバーオプションを設定できます。

    lxc config set <key> <value>

例えば、ポート8443でLXDサーバーにリモートからのアクセスを許可するには、以下のコマンドを入力します。

    lxc config set core.https_address :8443

クラスタ構成では、クラスタメンバーだけにサーバー設定を行うには`--target`フラグを追加してください。
例えば、特定のクラスタメンバーでイメージのtarballを保管する場所を設定するには、以下のようなコマンドを入力してください。

    lxc config set storage.images_volume my-pool/my-volume --target member02

## サーバー設定を表示する

現在のサーバー設定を表示するには、以下のコマンドを入力します。

    lxc config show

クラスタ構成では、クラスタメンバーだけにサーバー設定を行うには`--target`フラグを追加してください。

## サーバー設定全体を編集する

サーバー設定全体をYAMLファイルとして編集するには、以下のコマンドを入力します。

    lxc config edit

クラスタ構成では、クラスタメンバーだけにサーバー設定を行うには`--target`フラグを追加してください。

(lxc-alias)=
# コマンドエイリアスを追加するには

LXDコマンドラインクライアントでは良く使うコマンドのエイリアスを追加できます。
長いコマンドのショートカットとして、あるいは既存のコマンドに自動的にフラグを追加するために、エイリアスを使用できます。

コマンドエイリアスを管理するには、`lxc alias`コマンドを使用します。

例えば、インスタンスを削除する際に必ず確認を求めるようにするには`lxc delete`に常に`lxc delete -i`を実行するようにエイリアスを作成します。

    lxc alias add delete "delete -i"

登録された全てののエイリアスを表示するには`lxc alias list`を実行します。
全ての利用可能なサブコマンドを表示するには`lxc alias --help`を実行してください。

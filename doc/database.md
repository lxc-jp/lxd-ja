# データベース <!-- Database -->
# イントロダクション <!-- Introduction -->
<!--
So first of all, why a database?
-->
そもそも、なぜデータベースなのでしょう？

<!--
Rather than keeping the configuration and state within each container's
directory as is traditionally done by LXC, LXD has an internal database
which stores all of that information. This allows very quick queries
against all containers configuration.
-->
従来 LXC で行われていたように設定と状態をそれぞれのコンテナのディレクトリに
保存するのではなく、 LXD ではそれら全ての情報を保管する内部的なデータベースを
持っています。これによりすべてのコンテナの設定に対する問い合わせをとても
高速に行えます。


<!--
An example is the rather obvious question "what containers are using br0?".
To answer that question without a database, LXD would have to iterate
through every single container, load and parse its configuration and
then look at what network devices are defined in there.
-->
例えば、 「どのコンテナが br0 を使っているのか」というかなり分かりやすい問いが
あります。この問いにデータベース無しで答えるとすると、 LXD は一つ一つの
コンテナに対して、設定を読み込んでパースし、そこにどのネットワークデバイスが
定義されているかを見るということを繰り返し行わなければなりません。

<!--
While that may be quick with a few containers, imagine how many
filesystem access would be required for 2000 containers. Instead with a
database, it's only a matter of accessing the already cached database
with a pretty simple query.
-->
コンテナの数が少なければ、その処理は速いかもしれませんが、2000 個のコンテナに
対してどれだけ多くのファイルシステムへのアクセスが必要かを想像してみてください。
代わりにデータベースを使うことで、非常に単純なクエリでキャッシュ済みの
データベースにアクセスするだけで良くなるのです。


# データベースエンジン <!-- Database engine -->
<!--
Since LXD supports clustering, and all members of the cluster must share the
same database state, the database engine is based on a [distributed
version](https://github.com/CanonicalLtd/dqlite) of SQLite, which provides
replication, fault-tolerance and automatic failover without the need of external
database processes. We refer to this database as the "global" LXD database.
-->
LXD はクラスタリングをサポートし、クラスタの全てのメンバは同じデータベースの
状態を共有する必要があるため、データベースエンジンは SQLite の
[分散対応バージョン](https://github.com/CanonicalLtd/dqlite) をベースにしています。
それは外部のデータベースのプロセスを必要とせずに、 レプリケーション、
フォールトトレランス、自動フェールオーバーの機能を提供します。
このデータベースを「グローバル」 LXD データベースと呼びます。


<!--
Even when using LXD as single non-clustered node, the global database will still
be used, although in that case it effectively behaves like a regular SQLite
database.
-->
単一の非クラスターノードとして LXD を使う場合であっても、やはりグローバル
データベースを使用します。ただし、その場合は実質的には通常の SQLite
データベースとして振る舞います。

<!--
The files of the global database are stored under the ``./database/global``
sub-directory of your LXD data dir (e.g. ``/var/lib/lxd/database/global``).
-->
グローバルデータベースのファイルは LXD のデータディレクトリ
(例 ``/var/lib/lxd/database/global``) の ``./database/global`` サブディレクトリ
の下に格納されます。

<!--
Since each member of the cluster also needs to keep some data which is specific
to that member, LXD also uses a plain SQLite database (the "local" database),
which you can find in ``./database/local.db``.
-->
クラスタの各メンバもそのメンバ固有の何らかのデータを保持する必要があるため、
LXD は単なる SQLite のデータベース (「ローカル」データベース」も使用します。
これは ``./database/local.db`` に置かれます。

<!--
Backups of the global database directory and of the local database file are made
before upgrades, and are tagged with the ``.bak`` suffix. You can use those if
you need to revert the state as it was before the upgrade.
-->
アップグレードの前にはグローバルデータベースのディレクトリとローカルデータベースの
ファイルのバックアップが作成され、 ``.bak`` のサフィックス付きでタグ付けされます。
アップグレード前の状態に戻す必要がある場合は、このバックアップを使うことができます。

# データベースのデータとスキーマをダンプする <!-- Dumping the database content or schema -->
<!--
If you want to get a SQL text dump of the content or the schema of the databases,
use the ``lxd sql <local|global> [.dump|.schema]`` command, which produces the
equivalent output of the ``.dump`` or ``.schema`` directives of the sqlite3
command line tool.
-->
データベースのデータまたはスキーマの SQL テキスト形式でのダンプを取得したい場合は、
``lxd sql <local|global> [.dump|.schema]`` コマンドを使ってください。これにより
sqlite3 コマンドラインツールの ``.dump`` または ``.schema`` ディレクティブと同じ出力を
生成できます。

# コンソールからカスタムクエリを実行する <!-- Running custom queries from the console -->
<!--
If you need to perform SQL queries (e.g. ``SELECT``, ``INSERT``, ``UPDATE``)
against the local or global database, you can use the ``lxd sql`` command (run
``lxd sql \-\-help`` for details).
-->
ローカルまたはグローバルデータベースに SQL クエリ (例 ``SELECT``, ``INSERT``, ``UPDATE``) を
実行する必要がある場合、 ``lxd sql`` コマンドを使うことができます
(詳細は ``lxd sql --help`` を実行してください)。

<!--
You should only need to do that in order to recover from broken updates or bugs.
Please consult the LXD team first (creating a [GitHub
issue](https://github.com/lxc/lxd/issues/new) or
[forum](https://discuss.linuxcontainers.org/) post).
-->
ただ、これが必要になるのは壊れたアップデートかバグからリカバーするときだけでしょう。
その場合、まず LXD チームに相談してみてください (
[GitHubのイシュー](https://github.com/lxc/lxd/issues/new) を作成するか
[フォーラム](https://discuss.linuxcontainers.org/) に投稿)。

# LXD デーモン起動時にカスタムクエリを実行する <!-- Running custom queries at LXD daemon startup -->
<!--
In case the LXD daemon fails to start after an upgrade because of SQL data
migration bugs or similar problems, it's possible to recover the situation by
creating ``.sql`` files containing queries that repair the broken update.
-->
SQL のデータマイグレーションのバグあるいは関連する問題のために
アップグレード後に LXD デーモンが起動に失敗する場合、
壊れたアップデートを修復するクエリを含んだ ``.sql`` ファイルを
作成することで、その状況からリカバーできます。

<!--
To perform repairs against the local database, write a
``./database/patch.local.sql`` file containing the relevant queries, and
similarly a ``./database/patch.global.sql`` for global database repairs.
-->
ローカルデータベースに対して修復を実行するには、修復に必要なクエリを含む
``./database/patch.local.sql`` というファイルを作成してください。
同様にグローバルデータベースの修復には ``./database/patch.global.sql`` という
ファイルを作成してください。

<!--
Those files will be loaded very early in the daemon startup sequence and deleted
if the queries were successful (if they fail, no state will change as they are
run in a SQL transaction).
-->
これらのファイルはデーモンの起動シーケンスの非常に早い段階で読み込まれ、
クエリが成功したときは削除されます (クエリは SQL トランザクション内で実行されるので、
クエリが失敗したときにデータベースの状態が変更されることはありません)。

<!--
As above, please consult the LXD team first.
-->
上記の通り、まず LXD チームに相談してみてください。

# クラスタデータベースをディスクに同期 <!-- Syncing the cluster database to disk -->
クラスタデータベースの内容をディスクにフラッシュしたいなら、
``lxd sql global .sync`` コマンドを使ってください。これは SQLite そのままの
形式のデータベースのファイルを ``./database/global/db.bin`` に書き込みます。
その後 ``sqlite3`` コマンドラインツールを使って中身を見ることが出来ます。
<!--
If you want to flush the content of the cluster database to disk, use the ``lxd
sql global .sync`` command, that will write a plain SQLite database file into
``./database/global/db.bin``, which you can then inspect with the ``sqlite3``
command line tool.
-->

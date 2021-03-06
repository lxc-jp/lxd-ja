# デーモンの動作
<!-- Daemon behavior -->
## イントロダクション <!-- Introduction -->

<!--
This specification covers some of the daemon's behavior, such as
reaction to given signals, crashes, ...
-->
この仕様書は特定のシグナルに対する反応やクラッシュなどのデーモンの
振る舞いの一部を取り扱います。

## 起動 <!-- Startup -->
<!--
On every start, LXD checks that its directory structure exists. If it
doesn't, it'll create the required directories, generate a keypair and
initialize the database.
-->
起動する度に LXD はディレクトリ構造が存在することをチェックします。
もし存在しない場合は、必要なディレクトリを作成し、キーペアを生成し、
データベースを初期化します。

<!--
Once the daemon is ready for work, LXD will scan the instances table
for any instance for which the stored power state differs from the
current one. If an instance's power state was recorded as running and the
instance isn't running, LXD will start it.
-->
ひとたびデーモンが動作の準備が出来ると、 LXD はデータベース内の
インスタンスのテーブルから対象のテーブルを検索し、電源状態が実際の状態と
異なっていないかを確認します。もしインスタンスの電源状態が稼働中と記録
されているのにインスタンスが稼働していない場合は LXD はそのインスタンスを
開始します。

## シグナル処理 <!-- Signal handling -->
### SIGINT, SIGQUIT, SIGTERM
<!--
For those signals, LXD assumes that it's being temporarily stopped and
will be restarted at a later time to continue handling the instances.
-->
これらのシグナルについては LXD は一時的に停止し、後に再開してインスタンスの
処理を継続することを想定しています。

<!--
The instances will keep running and LXD will close all connections and
exit cleanly.
-->
インスタンスは稼働し続けて LXD は全ての接続を閉じ、クリーンな状態で終了する
でしょう。

### SIGPWR
<!--
Indicates to LXD that the host is going down.
-->
LXD にホストがシャットダウンしようとしていることを伝えます。

<!--
LXD will attempt a clean shutdown of all the instances. After 30s, it
will kill any remaining instance.
-->
LXD は全てのインスタンスをクリーンにシャットダウンしようと試みます。30秒後、
LXD は残りのインスタンスを kill します。

<!--
The instance `power_state` in the instances table is kept as it was so
that LXD after the host is done rebooting can restore the instances as
they were.
-->
ホストがリブートを完了後に LXD がインスタンスを元の状態に戻せるように、
データベース内のインスタンスのテーブルの `power_state` カラムにインスタンスの
元の電源状態を記録しておきます。

### SIGUSR1
<!--
Write a memory profile dump to the file specified with `\-\-memprofile`.
-->
メモリプロファイルを `--memprofile` で指定したファイルにダンプします。

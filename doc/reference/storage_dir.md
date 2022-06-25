(storage-dir)=
# ディレクトリ - `dir`

 - このバックエンドでは全ての機能を使えますが、他のバックエンドに比べて非常に時間がかかります。
   これは、イメージを展開したり、インスタンスやスナップショットやイメージのその時点のコピーを作成する必要があるからです。
 - ファイルシステムレベルでプロジェクトクォータが有効に設定されている ext4 もしくは XFS で実行している場合は、ディレクトリバックエンドでクォータがサポートされます。


## ストレージプール設定
キー              | 型     | デフォルト値 | 説明
:--               | :---   | :------      | :----------
rsync.bwlimit     | string | 0 (no limit) | ストレージエンティティの転送に rsync を使う必要があるときにソケット I/O に指定する上限を設定
rsync.compression | bool   | true         | ストレージブールのマイグレーションの際に圧縮を使うかどうか
source            | string | -            | ブロックデバイスかループファイルかファイルシステムエントリのパス

## ストレージボリューム設定
キー               | 型     | 条件               | デフォルト値       | 説明
:--                | :---   | :--------          | :------            | :----------
security.shifted   | bool   | custom volume      | false              | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped  | bool   | custom volume      | false              | ボリュームへの id マッピングを無効にする
size               | string | appropriate driver | volume.size と同じ | ストレージボリュームのサイズ
snapshots.expiry   | string | custom volume      | -                  | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern  | string | custom volume      | snap%d             | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule | string | custom volume      | -                  | {{snapshot_schedule_format}}

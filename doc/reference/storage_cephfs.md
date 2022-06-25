(storage-cephfs)=
# CephFS - `cephfs`

 - カスタムストレージボリュームにのみ利用可能
 - サーバサイドで許可されていればスナップショットもサポート

## ストレージプール設定
キー                   | 型     | デフォルト値 | 説明
:--                    | :---   | :------      | :----------
cephfs.cluster\_name   | string | ceph         | 新しいストレージプールを作成する Ceph クラスタの名前
cephfs.fscache         | bool   | false        | カーネルの fscache と cachefilesd を有効にするか
cephfs.path            | string | /            | CephFS をマウントするベースのパス
cephfs.user.name       | string | admin        | ストレージプールとボリュームを作成する際に用いる Ceph のユーザー
source                 | string | -            | 使用する既存のストレージプールかストレージプール内のパス
volatile.pool.pristine | string | true         | プールが作成時に空かどうか

## ストレージボリューム設定
キー               | 型     | 条件               | デフォルト値       | 説明
:--                | :---   | :--------          | :------            | :----------
security.shifted   | bool   | custom volume      | false              | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped  | bool   | custom volume      | false              | ボリュームへの id マッピングを無効にする
size               | string | appropriate driver | volume.size と同じ | ストレージボリュームのサイズ
snapshots.expiry   | string | custom volume      | -                  | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern  | string | custom volume      | snap%d             | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule | string | custom volume      | -                  | {{snapshot_schedule_format}}

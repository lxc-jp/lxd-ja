(storage-lvm)=
# LVM - `lvm`

 - イメージ用に LV を使うと、インスタンスとインスタンススナップショット用に LV のスナップショットを使います
 - LV で使われるファイルシステムは ext4 です（代わりに xfs を使うように設定できます）
 - デフォルトでは、すべての LVM ストレージプールは LVM thin pool を使います。すべての LXD ストレージエンティティ（イメージやインスタンスなど）のための論理ボリュームは、その LVM thin pool 内に作られます。
   この動作は、`lvm.use_thinpool` を "false" に設定して変更できます。
   この場合、LXD はインスタンススナップショットではないすべてのストレージエンティティ（イメージやインスタンスなど）に、通常の論理ボリュームを使います。
   Thinpool 以外の論理ボリュームは、スナップショットのスナップショットをサポートしていないので、ほとんどのストレージ操作を rsync にフォールバックする必要があります。
   これは、LVM ドライバがスピードとストレージ操作の両面で DIR ドライバに近づくため、必然的にパフォーマンスに重大な影響を与えることに注意してください。
   このオプションは、必要な場合のみに選択してください。
 - 頻繁にインスタンスとのやりとりが発生する環境（例えば継続的インテグレーション）では、`/etc/lvm/lvm.conf` 内の `retain_min` と `retain_days` を調整して、LXD とのやりとりが遅くならないようにすることが重要です。

## ストレージプール設定
キー                          | 型     | デフォルト値     | 説明
:--                           | :---   | :------          | :----------
lvm.thinpool\_name            | string | LXDThinPool      | イメージを作る thin pool 名
lvm.thinpool\_metadata\_size  | string | 0 (auto)         | thin pool メタデータボリュームのサイズ。デフォルトは LVM が適切なサイズを計算
lvm.use\_thinpool             | bool   | true             | ストレージプールは論理ボリュームに thin pool を使うかどうか
lvm.vg.force\_reuse           | bool   | false            | 既存の空でないボリュームグループの使用を強制
lvm.vg\_name                  | string | name of the pool | 作成するボリュームグループ名
rsync.bwlimit                 | string | 0 (no limit)     | ストレージエンティティーの転送にrsyncを使う場合、I/Oソケットに設定する上限を指定
rsync.compression             | bool   | true             | ストレージプールをマイグレートする際に圧縮を使用するかどうか
source                        | string | -                | ブロックデバイスかループファイルかファイルシステムエントリのパス

## ストレージボリューム設定
キー                 | 型     | 条件               | デフォルト値                       | 説明
:--                  | :---   | :--------          | :------                            | :----------
block.filesystem     | string | block based driver | volume.block.filesystem と同じ     | ストレージボリュームのファイルシステム
block.mount\_options | string | block based driver | volume.block.mount\_options と同じ | ブロックデバイスのマウントオプション
lvm.stripes          | string | lvm driver         | -                                  | 新しいボリューム (あるいは thinpool ボリューム) に使用するストライプ数
lvm.stripes.size     | string | lvm driver         | -                                  | 使用するストライプのサイズ (最低 4096 バイトで 512 バイトの倍数を指定)
security.shifted     | bool   | custom volume      | false                              | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped    | bool   | custom volume      | false                              | ボリュームへの id マッピングを無効にする
size                 | string | appropriate driver | volume.size と同じ                 | ストレージボリュームのサイズ
snapshots.expiry     | string | custom volume      | -                                  | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern    | string | custom volume      | snap%d                             | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule   | string | custom volume      | -                                  | {{snapshot_schedule_format}}

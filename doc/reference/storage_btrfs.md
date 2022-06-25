(storage-btrfs)=
# Btrfs - `btrfs`

 - インスタンス、イメージ、スナップショットごとにサブボリュームを使い、新しいオブジェクトを作成する際に btrfs スナップショットを作成します
 - btrfs は、親コンテナ自身が btrfs 上に作成されているときには、コンテナ内のストレージバックエンドとして使えます（ネストコンテナ）（qgroup を使った btrfs クオータについての注意を参照してください）
 - btrfs では qgroup を使ったストレージクオータが使えます。btrfs qgroup は階層構造ですが、新しいサブボリュームは自動的には親のサブボリュームの qgroup には追加されません。
   このことは、ユーザーが設定されたクオータをエスケープできるということです。
   もし、クオータを厳格に遵守させたいときは、ユーザーはこのことに留意し、refquota を使った zfs ストレージを使うことを検討してください。

 - クオータを使用する際は btrfs のエクステントはイミュータブルであるためブロックが書かれるときにブロックが新しいエクステントに書き込まれ古いブロックはその中のデータが全て参照されなくなるか再書き込みされるまで残ることを考慮することが非常に重要です。
   これはサブボリューム内の現在のファイルが使用中のスペースの合計量がクオータより小さいにもかかわらずクオータに達することがあり得ることを意味します。
   これは btrfs サブボリュームの上に生のディスクイメージファイルを使うランダム I/O の性質のため BTRFS 上で VM を使うときによく発生します。
   VM と btrfs のストレージプールの組み合わせは使わないことを私達は推奨します。
   もしそれでも使いたい場合は、ディスクイメージファイル内の全てのブロックが qgroup クオータの制限にかかること無く再書き込みできるように
   インスタンスのルートディスクの `size.state` プロパティをルートディスクサイズの 2 倍に設定してください。
   また `btrfs.mount_options=compress-force` ストレージオプションを使うことで圧縮を有効にする副作用として最大のエクステントサイズを縮小させブロックの再書き込みによりストレージの大部分が 2 倍の容量を消費するのを防ぐことができます。
   ただしこれはストレージプールのオプションですので、プール上の全てのボリュームに影響します。

## ストレージプール設定
キー                 | 型     | デフォルト値              | 説明
:--                  | :---   | :--------                 | :----------
btrfs.mount\_options | string | user\_subvol\_rm\_allowed | ブロックデバイスのマウントオプション
source               | string | -                         | ブロックデバイスまたはループファイルまたはファイルシステムエントリーのパス

## ストレージボリューム設定
キー               | 型     | 条件               | デフォルト値       | 説明
:--                | :---   | :--------          | :------            | :----------
security.shifted   | bool   | custom volume      | false              | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped  | bool   | custom volume      | false              | ボリュームへの id マッピングを無効にする
size               | string | appropriate driver | volume.size と同じ | ストレージボリュームのサイズ
snapshots.expiry   | string | custom volume      | -                  | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern  | string | custom volume      | snap%d             | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule | string | custom volume      | -                  | {{snapshot_schedule_format}}

## ループバックデバイスを使った btrfs プールの拡張
LXD では、ループバックデバイスの btrfs プールを直接は拡張できませんが、次のように拡張できます:

```bash
sudo truncate -s +5G /var/lib/lxd/disks/<POOL>.img
sudo losetup -c <LOOPDEV>
sudo btrfs filesystem resize max /var/lib/lxd/storage-pools/<POOL>/
```

(注意: snap のユーザーは `/var/lib/lxd/` の代わりに `/var/snap/lxd/common/mntns/var/snap/lxd/common/lxd/` を使ってください)
- LOOPDEV はストレージプールイメージに関連付けられたマウントされたループデバイス（例: `/dev/loop8`）を参照します。
- マウントされたループデバイスは次のコマンドで確認できます。
```bash
losetup -l
```

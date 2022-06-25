---
discourse: 1333
---

(storage-zfs)=
# ZFS - `zfs`

 - LXD が ZFS プールを作成した場合は、デフォルトで圧縮が有効になります
 - イメージ用に ZFS を使うと、インスタンスとスナップショットの作成にスナップショットとクローンを使います
 - ZFS でコピーオンライトが動作するため、すべての子のファイルシステムがなくなるまで、親のファイルシステムを削除できません。
   ですので、削除されたけれども、まだ参照されているオブジェクトを、LXD はランダムな `deleted/` なパスに自動的にリネームし、参照がなくなりオブジェクトを安全に削除できるようになるまで、そのオブジェクトを保持します。
 - 現時点では、ZFS では、プールの一部をコンテナユーザーに権限委譲できません。開発元では、この問題に積極的に取り組んでいます。
 - ZFS では最新のスナップショット以外からのリストアはできません。
   しかし、古いスナップショットから新しいインスタンスを作成することはできます。
   これにより、新しいスナップショットを削除する前に、スナップショットが確実にリストアしたいものかどうか確認できます。

   LXD はリストア中に新しいスナップショットを自動的に破棄するように設定することもできます。
   これは `volume.zfs.remove_snapshots` プールオプションを使って設定可能です。

   しかしインスタンスのコピーも ZFS スナップショットを使うこと、その結果として全ての子孫も消すことなしには最後のコピーより前に取られたスナップショットにインスタンスをリストアすることもできないことに注意してください。

   必要なスナップショットを新しいインスタンスにコピーした後に古いインスタンスを削除できますが、インスタンスが持っているかもしれない他のスナップショットを失ってしまいます。

 - LXD は ZFS プールとデータセットがフルコントロールできると仮定していることに注意してください。
   LXD の ZFS プールやデータセット内に LXD と関係ないファイルシステムエンティティを維持しないことをおすすめします。LXD がそれらを消してしまう恐れがあるからです。
 - ZFS データセットでクオータを使った場合、LXD は ZFS の "quota" プロパティを設定します。
   LXD に "refquota" プロパティを設定させるには、与えられたデータセットに対して "zfs.use\_refquota" を "true" に設定するか、
   ストレージプール上で "volume.zfs.use\_refquota" を "true" に設定するかします。
   前者のオプションは、与えられたストレージプールだけに refquota を設定します。
   後者のオプションは、ストレージプール内のストレージボリュームすべてに refquota を使うようにします。
   また、ボリュームに"zfs.reserve\_space"、ストレージプールに"volume.zfs.reserve\_space"を設定することで、ZFSの"quota"/"refquota"に加えて"reservation"/"refreservation"を使用することができます。
 - I/O クオータ（IOps/MBs）は ZFS ファイルシステムにはあまり影響を及ぼさないでしょう。
   これは、ZFS が（SPL を使った）Solaris モジュールの移植であり、
   I/O に対する制限が適用される Linux の VFS API を使ったネイティブな Linux ファイルシステムではないからです。

## ストレージプール設定
キー            | 型     | デフォルト値 | 説明
:--             | :---   | :------      | :----------
size            | string | 0            | ストレージプールのサイズ。バイト単位（suffixも使えます）（現時点では loop ベースのプールと ZFS で有効）
source          | string | -            | ブロックデバイスかループファイルかファイルシステムエントリのパス
zfs.clone\_copy | string | true         | boolean の文字列を指定した場合は ZFS のフルデータセットコピーの代わりに軽量なクローンを使うかどうかを制御し、 "rebase" という文字列を指定した場合は初期イメージをベースにコピーします。
zfs.export      | bool   | true         | アンマウントの実行中にzpoolのエクスポートを無効にする
zfs.pool\_name  | string | プールの名前 | Zpool 名

## ストレージボリューム設定

```{rst-class} dec-font-size
```
キー                  | 型     | 条件               | デフォルト値                        | 説明
:--                   | :---   | :--------          | :------                             | :----------
security.shifted      | bool   | custom volume      | false                               | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped     | bool   | custom volume      | false                               | ボリュームへの id マッピングを無効にする
size                  | string | appropriate driver | volume.size と同じ                  | ストレージボリュームのサイズ
snapshots.expiry      | string | custom volume      | -                                   | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern     | string | custom volume      | snap%d                              | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule      | string    | custom volume             | -                                     | {{snapshot_schedule_format}}
zfs.blocksize         | string | zfs driver         | volume.zfs.blocksize と同じ         | ZFSブロックのサイズを512～16MiBの範囲で指定します（2の累乗でなければなりません）。ブロックボリュームでは、より大きな値が設定されていても、最大値の128KiBが使用されます。
zfs.remove\_snapshots | string | zfs driver         | volume.zfs.remove\_snapshots と同じ | 必要に応じてスナップショットを削除するかどうか
zfs.use\_refquota     | string | zfs driver         | volume.zfs.use\_refquota と同じ     | 領域の quota の代わりに refquota を使うかどうか
zfs.reserve\_space    | string | zfs driver         | false                               | qouta/refquota に加えて reservation/refreservation も使用するかどうか

## ループバックの ZFS プールの拡張
LXD からは直接はループバックの ZFS プールを拡張できません。しかし、次のようにすればできます:

```bash
sudo truncate -s +5G /var/lib/lxd/disks/<POOL>.img
sudo zpool set autoexpand=on <POOL>
sudo zpool status -vg <POOL> # デバイス ID をメモしておきます
sudo zpool online -e <POOL> <device_ID>
sudo zpool set autoexpand=off <POOL>
```

(注意: snap のユーザーは `/var/lib/lxd/` の代わりに `/var/snap/lxd/common/lxd/` を使ってください)

## 既存のプールで TRIM を有効にする
LXD は ZFS 0.8 以降で新規に作成された全てのプールに TRIM サポートを自動で有効にします。

これによりコントローラーによるブロック再利用を改善し SSD の寿命を延ばすことができます。
さらにループバックの ZFS プールを使用している場合はルートファイルシステムの空きスペースを解放できます。

0.8 より古い ZFS を 0.8 にアップグレードしたシステムでは、以下の 1 度きりの操作で TRIM の自動実行を有効にできます。

 - zpool upgrade ZPOOL-NAME
 - zpool set autotrim=on ZPOOL-NAME
 - zpool trim ZPOOL-NAME

これにより現在未使用のスペースに TRIM を実行するだけでなく、将来 TRIM が自動的に実行されるようになります。

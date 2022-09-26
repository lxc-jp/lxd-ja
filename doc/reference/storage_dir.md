(storage-dir)=
# ディレクトリ - `dir`

```{youtube} https://www.youtube.com/watch?v=imWkPM9GjCY
```

ディレクトリストレージドライバは基本的なバックエンドで通常のファイルとディレクトリ構造にデータを保管します。
このドライバは素早くセットアップできディスク上のファイルを直接見ることができるので、テストには便利かもしれません。
しかし、 LXD の操作はこのドライバ用には {ref}`最適化されていません <storage-drivers-features>`。

## LXD の `dir` ドライバ

LXD の `dir` ドライバは完全に機能し、他のドライバと同じ機能セットを提供します。
しかし、他のドライバよりは圧倒的に遅いです。これはインスタンス、スナップ、ショットを一瞬でコピーする代わりにイメージの解凍を行う必要があるためです。

作成時に (`source` 設定オプションを使って) 別途指定されてない限り、データは `/var/snap/lxd/common/lxd/storage-pools/` (snap でインストールした場合) または `/var/lib/lxd/storage-pools/` ディレクトリに保管されます。

(storage-dir-quotas)=
### クォータ

`dir` ドライバは ext4 か XFS 上で動作しファイルシステムレベルでプロジェクトのクォータが有効な場合にストレージのクォータをサポートします。

## 設定オプション

`dir` ドライバを使うストレージプールとこれらのプール内のストレージボリュームには以下の設定オプションが利用できます。

## ストレージプール設定

キー                | 型     | デフォルト値 | 説明
:--                 | :---   | :------      | :----------
`rsync.bwlimit`     | string | `0` (no limit) | ストレージエンティティの転送に rsync を使う必要があるときにソケット I/O に指定する上限を設定
`rsync.compression` | bool   | `true`         | ストレージブールのマイグレーションの際に圧縮を使うかどうか
`source`            | string | -            | ブロックデバイスかループファイルかファイルシステムエントリのパス

{{volume_configuration}}

## ストレージボリューム設定

キー                 | 型     | 条件               | デフォルト値                                 | 説明
:--                  | :---   | :--------          | :------                                      | :----------
`security.shifted`   | bool   | カスタムボリューム | `volume.security.shifted` と同じか `false`   | {{enable_ID_shifting}}
`security.unmapped`  | bool   | カスタムボリューム | `volume.security.unmapped` と同じか `false`  | ボリュームの ID マッピングを無効にする
`size`               | string | 適切なドライバ     | `volume.size` と同じ                         | ストレージボリュームのサイズ/クォータ
`snapshots.expiry`   | string | カスタムボリューム | `volume.snapshots.expiry` と同じ             | {{snapshot_expiry_format}}
`snapshots.pattern`  | string | カスタムボリューム | `volume.snapshots.pattern` と同じか `snap%d` | {{snapshot_pattern_format}}
`snapshots.schedule` | string | カスタムボリューム | `volume.snapshots.schedule` と同じ           | {{snapshot_schedule_format}}

### ストレージバケット設定

ローカルのストレージプールドライバでストレージバケットを有効にし、 S3 プロトコル経由でアプリケーションがバケットにアクセスできるようにするには `core.storage_buckets_address` サーバ設定 ({ref}`server` 参照) を調整する必要があります。

ストレージバケットは `dir` プール用の設定はありません。
他のストレージプールドライバとは異なり、 `dir` ドライバは `size` 設定によるバケットクォータのサポートはありません。

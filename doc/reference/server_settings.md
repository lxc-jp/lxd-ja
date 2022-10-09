(server-settings)=
# LXD プロダクション環境のサーバ設定

あなたの LXD サーバで多数のインタンスを稼働できるようにするには、サーバの制限値にひっかからないよう以下の設定を調整してください。

`値` のカラムに各パラメータの推奨値を記載しています。

## `/etc/security/limits.conf`

```{note}
snap ユーザーの場合はこれらの制限は自動的に上げられます。
```

ドメイン | 種別 | 項目      | 値          | デフォルト | 説明
:-----   | :--- | :----     | :--------   | :--------  | :----------
`*`      | soft | `nofile`  | `1048576`   | 未設定     | オープンするファイルの最大数
`*`      | hard | `nofile`  | `1048576`   | 未設定     | オープンするファイルの最大数
`root`   | soft | `nofile`  | `1048576`   | 未設定     | オープンするファイルの最大数
`root`   | hard | `nofile`  | `1048576`   | 未設定     | オープンするファイルの最大数
`*`      | soft | `memlock` | `unlimited` | 未設定     | ロックされたメモリ内の最大のアドレス空間 (KB)
`*`      | hard | `memlock` | `unlimited` | 未設定     | ロックされたメモリ内の最大のアドレス空間 (KB)
`root`   | soft | `memlock` | `unlimited` | 未設定     | ロックされたメモリ内の最大のアドレス空間 (KB) (`bpf` システムコール監視にのみ必要)
`root`   | hard | `memlock` | `unlimited` | 未設定     | ロックされたメモリ内の最大のアドレス空間 (KB) (`bpf` システムコール監視にのみ必要)

## `/etc/sysctl.conf`

```{note}
これらのパラメータを変更した後、サーバを再起動してください。
```

パラメータ                          | 値        | デフォルト | 説明
:-----                              | :---      | :---       | :---
`fs.aio-max-nr`                     | `524288`  | `65536`    | これは並行に実行される非同期 I/O 操作の最大数です。 AIO サブシステムを使うワークロードが大量にある場合（例: MySQL ）、これを増やす必要があるかもしれません
`fs.inotify.max_queued_events`      | `1048576` | `16384`    | これは対応する `inotify` のインスタンスにキューイングされるイベント数の上限を指定します ([`inotify`](https://man7.org/linux/man-pages/man7/inotify.7.html) 参照)
`fs.inotify.max_user_instances`     | `1048576` | `128`      | これは実ユーザー ID ごとに作成可能な `inotify` のインスタンス数の上限を指定します ([`inotify`](https://man7.org/linux/man-pages/man7/inotify.7.html) 参照)
`fs.inotify.max_user_watches`       | `1048576` | `8192`     | これは実ユーザー ID ごとに作成可能な watch 数の上限を指定します ([`inotify`](https://man7.org/linux/man-pages/man7/inotify.7.html) 参照)
`kernel.dmesg_restrict`             | `1`       | `0`        | この設定を有効にするとコンテナがカーネルのリングバッファー内のメッセージにアクセスするのを拒否します。この設定はホスト・システム上の非 root ユーザーへのアクセスも拒否することに注意してください
`kernel.keys.maxbytes`              | `2000000` | `20000`    | 非 root ユーザーが使用できる key ring の最大サイズ
`kernel.keys.maxkeys`               | `2000`    | `200`      | 非 root ユーザーが使用できるキーの最大数で、コンテナ数より大きくなければなりません
`net.ipv4.neigh.default.gc_thresh3` | `8192`    | `1024`     | これは ARP テーブル (IPv4) 内のエントリーの最大数です。1024 個を超えるコンテナを作成するなら増やすべきです。増やさなければ ARP テーブルがフルになったときに `neighbour: ndisc_cache: neighbor table overflow!` というエラーが発生し、コンテナがネットワーク設定を取得できなくなります ([`ip-sysctl`](https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt) 参照)
`net.ipv6.neigh.default.gc_thresh3` | `8192`    | `1024`     | これは ARP テーブル (IPv6) 内のエントリーの最大数です。1024 個を超えるコンテナを作成するなら増やすべきです。増やさなければ ARP テーブルがフルになったときに `neighbour: ndisc_cache: neighbor table overflow!` というエラーが発生し、コンテナがネットワーク設定を取得できなくなります ([`ip-sysctl`](https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt) 参照)
`vm.max_map_count`                  | `262144`  | `65530`    | このファイルはプロセスが持つメモリマップ領域の最大数を含みます。`malloc` の呼び出しの副作用として、 直接的には `mmap` と `mprotect` によって、また、共有ライブラリーをロードすることによって、メモリマップ領域を使います

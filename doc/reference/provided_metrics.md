(provided-metrics)=
# 提供されるメトリクス

LXDは、数々のインスタンスメトリクスと内部メトリクスを提供します。
これらのメトリクスの取り扱い方については、{ref}`metrics`を参照してください。

## インスタンスメトリクス

以下のインスタンスメトリクスが提供されています：

```{list-table}
   :header-rows: 1

* - メトリック
  - 説明
* - `lxd_cpu_effective_total`
  - 使用可能なCPUの総数
* - `lxd_cpu_seconds_total{cpu="<cpu>", mode="<mode>"}`
  - 使用されたCPU時間の合計（秒単位）
* - `lxd_disk_read_bytes_total{device="<dev>"}`
  - 読み出されたバイト数合計
* - `lxd_disk_reads_completed_total{device="<dev>"}`
  - 完了した読み取り回数合計
* - `lxd_disk_written_bytes_total{device="<dev>"}`
  - 書き込まれたバイト数合計
* - `lxd_disk_writes_completed_total{device="<dev>"}`
  - 完了した書き込み回数合計
* - `lxd_filesystem_avail_bytes{device="<dev>",fstype="<type>"}`
  - 利用可能な領域（バイト単位）
* - `lxd_filesystem_free_bytes{device="<dev>",fstype="<type>"}`
  - 空き領域（バイト単位）
* - `lxd_filesystem_size_bytes{device="<dev>",fstype="<type>"}`
  - ファイルシステムのサイズ（バイト単位）
* - `lxd_memory_Active_anon_bytes`
  - アクティブLRUリスト上のアノニマスメモリの量
* - `lxd_memory_Active_bytes`
  - アクティブLRUリスト上のメモリの量
* - `lxd_memory_Active_file_bytes`
  - アクティブLRUリスト上のファイルでバックアップされたメモリの量
* - `lxd_memory_Cached_bytes`
  - キャッシュメモリの量
* - `lxd_memory_Dirty_bytes`
  - ディスクへの書き込み待ちのメモリの量
* - `lxd_memory_HugepagesFree_bytes`
  - `hugetlb`の空きメモリの量
* - `lxd_memory_HugepagesTotal_bytes`
  - `hugetlb`の使用メモリの量
* - `lxd_memory_Inactive_anon_bytes`
  - インアクティブLRUリスト上のアノニマスメモリの量
* - `lxd_memory_Inactive_bytes`
  - インアクティブLRUリスト上のメモリの量
* - `lxd_memory_Inactive_file_bytes`
  - インアクティブLRUリスト上のファイルでバックアップされたメモリの量
* - `lxd_memory_Mapped_bytes`
  - マップされたメモリの量
* - `lxd_memory_MemAvailable_bytes`
  - 利用可能なメモリの量
* - `lxd_memory_MemFree_bytes`
  - 空きメモリの量
* - `lxd_memory_MemTotal_bytes`
  - 使用中メモリの量
* - `lxd_memory_OOM_kills_total`
  - out-of-memoryでkillされた回数
* - `lxd_memory_RSS_bytes`
  - アノニマスとswapキャッシュメモリの量
* - `lxd_memory_Shmem_bytes`
  - swapでバックアップされたキャッシュされたファイルシステムデータの量
* - `lxd_memory_Swap_bytes`
  - 使用中のスワップメモリの量
* - `lxd_memory_Unevictable_bytes`
  - 再生不可のメモリの使用量
* - `lxd_memory_Writeback_bytes`
  - ディスクへの同期のためみキューに入れられているメモリの量
* - `lxd_network_receive_bytes_total{device="<dev>"}`
  - 指定のインタフェース上の受信したバイト数
* - `lxd_network_receive_drop_total{device="<dev>"}`
  - 指定のインタフェース上の受信でドロップしたバイト数
* - `lxd_network_receive_errs_total{device="<dev>"}`
  - 指定のインタフェース上の受信エラー数
* - `lxd_network_receive_packets_total{device="<dev>"}`
  - 指定のインタフェース上の受信パケット数
* - `lxd_network_transmit_bytes_total{device="<dev>"}`
  - 指定のインタフェース上の送信したバイト数
* - `lxd_network_transmit_drop_total{device="<dev>"}`
  - 指定のインタフェース上の送信でドロップしたバイト数
* - `lxd_network_transmit_errs_total{device="<dev>"}`
  - 指定のインタフェース上の送信エラー数
* - `lxd_network_transmit_packets_total{device="<dev>"}`
  - 指定のインタフェース上の送信パケット数
* - `lxd_procs_total`
  - 稼働中のプロセス数
```

## 内部メトリクス

以下の内部メトリクスが提供されています：

```{list-table}
   :header-rows: 1

* - メトリック
  - 説明
* - `lxd_go_alloc_bytes_total`
  - 割り当てられた（その後の解放された分も含む）バイト数累計
* - `lxd_go_alloc_bytes`
  - 割り当てられ使用中のバイト数
* - `lxd_go_buck_hash_sys_bytes`
  - プロファイルのバケットハッシュテーブルで使用されたバイト数
* - `lxd_go_frees_total`
  - 解放の合計回数
* - `lxd_go_gc_sys_bytes`
  - システムメタデータのガベージコレクションで使用されたバイト数
* - `lxd_go_goroutines`
  - 現在存在するgoroutine数
* - `lxd_go_heap_alloc_bytes`
  - 割り当てられ使用中のヒープのバイト数
* - `lxd_go_heap_idle_bytes`
  - 使用を待っているヒープのバイト数
* - `lxd_go_heap_inuse_bytes`
  - 使用中のヒープのバイト数
* - `lxd_go_heap_objects`
  - 割り当てられたオブジェクト数
* - `lxd_go_heap_released_bytes`
  - OSに開放されたヒープのバイト数
* - `lxd_go_heap_sys_bytes`
  - システムから取得されたヒープのバイト数
* - `lxd_go_lookups_total`
  - ポインタルックアップの合計回数
* - `lxd_go_mallocs_total`
  - `mallocs`の合計回数
* - `lxd_go_mcache_inuse_bytes`
  - `mcache`構造で使用されるバイト数
* - `lxd_go_mcache_sys_bytes`
  - システムから取得された`mcache`構造で使用されるバイト数
* - `lxd_go_mspan_inuse_bytes`
  - `mspan`構造で使用されるバイト数
* - `lxd_go_mspan_sys_bytes`
  - システムから取得された`mspan`構造で使用されるバイト数
* - `lxd_go_next_gc_bytes`
  - 次のガベージコレクションが発生する際のヒープのバイト数
* - `lxd_go_other_sys_bytes`
  - 他のシステム割当に使用されるバイト数
* - `lxd_go_stack_inuse_bytes`
  - スタックアロケータに使用されるバイト数
* - `lxd_go_stack_sys_bytes`
  - スタックアロケータ用にシステムから取得されたバイト数
* - `lxd_go_sys_bytes`
  - システムから取得されたバイト数
* - `lxd_operations_total`
  - 実行中の処理の数
* - `lxd_uptime_seconds`
  - デーモンのuptime（秒単位）
* - `lxd_warnings_total`
  - アクティブな警告の数
```

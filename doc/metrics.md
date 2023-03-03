---
discourse: 12281,11735
relatedlinks: https://grafana.com/grafana/dashboards/15726
---

(metrics)=
# メトリクス

```{youtube} https://www.youtube.com/watch?v=EthK-8hm_fY
```

<!-- Include start metrics intro -->
LXD は全ての実行中のインスタンスについてのメトリクスといくつかの内部メトリクスを収集します。
これは CPU、メモリー、ネットワーク、ディスク、プロセスの使用量を含みます。
Prometheus で読み取って Grafana でグラフを表示するのに使うことを想定しています。
<!-- Include end metrics intro -->

クラスタ環境では、 LXD はアクセスされているサーバ上で稼働中のインスタンスの値だけを返します。各クラスタメンバーから別々にデータを取得する想定です。

インスタンスメトリクスは `/1.0/metrics` エンドポイントを呼ぶと更新されます。
メトリクスは複数のスクレイパーに対応するため 8 秒キャッシュします。メトリクスの取得は比較的重い処理ですので、影響が大きすぎるようならデフォルトの間隔より長い間隔でスクレイピングすることを検討してください。

## メトリクス用証明書の作成

`1.0/metrics` エンドポイントは他の証明書に加えて `metrics` タイプの証明書を受け付けるという点で特別なエンドポイントです。
このタイプの証明書はメトリクス専用で、インスタンスや他のLXDのエンティティの操作には使用できません。

新しい証明書は以下のように作成します（この手順はメトリクス用の証明書に限ったものではありません）。

```bash
openssl req -x509 -newkey ec -pkeyopt ec_paramgen_curve:secp384r1 -sha384 -keyout metrics.key -nodes -out metrics.crt -days 3650 -subj "/CN=metrics.local"
```

*注意*: 上のコマンドが正しい証明書を生成するには OpenSSL 1.1.0+ が必要です。

作成後、証明書を信頼済みクライアントのリストに追加する必要があります。

```bash
lxc config trust add metrics.crt --type=metrics
```

## Prometheus にターゲットを追加

Prometheus が LXD からメトリクスを取得するためには、 LXD をターゲットに追加する必要があります。

まず、 LXD にネットワーク越しにアクセスできるように `core.https_address` を設定しているかを確認してください。
これは以下のコマンドを実行することで設定できます。

```bash
lxc config set core.https_address ":8443"
```

あるいは、メトリクス用途専用の `core.metrics_address` を使うことも出来ます。

次に、新しく作成した証明書と鍵を LXD のサーバ証明書とともに Prometheus からアクセスできるようにする必要があります。
これは以下の 3 つのファイルを `/etc/prometheus/tls` にコピーすればできます。

```bash
# tls ディレクトリーを新規に作成
mkdir /etc/prometheus/tls

# 新規に作成された証明書と鍵を tls ディレクトリーにコピー
cp metrics.crt metrics.key /etc/prometheus/tls

# LXD サーバ証明書を tls ディレクトリーにコピー
cp /var/snap/lxd/common/lxd/server.crt /etc/prometheus/tls

# これらのファイルを Prometheus が読めるようにする（通常 Prometheus は "prometheus" ユーザーで稼働しています）
chown -R prometheus:prometheus /etc/prometheus/tls
```

最後に、 LXD をターゲットに追加する必要があります。
これは `/etc/prometheus/prometheus.yaml` を編集する必要があります。
設定を以下のようにします。

```yaml
scrape_configs:
  - job_name: lxd
    metrics_path: '/1.0/metrics'
    scheme: 'https'
    static_configs:
      - targets: ['foo.example.com:8443']
    tls_config:
      ca_file: 'tls/server.crt'
      cert_file: 'tls/metrics.crt'
      key_file: 'tls/metrics.key'
      # XXX: server_name は targets のホスト名が証明書でカバーされない
      #      (証明書の SAN リストに含まれない) 場合は必須です
      server_name: 'foo'
```

上の例では `/etc/prometheus/tls/server.crt` は以下のような内容になっています。

```
$ openssl x509 -noout -text -in /etc/prometheus/tls/server.crt
...
            X509v3 Subject Alternative Name:
                DNS:foo, IP Address:127.0.0.1, IP Address:0:0:0:0:0:0:0:1
...
```

Subject Alternative Name (SAN) リストが `targets` リストのホスト名を含んでいないので、 `server_name` ディレクティブを使用して比較に使用する名前を上書きする必要があります。

以下は複数の LXD サーバのメトリックを収集するために複数のジョブを使用する `prometheus.yaml` の設定例です。

```yaml
scrape_configs:
  # abydos, langara, orilla は最初にabydosからブートストラップした単一クラスタで
  # (ここでは`hdc`と呼びます)、このため3ノードで`ca_file`と`server_name`を共有しています。
  # `ca_file`はLXDクラスタの各メンバ上に存在する`/var/snap/lxd/common/lxd/cluster.crt`
  # ファイルに対応しています。
  #
  # 注意: `project`パラメータは`default`プロジェクトを使用しないか複数のプロジェクトを
  #       使用する場合に提供されます。
  #
  # 注意: クラスタの各メンバーはローカルで稼働するインスタンスのメトリクスだけを提供します。
  #       これが`lxd-hdc`クラスタが3つのターゲットを一覧表示している理由です。
  - job_name: "lxd-hdc"
    metrics_path: '/1.0/metrics'
    params:
      project: ['jdoe']
    scheme: 'https'
    static_configs:
      - targets:
        - 'abydos.hosts.example.net:8444'
        - 'langara.hosts.example.net:8444'
        - 'orilla.hosts.example.net:8444'
    tls_config:
      ca_file: 'tls/abydos.crt'
      cert_file: 'tls/metrics.crt'
      key_file: 'tls/metrics.key'
      server_name: 'abydos'

  # jupiter, mars, saturn は3つのスタンドアロンの LXD サーバです。
  # 注意: これらでは`default`プロジェクトのみが使用されているため、プロジェクトの設定は省略しています。
  - job_name: "lxd-jupiter"
    metrics_path: '/1.0/metrics'
    scheme: 'https'
    static_configs:
      - targets: ['jupiter.example.com:9101']
    tls_config:
      ca_file: 'tls/jupiter.crt'
      cert_file: 'tls/metrics.crt'
      key_file: 'tls/metrics.key'
      server_name: 'jupiter'

  - job_name: "lxd-mars"
    metrics_path: '/1.0/metrics'
    scheme: 'https'
    static_configs:
      - targets: ['mars.example.com:9101']
    tls_config:
      ca_file: 'tls/mars.crt'
      cert_file: 'tls/metrics.crt'
      key_file: 'tls/metrics.key'
      server_name: 'mars'

  - job_name: "lxd-saturn"
    metrics_path: '/1.0/metrics'
    scheme: 'https'
    static_configs:
      - targets: ['saturn.example.com:9101']
    tls_config:
      ca_file: 'tls/saturn.crt'
      cert_file: 'tls/metrics.crt'
      key_file: 'tls/metrics.key'
      server_name: 'saturn'
```

## 提供されるインスタンスメトリクス

以下のインスタンスメトリクスが提供されます。

* `lxd_cpu_effective_total`
* `lxd_cpu_seconds_total{cpu="<cpu>", mode="<mode>"}`
* `lxd_disk_read_bytes_total{device="<dev>"}`
* `lxd_disk_reads_completed_total{device="<dev>"}`
* `lxd_disk_written_bytes_total{device="<dev>"}`
* `lxd_disk_writes_completed_total{device="<dev>"}`
* `lxd_filesystem_avail_bytes{device="<dev>",fstype="<type>"}`
* `lxd_filesystem_free_bytes{device="<dev>",fstype="<type>"}`
* `lxd_filesystem_size_bytes{device="<dev>",fstype="<type>"}`
* `lxd_memory_Active_anon_bytes`
* `lxd_memory_Active_bytes`
* `lxd_memory_Active_file_bytes`
* `lxd_memory_Cached_bytes`
* `lxd_memory_Dirty_bytes`
* `lxd_memory_HugepagesFree_bytes`
* `lxd_memory_HugepagesTotal_bytes`
* `lxd_memory_Inactive_anon_bytes`
* `lxd_memory_Inactive_bytes`
* `lxd_memory_Inactive_file_bytes`
* `lxd_memory_Mapped_bytes`
* `lxd_memory_MemAvailable_bytes`
* `lxd_memory_MemFree_bytes`
* `lxd_memory_MemTotal_bytes`
* `lxd_memory_OOM_kills_total`
* `lxd_memory_RSS_bytes`
* `lxd_memory_Shmem_bytes`
* `lxd_memory_Swap_bytes`
* `lxd_memory_Unevictable_bytes`
* `lxd_memory_Writeback_bytes`
* `lxd_network_receive_bytes_total{device="<dev>"}`
* `lxd_network_receive_drop_total{device="<dev>"}`
* `lxd_network_receive_errs_total{device="<dev>"}`
* `lxd_network_receive_packets_total{device="<dev>"}`
* `lxd_network_transmit_bytes_total{device="<dev>"}`
* `lxd_network_transmit_drop_total{device="<dev>"}`
* `lxd_network_transmit_errs_total{device="<dev>"}`
* `lxd_network_transmit_packets_total{device="<dev>"}`
* `lxd_procs_total`

## 提供される内部メトリクス

以下の内部メトリクスが提供されます。

* `lxd_go_alloc_bytes_total`
* `lxd_go_alloc_bytes`
* `lxd_go_buck_hash_sys_bytes`
* `lxd_go_frees_total`
* `lxd_go_gc_sys_bytes`
* `lxd_go_goroutines`
* `lxd_go_heap_alloc_bytes`
* `lxd_go_heap_idle_bytes`
* `lxd_go_heap_inuse_bytes`
* `lxd_go_heap_objects`
* `lxd_go_heap_released_bytes`
* `lxd_go_heap_sys_bytes`
* `lxd_go_lookups_total`
* `lxd_go_mallocs_total`
* `lxd_go_mcache_inuse_bytes`
* `lxd_go_mcache_sys_bytes`
* `lxd_go_mspan_inuse_bytes`
* `lxd_go_mspan_sys_bytes`
* `lxd_go_next_gc_bytes`
* `lxd_go_other_sys_bytes`
* `lxd_go_stack_inuse_bytes`
* `lxd_go_stack_sys_bytes`
* `lxd_go_sys_bytes`
* `lxd_operations_total`
* `lxd_uptime_seconds`
* `lxd_warnings_total`

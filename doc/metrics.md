# インスタンスメトリクス <!-- Instance metrics -->
LXD は全ての実行中のインスタンスについてのメトリクスを提供します。これは CPU、メモリー、ネットワーク、ディスク、プロセスの使用量を含み、Prometheus で読み取って Grafana でグラフを表示するのに使うことを想定しています。
クラスター環境では、 LXD はアクセスされているサーバー上で稼働中のインスタンスの値だけを返します。各クラスターメンバーから別々にデータを取得する想定です。
インスタンスメトリクスは `/1.0/metrics` エンドポイントを呼ぶと更新されます。
メトリクスは複数からデータ取得するのに対応するため 15 秒キャッシュします。メトリクスの取得は比較的重い処理ですので、影響を抑えるため 30 秒か 60 秒間隔でメトリクスを取得することをお勧めします。
<!--
LXD provides metrics for all running instances. Those covers CPU, memory, network, disk and process usage and are meant to be consumed by Prometheus and likely graphed in Grafana.
In cluster environments, LXD will only return the values for instances running on the server being accessed. It's expected that each cluster member will be scraped separately.
The instance metrics are updated when calling the `/1.0/metrics` endpoint.
They are cached for 15s to handle multiple scrapers. Fetching metrics is a relatively expensive operation for LXD to perform so we would recommend scraping at a 30s or 60s rate to limit impact.
-->

# メトリクス用証明書の作成 <!-- Create metrics certificate -->
`1.0/metrics` エンドポイントは他の証明書に加えて `metrics` タイプの証明書を受け付けるという点で特別なエンドポイントです。
このタイプの証明書はメトリクス専用で、インスタンスや他の LXD のオブジェクトの操作には使用できません。
<!--
The `/1.0/metrics` endpoint is a special one as it also accepts a `metrics` type certificate.
This kind of certificate is meant for metrics only, and won't work for interaction with instances or any other LXD objects.
-->

新しい証明書は以下のように作成します（この手順はメトリクス用の証明書に限ったものではありません）。
<!--
Here's how to create a new certificate (this is not specific to metrics):
-->

```bash
openssl req -x509 -newkey rsa:2048 -keyout ~/.config/lxc/metrics.key -nodes -out ~/.config/lxc/metrics.crt -subj "/CN=lxd.local"
```

作成後、証明書を信頼済みクライアントのリストに追加する必要があります。
<!--
Now, this certificate needs to be added to the list of trusted clients:
-->

```bash
lxc config trust add ~/.config/lxc/metrics.crt --type=metrics
```

# Prometheus にターゲットを追加 <!-- Add target to Prometheus -->
Prometheus が LXD からメトリクスを取得するためには、 LXD をターゲットに追加する必要があります。
<!--
In order for Prometheus to scrape from LXD, it has to be added to the targets.
-->

まず、 LXD にネットワーク越しにアクセスできるように `core.https_address` を設定しているかを確認してください。
これは以下のコマンドを実行することで設定できます。
<!--
First, one needs to ensure that `core.https_address` is set so LXD can be reached over the network.
This can be done by running:
-->

```bash
lxc config set core.https_address ":8443"
```

あるいは、メトリクス用途専用の `core.metrics_address` を使うことも出来ます。
<!--
Alternatively, one can use `core.metrics_address` which is intended for metrics only.
-->

次に、新しく作成した証明書と鍵を LXD のサーバー証明書とともに Prometheus からアクセスできるようにする必要があります。
これは以下の 3 つのファイルを `/etc/prometheus/tls` にコピーすればできます。
<!--
Second, the newly created certificate and key, as well as the LXD server certificate need to be accessible to Prometheus.
For this, these three files can be copied to `/etc/prometheus/tls`:
-->

```bash
# tls ディレクトリーを新規に作成
mkdir /etc/prometheus/tls

# 新規に作成された証明書と鍵を tls ディレクトリーにコピー
cp ~/.config/lxc/metrics.crt ~/.config/lxc/metrics.key /etc/prometheus/tls

# LXD サーバー証明書を tls ディレクトリーにコピー
cp /var/snap/lxd/common/lxd/server.crt /etc/prometheus/tls

# これらのファイルを Prometheus が読めるようにする（通常 Prometheus は "prometheus" ユーザーで稼働しています）
chown -R prometheus:prometheus /etc/prometheus/tls
```

<!--
```bash
# Create new tls directory
mkdir /etc/prometheus/tls

# Copy newly created certificate and key to tls directory
cp ~/.config/lxc/metrics.crt ~/.config/lxc/metrics.key /etc/prometheus/tls

# Copy LXD server certificate to tls directory
cp /var/snap/lxd/common/lxd/server.crt /etc/prometheus/tls

# Make sure Prometheus can read these files (usually, Prometheus is run as user "prometheus")
chown -R prometheus:prometheus /etc/prometheus/tls
```
-->

最後に、 LXD をターゲットに追加する必要があります。
これは `/etc/prometheus/prometheus.yaml` を編集する必要があります。
設定を以下のようにします。
<!--
Lastly, LXD has to be added as target.
For this, `/etc/prometheus/prometheus.yaml` needs to be edited.
Here's what the config needs to look like:
-->

```yaml
scrape_configs:
  - job_name: lxd
    tls_config:
      ca_file: 'tls/lxd.crt'
      key_file: 'tls/metrics.key'
      cert_file: 'tls/metrics.crt'
    static_configs:
      - targets: ['127.0.0.1:8443']
    metrics_path: '/1.0/metrics'
    scheme: 'https'
```

# インスタンスメトリクス
LXD は全ての実行中のインスタンスについてのメトリクスを提供します。これは CPU、メモリー、ネットワーク、ディスク、プロセスの使用量を含み、Prometheus で読み取って Grafana でグラフを表示するのに使うことを想定しています。
クラスタ環境では、 LXD はアクセスされているサーバ上で稼働中のインスタンスの値だけを返します。各クラスタメンバーから別々にデータを取得する想定です。
インスタンスメトリクスは `/1.0/metrics` エンドポイントを呼ぶと更新されます。
メトリクスは複数のスクレイパーに対応するため 8 秒キャッシュします。メトリクスの取得は比較的重い処理ですので、影響が大きすぎるようならデフォルトの間隔より長い間隔でスクレイピングすることを検討してください。

## メトリクス用証明書の作成
`1.0/metrics` エンドポイントは他の証明書に加えて `metrics` タイプの証明書を受け付けるという点で特別なエンドポイントです。
このタイプの証明書はメトリクス専用で、インスタンスや他の LXD のオブジェクトの操作には使用できません。

新しい証明書は以下のように作成します（この手順はメトリクス用の証明書に限ったものではありません）。

```bash
openssl req -x509 -newkey ec -pkeyopt ec_paramgen_curve:secp384r1 -sha384 -keyout metrics.key -nodes -out metrics.crt -days 3650 -subj "/CN=metrics.local"
```

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
      - targets: ['127.0.0.1:8443']
    tls_config:
      ca_file: 'tls/lxd.crt'
      cert_file: 'tls/metrics.crt'
      key_file: 'tls/metrics.key'
```

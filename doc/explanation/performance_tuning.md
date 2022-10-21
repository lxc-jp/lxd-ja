(performance-tuning)=
# パフォーマンスチューニング

お使いの LXD 環境を本番稼働に移行する準備が出来たら、システムのパフォーマンスを最適化するためにいくらか時間を取るほうが良いです。
パフォーマンスに影響を与えるいくつかの視点があります。
お使いの LXD 環境を改善するためにチューニングするべき選択肢と設定を決定するのに以下の手順が役立ちます。

## ベンチマークを実行する

LXD はシステムのパフォーマンスを評価するためにベンチマークツールを提供しています。
このツールを使って複数のコンテナを初期化・起動し、システムがコンテナを作成するのに必要な時間を計測できます。
異なる LXD の設定、システム設定、さらにはハードウェア構成に対して繰り返しツールを実行することで、パフォーマンスを比較し、どの設定が理想的か評価できます。

ツールを実行する手順については {ref}`benchmark-performance` を参照してください。

## インスタンスのメトリクスをモニターする

% Include content from [../metrics.md](../metrics.md)
```{include} ../metrics.md
    :start-after: <!-- Include start metrics intro -->
    :end-before: <!-- Include end metrics intro -->
```

あなたのインスタンスが使用しているリソースを見積もるために定期的にメトリクスをモニターするほうが良いです。
スパイクやボトルネックがある場合や、使用量のパターンが変化したり、設定を見直す必要がある場合に、これらの数値が役立ちます。

メトリクス収集についての詳細な情報は {ref}`instance-metrics` を参照してください。

## サーバ設定をチューニングする

ほとんどの Linux ディストリビューションのデフォルトのカーネル設定は大量のコンテナや仮想マシンを稼働させるのに最適化されていません。
ですので、デフォルトの設定で引き起こされる制限にひっかかるのを避けるため、関連する設定を確認、変更するほうが良いです。

これらの制限にひっかかかった場合の典型的なエラーは以下のようなものです。

* `Failed to allocate directory watch: Too many open files`
* `<Error> <Error>: Too many open files`
* `failed to open stream: Too many open files in...`
* `neighbour: ndisc_cache: neighbor table overflow!`

関連するサーバ設定と提案される値の一覧は {ref}`server-settings` を参照してください。

## ネットワーク帯域幅をチューニングする

インスタンス間あるいは LXD ホストとインスタンス間で大量のローカルなアクティビティがある場合、あるいは高速なインターネット接続をお持ちの場合、 LXD のセットアップのネットワーク帯域幅を増やすことを検討すると良いです。
これは送信と受信のキューの長さを拡張することで実現できます。

手順については {ref}`network-increase-bandwidth` を参照してください。

```{toctree}
:maxdepth: 1
:hidden:

パフォーマンスのベンチマーク <../howto/benchmark_performance>
帯域幅の拡大 <../howto/network_increase_bandwidth>
サーバ設定 <../reference/server_settings>
```
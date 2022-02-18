# プリシード YAML を使った非対話型設定
`lxd init` コマンドは `--preseed` コマンドラインフラグをサポートしており、
LXD デーモンの設定、ストレージプール、ネットワークデバイスとプロファイルを
非対話式にを完全に構成することができます。

例えば、 LXD を新規インストールした状態で、以下のコマンドを実行

```bash
    cat <<EOF | lxd init --preseed
config:
  core.https_address: 192.168.1.1:9999
  images.auto_update_interval: 15
networks:
- name: lxdbr0
  type: bridge
  config:
    ipv4.address: auto
    ipv6.address: none
EOF
```

すると 192.168.1.1 のアドレスのポート 9999 で HTTPS 接続をリッスンし、
15 時間ごとにイメージを自動的にアップデートし、 `lxdbr0` という名前の
ネットワークデバイスを作成し、 IPV4の アドレスを自動的に割り当てるように
LXD デーモンを構成します。

## 新規インストールした LXD を設定する

新規インストールした LXD のインスタンスを設定する場合、 preseed
コマンドは必ず成功し、希望の設定を適用できます (与えられた YAML が
正しいキーと値を含んでいる限り)。というのは、希望の状態と衝突する
既存の状態が存在しないからです。

## 既存の LXD を再構成する

既存の LXD インスタンスを preseed コマンドを使って再構成する場合、
指定された YAML 設定は既存の設定を完全に上書きすることを意味します
(新規インストールの LXD の場合と同様に、指定された設定が存在しない場合は
それらは単に作成されます)。

既存の設定を上書きする場合は、その設定の新しい希望の状態の完全な設定を
指定する必要があります (つまり [RESTful API](rest-api.md) での `PUT` と
同じ考え方です)。

### ロールバック

新しく希望する設定の一部が既存の設定と衝突する場合 (例えばストレージ
プールをドライバーを `dir` から `zfs` に変更しようとするなど)、
preseed コマンドは失敗し、それまでに適用したあらゆる変更を自動的に
ロールバックしようとベストを尽くします。

例えば新しい設定で作られた設定を削除したり、上書きされた設定を元の
状態に戻したりするでしょう。

設定を上書きするのに失敗した場合のモードは [RESTful API](rest-api.md)
の `PUT` リクエストの場合と同様です。

ただし、まれにではありますが (典型的にはバックエンドのバグか制限により)、
ロールバック自体も失敗する可能性があることにご注意ください。ですので
LXD デーモンをプリシードで再構成しようとするときは注意が必要です。


## デフォルト・プロファイル

対話的な初期化モードをは異なり、指定した YAML のペイロードで明示的に
変更を指示しない限り、 `lxd init --preseed` はデフォルト・プロファイルを
特定の状態に変更することはしません。

例えば、典型的にはデフォルトプロファイルにルート・ディスク・デバイスと
ネットワーク・インターフェースをアタッチしたいでしょう。以下の例を
参照してください。

## 設定の形式

さまざまな設定のサポートされるキーと値は [RESTful API](rest-api.md) に
ドキュメントされているものと同じです。ただし、 YAML が読みやすいように
変換されたものになっています (YAML は JSON のスーパーセットなので
JSON を使うこともできます)。

以下に設定可能な取っ手のほとんどを含むプリシードのペイロードの例を
示します。あなた自身のペイロードのテンプレートとして使い、ここに
必要な設定を追加、変更、削除して使うことができます。


```yaml

# デーモンの設定
config:
  core.https_address: 192.168.1.1:9999
  core.trust_password: sekret
  images.auto_update_interval: 6

# ストレージプール
storage_pools:
- name: data
  driver: zfs
  config:
    source: my-zfs-pool/my-zfs-dataset

# ネットワークデバイス
networks:
- name: lxd-my-bridge
  type: bridge
  config:
    ipv4.address: auto
    ipv6.address: none

# プロファイル
profiles:
- name: default
  devices:
    root:
      path: /
      pool: data
      type: disk
- name: test-profile
  description: "Test profile"
  config:
    limits.memory: 2GB
  devices:
    test0:
      name: test0
      nictype: bridged
      parent: lxd-my-bridge
      type: nic
```

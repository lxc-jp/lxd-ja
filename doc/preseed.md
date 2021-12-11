# presseed YAML を使った非対話型設定
<!-- Non-interactive configuration via preseed YAML -->

<!--
The `lxd init` command supports a `\-\-preseed` command line flag that
makes it possible to fully configure LXD daemon settings, storage
pools, network devices and profiles, in a non-interactive way.
-->
`lxd init` コマンドは `--preseed` コマンドラインフラグをサポートしており、
LXD デーモンの設定、ストレージプール、ネットワークデバイスとプロファイルを
非対話式にを完全に構成することができます。

<!--
For example, starting from a brand new LXD installation, the command
line:
-->
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

<!--
will configure the LXD daemon to listen for HTTPS connections on port
9999 of the 192.168.1.1 address, to automatically update images every
15 hours, and to create a network bridge device named `lxdbr0`, which
will get assigned an IPv4 address automatically.
-->
すると 192.168.1.1 のアドレスのポート 9999 で HTTPS 接続をリッスンし、 
15 時間ごとにイメージを自動的にアップデートし、 `lxdbr0` という名前の
ネットワークデバイスを作成し、 IPV4の アドレスを自動的に割り当てるように
LXD デーモンを構成します。

## 新規インストールした LXD を設定する <!-- Configure a brand new LXD -->

<!--
If you are configuring a brand new LXD instance, then the preseed
command will always succeed and apply the desired configuration (as
long as the given YAML contains valid keys and values), since there is
no existing state that might conflict with the desired one.
-->
新規インストールした LXD のインスタンスを設定する場合、 preseed
コマンドは必ず成功し、希望の設定を適用できます (与えられた YAML が
正しいキーと値を含んでいる限り)。というのは、希望の状態と衝突する
既存の状態が存在しないからです。

## 既存の LXD を再構成する <!-- Re-configuring an existing LXD -->

<!--
If you are re-configuring an existing LXD instance using the preseed
command, then the provided YAML configuration is meant to completely
overwrite existing entities (if the provided entities do not exist,
they will just be created, as in the brand new LXD case).
-->
既存の LXD インスタンスを preseed コマンドを使って再構成する場合、
指定された YAML 設定は既存の設定を完全に上書きすることを意味します
(新規インストールの LXD の場合と同様に、指定された設定が存在しない場合は
それらは単に作成されます)。

<!--
In case you are overwriting an existing entity you must provide the full
configuration of the new desired state for the entity (i.e. the semantics is
the same as a `PUT` request in the [RESTful API](rest-api.md)).
-->
既存の設定を上書きする場合は、その設定の新しい希望の状態の完全な設定を
指定する必要があります (つまり [RESTful API](rest-api.md) での `PUT` と
同じ考え方です)。

### ロールバック <!-- Rollback -->

<!--
If some parts of the new desired configuration conflict with the
existing state (for example they try to change the driver of a storage
pool from `dir` to `zfs`), then the preseed command will fail and will
automatically try its best to rollback any change that was applied so
far.
-->
新しく希望する設定の一部が既存の設定と衝突する場合 (例えばストレージ
プールをドライバーを `dir` から `zfs` に変更しようとするなど)、
preseed コマンドは失敗し、それまでに適用したあらゆる変更を自動的に
ロールバックしようとベストを尽くします。

<!--
For example it will delete entities that were created by the new
configuration and revert overwritten entities back to their original
state.
-->
例えば新しい設定で作られた設定を削除したり、上書きされた設定を元の
状態に戻したりするでしょう。

<!--
Failure modes when overwriting entities are the same as `PUT` requests
in the [RESTful API](rest-api.md).
-->
設定を上書きするのに失敗した場合のモードは [RESTful API](rest-api.md)
の `PUT` リクエストの場合と同様です。

<!--
Note however, that the rollback itself might potentially fail as well,
although rarely (typically due to backend bugs or limitations). Thus
care must be taken when trying to reconfigure a LXD daemon via
preseed.
-->
ただし、まれにではありますが (典型的にはバックエンドのバグか制限により)、
ロールバック自体も失敗する可能性があることにご注意ください。ですので
LXD デーモンを preseed で再構成しようとするときは注意が必要です。


## デフォルト・プロファイル <!-- Default profile-->

<!--
Differently from the interactive init mode, the `lxd init \-\-preseed`
command line will not modify the default profile in any particular
way, unless you explicitly express that in the provided YAML payload.
-->
対話的な初期化モードをは異なり、指定した YAML のペイロードで明示的に
変更を指示しない限り、 `lxd init --preseed` はデフォルト・プロファイルを
特定の状態に変更することはしません。

<!--
For instance, you will typically want to attach a root disk device and
a network interface to your default profile. See below for an example.
-->
例えば、典型的にはデフォルトプロファイルにルート・ディスク・デバイスと
ネットワーク・インターフェースをアタッチしたいでしょう。以下の例を
参照してください。

## 設定の形式 <!-- Configuration format -->

<!--
The supported keys and values of the various entities are the same as
the ones documented in the [RESTful API](rest-api.md), but converted
to YAML for easier reading (however you can use JSON too, since YAML
is a superset of JSON).
-->
さまざまな設定のサポートされるキーと値は [RESTful API](rest-api.md) に
ドキュメントされているものと同じです。ただし、 YAML が読みやすいように
変換されたものになっています (YAML は JSON のスーパーセットなので
JSON を使うこともできます)。

<!--
Here follows an example of a preseed payload containing most of the
possible configuration knobs. You can use it as a template for your
own one, and add, change or remove what you need:
-->
以下に設定可能な取っ手のほとんどを含む preseed のペイロードの例を
示します。あなた自身のペイロードのテンプレートとして使い、ここに
必要な設定を追加、変更、削除して使うことができます。

<!--
```yaml

# Daemon settings
config:
  core.https_address: 192.168.1.1:9999
  core.trust_password: sekret
  images.auto_update_interval: 6

# Storage pools
storage_pools:
- name: data
  driver: zfs
  config:     
    source: my-zfs-pool/my-zfs-dataset

# Network devices
networks:
- name: lxd-my-bridge
  type: bridge
  config:
    ipv4.address: auto
    ipv6.address: none

# Profiles
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
-->

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

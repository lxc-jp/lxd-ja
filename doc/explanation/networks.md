# ネットワークについて

あなたのインスタンスをインターネットに接続するにはいろいろな方法があります。
最も簡単な方法は LXD の初期化時にネットワークブリッジを作って全てのインスタンスでこのブリッジを使うことですが、 LXD はネットワークの多くの異なる高度な設定をサポートします。

## ネットワークデバイス

インスタンスへの直接のネットワークアクセスを許可するには、最低 1 つのネットワークデバイス、 {abbr}`NIC (Network Interface Controller)` とも呼ばれる、を割り当てる必要があります。
ネットワークデバイスは以下のどれかの方法で設定できます。

- LXD の初期化中にセットアップしたデフォルトのネットワークブリッジを使用する。
  デフォルトの設定を表示するにはデフォルトのプロファイルを確認します。

        lxc profile show default

  この方法はインスタンスのネットワークを指定しない場合に使用します。
- 既存のネットワークインタフェースをインスタンスにネットワークデバイスとして追加して使用する。
  このネットワークインタフェースは LXD の制御外です。
  そのため、ネットワークインタフェースを使用するために必要な全ての情報を LXD に指定する必要があります。

  以下のようなコマンドを使用します。

        lxc config device add <インスタンス名> <デバイス名> nic nictype=<NICタイプ> ...

  指定可能な NIC タイプの一覧とそれらの設定プロパティについては {ref}`instance_device_type_nic` を参照してください。

  例えば、既存の Linux ブリッジ (`br0`) を追加するには以下のコマンドを使えます。

        lxc config device add <instance_name> eth0 nic nictype=bridged parent=br0
- {doc}`マネージドネットワークを作成 </howto/network_create>` し、それをインスタンスにネットワークデバイスとして追加する。
  この方法では LXD は設定されるネットワークについての全ての必要な情報を持っているのでデバイスへの追加する際はネットワーク名するだけで良いです。

        lxc config device add <インスタンス名> <デバイス名> nic network=<ネットワーク名>

  必要であれば、ネットワークのデフォルト設定をオーバーライドする追加の設定をコマンドに追加できます。

## マネージドネットワーク

LXD でマネージドネットワークは `lxc network [create|edit|set]` コマンドで作成と設定をします。

ネットワークタイプによって、 LXD はネットワークを完全に制御するか、単に外部のネットワークインタフェースを管理するかのどちらかになります。

全ての {ref}`NIC タイプ <instance_device_type_nic>` がネットワークタイプとしてサポートされているわけではないことに注意してください。
LXD はいくつかのタイプのみマネージドネットワークとしてセットアップできます。

### 完全に制御されるネットワーク

完全に制御されるネットワークではネットワークインタフェースを作成し、例えば IP を管理する機能を含むほとんどの機能を提供します。

LXD は以下のネットワークタイプをサポートします。

{ref}`network-bridge`
: % Include content from [../reference/network_bridge.md](../reference/network_bridge.md)
  ```{include} ../reference/network_bridge.md
      :start-after: <!-- Include start bridge intro -->
      :end-before: <!-- Include end bridge intro -->
  ```

  LXD の文脈では、 `bridge` ネットワークタイプはインスタンスに接続し単一の L2 ネットワークセグメントにするような L2 ブリッジを作成します。
  これによりインスタンス間のトラフィックを通すことができます。
  ブリッジはさらにローカルの DHCP と DNS を提供することもできます。

  これがデフォルトのネットワークタイプです。

{ref}`network-ovn`
: % Include content from [../reference/network_ovn.md](../reference/network_ovn.md)
  ```{include} ../reference/network_ovn.md
      :start-after: <!-- Include start OVN intro -->
      :end-before: <!-- Include end OVN intro -->
  ```

  LXD の文脈では、 `ovn` ネットワークタイプは論理ネットワークを作成します。
  セットアップするには OVN ツールをインストールし設定する必要があります。
  さらに、OVN にネットワーク接続を提供するアップリンクのネットワークを作成する必要があります。
  アップリンクのネットワークとして、外部ネットワークタイプの 1 つかマネージドな LXD ブリッジを使う必要があります。

  ```{tip}
  他のネットワークタイプと違って、 OVN ネットワークは {ref}`プロジェクト <projects>` 内に作成・管理できます。
  これは、制限されたプロジェクトであっても、非管理者ユーザとして自身の OVN ネットワークを作成できることを意味します。
  ```

### 外部ネットワーク

% Include content from [../reference/network_external.md](../reference/network_external.md)
```{include} ../reference/network_external.md
    :start-after: <!-- Include start external intro -->
    :end-before: <!-- Include end external intro -->
```

{ref}`network-macvlan`
: % Include content from [../reference/network_macvlan.md](../reference/network_macvlan.md)
  ```{include} ../reference/network_macvlan.md
      :start-after: <!-- Include start macvlan intro -->
      :end-before: <!-- Include end macvlan intro -->
  ```

  LXD の文脈では、 `macvlan` ネットワークタイプは親の macvlan インタフェースへインスタンスを接続する際に使用するプリセット設定を提供します。

{ref}`network-sriov`
: % Include content from [../reference/network_sriov.md](../reference/network_sriov.md)
  ```{include} ../reference/network_sriov.md
      :start-after: <!-- Include start SR-IOV intro -->
      :end-before: <!-- Include end SR-IOV intro -->
  ```

  LXD の文脈では、 `sriov` ネットワークタイプは親の SR-IOV インタフェースへインスタンスを接続する際に使用するプリセット設定を提供します。

{ref}`network-physical`
: % Include content from [../reference/network_physical.md](../reference/network_physical.md)
  ```{include} ../reference/network_physical.md
      :start-after: <!-- Include start physical intro -->
      :end-before: <!-- Include end physical intro -->
  ```

  OVN ネットワークを親インタフェースに接続する際のプリセット設定を提供します。

## お勧めの設定

一般に、マネージドネットワークは設定が容易で設定を繰り返すこと無く複数のインスタンスで同じネットワークを再利用できるので、マネージドネットワークが使用できる場合はこれを使用すべきです。

どのネットワークタイプを使用すべきかはあなたの固有の使い方によります。
完全に制御されたネットワークを選ぶ場合は、ネットワークデバイスを使用するのに比べてより多くの機能を提供します。

一般的なお勧めとしては

- LXD を単一のシステム上かパブリッククラウドで動かしている場合は、 {ref}`network-bridge` を使用してください。場合によっては [Ubuntu Fan](https://www.youtube.com/watch?v=5cwd0vZJ5bw) と共に使用するのが良いかもしれません。
- あなた自身のプライベートクラウドで LXD を動かしている場合は、 {ref}`network-ovn` を使用してください。

  ```{note}
  OVN は適切な運用には共有された L2 のアップリンクネットワークが必要です。
  このため、パブリッククラウドで LXD を動かしている場合は通常 OVN は使用できません。
  ```
- インスタンス NIC をマネージドネットワークに接続するためには、可能であれば `parent` プロパティより `network` プロパティを使用してください。
  こうすることで、 NIC はネットワークの設定を引き継ぎ、 `nictype` を指定する必要がなくなります。

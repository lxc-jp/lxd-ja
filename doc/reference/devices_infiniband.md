(devices-infiniband)=
# タイプ: `infiniband`

サポートされるインスタンスタイプ: コンテナ、VM

LXD では、InfiniBand デバイスに対する 2 種類の異なったネットワークタイプが使えます:

- `physical`: ホストの物理デバイスをパススルーで直接使います。対象のデバイスはホスト上では見えなくなり、インスタンス内に出現します
- `sriov`: SR-IOV が有効な物理ネットワークデバイスの仮想ファンクション（virtual function）をインスタンスに与えます

ネットワークインターフェースの種類が異なると追加のプロパティが異なります。現時点のリストは次の通りです:

キー      | 型      | デフォルト値       | 必須 | 使用される種別      | 説明
:--       | :--     | :--                | :--  | :--                 | :--
`nictype` | string  | -                  | yes  | 全て                | デバイスタイプ。`physical` か `sriov` のいずれか
`name`    | string  | カーネルが割り当て | no   | 全て                | インスタンス内部でのインターフェース名
`hwaddr`  | string  | ランダムに割り当て | no   | 全て                | 新しいインターフェースの MAC アドレス。 20 バイト全てを指定するか短い 8 バイト (この場合親デバイスの最後の 8 バイトだけを変更) のどちらかを設定可能
`mtu`     | integer | 親の MTU           | no   | 全て                | 新しいインターフェースの MTU
`parent`  | string  | -                  | yes  | `physical`, `sriov` | ホスト上のデバイス、ブリッジの名前

`physical` な `infiniband` デバイスを作成するには次のように実行します:

```
lxc config device add <instance> <device-name> infiniband nictype=physical parent=<device>
```

## InfiniBand デバイスでの SR-IOV

InfiniBand デバイスは SR-IOV をサポートしますが、他の SR-IOV と違って、SR-IOV モードでの動的なデバイスの作成はできません。
つまり、カーネルモジュール側で事前に仮想ファンクション（virtual functions）の数を設定する必要があるということです。

`sriov` の `infiniband` でデバイスを作るには次のように実行します:

```
lxc config device add <instance> <device-name> infiniband nictype=sriov parent=<sriov-enabled-device>
```

(devices-infiniband)=
# タイプ: `infiniband`

```{youtube} https://www.youtube.com/watch?v=SDewhlRSOuM
:title: LXD InfiniBand devices - YouTube
```

```{note}
`infiniband`デバイスタイプはコンテナとVMの両方でサポートされます。
ホットプラグはコンテナのみでサポートし、VMではサポートしません。
```

LXDでは、InfiniBandデバイスに対する2種類の異なったネットワークタイプが使えます。

- `physical`: ホストの物理デバイスをインスタンスにパススルーします。
  対象のデバイスはホスト上では見えなくなり、インスタンス内に出現します。
- `sriov`: SR-IOVが有効な物理ネットワークデバイスの仮想ファンクション(virtual function)をインスタンスにパススルーします。

  ```{note}
  InfiniBandデバイスはSR-IOVをサポートしますが、他のSR-IOVが有効なデバイスと異なり、InfiniBandはSR-IOVモードの動的なデバイスの作成をサポートしません。
  このため、対応するカーネルモジュールを設定することで仮想ファンクションの数を事前に設定する必要があります。
  ```

`physical`な`infiniband`デバイスを作成するには、以下のコマンドを使用します。

    lxc config device add <instance_name> <device_name> infiniband nictype=physical parent=<device>

`sriov`の`infiniband`デバイスを作成するには、以下のコマンドを使用します。

    lxc config device add <instance_name> <device_name> infiniband nictype=sriov parent=<sriov_enabled_device>

## デバイスオプション

`infiniband`デバイスには以下のデバイスオプションがあります。

キー      | 型      | デフォルト値       | 必須 | 説明
:--       | :--     | :--                | :--  | :--
`hwaddr`  | string  | ランダムに割り当て | no   | 新しいインタフェースのMACアドレス。20バイト全てを指定するか短い8バイト(この場合親デバイスの最後の8バイトだけを変更)のどちらかを設定可能
`mtu`     | integer | 親の MTU           | no   | 新しいインタフェースのMTU
`name`    | string  | カーネルが割り当て | no   | インスタンス内部でのインタフェース名
`nictype` | string  | -                  | yes  | デバイスタイプ(`physical`か`sriov`のいずれか)
`parent`  | string  | -                  | yes  | ホスト上のデバイスまたはブリッジの名前

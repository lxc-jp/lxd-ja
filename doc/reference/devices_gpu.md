(devices-gpu)=
# タイプ: `gpu`

```{youtube} https://www.youtube.com/watch?v=T0aV2LsMpoA
```

GPUデバイスは、指定のGPUデバイスをインスタンス内に出現させます。

```{note}
コンテナでは、`gpu` デバイスは同時に複数のGPUにマッチングさせることができます。
VMでは、各デバイスは1つのGPUにしかマッチできません。
```

以下のタイプの GPU が `gputype` デバイスオプションを使って追加できます。

- [`physical`](#gpu-physical) (コンテナとVM): GPU全体をインスタンスにパススルーします。 
  `gputype` が指定されない場合これがデフォルトです。
- [`mdev`](#gpu-mdev) (VMのみ): 仮想GPUを作成しインスタンスにパススルーします。
- [`mig`](#gpu-mig) (コンテナのみ): MIG(Multi-Instance GPU)を作成しインスタンスにパススルーします。
- [`sriov`](#gpu-sriov) (VMのみ): SR-IOVを有効にしたGPUの仮想ファンクション(virtual function)をインスタンスに与えます。

利用可能なデバイスオプションはGPUタイプごとに異なり、以下のセクションの表に一覧表示されます。

(gpu-physical)=
## `gputype`: `physical`

```{note}
`physical` GPUタイプはコンテナとVMの両方でサポートされます。
ホットプラグはコンテナのみでサポートし、VMではサポートしません。
```

`physical` GPUデバイスはGPU全体をインスタンスにパススルーします。

### デバイスオプション

`physical` タイプのデバイスには以下のデバイスオプションがあります。

キー        | 型     | デフォルト値 | 説明
:--         | :--    | :--          | :--
`gid`       | int    | `0`          | インスタンス(コンテナのみ)内のデバイス所有者のGID
`id`        | string | -            | GPUデバイスのDRMカードID
`mode`      | int    | `0660`       | インスタンス(コンテナのみ)内のデバイスのモード
`pci`       | string | -            | GPUデバイスのPCIアドレス
`productid` | string | -            | GPUデバイスのプロダクトID
`uid`       | int    | `0`          | インスタンス(コンテナのみ)内のデバイス所有者のUID
`vendorid`  | string | -            | GPUデバイスのベンダーID

(gpu-mdev)=
## `gputype`: `mdev`

```{note}
`mdev` GPUタイプはVMでのみサポートされます。
ホットプラグはサポートしていません。
```

`mdev` GPUデバイスは仮想 GPU を作成しインスタンスにパススルーします。
利用可能な`mdev`プロファイルの一覧は `lxc info --resources` を実行すると確認できます。

### デバイスオプション

`mdev` タイプのデバイスには以下のデバイスオプションがあります。

キー        | 型     | デフォルト値 | 説明
:--         | :--    | :--          | :--
`id`        | string | -            | GPUデバイスのDRMカードID
`mdev`      | string | -            | 使用する`mdev`プロファイル(必須 - 例:`i915-GVTg_V5_4`)
`pci`       | string | -            | GPUデバイスのPCIアドレス
`productid` | string | -            | GPUデバイスのプロダクトID
`vendorid`  | string | -            | GPUデバイスのベンダーID

(gpu-mig)=
## `gputype`: `mig`

```{note}
`mig` GPUタイプはコンテナでのみサポートされます。
ホットプラグはサポートしていません。
```

`mig` GPUデバイスはMIGコンピュートインスタンスを作成しインスタンスにパススルーします。
現状これは NVIDIA MIG を事前に作成しておく必要があります。

### デバイスオプション

`mig` タイプのデバイスには以下のデバイスオプションがあります。

キー        | 型     | デフォルト値 | 説明
:--         | :--    | :--          | :--
`id`        | string | -            | GPUデバイスのDRMカードID
`mig.ci`    | int    | -            | 既存のMIGコンピュートインスタンスID
`mig.gi`    | int    | -            | 既存のMIG GPUインスタンスID
`mig.uuid`  | string | -            | 既存のMIGデバイスUUID(`MIG-`接頭辞は省略可)
`pci`       | string | -            | GPUデバイスのPCIアドレス
`productid` | string | -            | GPUデバイスのプロダクトID
`vendorid`  | string | -            | GPUデバイスのベンダーID

`mig.uuid`(NVIDIA drivers 470+)か、`mig.ci`と`mig.gi`(古いNVIDIAドライバ)の両方を設定する必要があります。

(gpu-sriov)=
## `gputype`: `sriov`

```{note}
`sriov` GPUタイプはVMでのみサポートされます。
ホットプラグはサポートしていません。
```

`sriov` GPUデバイスはSR-IOVが有効なGPUの仮想ファンクション(virtual function)をインスタンスにパススルーします。

### デバイスオプション

`sriov`タイプのデバイスには以下のデバイスオプションがあります。

キー        | 型     | デフォルト値 | 説明
:--         | :--    | :--          | :--
`id`        | string | -            | GPUデバイスのDRMカードID
`pci`       | string | -            | 親GPUデバイスのPCIアドレス
`productid` | string | -            | 親GPUデバイスのプロダクトID
`vendorid`  | string | -            | 親GPUデバイスのベンダーID

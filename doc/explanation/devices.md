(devices)=
# デバイス

デバイスはインスタンスにアタッチされます。
デバイスには、例えば、ネットワークインタフェース、マウントポイント、USB そして GPU デバイスがあります。
これらのデバイスはインスタンスデバイスの種別に応じてインスタンスデバイスオプションを持つことができます。

## 標準デバイス

LXD は、標準の POSIX システムが動作するのに必要な基本的なデバイスを常にインスタンスに提供します。これらはインスタンスやプロファイルの設定では見えず、上書きもできません。

このデバイスには次のようなデバイスが含まれます:

- `/dev/null` (キャラクターデバイス)
- `/dev/zero` (キャラクターデバイス)
- `/dev/full` (キャラクターデバイス)
- `/dev/console` (キャラクターデバイス)
- `/dev/tty` (キャラクターデバイス)
- `/dev/random` (キャラクターデバイス)
- `/dev/urandom` (キャラクターデバイス)
- `/dev/net/tun` (キャラクターデバイス)
- `/dev/fuse` (キャラクターデバイス)
- `lo` (ネットワークインタフェース)

これ以外に関しては、インスタンスの設定もしくはインスタンスで使われるいずれかのプロファイルで定義する必要があります。デフォルトのプロファイルには、インスタンス内で `eth0` になるネットワークインタフェースが通常は含まれます。

## デバイスを追加するには

インスタンスに追加でデバイスを追加する場合は、デバイスエントリーを直接インスタンスかプロファイルに追加できます。

デバイスはインスタンスの実行中に追加・削除できます。

各デバイスエントリーは一意な名前で識別されます。もし同じ名前が後続のプロファイルやインスタンス自身の設定で使われている場合、エントリー全体が新しい定義で上書きされます。

デバイスエントリーは次のようにインスタンスに追加するか:

```bash
lxc config device add <instance> <name> <type> [key=value]...
```

もしくは次のようにプロファイルに追加します:

```bash
lxc profile device add <profile> <name> <type> [key=value]...
```

(device-types)=
## デバイスタイプ

LXD では次のデバイスタイプが使えます:

| ID (データベース) | 名前                                 | 条件       | 説明                             |
|:------------------|:------------------------------------ |:-----------|:---------------------------------|
| 0                 | [`none`](#type-none)                 | -          | 継承ブロッカー                   |
| 1                 | [`nic`](#type-nic)                   | -          | ネットワークインタフェース     |
| 2                 | [`disk`](#type-disk)                 | -          | インスタンス内のマウントポイント |
| 3                 | [`unix-char`](#type-unix-char)       | コンテナ   | Unix キャラクターデバイス        |
| 4                 | [`unix-block`](#type-unix-block)     | コンテナ   | Unix ブロックデバイス            |
| 5                 | [`usb`](#type-usb)                   | -          | USB デバイス                     |
| 6                 | [`gpu`](#type-gpu)                   | -          | GPU デバイス                     |
| 7                 | [`infiniband`](#type-infiniband)     | コンテナ   | インフィニバンドデバイス         |
| 8                 | [`proxy`](#type-proxy)               | コンテナ   | プロキシデバイス                 |
| 9                 | [`unix-hotplug`](#type-unix-hotplug) | コンテナ   | Unix ホットプラグデバイス        |
| 10                | [`tpm`](#type-tpm)                   | -          | TPM デバイス                     |
| 11                | [`pci`](#type-pci)                   | 仮想マシン | PCI デバイス                     |

```{toctree}
:maxdepth: 1
:hidden:

../reference/devices_none.md
../reference/devices_nic.md
../reference/devices_disk.md
../reference/devices_unix_char.md
../reference/devices_unix_block.md
../reference/devices_usb.md
../reference/devices_gpu.md
../reference/devices_infiniband.md
../reference/devices_proxy.md
../reference/devices_unix_hotplug.md
../reference/devices_tpm.md
../reference/devices_pci.md
```

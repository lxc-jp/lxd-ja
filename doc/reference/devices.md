(devices)=
# デバイス

デバイスはインスタンス ({ref}`instances-configure-devices` 参照) またはプロファイル ({ref}`profiles-edit` 参照) にアタッチされます。

デバイスには、例えば、ネットワークインタフェース、マウントポイント、USB そして GPU デバイスがあります。
これらのデバイスはインスタンスデバイスの種別に応じてインスタンスデバイスオプションを持つことができます。

LXD では次のデバイスタイプが使えます。

| ID (データベース)   | 名前                                  | 条件         | 説明                               |
| :------------------ | :------------------------------------ | :----------- | :--------------------------------- |
| 0                   | [`none`](devices-none)                  | -            | 継承ブロッカー                     |
| 1                   | [`nic`](devices-nic)                    | -            | ネットワークインタフェース         |
| 2                   | [`disk`](devices-disk)                  | -            | インスタンス内のマウントポイント   |
| 3                   | [`unix-char`](devices-unix-char)        | コンテナ     | Unix キャラクタデバイス            |
| 4                   | [`unix-block`](devices-unix-block)      | コンテナ     | Unix ブロックデバイス              |
| 5                   | [`usb`](devices-usb)                    | -            | USB デバイス                       |
| 6                   | [`gpu`](devices-gpu)                    | -            | GPU デバイス                       |
| 7                   | [`infiniband`](devices-infiniband)      | コンテナ     | インフィニバンドデバイス           |
| 8                   | [`proxy`](devices-proxy)                | コンテナ     | プロキシデバイス                   |
| 9                   | [`unix-hotplug`](devices-unix-hotplug)  | コンテナ     | Unix ホットプラグデバイス          |
| 10                  | [`tpm`](devices-tpm)                    | -            | TPM デバイス                       |
| 11                  | [`pci`](devices-pci)                    | 仮想マシン   | PCI デバイス                       |

各インスタンスには一組の {ref}`standard-devices` が付属します。

```{toctree}
:maxdepth: 1
:hidden:

../reference/standard_devices.md
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

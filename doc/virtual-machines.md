# 仮想マシン <!-- Virtual Machines -->
## イントロダクション <!-- Introduction -->
仮想マシンはコンテナーと共に LXD でサポートされる新しいインスタンスタイプです。
<!--
Virtual machines are a new instance type supported by LXD alongside containers.
-->

仮想マシンは `qemu` を使って実装されています。
<!--
They are implemented through the use of `qemu`.
-->

現状はコンテナーで利用可能な全ての機能が VM には実装されているわけではないことにご注意ください。
しかし、私達はコンテナーと同等の機能を目指して引き続き努力します。
<!--
Please note, currently not all features that are available with containers have been implemented for VMs,
however we continue to strive for feature parity with containers.
-->

## 設定 <!-- Configuration -->
有効な設定項目については [インスタンスの設定](instances.md) を参照してください。
<!--
See [instance configuration](instances.md) for valid configuration options.
-->

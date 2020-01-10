# 仮想マシン <!-- Virtual Machines -->
## イントロダクション <!-- Introduction -->
仮想マシンはコンテナと共に LXD でサポートされる新しいインスタンスタイプです。
<!--
Virtual machines are a new instance type supported by LXD alongside containers.
-->

仮想マシンは `qemu` を使って実装されています。
<!--
They are implemented through the use of `qemu`.
-->

仮想マシンの機能は現状は実験的 (experimental) です。
コンテナと同程度の機能に達するには多くの機能を今後実装していく必要があります。
<!--
This feature is currently considered to be experimental with a lot of
functionality still yet to be implemented in order to reach feature
parity with containers.
-->

## 設定 <!-- Configuration -->
有効な設定項目については [インスタンスの設定](instances.md) を参照してください。
<!--
See [instance configuration](instances.md) for valid configuration options.
-->

---
discourse: 8767,7519,9281
relatedlinks: https://ubuntu.com/blog/lxd-virtual-machines-an-overview
---

(expl-instances)=
# インスタンスについて

LXD 以下のインスタンスタイプをサポートします。

コンテナ
: コンテナはデフォルトのインスタンスタイプです。
  コンテナは現状 LXD インスタンスの最も完全な実装であり、仮想マシンよりも多くの機能をサポートしています。

  コンテナは `liblxc` (LXC) を使って実装されています。

仮想マシン
: {abbr}`Virtual machines (VMs)` は LXD バージョン 4.0 以降ネイティブにサポートされています。
  ビルトインのエージェントのおかげで、ほぼコンテナと同様に使えます。

  LXD は仮想マシンの機能を提供するために `qemu` を使用しています。

  ```{note}
  現状、仮想マシンはコンテナよりサポートする機能が少ないですが、将来には両方のインスタンスタイプで同じ機能セットをサポートする計画です。

  仮想マシンでどの機能が利用可能かを見るには、{ref}`instance-options` ドキュメントの条件のカラムを確認してください。
  ```
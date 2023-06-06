(migration)=
# マイグレーション

```{youtube} https://www.youtube.com/watch?v=F9GALjHtnUU
```

LXD は異なる状況でインスタンスをマイグレートするためのツールと機能を提供します。

サーバー間で既存の LXD インスタンスをマイグレートする
: 最も基本的な種類のマイグレーションは 1 つのサーバー上に LXD インスタンスがあり、それを別の LXD サーバーに移動したいというものです。
  仮想マシンでは、ライブマイグレーションを行えます。これは稼働中にダウンタイムなしで VM をマイグレートできることを意味します。

  詳細は {ref}`move-instances` を参照してくだい。

物理または仮想マシンを LXD インスタンスにマイグレートする
: 物理または仮想(VMまたはコンテナ)の既存のマシンがある場合、既存のマシン上に LXD インスタンスを作成するために `lxd-migrate` ツールが使えます。
  このツールは提供されたパーティション、ディスクやイメージを LXD サーバーの LXD ストレージプールにコピーし、そのストレージを使ってインスタンスをセットアップします。新しいインスタンスの追加の設定を行うこともできます。

  詳細は {ref}`import-machines-to-instances` を参照してくだい。

LXD から LXD へインスタンスをマイグレートする
: LXC を使っていて全てまたは一部の LXC コンテナを同じマシン上の LXD に移動したい場合、 `lxc-to-lxd` ツールが使えます。
  このツールは LXC 設定を解析し、既存の LXC コンテナのデータと設定を新しい LXD コンテナにコピーします。

  詳細は {ref}`migrate-from-lxc` を参照してくだい。

```{toctree}
:maxdepth: 1
:hidden:

インスタンスの移動 <howto/move_instances>
既存のマシンのインポート <howto/import_machines_to_instances>
LXCからのマイグレート <howto/migrate_from_lxc>
```

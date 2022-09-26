---
discourse: 14579
relatedlinks: https://youtube.com/watch?v=kVLGbvRU98A
---

(storage-cephobject)=
# Ceph Object - `cephobject`

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph intro -->
    :end-before: <!-- Include end Ceph intro -->
```

[Ceph Object Gateway](https://docs.ceph.com/en/latest/radosgw/) は [`librados`](https://docs.ceph.com/en/latest/rados/api/librados-intro/) 上に構築されたオブジェクトストレージインタフェースであり [Ceph Storage Clusters](https://docs.ceph.com/en/latest/rados/) への RESTful ゲートウェイを持つアプリケーションを提供します。
Amazon S3 RESTful API の大きなサブセットと互換性を持つオブジェクトストレージの機能を提供します。

## 用語

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph terminology -->
    :end-before: <!-- Include end Ceph terminology -->
```

*Ceph Object Gateway* は複数の OSD プールとゲートウェイの機能を提供する 1 つ以上の *Ceph Object Gateway daemon* (`radosgw`) プロセスから構成されます。

## LXD の `cephobject` ドライバ

```{note}
`cephobject` ドライバはバケットのみに使用できます。

ストレージボリュームには {ref}`Ceph <storage-ceph>` または {ref}`CephFS <storage-cephfs>` ドライバを使用してください。
```

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph driver cluster -->
    :end-before: <!-- Include end Ceph driver cluster -->
```

事前に `radosgw` 環境をセットアップし、 HTTP/HTTPS エンドポイント URL が LXD サーバからアクセス可能なことを確認してください。
Ceph クラスタをどのようにセットアップするかの情報については [Manual Deployment](https://docs.ceph.com/en/latest/install/manual-deployment/) を、そして `radosgw` 環境をどのようにセットアップするかについては [Ceph Object Gateway](https://docs.ceph.com/en/latest/radosgw/) を参照してください。

`radosgw` URL はプールの作成時に [`cephobject.radosgw.endpoint`](storage-cephobject-pool-config) オプションを使って指定できます。
また LXD はバケットの管理に `radosgw-admin` コマンドを使用しています。ですのでこのコマンドが LXD サーバ上で利用可能で操作可能である必要があります。

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph driver remote -->
    :end-before: <!-- Include end Ceph driver remote -->
```

% Include content from [storage_ceph.md](storage_ceph.md)
```{include} storage_ceph.md
    :start-after: <!-- Include start Ceph driver control -->
    :end-before: <!-- Include end Ceph driver control -->
```

## 設定オプション

`cephobject` ドライバを使うストレージプールとこれらのプール内のストレージボリュームには以下の設定オプションが利用できます。

(storage-cephobject-pool-config)=
### ストレージプール設定

キー                                    | 型     | デフォルト値 | 説明
:--                                     | :---   | :------      | :----------
`cephobject.bucket.name_prefix`         | string | -            | Ceph 内のバケット名に追加する接頭辞
`cephobject.cluster_name`               | string | `ceph`       | 使用する Ceph クラスタ
`cephobject.radosgw.endpoint`           | string | -            | `radosgw` ゲートウェイプロセスのURL
`cephobject.radosgw.endpoint_cert_file` | string | -            | エンドポイント通信に使用する TLS クライアント証明書を含むファイルへのパス
`cephobject.user.name`                  | string | `admin`      | 使用する Ceph ユーザ
`volatile.pool.pristine`                | string | `true`       | 作成時に `radosgw` `lxd-admin` ユーザが存在したかどうか

### ストレージバケット設定

キー   | 型     | デフォルト値 | 説明
:--    | :---   | :------      | :----------
`size` | string | -            | ストレージバケットのクォータ

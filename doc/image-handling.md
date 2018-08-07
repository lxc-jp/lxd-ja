# イントロダクション <!-- Introduction -->
LXD はイメージをベースとしたワークフローを使用します。 LXD にはビルトイン
のイメージ・ストアがあり、ユーザが外部のツールがそこからイメージをインポート
できます。
<!--
LXD uses an image based workflow. It comes with a built-in image store
where the user or external tools can import images.
-->

その後、それらのイメージからコンテナが起動されます。
<!--
Containers are then started from those images.
-->

ローカルのイメージを使ってリモートのコンテナを起動できますし、リモートの
イメージを使ってローカルのコンテナを起動することもできます。こういった
ケースではイメージはターゲットの LXD にキャッシュされます。
<!--
It's possible to spawn remote containers using local images or local
containers using remote images. In such cases, the image may be cached
on the target LXD.
-->

# キャッシュ <!-- Caching -->
リモートのイメージからコンテナを起動する時、リモートのイメージが
ローカルのイメージ・ストアにキャッシュ・ビットをセットした状態で
ダウンロードされます。イメージは、 `images.remote_cache_expiry` に
設定された日数だけ使われない (新たなコンテナが起動されない) か、
イメージが期限を迎えるか、どちらか早いほうが来るまで、プライベートな
イメージとしてローカルに保存されます。
<!--
When spawning a container from a remote image, the remote image is
downloaded into the local image store with the cached bit set. The image
will be kept locally as a private image until either it's been unused
(no new container spawned) for the number of days set in
`images.remote_cache_expiry` or until the image's expiry is reached
whichever comes first.
-->

LXD はイメージから新しいコンテナが起動される度にイメージの `last_used_at` 
プロパティを更新することで、イメージの利用状況を記録しています。
<!--
LXD keeps track of image usage by updating the `last_used_at` image
property every time a new container is spawned from the image.
-->

# 自動更新 <!-- Auto-update -->
LXD はイメージを最新に維持できます。デフォルトではエイリアスで指定し
リモートサーバから取得したイメージは LXD によって自動更新されます。
これは `images.auto_update_cached` という設定で変更できます。
<!--
LXD can keep images up to date. By default, any image which comes from a
remote server and was requested through an alias will be automatically
updated by LXD. This can be changed with `images.auto_update_cached`.
-->

(`images.auto_update_interval` が設定されない限り) 起動時とその後
6 時間毎に、 LXD デーモンはイメージ・ストア内で自動更新対象となっていて
ダウンロード元のサーバが記録されている全てのイメージのより新しい
バージョンがあるかを確認します。
<!--
On startup and then every 6 hours (unless `images.auto_update_interval`
is set), the LXD daemon will go look for more recent version of all the
images in the store which are marked as auto-update and have a recorded
source server.
-->

新しいイメージが見つかったら、イメージ・ストアにダウンロードされ、
古いイメージを指していたエイリアスは新しいイメージを指すように変更され、
古いイメージはストアから削除されます。
<!--
When a new image is found, it is downloaded into the image store, the
aliases pointing to the old image are moved to the new one and the old
image is removed from the store.
-->

リモート・サーバからイメージを手動でコピーする際に、特定のイメージを
最新に維持するように設定することもできます。
<!--
The user can also request a particular image be kept up to date when
manually copying an image from a remote server.
-->

ユーザがイメージのキャッシュから新しいコンテナを作成しようとした時に、
アップストリームの新しいイメージ更新が公開されており、ローカルの LXD が
キャッシュに古いイメージを持っている場合は、 LXD はコンテナの作成を
遅らせるのではなく、古いバージョンのイメージを使います。
<!--
If a new upstream image update is published and the local LXD has the
previous image in its cache when the user requests a new container to be
created from it, LXD will use the previous version of the image rather
than delay the container creation.
-->

この振る舞いは現在のイメージが自動更新されるように設定されている時のみに
発生し、 `images.auto_update_interval` を 0 にすることで無効にできます。
<!--
This behavior only happens if the current image is scheduled to be
auto-updated and can be disabled by setting `images.auto_update_interval` to 0.
-->

# イメージの形式 <!-- Image format -->
LXD は現状 2 つの LXD に特有なイメージの形式をサポートします。
<!--
LXD currently supports two LXD-specific image formats.
-->

1 つめは統合された tarball で、単一の tarball がコンテナの rootfs と
必要なメタデータの両方を含みます。
<!--
The first is a unified tarball, where a single tarball
contains both the container rootfs and the needed metadata.
-->

2 つめは分離されたモデルで、 2 つの tarball を使い、 1 つは rootfs を
含み、もう一つはコンテナのメタデータを含みます。
<!--
The second is a split model, using two tarballs instead, one containing
the rootfs, the other containing the metadata.
-->

LXD 自身によって生成されるのは前者の形式で、 LXD 特有のイメージを使う
際はこちらの形式を使うべきです。
<!--
The former is what's produced by LXD itself and what people should be
using for LXD-specific images.
-->

後者は、今日既に利用可能なものとして存在している LXD 以外の rootfs tarball 
を使ってイメージを簡単に作成できるように想定されているものです。
<!--
The latter is designed to allow for easy image building from existing
non-LXD rootfs tarballs already available today.
-->

## 統合された tarball <!-- Unified tarball -->
tarball は圧縮できます。そして次のものを含みます。
<!--
Tarball, can be compressed and contains:
-->

 - `rootfs/`
 - `metadata.yaml`
 - `templates/` (省略可能) <!-- (optional) -->

このモードではイメージの識別子は tarball の SHA-256 です。
<!--
In this mode, the image identifier is the SHA-256 of the tarball.
-->

## 分離された tarball <!-- Split tarballs -->
2 つの (圧縮しても良い) tarball 。 1 つはメタデータ、もう 1 つは rootfs です。
<!--
Two (possibly compressed) tarballs. One for metadata, one for the rootfs.
-->

`metadata.tar` は以下のものを含みます。
<!--
`metadata.tar` contains:
-->

 - `metadata.yaml`
 - `templates/` (省略可能) <!-- (optional) -->

`rootfs.tar` は、そのルートに Linux の root ファイルシステムを含みます。
<!--
`rootfs.tar` contains a Linux root filesystem at its root.
-->

このモードではイメージの識別子はメタデータと rootfs の tarball を(この順番で)
結合したものの SHA-256 です。
<!--
In this mode the image identifier is the SHA-256 of the concatenation of
the metadata and rootfs tarball (in that order).
-->

## サポートされている圧縮形式 <!-- Supported compression -->
tarball は bz2, gz, xz, lzma, tar (非圧縮) で圧縮することができ、あるいは
squashfs のイメージでも構いません。
<!--
The tarball(s) can be compressed using bz2, gz, xz, lzma, tar (uncompressed) or
it can also be a squashfs image.
-->

## 中身 <!-- Content -->
rootfs のディレクトリ (あるいは tarball) は完全なファイルシステムのツリーを含み、
それがコンテナの `/` になります。
<!--
The rootfs directory (or tarball) contains a full file system tree of what will become the container's `/`.
-->

テンプレートのディレクトリはコンテナ内で使用される pongo2 形式のテンプレート・ファイルを含みます。
<!--
The templates directory contains pongo2-formatted templates of files inside the container.
-->

`metadata.yaml` はイメージを (現状は) LXD で稼働されるために必要な情報を
含んでおり、これは以下のものを含みます。
<!--
`metadata.yaml` contains information relevant to running the image under
LXD, at the moment, this contains:
-->

```yaml
architecture: x86_64
creation_date: 1424284563
properties:
  description: Ubuntu 14.04 LTS Intel 64bit
  os: Ubuntu
  release:
    - trusty
    - 14.04
templates:
  /etc/hosts:
    when:
      - create
      - rename
    template: hosts.tpl
    properties:
      foo: bar
  /etc/hostname:
    when:
      - start
    template: hostname.tpl
  /etc/network/interfaces:
    when:
      - create
    template: interfaces.tpl
    create_only: true
```

`architecture` と `creation_date` の項目は必須です。 `properties` は
単にイメージのデフォルト・プロパティの組です。 `os`, `release`, `name`
と `description` の項目は必須ではないですが、記載されることが多いでしょう。
<!--
The `architecture` and `creation_date` fields are mandatory, the properties
are just a set of default properties for the image. The os, release,
name and description fields while not mandatory in any way, should be
pretty common.
-->

テンプレートで `when` キーは以下の 1 つあるいは複数が指定可能です。
<!--
For templates, the `when` key can be one or more of:
-->

 - `create` (そのイメージから新しいコンテナが作成されたときに実行される) <!-- (run at the time a new container is created from the image) -->
 - `copy` (既存のコンテナから新しいコンテナが作成されたときに実行される) <!-- (run when a container is created from an existing one) -->
 - `start` (コンテナが開始される度に実行される) <!-- (run every time the container is started) -->

テンプレートは常に以下のコンテキストを受け取ります。
<!--
The templates will always receive the following context:
-->

 - `trigger`: テンプレートを呼び出したイベントの名前 <!-- name of the event which triggered the template --> (string)
 - `path`: テンプレート出力先のファイルのパス <!-- path of the file being templated --> (string)
 - `container`: コンテナのプロパティ (name, architecture, privileged そして ephemeral) の key/value の map <!-- key/value map of container properties (name, architecture, privileged and ephemeral) --> (map[string]string)
 - `config`: コンテナの設定の key/value の map <!-- key/value map of the container's configuration --> (map[string]string)
 - `devices`: コンテナに割り当てられたデバイスの key/value の map <!-- key/value map of the devices assigned to this container --> (map[string]map[string]string)
 - `properties`: metadata.yaml に指定されたテンプレートのプロパティの key/value の map <!-- key/value map of the template properties specified in metadata.yaml --> (map[string]string)

`create_only` キーを設定すると LXD が存在しないファイルだけを生成し、
既存のファイルを上書きしないようにできます。
<!--
The `create_only` key can be set to have LXD only only create missing files but not overwrite an existing file.
-->

一般的な規範として、パッケージで管理されているファイルをテンプレートの
生成対象とすべきではないです。そうしてしまうとコンテナの通常の操作で
上書きされてしまうでしょう。
<!--
As a general rule, you should never template a file which is owned by a
package or is otherwise expected to be overwritten by normal operation
of the container.
-->

利便性のため、以下の関数が pongo のテンプレートで利用可能となっています。
<!--
For convenience the following functions are exported to pongo templates:
-->

 - `config_get("user.foo", "bar")` => `user.foo` の値か、未設定の場合は `"bar"` を返します。 <!-- Returns the value of `user.foo` or `"bar"` if unset. -->

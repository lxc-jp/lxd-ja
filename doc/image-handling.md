# イメージの扱い
<!-- Image handling -->
## イントロダクション <!-- Introduction -->
LXD はイメージをベースとしたワークフローを使用します。 LXD にはビルトイン
のイメージ・ストアがあり、ユーザーが外部のツールがそこからイメージをインポート
できます。
<!--
LXD uses an image based workflow. It comes with a built-in image store
where the user or external tools can import images.
-->

その後、それらのイメージからコンテナーが起動されます。
<!--
Containers are then started from those images.
-->

ローカルのイメージを使ってリモートのインスタンスを起動できますし、リモートの
イメージを使ってローカルのインスタンスを起動することもできます。こういった
ケースではイメージはターゲットの LXD にキャッシュされます。
<!--
It's possible to spawn remote instances using local images or local
instances using remote images. In such cases, the image may be cached
on the target LXD.
-->

## キャッシュ <!-- Caching -->
リモートのイメージからインスタンスを起動する時、リモートのイメージが
ローカルのイメージ・ストアにキャッシュ・ビットをセットした状態で
ダウンロードされます。イメージは、 `images.remote_cache_expiry` に
設定された日数だけ使われない (新たなインスタンスが起動されない) か、
イメージが期限を迎えるか、どちらか早いほうが来るまで、プライベートな
イメージとしてローカルに保存されます。
<!--
When spawning an instance from a remote image, the remote image is
downloaded into the local image store with the cached bit set. The image
will be kept locally as a private image until either it's been unused
(no new instance spawned) for the number of days set in
`images.remote_cache_expiry` or until the image's expiry is reached
whichever comes first.
-->

LXD はイメージから新しいインスタンスが起動される度にイメージの `last_used_at` 
プロパティを更新することで、イメージの利用状況を記録しています。
<!--
LXD keeps track of image usage by updating the `last_used_at` image
property every time a new instance is spawned from the image.
-->

## 自動更新 <!-- Auto-update -->
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

ユーザーがイメージのキャッシュから新しいインスタンスを作成しようとした時に、
アップストリームの新しいイメージ更新が公開されており、ローカルの LXD が
キャッシュに古いイメージを持っている場合は、 LXD はインスタンスの作成を
遅らせるのではなく、古いバージョンのイメージを使います。
<!--
If a new upstream image update is published and the local LXD has the
previous image in its cache when the user requests a new instance to be
created from it, LXD will use the previous version of the image rather
than delay the instance creation.
-->

この振る舞いは現在のイメージが自動更新されるように設定されている時のみに
発生し、 `images.auto_update_interval` を 0 にすることで無効にできます。
<!--
This behavior only happens if the current image is scheduled to be
auto-updated and can be disabled by setting `images.auto_update_interval` to 0.
-->

## プロファイル <!-- Profiles -->
`lxc image edit` コマンドを使ってイメージにプロファイルのリストを関連付けできます。
イメージにプロファイルを関連付けた後に起動したインスタンスはプロファイルを順番に適用します。
プロファイルのリストとして `nil` を指定すると `default` プロファイルのみがイメージに関連付けされます。
空のリストを指定すると、 `default` プロファイルも含めて一切のプロファイルをイメージに適用しません。
イメージに関連付けされたプロファイルは `lxc launch` の `--profile` と `--no-profiles` オプションを使ってインスタンス起動時にオーバーライドできます。
<!--
A list of profiles can be associated with an image using the `lxc image edit`
command. After associating profiles with an image, an instance launched
using the image will have the profiles applied in order. If `nil` is passed
as the list of profiles, only the `default` profile will be associated with 
the image. If an empty list is passed, then no profile will be associated
with the image, not even the `default` profile. An image's associated
profiles can be overridden when launching an instance by using the 
`-\-profile` and the `-\-no-profiles` flags to `lxc launch`.
-->

## イメージの形式 <!-- Image format -->
LXD は現状 2 つの LXD に特有なイメージの形式をサポートします。
<!--
LXD currently supports two LXD-specific image formats.
-->

1 つめは統合された tarball で、単一の tarball がインスタンスの root と
必要なメタデータの両方を含みます。
<!--
The first is a unified tarball, where a single tarball
contains both the instance root and the needed metadata.
-->

2 つめは分離されたモデルで、 2 つのファイルを使い、 1 つは root を
含み、もう一つはメタデータを含みます。
<!--
The second is a split model, using two files instead, one containing
the root, the other containing the metadata.
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

### 統合された tarball <!-- Unified tarball -->
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

### 分離された tarball <!-- Split tarballs -->
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

### サポートされている圧縮形式 <!-- Supported compression -->
tarball は bz2, gz, xz, lzma, tar (非圧縮) で圧縮することができ、あるいは
squashfs のイメージでも構いません。
<!--
The tarball(s) can be compressed using bz2, gz, xz, lzma, tar (uncompressed) or
it can also be a squashfs image.
-->

### 中身 <!-- Content -->
コンテナーでは rootfs のディレクトリ (あるいは tarball) は完全なファイルシステムのツリーを含み、それが `/` になります。
VM ではこれは代わりに `root.img` ファイルでメインのディスクデバイスになります。
<!--
For containers, the rootfs directory (or tarball) contains a full file system tree of what will become the `/`.
For VMs, this is instead a `root.img` file which becomes the main disk device.
-->

テンプレートのディレクトリはコンテナー内で使用される pongo2 形式のテンプレート・ファイルを含みます。
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
  description: Ubuntu 18.04 LTS Intel 64bit
  os: Ubuntu
  release: bionic 18.04
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

 - `create` (そのイメージから新しいインスタンスが作成されたときに実行される) <!-- (run at the time a new instance is created from the image) -->
 - `copy` (既存のインスタンスから新しいインスタンスが作成されたときに実行される) <!-- (run when an instance is created from an existing one) -->
 - `start` (インスタンスが開始される度に実行される) <!-- (run every time the instance is started) -->

テンプレートは常に以下のコンテキストを受け取ります。
<!--
The templates will always receive the following context:
-->

 - `trigger`: テンプレートを呼び出したイベントの名前 <!-- name of the event which triggered the template --> (string)
 - `path`: テンプレート出力先のファイルのパス <!-- path of the file being templated --> (string)
 - `container`: インスタンスのプロパティ (name, architecture, privileged そして ephemeral) の key/value の map <!-- key/value map of instance properties (name, architecture, privileged and ephemeral) --> (map[string]string) (廃止予定。代わりに `instance` を使用してください) <!-- (deprecated in favor of `instance`) -->
 - `instance`: インスタンスのプロパティ (name, architecture, privileged そして ephemeral) の key/value の map <!-- key/value map of instance properties (name, architecture, privileged and ephemeral) --> (map[string]string)
 - `config`: インスタンスの設定の key/value の map <!-- key/value map of the instance's configuration --> (map[string]string)
 - `devices`: インスタンスに割り当てられたデバイスの key/value の map <!-- key/value map of the devices assigned to this instance --> (map[string]map[string]string)
 - `properties`: metadata.yaml に指定されたテンプレートのプロパティの key/value の map <!-- key/value map of the template properties specified in metadata.yaml --> (map[string]string)

`create_only` キーを設定すると LXD が存在しないファイルだけを生成し、
既存のファイルを上書きしないようにできます。
<!--
The `create_only` key can be set to have LXD only only create missing files but not overwrite an existing file.
-->

一般的な規範として、パッケージで管理されているファイルをテンプレートの
生成対象とすべきではないです。そうしてしまうとインスタンスの通常の操作で
上書きされてしまうでしょう。
<!--
As a general rule, you should never template a file which is owned by a
package or is otherwise expected to be overwritten by normal operation
of the instance.
-->

利便性のため、以下の関数が pongo のテンプレートで利用可能となっています。
<!--
For convenience the following functions are exported to pongo templates:
-->

 - `config_get("user.foo", "bar")` => `user.foo` の値か、未設定の場合は `"bar"` を返します。 <!-- Returns the value of `user.foo` or `"bar"` if unset. -->

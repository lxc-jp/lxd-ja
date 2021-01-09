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

## ソース <!-- Sources -->
LXD は 3 つの異なるソースからのイメージのインポートをサポートします。
<!--
LXD supports importing images from three different sources:
-->

 - リモートのイメージサーバー (LXD か simplestreams) <!-- Remote image server (LXD or simplestreams) -->
 - イメージファイルの direct push <!-- Direct pushing of the image files -->
 - リモートのウェブサーバー上のファイル <!-- File on a remote web server -->

### リモートのイメージサーバー (LXD か simplestreams) <!-- Remote image server (LXD or simplestreams) -->
これは最も一般的なイメージソースで 3 つの選択肢のうちインスタンスの作成時に直接サポートされている唯一の選択肢です。
<!--
This is the most common source of images and the only one of the three
options which is supported directly at instance creation time.
-->

この選択肢では、イメージサーバーは検証されるために必要な証明書（HTTPS のみがサポートされます）と共にターゲットの LXD サーバーに提供されます。
<!--
With this option, an image server is provided to the target LXD server
along with any needed certificate to validate it (only HTTPS is supported).
-->

次にイメージそのものがフィンガープリント (SHA256) あるいはエイリアスによって選択されます。
<!--
The image itself is then selected either by its fingerprint (SHA256) or
one of its aliases.
-->

CLI の視点では、これは以下の一般的なアクションによって実行されます。
<!--
From a CLI point of view, this is what's done behind those common actions:
-->

 - lxc launch ubuntu:20.04 u1
 - lxc launch images:centos/8 c1
 - lxc launch my-server:SHA256 a1
 - lxc image copy images:gentoo local: --copy-aliases --auto-update

上記の `ubuntu` と `images` のケースではリモートは simplestreams を読み取り専用のサーバープロトコルとして使用し、イメージの複数のエイリアスの 1 つによりイメージを選択します。
<!--
In the cases of `ubuntu` and `images` above, those remotes use
simplestreams as a read-only image server protocol and select images by
one of their aliases.
-->

`my-server` リモートのケースでは別の LXD サーバーがあり、上記の例ではフィンガープリントによってイメージを選択します。
<!--
The `my-server` remote there is another LXD server and in that example
selects an image based on its fingerprint.
-->

### イメージファイルの direct push <!-- Direct pushing of the image files -->
これは主に外部サーバーから直接イメージを取得できない隔離された環境で有用です。
<!--
This is mostly useful for air-gapped environments where images cannot be
directly retrieved from an external server.
-->

そのような状況ではイメージファイルは他のシステムで以下のコマンドを使ってダウンロードできます。
<!--
In such a scenario, image files can be downloaded on another system using:
-->

 - lxc image export ubuntu:20.04

その後ターゲットのシステムにイメージを転送してローカルイメージストアーに手動でインポートします。
<!--
Then transferred to the target system and manually imported into the
local image store with:
-->

 - lxc image import META ROOTFS --alias ubuntu-20.04

`lxc image import` は統合イメージ (単一ファイル) と分割イメージ (2つのファイル) の両方をサポートします。
上の例では後者を使用しています。
<!--
`lxc image import` supports both unified images (single file) and split
images (two files) with the example above using the latter.
-->

### リモートのウェブサーバー上のファイル <!-- File on a remote web server -->
単一のイメージをユーザーに配布するためだけにフルのイメージサーバーを動かすことの代替として、 LXD は URL を指定してイメージをインポートするのもサポートしています。
<!--
As an alternative to running a full image server only to distribute a
single image to users, LXD also supports importing images by URL.
-->

ただし、この方法にはいくつか制限があります。
<!--
There are a few limitations to that method though:
-->

 - 統合ファイル（単一ファイル）のみがサポートされます <!-- Only unified (single file) images are supported -->
 - リモートサーバーが追加の http ヘッダーを返す必要があります <!-- Additional http headers must be returned by the remote server -->

LXD はサーバーに問い合わせをする際に以下のヘッダーを設定します。
<!--
LXD will set the following headers when querying the server:
-->

 - `LXD-Server-Architectures` にはクライアントがサポートするアーキテクチャーのカンマ区切りリストを設定します <!-- `LXD-Server-Architectures` to a comma separate list of architectures the client supports -->
 - `LXD-Server-Version` には使用している LXD のバージョンを設定します <!-- `LXD-Server-Version` to the version of LXD in use -->


リモートサーバーが `LXD-Image-Hash` と `LXD-Image-URL` を設定することを期待します。
前者はダウンロードされるイメージの SHA256 ハッシュで後者はイメージをダウンロードする URL です。
<!--
And expects `LXD-Image-Hash` and `LXD-Image-URL` to be set by the remote server.
The former being the SHA256 of the image being downloaded and the latter
the URL to download the image from.
-->

これによりかなり複雑なイメージサーバーがカスタムヘッダーをサポートする基本的なウェブサーバーだけで実装できます。
<!--
This allows for reasonably complex image servers to be implemented using
only a basic web server with support for custom headers.
-->

クライアント側では以下のように使用できます。
<!--
On the client side, this is used with:
-->

`lxc image import URL --alias some-name`

### インスタンスやスナップショットを新しいイメージとして公開する <!-- Publishing an instance or snapshot as a new image -->
インスタンスやスナップショットの 1 つを新しいイメージに変換できます。
これは `lxc publish` で CLI 上で実行できます。
<!--
An instance or one of its snapshots can be turned into a new image.
This is done on the CLI with `lxc publish`.
-->

これを行う際には、たいていの場合公開する前にインスタンスのメタデータやテンプレートを `lxc config metadata` と `lxc config template` コマンドを使って整理するのが良いでしょう。
さらにホストの SSH キーや dbus/systemd の machine-id などインスタンスに固有な状態も削除するのが良いでしょう。
<!--
When doing this, you will most likely first want to cleanup metadata and
templates on the instance you're publishing using the `lxc config metadata`
and `lxc config template` commands. You will also want to remove any
instance-specific state like host SSH keys, dbus/systemd machine-id, ...
-->

インスタンスから tarball を生成した後圧縮する必要があるので、公開のプロセスはかなり時間がかかるかもしれません。
この操作は特に I/O と CPU の負荷が高いので、公開操作は LXD により 1 つずつ順に実行されます。
<!--
The publishing process can take quite a while as a tarball must be
generated from the instance and then be compressed. As this can be
particularly I/O and CPU intensive, publish operations are serialized by LXD.
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
LXD は広範な tarball の圧縮アルゴリズムをサポートしますが、互換性のために gzip か xz が望ましいです。
<!--
LXD supports a wide variety of compression algorithms for tarballs
though for compatibility purposes, gzip or xz should be preferred.
-->

分離されたイメージではコンテナーの場合は rootfs ファイルはさらに squashfs 形式でフォーマットすることもできます。
仮想マシンでは `rootfs.img` ファイルは常に qcow2 であり、オプションで qcow2 のネイティブ圧縮を使って圧縮することもできます。
<!--
For split images, the rootfs file can also be squashfs formatted in the
container case. For virtual machines, the `rootfs.img` file is always
qcow2 and can optionally be compressed using qcow2's native compression.
-->

### 中身 <!-- Content -->
コンテナーでは rootfs のディレクトリ (あるいは tarball) は完全なファイルシステムのツリーを含み、それが `/` になります。
VM ではこれは代わりに `rootfs.img` ファイルでメインのディスクデバイスになります。
<!--
For containers, the rootfs directory (or tarball) contains a full file system tree of what will become the `/`.
For VMs, this is instead a `rootfs.img` file which becomes the main disk device.
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

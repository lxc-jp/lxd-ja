(image-format)=
# イメージ形式

イメージはルートファイルシステムとイメージを記述するメタデータファイルを含みます。
またイメージを使用するインスタンス内部でファイルを生成するためのテンプレートも含められます。

イメージは統合イメージ(単一ファイル)か分離イメージ(2つのファイル)としてパッケージできます。

## 中身

コンテナのイメージは以下のディレクトリ構造を持ちます。

```
metadata.yaml
rootfs/
templates/
```

VMのイメージは以下のディレクトリ構造を持ちます。

```
metadata.yaml
rootfs.img
templates/
```

どちらのインスタンスタイプでも、`templates/`ディレクトリは省略可能です。

### メタデータ

`metadata.yaml`ファイルはイメージがLXD内で稼働するために関連する情報を含みます。
以下の情報を含んでいます。

```yaml
architecture: x86_64
creation_date: 1424284563
properties:
  description: Ubuntu 22.04 LTS Intel 64bit
  os: Ubuntu
  release: jammy 22.04
templates:
  ...
```

`architecture`と`creation_date`フィールドは必須です。
`properties`フィールドはイメージのデフォルトプロパティのセットを含みます。
`os`, `release`, `name`, `description`フィールドはよく使われますが、必須ではありません。

`templates`フィールドは省略可能です。
テンプレートをどのように設定するかの情報は{ref}`image_format_templates`を参照してください。

### ルートファイルシステム

コンテナでは、`rootfs/`ディレクトリがコンテナ内のルートディレクトリ(`/`)の完全なファイルシステムツリーを含みます。

仮想マシンは`rootfs/`ディレクトリの代わりに`rootfs.img` `qcow2`ファイルを使います。
このファイルはメインのディスクデバイスになります。

(image_format_templates)=
### テンプレート (省略可能)

インスタンス内部でファイルを動的に作成するのにテンプレートを使用できます。
そのためには、`metadata.yaml`ファイル内でテンプレートルールを設定し、`templates/`ディレクトリ内にテンプレートファイルを配置します。

一般的なルールとして、パッケージに所有されるファイルはテンプレート化は決してするべきではないです。そうでないとインスタンスの通常のオペレーションで上書きされてしまうでしょう。

#### テンプレートルール

生成すべき各ファイルに対して、`metadata.yaml`ファイル内でルールを作成します。
例:

```yaml
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

`when`キーは以下の1つ以上を指定できます。

- `create` - 新規インスタンスがイメージから作成された時に実行
- `copy` - 既存インスタンスからインスタンスが作成されたときに実行
- `start` - インスタンスが開始する度に毎回実行

`template`キーは`templates/`ディレクトリ内のテンプレートファイルを指します。

`properties`キーでユーザ定義のテンプレートプロパティをテンプレートファイルに渡せます。

ファイルが存在しない場合にのみLXDにファイルを作らせ、ファイルが存在する場合は上書きしてほしくない場合は、`create_only`キーをセットします。

#### テンプレートファイル

テンプレートファイルは[Pongo2](https://www.schlachter.tech/solutions/pongo2-template-engine/)形式を使います。

テンプレートファイルは常に以下のコンテキストを受け取ります。

| 変数     | 型                           | 説明                                                                         |
|--------------|--------------------------------|-------------------------------------------------------------------------------------|
| `trigger`    | `string`                       | テンプレートをトリガーしたイベント名                                       |
| `path`       | `string`                       | テンプレートを使用するファイルのパス                                             |
| `instance`   | `map[string]string`            | インスタンスプロパティのキー/値マップ(名前、アーキテクチャ、特権、一時的) |
| `config`     | `map[string]string`            | インスタンス設定のキー/値マップ                                       |
| `devices`    | `map[string]map[string]string` | インスタンスに割り当てられたデバイスのキー/値マップ                               |
| `properties` | `map[string]string`            | `metadata.yaml`で指定されたテンプレートプロパティのキー/値マップ               |

利便性のため、以下の関数がPongo2テンプレートにエクスポートされます。

- `config_get("user.foo", "bar")` - `user.foo`の値か、未設定の場合は`"bar"`を返します。

## イメージのtarball

LXDは2種類のLXD固有のイメージ形式、統合tarballと分離tarballをサポートします。

これらのtarballは圧縮されていても構いません。
LXDはtarballの広範囲の圧縮アルゴリズムをサポートします。
しかし、互換性のためには`gzip`または`xz`を使うのが良いです。

(image-format-unified)=
### 統合tarball

統合tarballは単一のtarball(通常`*.tar.xz`)で、イメージの完全な中身を含みます。それにはメタデータ、ルートファイルシステムと省略可能なテンプレートファイルが含まれます。

これがLXD自身がイメージを公開する際に内部的に使用している形式です。
通常こちらのほうが扱いやすいので、LXD固有のイメージを作る際は統合形式を使うのが良いです。

この形式のイメージの識別子はtarballのSHA-256ハッシュ値です。

(image-format-split)=
### 分離tarball

分離イメージは2つの分離したtarballから構成されます。
1つのtarball(通常`*.tar.xz`)はメタデータと省略可能なテンプレートファイルを含み、もう1つ(通常、コンテナでは`*.squashfs`で仮想マシンでは`*.qcow2`)はルートファイルシステムを含みます。

コンテナでは、ルートファイルシステムのtarballはSquashFSでフォーマットされていても構いません。
仮想マシンでは、`rootfs.img`ファイルは常に`qcow2`形式を使用します。
任意で`qcow2`のネイティブ圧縮を使って圧縮しても構いません。

この形式は既に利用可能である既存のLXD以外のrootfs tarballから簡単にイメージをビルドできるように設計されています。
LXDと他のツールの両方で使用するイメージを作りたい場合もこの形式を使うのが良いです。

この形式のイメージの識別子はメタデータとルートファイルシステムtarballを(この順で)結合したもののSHA-256ハッシュ値です。

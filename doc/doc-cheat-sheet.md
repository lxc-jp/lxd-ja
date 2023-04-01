---
orphan: true
myst:
  substitutions:
    reuse_key: "これは **インクルードされた** テキストです"
    advanced_reuse_key: "これはコードブロックを含む置換文です。
                         ```
                         コードブロック
                         ```"
---

# ドキュメントチートシート

ドキュメントファイルには、[Markdown](https://commonmark.org/)と[MyST](https://myst-parser.readthedocs.io/)の構文が混在しています。

シンタックスのヘルプと規約については、以下のセクションを参照してください。

## 見出し

``{list-table}.
   :header-rows: 1

* - 入力
  - 説明
* - `#Title` (タイトル)
  - ページのタイトルとH1の見出し
* - `## ヘディング` (見出し)
  - H2の見出し
* - `### ヘディング` (見出し)
  - H3の見出し
* - `#### 見出し` のようになります。
  - H4の見出し
* - ...
  - その他の見出し
```

以下の規則に従ってください。

- 間にテキストを挟まずに連続した見出しを使用しない。
- 見出しには文型を使用する（最初の単語だけを大文字にする）。
- レベルを飛ばさない（例えば、H2の後にはH4ではなく必ずH3をつける）。

## インラインフォーマット

``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
* - `*Italic*`
  - **イタリック*
* - `**Bold**` (太字)
  - **ボールド**
* - `` `code` ``
  - `code` (コードネーム)

```

以下の規則に従ってください。

- イタリック体は控えめに使用してください。一般的にイタリック体を使うのは、タイトルや名前です（例えば、リンクできないセクションのタイトルを参照する場合や、コンセプトの名前を紹介する場合など）。
- 太字は控えめに使いましょう。太字の一般的な使用方法は、UI要素です（「**OK**をクリックしてください」）。強調のために太字を使うことは避け、むしろ言いたいことが伝わるように文章を書き換えましょう。

## コードブロック

コードブロックの開始と終了は3つのバックティックで行います。

    ```

バックティックの後にコード言語を指定して、特定のレキサーを強制することもできますが、多くの場合、デフォルトのレキサーで十分に機能します。


``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
* - ````

    ```
    # コードブロックのデモ
    コードを表示します。
    - 例：真
    ```

    ````

  - ```
    # コードブロックのデモ
    コードを表示します。
    - 例：真
    ```

* - ````
    ``yaml
    # コードブロックのデモ
    コードを表示します。
    - 例：真
    ```

    ````

  - ```yaml
    # コードブロックのデモ
    コードを表示します。
    - 例：真
    ```

```

コードブロックにバックティックを含めるには、周囲のバックティックの数を増やします。

``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
* -
    `````

    ````
    ```
    ````

    `````

  -
    ````

    ```

    ````

```

## リンク

リンクの方法は、外部のURLにリンクするのか、ドキュメントの他のページにリンクするのかによって異なります。

### 外部リンク

外部リンクの場合、URLのみを使用し、リンクテキストを上書きしたい場合はMarkdownの構文を使用します。

``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
* - `https://linuxcontainers.org`
  - [`https://linuxcontainers.org`](https://linuxcontainers.org)
* - `[Linux コンテナ](https://linuxcontainers.org)`
  - [Linux コンテナ](https://linuxcontainers.org)
```

URL をテキストとして表示し、リンクされないようにするには、`<span></span>` を追加します。

``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
* - `https:/<span></span>/linuxcontainers.org` (日本語)
  - `https:/<span></span>/linuxcontainers.org`

```

### 内部参照

内部参照の場合、Markdown と MyST の両方の構文がサポートされています。リンクテキストを自動的に解決し、GitHubのレンダリングでリンクを示すことができるので、ほとんどの場合、MyST構文を使用するべきです。

#### ページの参照

ドキュメントのページを参照するには、MyST 構文を使ってリンクテキストを自動的に抽出します。リンクテキストを上書きする場合は、Markdownの構文を使用します。

``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
  - GitHubでの出力
  - ステータス
* - `` {doc}`index` ``
  - {doc}`index`
  - {doc}<span></span>`index`
  - 望ましいのは
* - `[](インデクス)`
  - [](インデクス)
  -
  - 使用しないでください。
* - `[LXDドキュメント](インデクス)`
  - [LXDドキュメント](インデクス)
  - LXDドキュメンテーション](index)
  - リンクテキストをオーバーライドするときに好ましい。
* - `` {doc}`LXD documentation <index>` ``
  - {doc}`LXDドキュメント <index>`
  - {doc}<span></span>`LXD ドキュメント <index>` {doc}<span></span>`LXD ドキュメント <index>`
  - リンクテキストをオーバーライドする場合の代替手段

```
以下の規則に従ってください。
- リンクテキストを上書きするのは、必要なときだけにしてください。ドキュメントのタイトルをリンクテキストとして使用できる場合は、そうしてください。なぜなら、タイトルが変更された場合、テキストは自動的に更新されるからです。
- リンクテキストを、自動生成されるテキストと同じもので「上書き」してはいけません。

(a_section_target)=
#### セクションの参照

ドキュメント内のセクション（同じページまたは別のページ）を参照するには、セクションにターゲットを追加してそのターゲットを参照するか、自動生成されたアンカーとファイル名を組み合わせて使用します。

以下のような規約を守ってください。

- セクションの中心的な部分や、頻繁にリンクされることが予想される「典型的な」場所にターゲットを追加してください。一度きりのリンクには、自動生成されるアンカーを使用できます。
- 必要な場合のみリンクテキストを上書きする。セクションのタイトルをリンクテキストとして使用できる場合は、そうしてください。タイトルが変更された場合、テキストは自動的に更新されます。
- リンクテキストを、自動生成されるテキストと同じもので「上書き」してはいけません。

##### ターゲットの使用

ドキュメント内の任意の場所にターゲットを追加することができます。ただし、対象となる要素に見出しやタイトルがない場合は、リンクテキストを指定する必要があります。

(a_random_target)=
```{list-table}.
   :header-rows: 1

* - 入力
  - 出力
  - GitHubでの出力
  - 説明
* - `(ターゲット_ID)=`
  -
  - (target_ID\)=
  - ターゲットである ``target_ID`` を追加します。
* - `` {ref}`a_section_target` `` を追加します。
  - a_section_target` {ref}`a_section_target`
  - a_section_target` `` {ref}`a_section_target`
  - タイトルを持っているターゲットを参照します。
* - `` {ref}`リンクテキスト <a_random_target>` ``
  - {ref}`リンクテキスト <a_random_target>`
  - \{ref\}`リンクテキスト <a_random_target>`
  - ターゲットを参照して、タイトルを指定します。
* - ``[`option name\](a_random_target)``
  - [`option name`](a_random_target)
  - [`option name`](a_random_target) (リンク切れ)
  - リンクテキストをマークアプする必要がある場合は Markdown の文法を使ってください。
```

##### 自動生成アンカーの使用

自動生成されたアンカーを使用するには、Markdownの構文を使用する必要があります。
同じファイル内でリンクする場合は、ファイル名を省略できます。

``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
  - GitHubでの出力
  - 説明
* - `[](#referencing-a-section)`
  - [](#referencing-a-section)
  -
  - 使用しないでください。
* - `[リンクテキスト](#referencing-a-section)`
  - [リンクテキスト](#referencing-a-section)
  - [リンクテキスト](#referencing-a-section)
  - リンクテキストをオーバーライドする場合に好ましい。
```

## ナビゲーション

すべてのドキュメントページは、ナビゲーションの中の他のページのサブページとして含まれていなければなりません。

これは、親ページの[`toctree`](https://www.sphinx-doc.org/en/master/usage/restructuredtext/directives.html#directive-toctree)ディレクティブで実現します。 <!-- wokeignore:rule=master -->

````
``{toctree}
:hidden:

サブページ1
subpage1
subpage2
```
````

ナビゲーションに含めてはいけないページがある場合、ファイルの先頭に次のような命令を記述することで、結果として生じるビルド警告を抑制することができます。

```
---
orphan: true
---
```

孤児ページを使うのは、明確な理由がある場合に限られます。

## リスト

``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
* - ```
    - 項目1
    - 項目2
    - 項目3
    ```
  - 項目1
    - 項目2
    - 項目3
* - ```
    1. ステップ1
    1. ステップ2
    1. ステップ3
    ```
  - 1. ステップ1
    1. ステップ2
    1. ステップ3
* - ```
    1. ステップ1
       - アイテム1
         * 小項目
       - 項目2
    1. ステップ2
       1. サブステップ1
       1. サブステップ2
    ```
  - 1.ステップ1
       - 項目1
         * 小項目
       - 項目2
    1. ステップ2
       1. サブステップ1
       1. サブステップ2
```

以下の規則に従ってください。

- 番号付きリストでは、ステップ番号を自動的に生成するために、すべての項目に ``1.`` を使用してください。
- 順不同のリストには`-`を使用してください。入れ子のリストを使用する場合は、入れ子のレベルに `*` を使用します。

### 定義リスト

``{list-table}.
   :header-rows: 1

* - 入力
  - 出力
* - ```
    項番
    : 定義

    用語2
    : 定義
    ```
  - 用語1
    : 定義

    用語2
    : 定義
```

## テーブル

標準的なMarkdownのテーブルを使用することができます。しかし、rST [list table](https://docutils.sourceforge.io/docs/ref/rst/directives.html#list-table)の構文を使用する方が通常ははるかに簡単です。

どちらのマークアップも以下のような出力になります。

``{list-table}.
   :header-rows: 1

* - ヘッダ1
  - ヘッダー2
* - セル1

    第2段落 セル1
  - セル2
* - セル3
  - セル4
```

### マークダウンテーブル

```
| Header 1                           | Header 2 |
|------------------------------------|----------|
| Cell 1<br><br>2nd paragraph cell 1 | Cell 2   |
| Cell 3                             | Cell 4   |
```

### リストテーブル

````
```{list-table}
   :header-rows: 1

* - Header 1
  - Header 2
* - Cell 1

    2nd paragraph cell 1
  - Cell 2
* - Cell 3
  - Cell 4
```
````

## ノート

```{list-table}
   :header-rows: 1

* - 入力
  - 出力
* - ````
    ```{note}
    A note.
    ```
    ````
  - ```{note}
    A note.
    ```
* - ````
    ```{tip}
    A tip.
    ```
    ````
  - ```{tip}
    A tip.
    ```
* - ````
    ```{important}
    Important information
    ```
    ````
  - ```{important}
    Important information.
    ```
* - ````
    ```{caution}
    This might damage your hardware!
    ```
    ````
  - ```{caution}
    This might damage your hardware!
    ```


```

以下の規則に従ってください。
- メモの使用は控えめに。
- 以下のタイプのノートのみを使用してください。note`, `tip`, `important`, `caution` です。
- `caution` は、ハードウェアの破損やデータの損失の危険性が明らかな場合にのみ使用してください。

## 画像

```{list-table}
   :header-rows: 1

* - 入力
  - 出力
* - ```
    ![Altテキスト](https://linuxcontainers.org/static/img/containers.png)
    ```
  - ![Altテキスト](https://linuxcontainers.org/static/img/containers.png)
* - ````
    ```{figure} https://linuxcontainers.org/static/img/containers.png
       :width: 100px
       :alt: Altテキスト

       図のキャプション
    ```
    ````
  - ```{figure} https://linuxcontainers.org/static/img/containers.png
       :width: 100px
       :alt: Altテキスト

       図のキャプション
    ```
```

以下のような規約があります。

- `doc` ディレクトリ内の画像は、パスを `/` で始めてください (例: `/images/image.png`)。
- スクリーンショットはPNG形式、グラフィックはSVG形式を使用してください。

### 代用

あまり多くのマークアップや特殊なフォーマットをせずに文や段落を再利用するには、置換を使用します。

置換機能は以下の場所で定義できます。

- `substitutions.yaml` というファイルがあります。このファイルで定義された置換は、すべてのドキュメントページで利用できます。
- 次のような形式の単一のファイルの先頭に。

  ````
  ---
  orphan: true
  myst:
    substitutions:
      reuse_key: "これは **インクルードされた** テキストです"
      advanced_reuse_key: "これはコードブロックを含む置換文です。
                           ```
                           コードブロック
                           ```"
  ---
  ````

`reuse/substitutions.py` でデフォルトの置換を定義し、ファイルの先頭でそれをオーバーライドすることで、両方のオプションを組み合わせることができます。

```{list-table}
   :header-rows: 1

* - 入力
  - 出力
* - `{{reuse_key}}`
  - {{reuse_key}}
* - `{{advanced_reuse_key}}`
  - {{advanced_reuse_key}}
```

以下の規約に従ってください。
- 置換はGitHubでは機能しません。そのため、インクルードされたテキストを示すキー名を使用してください (例えば、`reuse_note` の代わりに `note_not_supported` とします)。

### ファイルのインクルード

長いセクションや高度なマークアップを施したテキストを再利用するには、コンテンツを別のファイルに置き、そのファイルまたはファイルの一部を複数の場所にインクルードすることができます。

再利用するコンテンツにターゲットを入れることはできません (このターゲットへの参照が曖昧になるため)。ただし、ファイルをインクルードする直前にターゲットを置くことはできます。

ファイルのインクルードと置換を組み合わせれば、インクルードされたテキストの一部を置き換えることもできます。

`````{list-table}
   :header-rows: 1

* - 入力
  - 出力
* - ````

    % Include parts of the content from file [../README.md](../README.md)
    ```{include} ../README.md
       :start-after: Installing LXD from packages
       :end-before: <!-- Include end installing -->
    ```

    ````

  -
    % Include parts of the content from file [../README.md](../README.md)
    ```{include} ../README.md
       :start-after: Installing LXD from packages
       :end-before: <!-- Include end installing -->
    ```

`````

以下の規則に従ってください。

- ファイルのインクルードはGitHubでは機能しません。そのため、必ずインクルードしたファイルにリンクするコメントを追加してください。
- テキストの一部を選択するには、開始点と終了点にHTMLコメントを追加し、可能であれば `:start-after:` と `:end-before:` を使用します。必要に応じて、`:start-after:`と`:end-before:`を`:start-line:`と`:end-line:`と組み合わせることができます。ただし`:start-line:`と`:end-line:`だけの使用はエラーになりやすいです。

## タブ

``````{list-table}
   :header-rows: 1

* - 入力
  - 出力
* - `````

    ````{tabs}

    ```{group-tab} Tab 1

    Content Tab 1
    ```

    ```{group-tab} Tab 2

    Content Tab 2
    ```

    ````

    `````

  - ````{tabs}

    ```{group-tab} Tab 1

    Content Tab 1
    ```

    ```{group-tab} Tab 2

    Content Tab 2
    ```
    ````
``````

## 折りたたみ可能なセクション

rSTには詳細セクションのサポートはありませんが、HTMLを挿入してセクションを作成することができます。

```{list-table}
   :header-rows: 1

* - 入力
  - 出力
* - ```
    <details>
    <summary>Details</summary>

    Content
    </details>
    ```

  - <details>
    <summary>Details</summary>

    Content
    </details>

```

## 用語集

用語集はどのファイルでも定義することができます。理想的には、すべての用語を1つの用語集ファイルにまとめ、どのファイルからでも参照できるようにすることです。

`````{list-table}
   :header-rows: 1

* - 入力
  - 出力
* - ````

    ```{glossary}

    example term
      Definition of the example term.
    ```

    ````

  - ```{glossary}

    example term
      Definition of the example term.
    ```

* - ``{term}`example term` ``
  - {term}`example term`
`````

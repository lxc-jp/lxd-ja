# LXD の ドキュメント <!-- LXD documentation -->

LXD のドキュメントは http://lxd-ja.readthedocs.io/ja/latest/ (原文は https://linuxcontainers.org/lxd/docs/master/) で見られます。
<!--
The LXD documentation is available at: https://linuxcontainers.org/lxd/docs/master/
-->

GitHub でもドキュメントの基本的なレンダリングは行いますが、インクルードやクリック可能なリンクなどの重要な機能は含まれていません。ですので、 [公開されたドキュメント](http://lxd-ja.readthedocs.io/ja/latest/) を読むことをお勧めします。
<!--
GitHub provides a basic rendering of the documentation as well, but important features like includes and clickable links are missing. Therefore, we recommend reading the [published documentation](https://linuxcontainers.org/lxd/docs/master/).
-->

## ドキュメントの形式 <!-- Documentation format -->

日本語訳のドキュメントは [Markdown](https://commonmark.org/) で書かれています（原文のドキュメントは [MyST](https://myst-parser.readthedocs.io/) 拡張を使った [Markdown](https://commonmark.org/) で書かれています）。
<!--
The documentation is written in [Markdown](https://commonmark.org/) with [MyST](https://myst-parser.readthedocs.io/) extensions.
-->

文法のヘルプとガイドラインについては [documentation cheat sheet](https://linuxcontainers.org/lxd/docs/master/doc-cheat-sheet/) ([source](doc-cheat-sheet.md?plain=1)) を参照してください。
<!--
For syntax help and guidelines, see the [documentation cheat sheet](https://linuxcontainers.org/lxd/docs/master/doc-cheat-sheet/) ([source](doc-cheat-sheet.md?plain=1)).
-->

## ドキュメントのビルド <!-- Building the documentation -->

原文のドキュメントをビルドするには、レポジトリーのルートフォルダーで `make doc` を実行してください。このコマンドは必要なツールをインストールし、レンダリング結果を `doc/html/` フォルダーに出力します。（ツールを再インストールすることなく）変更したファイルだけを更新するには `make doc-incremental` を実行してください。
<!--
To build the documentation, run `make doc` from the root folder of the repository. This command installs the required tools and renders the output to the `doc/html/` folder. To update the documentation for changed files only (without re-installing the tools), run `make doc-incremental`.
-->

ビルドした後、`make doc-serve` を実行して http://localhost:8001 へ行くとレンダリングされたドキュメントを見ることができます。
<!--
After building, run `make doc-serve` and go to http://localhost:8001 to view the rendered documentation.
-->

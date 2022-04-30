# LXD ドキュメント

LXDの日本語ドキュメントは、https://lxd-ja.readthedocs.io/ja/latest/ で閲覧できます。

GitHubでもドキュメントの基本的なレンダリングを提供していますが、includeやクリッカブルリンクなどの重要な機能が欠落しています。そのため、[公開ドキュメント](https://lxd-ja.readthedocs.io/ja/latest/)を読むことをお勧めします。

## ドキュメントのフォーマット

ドキュメントは[Markdown](https://commonmark.org/)と[MyST](https://myst-parser.readthedocs.io/)の拡張で書かれています。

構文のヘルプやガイドラインについては、[ドキュメントチートシート](https://lxd-ja.readthedocs.io/ja/latest/doc-cheat-sheet/) ([ソース](doc-cheat-sheet.md?plain=1))を参照してください。

## ドキュメンテーションの構築

ドキュメントをビルドするには、リポジトリのルートフォルダから `make doc` を実行します。このコマンドは必要なツールをインストールして、出力を `doc/html/` フォルダにレンダリングします。変更されたファイルのみを対象にドキュメントを更新するには（ツールを再インストールすることなく）、`make doc-incremental`を実行します。

ビルド後、`make doc-serve`を実行して、http://localhost:8001、レンダリングされたドキュメントを見ることができます。

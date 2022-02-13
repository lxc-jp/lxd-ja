# コントリビュート
このプロジェクトに対する変更は Github: <https://github.com/lxc/lxd> 上でのプルリクエストで提案する必要があります。

そのあと、提案はコードレビューを経て承認され、メインブランチにマージされます。

## コミットの構造
コミットを次のように分類する必要があります:


 - API 拡張 (`doc/api-extensions.md` と `shared/version.api.go` を含む変更に対して `api: Add XYZ extension`)
 - ドキュメント (`doc/` 内のファイルに対して `doc: Update XYZ`)
 - API 構造 (`shared/api/` の変更に対して `shared/api: Add XYZ`)
 - Go クライアントパッケージ (`client/` の変更に対して `client: Add XYZ`)
 - CLI (`lxc/` の変更に対して `lxc/<command>: Change XYZ`)
 - スクリプト (`scripts/` の変更に対して `scripts: Update bash completion for XYZ`)
 - LXD デーモン (`lxd/` の変更に対して `lxd/<package>: Add support for XYZ`)
 - テスト (`tests/` の変更に対して `tests: Add test for XYZ`)

同様のパターンが LXD コードツリーの他のツールにも適用されます。そして複雑さによっては、さらに小さな単位に分けられるかもしれません。

CLI ツール (`lxc/`) 内の文字列を更新する際は、テンプレートを更新してコミットする必要があるでしょう:

 - make i18n
 - git commit -a -s -m "i18n: Update translation templates" po/

このようにすることで、コントリビューションに対するレビューが容易になり、stable ブランチへバックポートするプロセスが大幅に簡素化されます。

## ライセンスと著作権

デフォルトで、このプロジェクトに対するいかなる貢献も Apache 2.0 ライセンスの下で行われます。

変更の著者は、そのコードに対する著作権を保持します（著作権の割り当てはありません）。

## Developer Certificate of Origin
このプロジェクトに対する貢献へのトラッキングを改善するために DCO 1.1 を採用します。そして、ブランチに対するすべての変更に対する "sign-off" 手順を使います。

sign-off はコミットに対する説明の最後に付けるシンプルな行です。この行は、オープンソースへの貢献として、それを書いた本人であることを証明するか、それを渡す権利を有することを証明します。

> Developer Certificate of Origin
> Version 1.1
>
> Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
> 660 York Street, Suite 102,
> San Francisco, CA 94110 USA
>
> Everyone is permitted to copy and distribute verbatim copies of this
> license document, but changing it is not allowed.
>
> Developer's Certificate of Origin 1.1
>
> By making a contribution to this project, I certify that:
>
> (a) The contribution was created in whole or in part by me and I
>     have the right to submit it under the open source license
>     indicated in the file; or
>
> (b) The contribution is based upon previous work that, to the best
>     of my knowledge, is covered under an appropriate open source
>     license and I have the right under that license to submit that
>     work with modifications, whether created in whole or in part
>     by me, under the same open source license (unless I am
>     permitted to submit under a different license), as indicated
>     in the file; or
>
> (c) The contribution was provided directly to me by some other
>     person who certified (a), (b) or (c) and I have not modified
>     it.
>
> (d) I understand and agree that this project and the contribution
>     are public and that a record of the contribution (including all
>     personal information I submit with it, including my sign-off) is
>     maintained indefinitely and may be redistributed consistent with
>     this project or the open source license(s) involved.

有効な sign-off 行は次のようなものです:

```
Signed-off-by: Random J Developer <random@developer.org>
```

本名と有効なメールアドレスを使ってください。
申し訳ありませんが、仮名や匿名の貢献は許可されていません。

各コミットは、それが大きなセットの一部であっても、それぞれの著者によって個別に signed-off される必要があります。
`git commit -s` が役に立つでしょう。

## 開発を始める

開発環境をセットアップし LXD の新機能に取り組みを開始するには以下の手順に従ってください。

### 依存ライブラリーのビルド

依存ライブラリーをビルドするには [README.md](index.md) の「LXD のソースからのインストール」のセクションの手順に従ってください。

### あなたの fork の remote を追加

依存ライブラリーをビルドし終わったら、 GitHub の fork を remote として追加しその fork  にスイッチできます。
```bash
git remote add myfork git@github.com:<your_username>/lxd.git
git remote update
git checkout myfork/master
```

### LXD のビルド

最後にレポジトリ内で `make` を実行すれば LXD のあなたの fork をビルドできます。

この時点であなたが最も行いたいであろうことはあなたの fork 上にあなたの変更のための新しいブランチを作ることです。

```bash
git checkout -b [name_of_your_new_branch]
git push myfork [name_of_your_new_branch]
```

### LXD の新しいコントリビュータのための重要な注意事項

- 永続データは `LXD_DIR` ディレクトリに保管されます。これは `lxd init` で作成されます。 `LXD_DIR` のデフォルトは `/var/lib/lxd` か snap ユーザーは `/var/snap/lxd/common/lxd` です。
- 開発中はバージョン衝突を避けるため LXD のあなたの fork 用に `LXD_DIR` の値を変更すると良いでしょう。
- あなたのソースからコンパイルされる実行ファイルはデフォルトでは `$(go env GOPATH)/bin` に生成されます。
    - あなたの変更をテストするときはこれらの実行ファイル（インストール済みかもしれないグローバルの `lxd` ではなく）を明示的に起動する必要があります。
    - これらの実行ファイルを適切なオプションを指定してもっと便利に呼び出せるように `~/.bashrc` にエイリアスを作るという選択も良いでしょう。
- 既存のインストール済み LXD のデーモンを実行するための systemd サービスが設定されている場合はバージョン衝突を避けるためにサービスを無効にすると良いでしょう。

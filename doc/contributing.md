% Include content from [../CONTRIBUTING.md](../CONTRIBUTING.md)
```{include} ../CONTRIBUTING.md
    :end-before: <!-- Include end contributing -->
```

## 開発を始める

開発環境をセットアップし LXD の新機能に取り組みを開始するには以下の手順に従ってください。

### 依存ライブラリーのビルド

依存ライブラリーをビルドするには {ref}`installing_from_source` の手順に従ってください。

### あなたの fork の remote を追加

依存ライブラリーをビルドし終わったら、 GitHub の fork を remote として追加できます。

    git remote add myfork git@github.com:<your_username>/lxd.git
    git remote update

次にこちらに切り替えます。

<!-- wokeignore:rule=master -->
    git checkout myfork/master

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

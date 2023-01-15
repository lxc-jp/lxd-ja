(instances-snapshots)=
# インスタンススナップショットを作成するには

インスタンスのスナップショットを作成することでその時点のインスタンスを保存でき、これによりインスタンスを元の状態に簡単に復元できます。

インスタンススナップショットはインスタンスボリュームと同じストレージプールに保管されます。

## スナップショットを作成する

インスタンスのスナップショットを作成するには以下のコマンドを使用します。

    lxc snapshot <instance_name> [<snapshot name>]

% Include content from [storage_backup_volume.md](storage_backup_volume.md)
```{include} storage_backup_volume.md
    :start-after: <!-- Include start create snapshot options -->
    :end-before: <!-- Include end create snapshot options -->
```

仮想マシンの場合は、`--stateful`フラグを追加すると、インスタンスボリュームに含まれるデータだけでなくインスタンスの稼働状態もキャプチャーします。
この機能はCRIUの制限により、コンテナでは完全にはサポートされていないことに注意してください。

## スナップショットを表示、編集、削除する

インスタンスのスナップショットを表示するには以下のコマンドを使用します。

    lxc info <instance_name>

スナップショットを `<instance_name>/<snapshot_name>` で参照することで、インスタンスと同様にスナップショットを表示や変更できます。

スナップショットについての設定情報を表示するには、以下のコマンドを使用します。

    lxc config show <instance_name>/<snapshot_name>

スナップショットの有効期限を変更するには、以下のコマンドを使用します。

    lxc config edit <instance_name>/<snapshot_name>

```{note}
一般には、スナップショットはインスタンスの状態を保存しているので、編集できません。
唯一の例外が有効期限です。
設定に対する他の変更は黙って無視されます。
```

スナップショットを削除するには、以下のコマンドを使用します。

    lxc delete <instance_name>/<snapshot_name>

## インスタンススナップショットをスケジュールする

指定の日時 (最大で毎分1回) にスナップショットを自動的に作成するようにインスタンスを設定できます。
そうするには [`snapshots.schedule`](instance-options-snapshots) インスタンスオプションを設定します。

例えば、日次スナップショットを設定するには、以下のコマンドを使用します。

    lxc config set <instance_name> snapshots.schedule @daily

毎日午前6時にスナップショットを取得するように設定するには、以下のコマンドを使用します。

    lxc config set <instance_name> snapshots.schedule "0 6 * * *"

定期的なスナップショットをスケジュールする際は、自動的な削除 ([`snapshots.expiry`](instance-options-snapshots)) とスナップショットの命名パターン ([`snapshots.pattern`](instance-options-snapshots)) の設定を検討してください。
また稼働中でないインスタンスのスナップショットを取得するかどうか ([`snapshots.schedule.stopped`](instance-options-snapshots)) も設定すると良いかもしれません。

## インスタンススナップショットを復元する

任意のスナップショットにインスタンスを復元できます。

そうするには、以下のコマンドを使用します。

    lxc restore <instance_name> <snapshot_name>

スナップショットがステートフル (インスタンスの稼働状態についての情報を含むという意味) な場合、`--stateful`　フラグを追加すると状態も復元できます。

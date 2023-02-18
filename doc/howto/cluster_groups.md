---
discourse: 12716
---

(howto-cluster-groups)=
# クラスタグループをセットアップするには

```{youtube} https://www.youtube.com/watch?v=t_3YJo_xItM
```

クラスタメンバーは {ref}`cluster-groups` にアサインできます。
デフォルトでは、全てのクラスタメンバーは `default` グループに属しています。

クラスタグループを作成するには、`lxc cluster group create` コマンドを使用します。
例えば以下のようにします。

    lxc cluster group create gpu

クラスタメンバーを特定のグループに割り当てるには、`lxc cluster group assign` コマンドを使用します。
例えば以下のようにします。

    lxc cluster group assign server1 gpu

## クラスタグループメンバー上でインスタンスを起動する

クラスタグループがある場合、インスタンスを、特定のメンバー上で動かすようにターゲットする代わりに、クラスタグループのいずれかのメンバー上で動かすようにターゲットできます。

```{note}
クラスタグループにインスタンスをターゲットできるようにするには [`scheduler.instance`](cluster-member-config) は `all` (デフォルト) または `group` に設定する必要があります。

詳細は{ref}`clustering-instance-placement`を参照してください。
```

クラスタグループのメンバー上でインスタンスを起動するには、{ref}`cluster-target-instance` の指示に従ってください。ただし `--target` フラグではグループ名の前に `@` をつけて指定してください。
例えば以下のようにします。

    lxc launch images:ubuntu/22.04 c1 --target=@gpu

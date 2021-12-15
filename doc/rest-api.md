# REST API
## イントロダクション <!-- Introduction -->
LXD とクライアントの間の全ての通信は HTTP 上の RESTful API を使って
行います。リモートの操作は SSL で暗号化して通信し、ローカルの操作は
Unix ソケットを使って通信します。
<!--
All the communications between LXD and its clients happen using a
RESTful API over http which is then encapsulated over either SSL for
remote operations or a unix socket for local operations.
-->

## API のバージョニング <!-- API versioning -->
サポートされている API のメジャーバージョンのリストは `GET /` を使って
取得できます。
<!--
The list of supported major API versions can be retrieved using `GET /`.
-->

後方互換性を壊す場合は API のメジャーバージョンが上がります。
<!--
The reason for a major API bump is if the API breaks backward compatibility.
-->

後方互換性を壊さずに追加される機能は `api_extensions` の追加という形になり、
特定の機能がサーバでサポートされているかクライアントがチェックすることで
利用できます。
<!--
Feature additions done without breaking backward compatibility only
result in addition to `api_extensions` which can be used by the client
to check if a given feature is supported by the server.
-->

## 戻り値 <!-- Return values -->
次の 3 つの標準的な戻り値の型があります。
<!--
There are three standard return types:
-->

 * 標準の戻り値 <!-- Standard return value -->
 * バックグラウンド操作 <!-- Background operation -->
 * エラー <!-- Error -->

### 標準の戻り値 <!-- Standard return value -->
標準の同期的な操作に対しては以下のような dict が返されます。
<!--
For a standard synchronous operation, the following dict is returned:
-->

```js
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "metadata": {}                          // リソースやアクションに固有な追加のメタデータ
}
```

<!--
```js
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "metadata": {}                          // Extra resource/action specific metadata
}
```
-->

HTTP ステータスコードは必ず 200 です。
<!--
HTTP code must be 200.
-->

### バックグラウンド操作 <!-- Background operation -->
リクエストの結果がバックグラウンド操作になる場合、 HTTP ステータスコードは 202 (Accepted)
になり、操作の URL を指す HTTP の Location ヘッダが返されます。
<!--
When a request results in a background operation, the HTTP code is set to 202 (Accepted)
and the Location HTTP header is set to the operation URL.
-->

レスポンスボディは以下のような構造を持つ dict です。
<!--
The body is a dict with the following structure:
-->

```js
{
    "type": "async",
    "status": "OK",
    "status_code": 100,
    "operation": "/1.0/instances/<id>",                     // バックグラウンド操作の URL
    "metadata": {}                                          // 操作のメタデータ (下記参照)
}
```

<!--
```js
{
    "type": "async",
    "status": "OK",
    "status_code": 100,
    "operation": "/1.0/instances/<id>",                     // URL to the background operation
    "metadata": {}                                          // Operation metadata (see below)
}
```
-->

操作のメタデータの構造は以下のようになります。
<!--
The operation metadata structure looks like:
-->

```js
{
    "id": "a40f5541-5e98-454f-b3b6-8a51ef5dbd3c",           // 操作の UUID
    "class": "websocket",                                   // 操作の種別 (task, websocket, token のいずれか)
    "created_at": "2015-11-17T22:32:02.226176091-05:00",    // 操作の作成日時
    "updated_at": "2015-11-17T22:32:02.226176091-05:00",    // 操作の最終更新日時
    "status": "Running",                                    // 文字列表記での操作の状態
    "status_code": 103,                                     // 整数表記での操作の状態 (status ではなくこちらを利用してください。訳注: 詳しくは下記のステータスコードの項を参照)
    "resources": {                                          // リソース種別 (container, snapshots, images のいずれか) の dict を影響を受けるリソース
      "containers": [
        "/1.0/instances/test"
      ]
    },
    "metadata": {                                           // 対象となっている (この例では exec) 操作に固有なメタデータ
      "fds": {
        "0": "2a4a97af81529f6608dca31f03a7b7e47acc0b8dc6514496eb25e325f9e4fa6a",
        "control": "5b64c661ef313b423b5317ba9cb6410e40b705806c28255f601c0ef603f079a7"
      }
    },
    "may_cancel": false,                                    // (REST で DELETE を使用して) 操作がキャンセル可能かどうか
    "err": ""                                               // 操作が失敗した場合にエラー文字列が設定されます
}
```

<!--
```js
{
    "id": "a40f5541-5e98-454f-b3b6-8a51ef5dbd3c",           // UUID of the operation
    "class": "websocket",                                   // Class of the operation (task, websocket or token)
    "created_at": "2015-11-17T22:32:02.226176091-05:00",    // When the operation was created
    "updated_at": "2015-11-17T22:32:02.226176091-05:00",    // Last time the operation was updated
    "status": "Running",                                    // String version of the operation's status
    "status_code": 103,                                     // Integer version of the operation's status (use this rather than status)
    "resources": {                                          // Dictionary of resource types (container, snapshots, images) and affected resources
      "containers": [
        "/1.0/instances/test"
      ]
    },
    "metadata": {                                           // Metadata specific to the operation in question (in this case, exec)
      "fds": {
        "0": "2a4a97af81529f6608dca31f03a7b7e47acc0b8dc6514496eb25e325f9e4fa6a",
        "control": "5b64c661ef313b423b5317ba9cb6410e40b705806c28255f601c0ef603f079a7"
      }
    },
    "may_cancel": false,                                    // Whether the operation can be canceled (DELETE over REST)
    "err": ""                                               // The error string should the operation have failed
}
```
-->

対象の操作に対して追加のリクエストを送って情報を取り出さなくても、
何が起こっているかユーザーにとってわかりやすい形でボディは構成されています。
ボディに含まれる全ての情報はバックグラウンド操作の URL から取得する
こともできます。
<!--
The body is mostly provided as a user friendly way of seeing what's
going on without having to pull the target operation, all information in
the body can also be retrieved from the background operation URL.
-->

### エラー <!-- Error -->
さまざまな状況によっては操作を行う前に直ぐに問題が起きる場合があり、
そういう場合には以下のような値が返されます。
<!--
There are various situations in which something may immediately go
wrong, in those cases, the following return value is used:
-->

```js
{
    "type": "error",
    "error": "Failure",
    "error_code": 400,
    "metadata": {}                      // エラーについてのさらなる詳細
}
```

<!--
```js
{
    "type": "error",
    "error": "Failure",
    "error_code": 400,
    "metadata": {}                      // More details about the error
}
```
-->

HTTP ステータスコードは 400, 401, 403, 404, 409, 412, 500 のいずれかです。
<!--
HTTP code must be one of of 400, 401, 403, 404, 409, 412 or 500.
-->

## ステータスコード <!-- Status codes -->
LXD REST API はステータス情報を返す必要があります。それはエラーの理由だったり、
操作の現在の状態だったり、 LXD が提供する様々なリソースの状態だったりします。
<!--
The LXD REST API often has to return status information, be that the
reason for an error, the current state of an operation or the state of
the various resources it exports.
-->

デバッグをシンプルにするため、ステータスは常に文字列表記と整数表記で
重複して返されます。ステータスの整数表記の値は将来に渡って不変なので
API クライアントが個々の値に依存できます。文字列表記のステータスは
人間が API を手動で実行したときに何が起きているかをより簡単に判断
できるように用意されています。
<!--
To make it simple to debug, all of those are always doubled. There is a
numeric representation of the state which is guaranteed never to change
and can be relied on by API clients. Then there is a text version meant
to make it easier for people manually using the API to figure out what's
happening.
-->

ほとんどのケースでこれらは `status` と `status_code` と呼ばれ、前者は
ユーザーフレンドリーな文字列表記で後者は固定の数値です。
<!--
In most cases, those will be called status and `status_code`, the former
being the user-friendly string representation and the latter the fixed
numeric value.
-->

整数表記のコードは常に 3 桁の数字で以下の範囲の値となっています。
<!--
The codes are always 3 digits, with the following ranges:
-->

 * 100 to 199: リソースの状態 (started, stopped, ready, ...) <!-- resource state (started, stopped, ready, ...) -->
 * 200 to 399: 成功したアクションの結果 <!-- positive action result -->
 * 400 to 599: 失敗したアクションの結果 <!-- negative action result -->
 * 600 to 999: 将来使用するために予約されている番号の範囲 <!-- future use -->

### 現在使用されているステータスコード一覧 <!-- List of current status codes -->

コード <!-- Code -->  | 意味 <!-- Meaning -->
:---  | :------
100   | 操作が作成された <!-- Operation created -->
101   | 開始された <!-- Started -->
102   | 停止された <!-- Stopped -->
103   | 実行中 <!-- Running -->
104   | キャンセル中 <!-- Cancelling -->
105   | ペンディング <!-- Pending -->
106   | 開始中 <!-- Starting -->
107   | 停止中 <!-- Stopping -->
108   | 中断中 <!-- Aborting -->
109   | 凍結中 <!-- Freezing -->
110   | 凍結された <!-- Frozen -->
111   | 解凍された <!-- Thawed -->
112   | エラー <!-- Error -->
200   | 成功 <!-- Success -->
400   | 失敗 <!-- Failure -->
401   | キャンセルされた <!-- Cancelled -->

## 再帰 <!-- Recursion -->
巨大な一覧のクエリを最適化するために、コレクションに対して再帰が実装されています。
コレクションに対するクエリの GET リクエストに `recursion` パラメータを指定できます。
<!--
To optimize queries of large lists, recursion is implemented for collections.
A `recursion` argument can be passed to a GET query against a collection.
-->

デフォルト値は 0 でコレクションのメンバーの URL が返されることを意味します。
1 を指定するとこれらの URL がそれが指すオブジェクト (通常は dict 形式) で
置き換えられます。
<!--
The default value is 0 which means that collection member URLs are
returned. Setting it to 1 will have those URLs be replaced by the object
they point to (typically a dict).
-->

再帰はジョブへのポインタ (URL) をオブジェクトそのもので単に置き換えるように
実装されています。
<!--
Recursion is implemented by simply replacing any pointer to an job (URL)
by the object itself.
-->

## フィルタ <!-- Filtering -->
検索結果をある値でフィルタするために、コレクションにフィルタが実装されています。
コレクションに対する GET クエリに `filter` 引数を渡せます。
<!--
To filter your results on certain values, filter is implemented for collections.
A `filter` argument can be passed to a GET query against a collection.
-->

フィルタはインスタンスとイメージのエンドポイントに提供されています。
<!--
Filtering is available for the instance and image endpoints.
-->

フィルタにはデフォルト値はありません。これは見つかった全ての結果が返されることを意味します。
フィルタの引数には以下のような言語を設定します。
<!--
There is no default value for filter which means that all results found will
be returned. The following is the language used for the filter argument:
-->

?filter=field\_name eq desired\_field\_assignment

この言語は REST API のフィルタロジックを構成するための OData の慣習に従います。
フィルタは下記の論理演算子もサポートします。
not(not), equals(eq), not equals(ne), and(and), or(or)
フィルタは左結合で評価されます。
空白を含む値はクォートで囲むことができます。
ネストしたフィルタもサポートされます。
例えば config 内のフィールドに対してフィルタするには以下のように指定します。
<!--
The language follows the OData conventions for structuring REST API filtering
logic. Logical operators are also supported for filtering: not(not), equals(eq),
not equals(ne), and(and), or(or). Filters are evaluated with left associativity.
Values with spaces can be surrounded with quotes. Nesting filtering is also supported. 
For instance, to filter on a field in a config you would pass:
-->

?filter=config.field\_name eq desired\_field\_assignment

device の属性についてフィルタするには以下のように指定します。
<!--
For filtering on device attributes you would pass:
-->

?filter=devices.device\_name.field\_name eq desired\_field\_assignment

以下に上記の異なるフィルタの方法を含む GET クエリをいくつか示します。
<!--
Here are a few GET query examples of the different filtering methods mentioned above:
-->

containers?filter=name eq "my container" and status eq Running

containers?filter=config.image.os eq ubuntu or devices.eth0.nictype eq bridged

images?filter=Properties.os eq Centos and not UpdateSource.Protocol eq simplestreams

## 非同期操作 <!-- Async operations -->
完了までに 1 秒以上かかるかもしれない操作はバックグラウンドで実行しなければ
なりません。そしてクライアントにはバックグラウンド操作 ID を返します。
<!--
Any operation which may take more than a second to be done must be done
in the background, returning a background operation ID to the client.
-->

クライアントは操作のステータス更新をポーリングするか long-poll API を使って
通知を待つことが出来ます。
<!--
The client will then be able to either poll for a status update or wait
for a notification using the long-poll API.
-->

## 通知 <!-- Notifications -->
通知のために Websocket ベースの API が利用できます。クライアントへ送られる
トラフィックを制限するためにいくつかの異なる通知種別が存在します。
<!--
A websocket based API is available for notifications, different notification
types exist to limit the traffic going to the client.
-->

リモート操作の状態をポーリングしなくて済むように、リモート操作を開始する
前に操作の通知をクライアントが常に購読しておくのがお勧めです。
<!--
It's recommended that the client always subscribes to the operations
notification type before triggering remote operations so that it doesn't
have to then poll for their status.
-->

## PUT と PATCH の使い分け <!-- PUT vs PATCH -->
LXD API は既存のオブジェクトを変更するのに PUT と PATCH の両方をサポートします。
<!--
The LXD API supports both PUT and PATCH to modify existing objects.
-->

PUT はオブジェクト全体を新しい定義で置き換えます。典型的には GET で現在の
オブジェクトの状態を取得した後に PUT が呼ばれます。
<!--
PUT replaces the entire object with a new definition, it's typically
called after the current object state was retrieved through GET.
-->

レースコンディションを避けるため、 GET のレスポンスから ETag ヘッダを読み取り
PUT リクエストの If-Match ヘッダに設定するべきです。こうしておけば GET と
PUT の間にオブジェクトが他から変更されていた場合は更新が失敗するようになります。
<!--
To avoid race conditions, the Etag header should be read from the GET
response and sent as If-Match for the PUT request. This will cause LXD
to fail the request if the object was modified between GET and PUT.
-->

PATCH は変更したいプロパティだけを指定することでオブジェクト内の単一の
フィールドを変更するのに用いられます。キーを削除するには通常は空の値を
設定すれば良いようになっていますが、 PATCH ではキーの削除は出来ず、代わりに
PUT を使う必要がある場合もあります。
<!--
PATCH can be used to modify a single field inside an object by only
specifying the property that you want to change. To unset a key, setting
it to empty will usually do the trick, but there are cases where PATCH
won't work and PUT needs to be used instead.
-->

## インスタンス、コンテナーと仮想マシン <!-- Instances, containers and virtual-machines -->
このドキュメントでは `/1.0/instances/...` のようなパスを常に示します。
これらはかなり新しく、仮想マシンがサポートされた LXD 3.19 で導入されました。
<!--
This documentation will always show paths such as `/1.0/instances/...`.
Those are fairly new, introduced with LXD 3.19 when virtual-machine support.
-->

コンテナーのみをサポートする古いリリースでは全く同じ API を `/1.0/containers/...` で利用します。
<!--
Older releases that only supported containers will instead use the exact same API at `/1.0/containers/...`.
-->

後方互換性の理由で LXD は `/1.0/containers` API を引き続き公開しサポートしますが、簡潔さのため以下では両方をドキュメントはしないことにしました。
<!--
For backward compatibility reasons, LXD does still expose and support
that `/1.0/containers` API, though for the sake of brevity, we decided
not to double-document everything below.
-->

`/1.0/virtual-machines` に追加のエンドポイントも存在し、 `/1.0/containers` とほぼ同様ですが、仮想マシンのタイプのインスタンスのみを表示します。
<!--
An additional endpoint at `/1.0/virtual-machines` is also present and
much like `/1.0/containers` will only show you instances of that type.
-->

## API 構造 <!-- API structure -->
LXD は API エンドポイントを記述する [Swagger](https://swagger.io/) 仕様を自動生成しています。
この API 仕様の YAML 版が [rest-api.yaml](https://github.com/lxc/lxd/blob/master/doc/rest-api.yaml) にあります。
手軽にウェブで見る場合は [https://linuxcontainers.org/lxd/api/master/](https://linuxcontainers.org/lxd/api/master/) を参照してください。
<!--
LXD has an auto-generated [Swagger](https://swagger.io/) specification describing its API endpoints.
The YAML version of this API specification can be found in [rest-api.yaml](https://github.com/lxc/lxd/blob/master/doc/rest-api.yaml).
A convenient web rendering of it can be found here: [https://linuxcontainers.org/lxd/api/master/](https://linuxcontainers.org/lxd/api/master/)
-->

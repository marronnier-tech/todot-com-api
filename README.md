# 目次
- [API一覧](#API一覧)
- [APIリクエスト・レスポンス詳細](#APIリクエスト・レスポンス詳細)

# API一覧

## 一覧系列（ログイン未定）
- [GET-全ユーザーのTODO一覧を表示](#GET-全ユーザーのTODO一覧を表示)；非優先：フォローに限定
- [GET-該当ユーザーのTODO一覧を表示](#GET-該当ユーザーのTODO一覧を表示)
- [GET-全ユーザーのゴールしたTODOリスト](#GET-全ユーザーのゴールTODOリスト)
- [GET-該当ユーザーのゴールTODOリスト](#GET-該当ユーザーのゴールTODOリスト)

## ログイン必須系列
- [GET-本人のTODO一覧を表示](#GET-本人のTODO一覧を表示)
- [GET-本人のゴール一覧を表示](#GET-本人のゴール一覧を表示)
- [POST-TODOを登録](#POST-TODOを登録)
- [DELETE-TODOを削除（論理削除）](#DELETE-TODOを削除（論理削除）)
- [POST-当日TODO完了](#POST-当日TODO完了)
- [DELETE-当日TODO完了取消](#DELETE-当日TODO完了取消)
- [PATCH-TODOをゴールに変更](#PATCH-TODOをゴールに変更)
- [GET-該当ユーザーの月別TODO達成状況取得](#GET-該当ユーザーの月別TODO達成状況取得) ＜＜未実装＞＞

## ユーザー情報詳細系列
- [GET-本人情報詳細表示](#GET-本人情報詳細表示)
- [PATCH-ユーザー情報の更新](#PATCH-ユーザー情報の更新)
- [GET-該当ユーザー情報詳細表示](#GET-該当ユーザー情報詳細表示)


## ユーザー登録・大会
- [POST-ユーザー登録](#POST-ユーザー登録)
- [POST-ユーザーログイン](#POST-ユーザーログイン)
- [DELETE-ユーザーログアウト](#DELETE-ユーザーログアウト)
- [DELETE-ユーザー退会（論理削除）](#DELETE-ユーザー退会（論理削除）)


## ユーザー秘匿情報系列 - ＜＜未実装＞＞
- [GET-ユーザー秘匿情報表示](#GET-ユーザー秘匿情報表示)
- [PATCH-メールアドレス更新](#PATCH-メールアドレス更新)

## フォロー系列 - ＜＜未実装＞＞
- [GET-フォロー一覧](#GET-フォロー一覧)
- [DELETE-フォロー削除（物理削除）](#DELETE-フォロー削除（物理削除）)

## ログイン判定
- [GET-ログインユーザー識別フラグ](#GET-ログインユーザー識別フラグ)

※ゴール……完全に自分の達成したい目標に達成し、そのTODOをやる必要がなくなったこと

-----

# APIリクエスト・レスポンス詳細

## GET-全ユーザーのTODO一覧を表示
### URI
```
GET /todo{?page,limit,order}
```
### 処理概要
- 全ユーザーのTODO一覧を表示する。
- limitで各ページ上限、pageでページ数を指定できる。
デフォルトはlimit:100、page:1。クエリパラメータで取得する。
- orderは最終達成日順（last_achieved）、達成回数順（count）、最近設定された順（created_at）にできる。デフォルトは最終達成日順。
- ゴールは含まない


### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| page | numeric | ページ数 |  o |
| limit | numeric | ページ内表示Todo数 | o |
| order | string | 順序 |  o |


### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoArray | array | 全todoのリスト| 
| TodoObj[TodoArray] | object | todo内容 |
| TodoID{TodoObj} | numeric | todoのID |
| IsDeleted{TodoObj} | boolean | 削除されたか（falseのみ表示） |
| Content{TodoObj} | string | todoの詳細 |
| CreatedAt{TodoObj} | string | todo登録日 | 
| LastAchieved{TodoObj} | string | 最終達成日（n日前） |
| User[TodoArray] | list | 所有user情報|
| UserId{User} | numeric | 所有userのID |
| UserName{User} | string | 所有ユーザー名 |
| UserHN{User} | string | 所有ユーザーのハンドルネーム |
| UserImg{User} | string | 所有ユーザーのアイコン；非優先 |
| GoaledCount{User} | count | ゴール数 |
| limit | numeric | ページ内表示Todo数 |
| page | numeric | ページ数 | 
| order | string | 順序 | 

### 正常レスポンス

```json
HTTP/1.1 200 OK
{
    "TodoArray" :[
        {
            "TodoObj":{
                "TodoID": 1,
                "IsDeleted": false,
                "Content": "プログラミング",
                "CreatedAt": "2020-10-31",
                "Count": 1,
                "LastAchieved": "2日前"
            },
            "User":{
                "UserId": 1,
                "UserName": "gopher0120",
                "UserHN": "Gopherくん",
                "UserImg": "cutiegopher.jpg",
                "GoaledCount": 1
            },
        },
    ],
    "limit": 100,
    "page": 1,
    "order": "count"
}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## GET-該当ユーザーのTODO一覧を表示
★ログイン必須
### URI
```
GET /todo/:name{?order}
```
### 処理概要
- キーで取得したユーザーのTODO一覧を表示する。
- orderは最終達成日順（last_achieved）、達成回数順（count）、最近設定された順（created_at）にできる。デフォルトは最終達成日順。

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| name | string | ユーザー名 | x |
| order | string | 表示順 | o |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| User | list | ユーザー情報 | 
| UserID{User} | numeric | ユーザーID | 
| UserName{User} | string | ユーザー名 |
| UserHN{User} | string | ユーザーのハンドルネーム |
| UserImg{User} | string | ユーザー画像 |
| GoaledCount{User} | count | ゴール数 |
| TodoArray | Array | todo内容 |
| TodoID[TodoArray] | numeric | todoのID |
| IsDeleted[TodoArray] | boolean | 削除されたか（falseのみ表示） |
| Content[TodoArray] | string | todoの詳細 |
| CreatedAt[TodoArray] | string | todo登録日 |
| LastAchieved[TodoArray] | string | 最終達成日（n日前） | 
| TodayAchieved[TodoArray] | boolean | 本日達成したか |
| order | string | 表示順序 |
| owner | boolean | ログイン中のユーザーと一致しているか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "Todo": {
        "User": {
            "UserID": 1,
            "UserName": "gopher0120",
            "UserHN": "Gopherくん",
            "UserImg": "cutiegopher.jpg",
            "GoaledCount": 1
        },
        "TodoArray": [
            {
                "TodoID" : 1,
                "IsDeleted": false,
                "Content": "プログラミング",
                "CreatedAt": "2020-10-30",
                "Count": 0,
                "LastAchieved": "達成日はありません",
                "TodayAchieved": false
            }
        ]
    },
    "order": "last_achieved",
    "owner": false

}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error


## GET-全ユーザーのゴールTODOリスト

### URI
```
GET /goal{?page,limit}
```
### 処理概要
- 全ユーザーのゴールしたTODO一覧を表示する。
- limitで各ページ上限、pageでページ数、orderで表示順序を指定できる。
デフォルトはlimit:100、page:1、order:"goaled_at"。クエリパラメータで取得する。
- orderはゴール達成日順（goaled_at）、達成回数順（count）にできる。デフォルトは最終達成日順。
- ゴールのみ表示


### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| page | numeric | ページ数 |  o |
| limit | numeric | ページ内表示Todo数 | o |
| order | string | 表示順序 | o |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoArray | array | 全ゴールリスト| 
| TodoObj[TodoArray] | object | ゴール内容 |
| TodoID{TodoObj} | numeric | ゴールしたtodoのID |
| IsDeleted{TodoObj} | boolean | 削除されたか（falseのみ表示） |
| Content{TodoObj} | string | ゴールしたtodoの詳細 |
| GoaledAt{TodoObj} | numeric | ゴール日 |
| User[TodoArray] | list | 所有user情報|
| UserId{User} | numeric | 所有userのID |
| UserName{User} | string | 所有ユーザーの名前 |
| UserHN{User} | string | 所有ユーザーのハンドルネーム |
| UserImg{User} | string | 所有ユーザーのアイコン；非優先 |
| GoaledCount{User} | count | ゴール数 |
| limit | numeric | ページ内表示Todo数 |
| page | numeric | ページ数 | 
| order | string | 表示順序 |

### 正常レスポンス
- ステータス：200
```json
HTTP/1.1 200 OK
{
    "TodoArray" :[
        {
            "TodoObj":{
                "TodoID": 1,
                "Content": "プログラミング",
                "GoaledAt": "2020-11-01",
                "AchievedCount": 1
            },
            "User":{
                "UserId": 1,
                "UserName": "gopher0120",
                "UserHN": "Gopherくん",
                "UserImg": "cutiegopher.jpg",
                "GoaledCount": 1
            },
        },
    ],
    "limit": 100,
    "page": 1,
    "order": "goaled_at"
}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## GET-該当ユーザーのゴールTODOリスト
★ログイン必須
### URI
```
GET /goal/:name
```
### 処理概要
- キーで取得したユーザーのTODO一覧を表示する。
- ゴール日順に表示する。
- orderはゴール達成日順（goaled_at）、達成回数順（count）にできる。デフォルトはゴール達成日順。

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| name | string | ユーザー名 | x |
| order | string | 表示順序 | o |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| User | list | ユーザー情報 | 
| UserID{User} | numeric | ユーザーID | 
| UserName{User} | string | ユーザー名 |
| UserHN{User} | string | ユーザーのハンドルネーム |
| UserImg{User} | string | ユーザー画像 |
| GoaledCount{User} | count | ゴール数 |
| TodoObj | array | todo取得 |
| TodoID[GoalArray] | numeric | todoのID |
| Content[GoalArray] | string | todoの詳細 |
| GoaledAt[GoalArray] | string | ゴール日 | 
| AchievedDays[GoalArray] | numeric | ゴールまでに達成した回数 |
| order | string | 表示順序 |
| owner | boolean | ログイン中のユーザーと該当ユーザーが一致しているか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "Goal":{
        "User": {
            "UserID": 1,
            "UserName": "gopher0120",
            "UserHN": "Gopherくん",
            "UserImg": "cutiegopher.jpg",
            "GoaledCount": 1
        },
        "GoalArray": [
            {
                "TodoID" : 1,
                "Content": "プログラミング",
                "GoaledAt": "2020-11-10",
                "AchievedCount": 10,
            }
        ]
    },
    "order": "goaled_at",
    "owner": false
}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## GET-本人のTODO一覧を表示
★ログイン必須
### URI
```
GET /mypage{?order}
```

### 処理概要
- [GET-該当ユーザーのTODO一覧を表示](#GET-該当ユーザーのTODO一覧を表示)へリダイレクト(302)
- ただしname = 本人UserName

## GET-本人のゴール一覧を表示
★ログイン必須
### URI
```
GET /mypage/goal
```

### 処理概要
- [GET-該当ユーザーのゴールTODO一覧を表示](#GET-該当ユーザーのゴールTODOリスト)へリダイレクト(302)
- ただしname = 本人UserName

## POST-TODOを登録
★ログイン必須
### URI
```
POST /mypage
```
### 処理概要
- Todoリストに内容を登録する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Content | string | Todoの詳細 | x |

### 入力例
```json
{
    "Content": "プログラミング"
}
```

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoID | numeric | todoのID |
| Content | string | todoの詳細 |
| CreatedAt | string | todo登録日 | 
| LastAchieved | string | 最終達成日（0日前） |
| Count | numeric | 累計達成回数 |
| TodayAchieved | bool | 本日達成したか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
        "TodoID" : 1,
        "Content": "プログラミング",
        "CreatedAt": "2020-10-31",
        "LastAchieved": "2020-11-03",
        "Count": 3,
        "TodayAchieved": true
}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## DELETE-TODOを削除（論理削除）
★ログイン必須
### URI
```
DELETE /mypage/:id
```
### 処理概要
- 取得したキーのTodo項目をTodoリストから削除する
- ゴールしたTodoは退会しない限りDeleteできない

### リクエストパラメータ

| key | type | content |  null |
| ---- | ---- | ---- | ---- |
| TodoID | numeric | TodoのID | x |

### ステータスコード

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- | 
| TodoID | numeric | TodoのID |
| DeletedTodo | boolean | Todoが削除されたか |

### 正常レスポンス
```json
HTTP/1.1 201 Created
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## POST-当日TODO完了
★ログイン必須
### URI
```
POST /mypage/:id/today
```
### 処理概要
- Todoリストに当日のtodo達成を登録する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| id | string | TodoのID | x |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoID | numeric | todoのID |
| IsDeleted | bool | 削除されていない |
| Content | string | todoの詳細 |
| CreatedAt | string | todo登録日 | 
| LastAchieved | string | 最終達成日（0日前） |
| Count | numeric | 累計達成回数 |
| TodayAchieved | bool | 本日達成したか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
        "TodoID" : 1,
        "IsDeleted": false,
        "Content": "プログラミング",
        "CreatedAt": "2020-10-31",
        "LastAchieved": "2020-11-03",
        "Count": 3,
        "TodayAchieved": true
}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error



## DELETE-当日TODO完了取消
★ログイン必須
### URI
```
DELETE /mypage/:id/today
```
### 処理概要
- Todoリストの当日のtodo達成を取り消す

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| id | string | TodoのID | x |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoID | numeric | todoのID |
| IsDeleted | bool | 削除されていない |
| Content | string | todoの詳細 |
| CreatedAt | string | todo登録日 | 
| LastAchieved | string | 最終達成日（0日前） |
| Count | numeric | 累計達成回数 |
| TodayAchieved | bool | 本日達成したか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
        "TodoID" : 1,
        "IsDeleted": false,
        "Content": "プログラミング",
        "CreatedAt": "2020-10-31",
        "LastAchieved": "2020-11-02",
        "Count": 2,
        "TodayAchieved": false
}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error


## PATCH-TODOをゴールに変更
★ログイン必須
### URI
```
PATCH /mypage/:id/goal
```

### 処理概要
- Todoリストの達成状況を「ゴール」に変更する

### リクエストパラメータ

| key | type | content |  null |
| ---- | ---- | ---- | ---- |
| id | int | TodoのID | x |

### 入力例
```json
{
    "id": 1
}
```

### レスポンスパラメータ
| key | type | content | 
| ---- | ---- | ---- |
| TodoID | numeric | TodoのID |
| Goaled | boolean | Todoがゴールか |
| GoaledAt | datetime | ゴールした日 |

### 正常レスポンス
```json
HTTP/1.1 201 Created
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## GET-該当ユーザーの月別TODO達成状況取得
★ログイン必須
※未実装
### URI
```
GET /mypage/achieved
```
### 処理概要
- キーで取得したユーザーのTODO達成状況を確認する
- TODO達成状況をグラフで表示する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| name | string | ユーザー名 | x |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| User | list | ユーザー情報 | 
| UserID{User} | numeric | ユーザーID | 
| UserName{User} | string | ユーザー名 |
| UserHN{User} | string | ユーザーのハンドルネーム |
| UserImg{User} | string | ユーザー画像 |
| GoaledCount{User} | count | ゴール数 |
| TodoArray | array | todo取得 |
| TodoID[TodoArray] | numeric | todoのID |
| Content[TodoArray] | string | todoの詳細 |
| CreatedAt[TodoArray] | string | todo登録日 |
| AchievedAll[TodoArray] | numeric | 達成回数 |
| Achieved[TodoArray] | Object | 達成詳細 |
| ByYear{Achieved} | list | 年ごとの達成 |
| Year[ByYear] | numeric | 達成年 |
| TimesByYear[ByYear]
| ByMonth[ByYear] | array | 月ごとの達成 |
| Month[ByMonth] | numeric | 達成月 |
| TimesByMonth[ByMonth] | numeric | 該当月の達成回数 |


### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "User": {
        "UserID": 1,
        "UserName": "gopher0120",
        "UserHN": "Gopherくん",
        "UserImg": "cutiegopher.jpg",
        "GoaledCount": 1
    },
    "TodoArray": [
        {
            "TodoID" : 1,
            "Content": "プログラミング",
            "CreatedAt": "2020-10-30",
            "AchievedAll": 30,
            "Achieved": {
                "ByYear": [
                    {
                        "Year": 2020,
                        "TimesByYear": 20,
                        "ByMonth": [                    
                            {
                                "Month": 11,
                                "TimesByMonth": 5
                            },
                            {
                                "Month": 12,
                                "TimesByMonth": 15
                            }
                        ]
                    },
                    {
                        "Year": 2021,
                        "TimesByMonth": 10,
                        "ByMonth" :[
                            {
                                "Month": 1,
                                "TimesByMonth": 10
                            }
                        ]
                    }
                ]
            }
        }
    ]
}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## GET-本人情報詳細表示
★ログイン必須
### URI
```
GET /profile
```
### 処理概要
- [GET-ユーザー情報詳細表示](#GET-ユーザー情報詳細表示)へリダイレクト
- ただしname = UserName

## PATCH-ユーザー情報の更新

### URI
```
PATCH /profile
```
### 処理概要
- ユーザー情報を更新する（秘匿情報以外）

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Name | string | ユーザーの名前 | x |
| HN | string | ユーザーのハンドルネーム | x |
| Img | string | ユーザーのアイコン；非優先 | o |
| FinalGoal | string | ユーザーの目標 | o |
| Profile | string | ユーザーのプロフィール（自由記述） | o |
| Twitter | string | ユーザーのTwitterアカウント | o |
| Instagram | string | ユーザーのInstagramアカウント | o |
| Facebook | string | ユーザーのFacebookアカウント | o |
| GitHub | string | ユーザーのGitHubアカウント | o |
| URL | string | その他ユーザーが載せたいURL | o |

### 正常レスポンス
```json
HTTP/1.1 201 Created
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error



## GET-該当ユーザー情報詳細表示
★ログイン必須
### URI
```
GET /profile/:name
```
### 処理概要
- ユーザー情報の詳細を取得する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| name | string | ユーザー名 | x |

### 入力例
```json
{
    "name": "gopher0120"
}
```

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| UserInfo | list | 該当ユーザー情報 |
| ID{UserInfo} | numeric | ユーザーID |
| Name{UserInfo} | string | ユーザーの名前 |
| HN{UserInfo} | string | ユーザーのハンドルネーム |
| Img{UserInfo} | string | ユーザーのアイコン；非優先 |
| FinalGoal{UserInfo} | string | ユーザーの目標 |
| Profile{UserInfo} | string | ユーザーのプロフィール（自由記述） |
| Twitter{UserInfo} | string | ユーザーのTwitterアカウント |
| Instagram{UserInfo} | string | ユーザーのInstagramアカウント |
| Facebook{UserInfo} | string | ユーザーのFacebookアカウント |
| GitHub{UserInfo} | string | ユーザーのGitHubアカウント |
| URL{UserInfo} | string | その他ユーザーが載せたいURL |
| owner{UserInfo} | boolean | ログイン中のユーザーと該当ユーザーが一致するか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "UserInfo": {
        "ID": 1,
        "Name": "gopher0120",
        "HN": "Gopherくん",
        "Img": "cutiegopher.jpg",
        "FinalGoal": "Golangの神になりたい！！",
        "Profile": "僕はGopher。Golangが大好き！最近Goで参加する競技プログラミングのYouTubeチャンネル始めました。Golangがもっと広まると嬉しいな！",
        "Twitter": "go",
        "Instagram": "go",
        "Facebook": "go",
        "Github": "go",
        "URL": "http://www.cutiegophergogogo.com/"
    },
    "owner": false
}
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## POST-ユーザー登録
### URI
```
POST /register
```
### 処理概要
- ユーザー登録をする。
- 固有のユーザー名、メールアドレス、パスワードが必須。
- HNがない場合は、ユーザー名がそのまま使用される。
- 詳細情報は空になる。

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Name | string | ユーザー名 | x |
| HN | string | ハンドルネーム | o |
| MailAddress | string | メールアドレス | o |
| Password | string | パスワード | x | 

### 入力例
```json
{
    "Name": "gopher0120",
    "HN": "Gopherくん",
    "MailAddress": "cutegopher@gophergogo.com",
    "Password": "golanggggggg"
}
```

### レスポンスパラメータ

```json
HTTP/1.1 302 Found
GET /mypage
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error


## POST-ユーザーログイン

### URI
```
POST /login
```

### 処理概要
- ユーザーのログイン
- nameとpasswordで認証する
- Cookieでログイン情報を保存する
- セッション中はログインをキープする

### リクエストパラメータ

| key | type | content | null |
| --- | --- | --- | --- | 
| name | string | ユーザーネーム | x |
| password | string | パスワード | x |

### 正常レスポンス

```json
HTTP/1.1 302 Found
GET /mypage
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error



## DELETE-ユーザーログアウト
★ログイン必須
### URI
```
DELETE /logout
```

### 処理概要
- ログアウトする
- Cookieからログイン情報を削除する

### 正常レスポンス
```
HTTP/1.1 204 No Content
GET /todo
```

### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## DELETE-ユーザー退会（論理削除）
★ログイン必須
### URI
```
DELETE /delete
```
### 処理概要
- ユーザーを削除する（論理削除）
- Todoはゴール含めてすべて削除する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Password | string | パスワード | x |

### 正常レスポンス
```json
HTTP/1.1 302 Redirect
GET /todo
```
### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## GET-ユーザー秘匿情報表示
★ログイン必須
※未実装
### URI
```
GET /secret
```
### 処理概要
- 秘匿情報を表示する（本人のみ）

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Password | string | パスワード | x |

### 入力例
```json
{
    "Password": "golanggggggg"
}
```

### レスポンスパラメータ

| key | type | content |
| ---- | ---- | ---- |
| ID | numeric | ユーザーID |
| Name | string | ユーザー名 |
| HN | string | ハンドルネーム |
| MailAddress | string | メールアドレス |

### 正常レスポンス
```json
HTTP/1.1 200 OK
```json
{
    "ID": 1,
    "Name": "gopher0120",
    "HN": "Gopherくん",
    "MailAddress": "cutegopher@gophergogo.com",
}
```
### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

## PATCH-メールアドレス更新
★ログイン必須
※未実装
### URI
```
PATCH /secret
```
### 処理概要
- メールアドレス を更新する（本人のみ）

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| MailAddress | string | メールアドレス | x |
| Password | string | パスワード | x |

### 入力例
```json
{
    "MailAddress": "go@go.com",
    "Password": "golanggggggg"
}
```

### レスポンスパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| ID | numeric | ユーザーID |
| Name | string | ユーザー名 |
| HN | string | ハンドルネーム |
| MailAddress | string | メールアドレス |

### 正常レスポンス
```json
HTTP/1.1 200 OK
```json
{
    "ID": 1,
    "Name": "gopher0120",
    "HN": "Gopherくん",
    "MailAddress": "cutegopher@gophergogo.com",
}
```
### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error


## GET-フォロー一覧
※未実装

## DELETE-フォロー削除（物理削除）
※未実装

## GET-ログインユーザー識別フラグ

### URI
```
GET /adminflag
```
### 処理概要
- ログインしているユーザーはtrue、それ以外はfalseを返す

### レスポンスパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| LoginFlag | boolean | ログインフラグ |

### 正常レスポンス
```json
HTTP/1.1 200 OK
```json
{
    "LoginFlag": true, 
}
```
### 異常レスポンス
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

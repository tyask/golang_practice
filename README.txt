1) hello (2min): UT及び課題の解き方デモ
Hello関数は、引数に与えられた人に対して挨拶する関数です。UTが通るように実装してください。

2) is_prime (5min): for/ifの練習
IsPrime関数は、素数判定をする関数です。UTが通るように実装してください。
なお、素数とは、2以上の自然数で正の約数が1と自分自身のみであるもののことを言います。

3) book (5min): 構造体/関数の練習
本を表す構造体Bookを定義し、UTが通るようにメソッドを実装してください。
・フィールドには以下の4フィールドを定義してください。
  Title     (string)
  Outher    (string)
  Publisher (string)
  Price     (int)
・Stringメソッドは、本の表示形式を返すメソッドです。表示形式は以下の通りです。
  {Title}/{Outher}/{Publisher} (¥{Price})
・SetPriceメソッドは、Priceの値を変更するメソッドです。

4) book_search_engine (10min): slice/mapの練習
本検索エンジンを作成し、UTが通るようにメソッドを実装してください。
・構造体BookSearchEngineを定義し、フィールドとしてBookのスライスを持たせてください。
・SearchWithPrefixメソッドは、タイトルを前方一致で検索するメソッドです。
・GroupByPublisherメソッドは、本の一覧をPublisherごとに分類して返すメソッドです。

5) book_search_engine_goroutine (20min): goroutine, channelの練習
本検索エンジンBookSearchEngineを拡張し、UTが通るようにメソッドを実装してください。
・SearchWithPrefixMultiメソッドは、SearchWithPrefixの引数を複数渡せるように改良したメソッドです。
  ある本が複数のprefixに引っかかったとしても、同じタイトルの本はただ1つだけ返すようにしてください。
  また、探索結果は本のタイトルの昇順にソートしてください。
・SearchWithPrefixMultiAsyncメソッドは、SearchWithPrefixMultiの性能を改善したメソッドです。
  引数に与えられた各prefixによる探索処理を別スレッド(goroutine)で実行するようにし、最後にその結果を集約して返すようにしてください。

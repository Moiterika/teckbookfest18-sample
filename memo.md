# domainパッケージのテストカバレッジを向上させる方針

現在のdomainパッケージのテストカバレッジは27.1%と低めです。以下の方針でテストカバレッジを向上させることを提案します。

## 1. カバレッジ詳細の確認

カバレッジ情報を取得するコマンド：

```bash
go test -cover -coverprofile=coverage.out ./domain && go tool cover -func=coverage.out
```

## 2. 現在のカバレッジ状況

現在、以下のファイルのみがテストされています：
1. `按分_calc_test.go`
2. `按分結果_entity_test.go`
3. `ordered_map_test.go`

## 3. カバレッジ向上の戦略

カバレッジレポートと未テストファイルを分析した結果、以下の戦略を提案します：

### 優先順位の高い実装すべきテスト

1. **サービスレイヤーのテスト**
   - `仕訳_service.go` (0%)
   - `工数集計_service.go` (0%)
   - `配賦_service.go` (0%)

2. **エンティティやドメインモデルのテスト**
   - `仕訳_entity.go` (0%)
   - `勘定科目_entity.go` (0%)
   - `按分ルール_entity.go` (0%)
   - `集計仕訳_entity.go` (未カバー)

3. **データ構造のテスト**
   - `按分結果_list.go` (0%)
   - `集計仕訳_list.go` (0%)

### テスト実装アプローチ

#### 1. モックの作成

サービスのテストにはリポジトリのモックが必要です。モックを実装するための方針：

- [testify/mock](https://github.com/stretchr/testify) や [gomock](https://github.com/golang/mock) などのモックライブラリを使用
- インターフェースに対するモックを実装
- リポジトリの振る舞いをモックで制御

#### 2. テーブル駆動テスト

現在の `按分_calc_test.go` のようなテーブル駆動テストを他のコンポーネントにも実装：

- 様々な入力ケース
- エッジケース
- エラーケース

#### 3. 具体的なテストファイル作成計画

テストファイルを以下の順で作成することを推奨します：

1. **基本的なエンティティのテスト**
   - `仕訳_entity_test.go`
   - `勘定科目_entity_test.go`

2. **データ構造のテスト**
   - `按分結果_list_test.go`
   - `集計仕訳_list_test.go`

3. **複雑なドメインロジック**
   - `工数集計_service_test.go`
   - `仕訳_service_test.go`
   - `配賦_service_test.go`

## 4. サンプルテストコード

例として、`仕訳_entity_test.go` の実装例を示します：

```go
package domain

import (
    "testing"
    "time"
    
    "github.com/shopspring/decimal"
    "github.com/stretchr/testify/assert"
)

func Test仕訳Entity(t *testing.T) {
    // エンティティの基本データ
    発生日 := time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC)
    取引先 := "株式会社テスト"
    摘要 := "テスト仕訳"
    税率 := decimal.NewFromInt(10)
    
    // 仕訳詳細のテストデータ
    詳細 := []詳細{
        {
            勘定科目: "売上高",
            金額:    decimal.NewFromInt(1000),
            部門:    "営業部",
            税区分:   "課税",
        },
        {
            勘定科目: "売掛金",
            金額:    decimal.NewFromInt(1100),
            部門:    "営業部", 
            税区分:   "対象外",
        },
    }
    
    // テスト実行
    仕訳 := New仕訳(発生日, 取引先, 摘要, 詳細)
    
    // アサーション
    assert.Equal(t, 発生日, 仕訳.発生日)
    assert.Equal(t, 取引先, 仕訳.取引先)
    assert.Equal(t, 摘要, 仕訳.摘要)
    assert.Len(t, 仕訳.詳細, 2)
    
    // 詳細のテスト
    assert.Equal(t, "売上高", 仕訳.詳細[0].勘定科目)
    assert.Equal(t, decimal.NewFromInt(1000), 仕訳.詳細[0].金額)
    
    // Keyメソッドのテスト
    key := 仕訳.Key()
    assert.NotEmpty(t, key)
}
```

## 5. 実行計画

1. 最初に `仕訳_entity_test.go` と `按分結果_list_test.go` の実装から始める
2. テストを実装後、カバレッジを確認
3. カバレッジが低いコンポーネントへの対応を継続

このように段階的にテストを実装することで、全体のカバレッジを向上させることができます。モックの作成や複雑なテストケースの設計などは、実装を進めながら具体化していきましょう。

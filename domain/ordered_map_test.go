package domain

import (
	"testing"
)

// 基本機能のテスト：NewOrderedMap関数が正常に初期化されることを確認
func TestNewOrderedMap(t *testing.T) {
	// stringをキー、intを値とするOrderedMapを作成
	om := NewOrderedMap[string, int]()
	
	// 新しいマップは空であるべき
	if om.Count() != 0 {
		t.Errorf("新しいOrderedMapが空ではありません。Count: %d, 期待値: 0", om.Count())
	}
	
	// キーと値の取得を試みるが、存在しないはず
	val, exists := om.Get("存在しないキー")
	if exists {
		t.Errorf("存在しないキーが存在すると報告されました")
	}
	if val != 0 {
		t.Errorf("存在しないキーに対して、ゼロ値でない値が返されました: %d", val)
	}
}

// Set/Get機能のテスト
func TestOrderedMapSetAndGet(t *testing.T) {
	om := NewOrderedMap[string, int]()
	
	// キーと値をセット
	om.Set("キー1", 100)
	om.Set("キー2", 200)
	om.Set("キー3", 300)
	
	// 正しい値が取得できることを確認
	testCases := []struct {
		key      string
		expected int
		exists   bool
	}{
		{"キー1", 100, true},
		{"キー2", 200, true},
		{"キー3", 300, true},
		{"存在しないキー", 0, false},
	}
	
	for _, tc := range testCases {
		val, exists := om.Get(tc.key)
		if exists != tc.exists {
			t.Errorf("キー '%s' の存在チェック結果が期待と異なります。取得結果: %t, 期待値: %t", 
				tc.key, exists, tc.exists)
		}
		if val != tc.expected {
			t.Errorf("キー '%s' の値が期待と異なります。取得値: %d, 期待値: %d", 
				tc.key, val, tc.expected)
		}
	}
	
	// キーの上書きテスト
	om.Set("キー2", 250)
	val, exists := om.Get("キー2")
	if !exists {
		t.Errorf("上書きされたキーが存在しません")
	}
	if val != 250 {
		t.Errorf("上書きされた値が期待と異なります。取得値: %d, 期待値: 250", val)
	}
	
	// Count関数のテスト
	if om.Count() != 3 {
		t.Errorf("要素数が期待と異なります。Count: %d, 期待値: 3", om.Count())
	}
}

// Keys関数のテスト：キーが挿入順に返されることを確認
func TestOrderedMapKeys(t *testing.T) {
	om := NewOrderedMap[string, int]()
	
	// キーと値をセット
	expectedKeys := []string{"キー1", "キー2", "キー3"}
	for i, key := range expectedKeys {
		om.Set(key, (i+1)*100)
	}
	
	// キーが挿入順に取得できるか確認
	keys := om.Keys()
	if len(keys) != len(expectedKeys) {
		t.Errorf("キー数が期待と異なります。取得数: %d, 期待値: %d", len(keys), len(expectedKeys))
	}
	
	for i, expectedKey := range expectedKeys {
		if keys[i] != expectedKey {
			t.Errorf("キーの順序が期待と異なります。位置 %d で、取得値: %s, 期待値: %s", 
				i, keys[i], expectedKey)
		}
	}
	
	// キーの上書きが順序を変更しないことを確認
	om.Set("キー2", 999)
	keysAfterUpdate := om.Keys()
	
	for i, expectedKey := range expectedKeys {
		if keysAfterUpdate[i] != expectedKey {
			t.Errorf("上書き後のキーの順序が期待と異なります。位置 %d で、取得値: %s, 期待値: %s", 
				i, keysAfterUpdate[i], expectedKey)
		}
	}
}

// Values関数のテスト：値が挿入順に返されることを確認
func TestOrderedMapValues(t *testing.T) {
	om := NewOrderedMap[string, int]()
	
	// キーと値をセット
	keys := []string{"キー1", "キー2", "キー3"}
	expectedValues := []int{100, 200, 300}
	
	for i, key := range keys {
		om.Set(key, expectedValues[i])
	}
	
	// 値が挿入順に取得できるか確認
	values := om.Values()
	if len(values) != len(expectedValues) {
		t.Errorf("値の数が期待と異なります。取得数: %d, 期待値: %d", len(values), len(expectedValues))
	}
	
	for i, expectedValue := range expectedValues {
		if values[i] != expectedValue {
			t.Errorf("値の順序が期待と異なります。位置 %d で、取得値: %d, 期待値: %d", 
				i, values[i], expectedValue)
		}
	}
	
	// キーの上書きが値を更新して、順序が維持されることを確認
	om.Set("キー2", 999)
	expectedValuesAfterUpdate := []int{100, 999, 300}
	valuesAfterUpdate := om.Values()
	
	for i, expectedValue := range expectedValuesAfterUpdate {
		if valuesAfterUpdate[i] != expectedValue {
			t.Errorf("上書き後の値の順序が期待と異なります。位置 %d で、取得値: %d, 期待値: %d", 
				i, valuesAfterUpdate[i], expectedValue)
		}
	}
}

// Delete関数のテスト
func TestOrderedMapDelete(t *testing.T) {
	om := NewOrderedMap[string, int]()
	
	// キーと値をセット
	om.Set("キー1", 100)
	om.Set("キー2", 200)
	om.Set("キー3", 300)
	
	// キー2を削除
	result := om.Delete("キー2")
	if !result {
		t.Errorf("存在するキーの削除が失敗しました")
	}
	
	// 削除後の存在確認
	_, exists := om.Get("キー2")
	if exists {
		t.Errorf("削除されたはずのキーが存在しています")
	}
	
	// カウントが減少していることを確認
	if om.Count() != 2 {
		t.Errorf("削除後の要素数が期待と異なります。Count: %d, 期待値: 2", om.Count())
	}
	
	// 存在しないキーの削除は失敗するはず
	result = om.Delete("存在しないキー")
	if result {
		t.Errorf("存在しないキーの削除が成功してしまいました")
	}
	
	// 削除後のキーが正しい順序で残っていることを確認
	expectedKeysAfterDelete := []string{"キー1", "キー3"}
	keysAfterDelete := om.Keys()
	
	for i, expectedKey := range expectedKeysAfterDelete {
		if keysAfterDelete[i] != expectedKey {
			t.Errorf("削除後のキーの順序が期待と異なります。位置 %d で、取得値: %s, 期待値: %s", 
				i, keysAfterDelete[i], expectedKey)
		}
	}
}

// Range関数のテスト
func TestOrderedMapRange(t *testing.T) {
	om := NewOrderedMap[string, int]()
	
	// キーと値をセット
	keys := []string{"キー1", "キー2", "キー3", "キー4", "キー5"}
	for i, key := range keys {
		om.Set(key, (i+1)*100)
	}
	
	// すべての要素を集計
	count := 0
	sum := 0
	visitedKeys := make([]string, 0, len(keys))
	
	om.Range(func(key string, value int) bool {
		count++
		sum += value
		visitedKeys = append(visitedKeys, key)
		return true // 継続
	})
	
	// すべての要素が処理されたことを確認
	if count != 5 {
		t.Errorf("Range関数で処理された要素数が期待と異なります。処理数: %d, 期待値: 5", count)
	}
	
	// 合計値の確認
	expectedSum := 100 + 200 + 300 + 400 + 500
	if sum != expectedSum {
		t.Errorf("Range関数で計算された合計値が期待と異なります。合計値: %d, 期待値: %d", sum, expectedSum)
	}
	
	// キーの順序が維持されていることを確認
	for i, expectedKey := range keys {
		if visitedKeys[i] != expectedKey {
			t.Errorf("Range関数でのキーの順序が期待と異なります。位置 %d で、取得値: %s, 期待値: %s", 
				i, visitedKeys[i], expectedKey)
		}
	}
	
	// 中断機能のテスト
	count = 0
	om.Range(func(key string, value int) bool {
		count++
		return count < 3 // 3要素目より前で中断
	})
	
	if count != 3 {
		t.Errorf("Range関数の中断が期待通りに動作していません。処理数: %d, 期待値: 3", count)
	}
}

// 複雑な型でのテスト
type TestKey struct {
	ID   int
	Name string
}

type TestValue struct {
	Data   string
	Number float64
}

func TestOrderedMapWithComplexTypes(t *testing.T) {
	// 複雑な型を使ったOrderedMapを作成
	om := NewOrderedMap[TestKey, TestValue]()
	
	// キーと値をセット
	key1 := TestKey{ID: 1, Name: "テスト1"}
	key2 := TestKey{ID: 2, Name: "テスト2"}
	
	value1 := TestValue{Data: "データ1", Number: 1.23}
	value2 := TestValue{Data: "データ2", Number: 4.56}
	
	om.Set(key1, value1)
	om.Set(key2, value2)
	
	// 取得テスト
	val1, exists1 := om.Get(key1)
	if !exists1 {
		t.Errorf("複雑なキー1が存在しません")
	}
	if val1.Data != "データ1" || val1.Number != 1.23 {
		t.Errorf("複雑な値1が期待と異なります。取得値: %+v, 期待値: %+v", val1, value1)
	}
	
	// 削除テスト
	om.Delete(key1)
	_, exists1After := om.Get(key1)
	if exists1After {
		t.Errorf("削除された複雑なキー1が存在しています")
	}
	
	// 存在するキー2の確認
	val2, exists2 := om.Get(key2)
	if !exists2 {
		t.Errorf("複雑なキー2が存在しません")
	}
	if val2.Data != "データ2" || val2.Number != 4.56 {
		t.Errorf("複雑な値2が期待と異なります。取得値: %+v, 期待値: %+v", val2, value2)
	}
}

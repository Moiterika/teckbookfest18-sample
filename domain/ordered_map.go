package domain

// OrderedMap は挿入順を保持するジェネリックなマップ構造体です。
// K は比較可能な型 (マップのキーとして使用できる型) である必要があります。
// V は任意の値の型です。
type OrderedMap[K comparable, V any] struct {
	keys []K     // キーの挿入順を保持するスライス
	m    map[K]V // キーと値を保持するマップ
}

// NewOrderedMap は新しいジェネリックな OrderedMap を作成します。
// 使用時にキーと値の型を指定します (例: NewOrderedMap[string, int]())。
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys: make([]K, 0),  // キーの型のスライスを初期化
		m:    make(map[K]V), // キーと値の型のマップを初期化
	}
}

// Set はキーと値を追加/更新し、挿入順を記憶します。
// もしキーが存在しなければ、挿入順リストの末尾に追加されます。
func (om *OrderedMap[K, V]) Set(key K, value V) {
	if _, exists := om.m[key]; !exists {
		om.keys = append(om.keys, key) // 新規キーの場合のみ順序リストに追加
	}
	om.m[key] = value // マップに値を追加または更新
}

// Get はキーに対応する値を取得します。
// キーが存在しない場合は、V 型のゼロ値と false を返します。
func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	val, exists := om.m[key]
	// exists が false の場合、val は V 型のゼロ値になります。
	return val, exists
}

func (om *OrderedMap[K, V]) Count() int {
	return len(om.keys)
}

// Keys は挿入された順序でキーのスライスを返します。
// 返されるスライスは内部状態のコピーなので、変更しても元の OrderedMap には影響しません。
func (om *OrderedMap[K, V]) Keys() []K {
	keysCopy := make([]K, len(om.keys))
	copy(keysCopy, om.keys)
	return keysCopy
}

// Values は挿入された順序で値のスライスを返します。
// 返されるスライスは内部状態のコピーなので、変更しても元の OrderedMap には影響しません。
func (om *OrderedMap[K, V]) Values() []V {
	values := make([]V, 0, len(om.keys))
	for _, key := range om.keys {
		values = append(values, om.m[key]) // マップから値を取得して追加
	}
	return values
}

// Delete は指定されたキーを持つエントリをマップから削除します。
// 削除に成功した場合は true を、キーが存在しなかった場合は false を返します。
// 注意: スライスからの削除は要素数が多い場合に効率が良くありません。
func (om *OrderedMap[K, V]) Delete(key K) bool {
	if _, exists := om.m[key]; !exists {
		return false // キーが存在しない
	}
	// マップから削除
	delete(om.m, key)

	// 順序リスト (スライス) からキーを削除
	newKeys := make([]K, 0, len(om.keys)-1)
	for _, k := range om.keys {
		// comparable 制約により、キー同士を比較可能
		if k != key {
			newKeys = append(newKeys, k)
		}
	}
	om.keys = newKeys // 更新されたスライスで置き換え
	return true
}

// Range は挿入された順序で、各キーと値に対して指定された関数 f を実行します。
// 関数 f が false を返した場合、反復処理は中断されます。
func (om *OrderedMap[K, V]) Range(f func(key K, value V) bool) {
	// 内部の keys スライスを直接使うことで、Keys() でのコピーを避ける
	for _, key := range om.keys {
		// マップにはキーが存在するはずなので、直接アクセスする
		value := om.m[key]
		if !f(key, value) { // コールバック関数を実行
			break // false が返されたら中断
		}
	}
}

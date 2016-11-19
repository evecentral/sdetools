package sdetools

import (
	"testing"
)

var yamlGroups = []byte(`
27:
    anchorable: false
    anchored: false
    categoryID: 6
    fittableNonSingleton: false
    name:
        de: Schlachtschiff
        en: Battleship
        fr: Cuirassé
        ja: 戦艦
        ru: Линкор
        zh: 战列舰
    published: true
    useBasePrice: false
28:
    anchorable: false
    anchored: false
    categoryID: 6
    fittableNonSingleton: false
    name:
        de: Industrial
        en: Industrial
        fr: Vaisseau industriel
        ja: 輸送艦
        ru: Грузовой корабль
        zh: 工业舰
    published: true
    useBasePrice: false
29:
    anchorable: false
    anchored: false
    categoryID: 6
    fittableNonSingleton: false
    iconID: 73
    name:
        de: Kapsel
        en: Capsule
        fr: Capsule
        ja: カプセル
        ru: Капсула
        zh: 太空舱
    published: false
    useBasePrice: false
30:
    anchorable: false
    anchored: false
    categoryID: 6
    fittableNonSingleton: false
    name:
        de: Titan
        en: Titan
        fr: Titan
        ja: タイタン
        ru: Титан
        zh: 泰坦
    published: true
    useBasePrice: false


`)

func TestLoadGroup(t *testing.T) {
	groups, err := LoadGroups(yamlGroups)
	if err != nil {
		t.Error(err)
	}
	if len(*groups) < 3 {
		t.Errorf("Decoded list was too small")
	}
}

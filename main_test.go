package handles

import (
	"testing"
)

func TestGenerateUniqueHandle(t *testing.T) {
	tests := []struct {
		existingHandles map[string]bool
		firstName       string
		lastName        string
		expected        string
	}{
		{map[string]bool{"NaDu": true}, "Nathalie", "Du", "NatDu"},
		{map[string]bool{"NaDu": true, "NatDu": true}, "Nathalie", "Du", "NathDu"},
		{map[string]bool{"NaDu": true, "NatDu": true, "NathDu": true}, "Nathalie", "Du", "NathaDu"},
		{map[string]bool{"NDu": true}, "N", "Du", "NDu1"},
		{map[string]bool{"NaD": true}, "Na", "D", "NaD1"},
		{map[string]bool{"NaDu": true, "NatDu": true, "NathDu": true, "NathaDu": true, "NathalDu": true, "NathaliDu": true}, "Nathalie", "Du", "NathalieDu"},
		{map[string]bool{"NaD": true}, "Na", "D", "NaD1"},
		{map[string]bool{}, "Malte", "Meiners", "MaMe"},
		{map[string]bool{"MaMe": true}, "Malte", "Meiners", "MaMei"},
		{map[string]bool{"MaMe": true, "MaMei": true}, "Malte", "Meiners", "MalMei"},
		{map[string]bool{"MaMe": true, "MaMei": true, "MalMei": true}, "Malte", "Meiners", "MalMein"},
		{map[string]bool{"MaMe": true, "MaMei": true, "MalMei": true, "MalMein": true}, "Malte", "Meiners", "MaltMein"},
		{map[string]bool{"MaMe": true, "MaMei": true, "MalMei": true, "MalMein": true, "MaltMein": true}, "Malte", "Meiners", "MaltMeine"},
		{map[string]bool{"MaMe": true, "MaMei": true, "MalMei": true, "MalMein": true, "MaltMein": true, "MaltMeine": true}, "Malte", "Meiners", "MalteMeine"},
		{map[string]bool{"MaMe": true, "MaMei": true, "MalMei": true, "MalMein": true, "MaltMein": true, "MaltMeine": true, "MalteMeine": true}, "Malte", "Meiners", "MalteMeiner"},
		{map[string]bool{"MaMe": true, "MaMei": true, "MalMei": true, "MalMein": true, "MaltMein": true, "MaltMeine": true, "MalteMeine": true, "MalteMeiner": true}, "Malte", "Meiners", "MalteMeiners"},
		{map[string]bool{"MaMe": true, "MaMei": true, "MalMei": true, "MalMein": true, "MaltMein": true, "MaltMeine": true, "MalteMeine": true, "MalteMeiner": true, "MalteMeiners": true}, "Malte", "Meiners", "MalteMeiners1"},
		{map[string]bool{}, "Bo", "Chang", "BoCh"},
		{map[string]bool{"BoCh": true}, "Bo", "Chang", "BoCha"},
		{map[string]bool{"BoCh": true, "BoCha": true}, "Bo", "Chang", "BoChan"},
		{map[string]bool{"BoCh": true, "BoCha": true, "BoChan": true}, "Bo", "Chang", "BoChang"},
		{map[string]bool{"BoCh": true, "BoCha": true, "BoChan": true, "BoChang": true}, "Bo", "Chang", "BoChang1"},
		{map[string]bool{"BoCh": true, "BoCha": true, "BoChan": true, "BoChang": true, "BoChang1": true}, "Bo", "Chang", "BoChang2"},
		{map[string]bool{}, "X", "Chang", "XCha"},
		{map[string]bool{"XCha": true}, "X", "Chang", "XChan"},
		{map[string]bool{"XCha": true, "XChan": true}, "X", "Chang", "XChang"},
		{map[string]bool{"XCha": true, "XChan": true, "XChang": true}, "X", "Chang", "XChang1"},
		{map[string]bool{"XCha": true, "XChan": true, "XChang": true, "XChang1": true}, "X", "Chang", "XChang2"},
		{map[string]bool{}, "Chang", "X", "ChaX"},
		{map[string]bool{"ChaX": true}, "Chang", "X", "ChanX"},
		{map[string]bool{"ChaX": true, "ChanX": true}, "Chang", "X", "ChangX"},
		{map[string]bool{"ChaX": true, "ChanX": true, "ChangX": true}, "Chang", "X", "ChangX1"},
		{map[string]bool{"ChaX": true, "ChanX": true, "ChangX": true, "ChangX1": true}, "Chang", "X", "ChangX2"},
		{map[string]bool{}, "X", "Y", "XY"},
		{map[string]bool{"XY": true}, "X", "Y", "XY1"},
		{map[string]bool{"XY": true, "XY1": true}, "X", "Y", "XY2"},
		{map[string]bool{}, "Jean-Claude", "Baumann", "JeBa"},
		{map[string]bool{"JeBa": true}, "Jean-Claude", "Baumann", "JeBau"},
		{map[string]bool{"JeBa": true, "JeBau": true}, "Jean-Claude", "Baumann", "JeaBau"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true}, "Jean-Claude", "Baumann", "JeaBaum"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true}, "Jean-Claude", "Baumann", "JeanBaum"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true}, "Jean-Claude", "Baumann", "JeanBauma"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true}, "Jean-Claude", "Baumann", "JeanCBauma"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true}, "Jean-Claude", "Baumann", "JeanCBauman"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true, "JeanCBauman": true}, "Jean-Claude", "Baumann", "JeanClBauman"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true, "JeanCBauman": true, "JeanClBauman": true}, "Jean-Claude", "Baumann", "JeanClBaumann"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true, "JeanCBauman": true, "JeanClBauman": true, "JeanClBaumann": true}, "Jean-Claude", "Baumann", "JeanClaBaumann"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true, "JeanCBauman": true, "JeanClBauman": true, "JeanClBaumann": true, "JeanClaBaumann": true}, "Jean-Claude", "Baumann", "JeanClauBaumann"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true, "JeanCBauman": true, "JeanClBauman": true, "JeanClBaumann": true, "JeanClaBaumann": true, "JeanClauBaumann": true}, "Jean-Claude", "Baumann", "JeanClaudBaumann"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true, "JeanCBauman": true, "JeanClBauman": true, "JeanClBaumann": true, "JeanClaBaumann": true, "JeanClauBaumann": true, "JeanClaudBaumann": true}, "Jean-Claude", "Baumann", "JeanClaudeBaumann"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true, "JeanCBauman": true, "JeanClBauman": true, "JeanClBaumann": true, "JeanClaBaumann": true, "JeanClauBaumann": true, "JeanClaudBaumann": true, "JeanClaudeBaumann": true}, "Jean-Claude", "Baumann", "JeanClaudeBaumann1"},
		{map[string]bool{"JeBa": true, "JeBau": true, "JeaBau": true, "JeaBaum": true, "JeanBaum": true, "JeanBauma": true, "JeanCBauma": true, "JeanCBauman": true, "JeanClBauman": true, "JeanClBaumann": true, "JeanClaBaumann": true, "JeanClauBaumann": true, "JeanClaudBaumann": true, "JeanClaudeBaumann": true, "JeanClaudeBaumann1": true}, "Jean-Claude", "Baumann", "JeanClaudeBaumann2"},
	}
	for _, tt := range tests {
		actual := createHandle(tt.existingHandles, tt.firstName, tt.lastName)
		if actual != tt.expected {
			t.Errorf("expected %s; got %s", tt.expected, actual)
		}
		if !tt.existingHandles[actual] {
			t.Errorf("new handle was not added to existingHandles map")
		}
	}
}

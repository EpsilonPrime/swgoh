package swgohgg

import (
	"testing"
)

func TestArena(t *testing.T) {
	gg := NewClient("ronoaldo")
	team, update, err := gg.Arena()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ally Code %s", gg.AllyCode())
	if gg.AllyCode() == "" {
		t.Errorf("Unexpected empty ally code returned.")
	}
	t.Logf("Player/Guild: %s/%s", gg.PlayerName(), gg.GuildName())
	count := 0
	for i := range team {
		char := team[i]
		t.Logf("Team member: %v", char)
		count++
	}
	if count != 5 {
		t.Errorf("Got %d arena characters. Expected 5", count)
	}
	t.Logf("Last update: %v", update)
	if update.IsZero() {
		t.Errorf("Zero last update timstamp.")
	}

	_, _, err = gg.Profile("").SetAllyCode("335983287").Arena()
	if err != nil {
		t.Errorf("Calling API with only allyCode fails! (err=%v)", err)
	}
}

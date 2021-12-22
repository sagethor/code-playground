package main

import (
	"bytes"
	"encoding/gob"
	"os"
)

// abandoning the idea of uint64/uint32 <-> uint8 (for now?)
type Player struct {
	name string
	w, x, y, z uint8
}

func(player *Player) newPlayer() {

}


// to-do:
// 1) check if player already exists in active array
// 2) load player to active array
// 3) function to unload if array full / logout & consider timed saves
// 4) verify if this is best practice for file opening and interchangability with network / non-flatfile methods.
func loadPlayer(name string) *player {
	p := player{name: name}

	file, _ := os.Open(name + ".gob");
	defer file.Close();
	// create dedicated encoder/decoder for goroutine?
	dec := gob.NewDecoder(file);
	dec.Decode(&player);

	if err := dec.Decode(&m); err != nil {
		log.Fatal(err);
	}

	return &p
}

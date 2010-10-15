package builder

import "os"
import "path"
import "json"
import "io/ioutil"

func LoadWorld(file string) (world *World, err os.Error) {
	var data []byte

	if data, err = ioutil.ReadFile(path.Clean(file)); err != nil {
		return
	}

	world = new(World)
	err = json.Unmarshal(data, &world)
	return
}

func SaveWorld(file string, world *World) (err os.Error) {
	var data []byte

	if data, err = json.Marshal(world); err != nil {
		return
	}

	return ioutil.WriteFile(path.Clean(file), data, 0600)
}

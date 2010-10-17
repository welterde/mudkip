package builder

import "os"
import "path"
import "json"
import "io/ioutil"

// Load a world from a JSON formatted data file.
func LoadWorld(file string) (world *World, err os.Error) {
	var data []byte

	if data, err = ioutil.ReadFile(path.Clean(file)); err != nil {
		return
	}

	world = new(World)
	err = json.Unmarshal(data, &world)
	return
}

// Saves the world to a JSON formatted data file. Optionally with indentation
// for easy reading/modification.
func SaveWorld(file string, world *World, compact bool) (err os.Error) {
	var data []byte

	if compact {
		if data, err = json.Marshal(world); err != nil {
			return
		}
	} else {
		if data, err = json.MarshalIndent(world, "", "  "); err != nil {
			return
		}
	}

	return ioutil.WriteFile(path.Clean(file), data, 0600)
}

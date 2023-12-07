package main

import (
	"advent-of-code/helpers"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"sync"
)

var almanacDocument string = "/05-almanac/almanac.txt"

type Almanac struct {
	Seeds []int
	//SeedToSoilRows            []MapItem
	SeedToSoilMap map[int]int
	//SoilToFertilizerRows      []MapItem
	SoilToFertilizerMap map[int]int
	//FertilizerToWaterRows     []MapItem
	FertilizerToWaterMap map[int]int
	//WaterToLightRows          []MapItem
	WaterToLightMap map[int]int
	//LightToTemperatureRows    []MapItem
	LightToTemperatureMap map[int]int
	//TemperatureToHumidityRows []MapItem
	TemperatureToHumidityMap map[int]int
	//HumidityToLocationRows    []MapItem
	HumidityToLocationMap map[int]int
	MinLocation           int
}

type MapItem struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

const (
	SeedToSoil            string = "seed-to-soil map:"
	SoilToFertilizer      string = "soil-to-fertilizer map:"
	FertilizerToWater     string = "fertilizer-to-water map:"
	WaterToLight          string = "water-to-light map:"
	LightToTemperature    string = "light-to-temperature map:"
	TemperatureToHumidity string = "temperature-to-humidity map:"
	HumidityToLocation    string = "humidity-to-location map:"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	almanac, err := run(wd + almanacDocument)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Lowest Location: %d\n", almanac.MinLocation)
}

func run(fileName string) (Almanac, error) {
	lines, err := getLinesFromFile(fileName)
	if err != nil {
		panic(err)
	}

	almanac, err := getAlmanacForLines(lines)
	if err != nil {
		panic(err)
	}

	return almanac, nil
}

func getLinesFromFile(fileName string) ([]string, error) {
	return helpers.GetFileLines(fileName)
}

func getAlmanacForLines(lines []string) (Almanac, error) {
	almanac := Almanac{
		SeedToSoilMap:            map[int]int{},
		SoilToFertilizerMap:      map[int]int{},
		FertilizerToWaterMap:     map[int]int{},
		WaterToLightMap:          map[int]int{},
		LightToTemperatureMap:    map[int]int{},
		TemperatureToHumidityMap: map[int]int{},
		HumidityToLocationMap:    map[int]int{},
		MinLocation:              math.MaxInt,
	}
	if len(lines) < 22 {
		return almanac, fmt.Errorf("not enough lines")
	}

	re := regexp.MustCompile(`(\d+)`)
	seedsSlice := re.FindAllString(lines[0], -1)

	var err error
	almanac.Seeds, err = helpers.ParseStringSliceToIntSlice(seedsSlice)
	if err != nil {
		return almanac, err
	}

	currentLineType := ""

	var wg sync.WaitGroup

	// loop over lines skipping first line
	for i, line := range lines[1:] {
		switch line {
		case
			SeedToSoil,
			SoilToFertilizer,
			FertilizerToWater,
			WaterToLight,
			LightToTemperature,
			TemperatureToHumidity,
			HumidityToLocation,
			"":
			currentLineType = line
			continue
		default:
		}

		wg.Add(1)

		stringSlice := re.FindAllString(line, -1)
		intSlice, err := helpers.ParseStringSliceToIntSlice(stringSlice)
		if err != nil {
			return almanac, err
		}

		if len(intSlice) != 3 {
			return almanac, fmt.Errorf("expected 3 numbers per row at row %d, line '%s'", i+1, line)
		}

		mapItem := MapItem{
			DestinationRangeStart: intSlice[0],
			SourceRangeStart:      intSlice[1],
			RangeLength:           intSlice[2],
		}

		go func(l string, m MapItem) {
			almanac.updateAlmanacWithMapItem(l, m)
			wg.Done()
		}(currentLineType, mapItem)
	}

	wg.Wait()

	//almanac.fillEmptyListIndicies()

	almanac.updateLowestLocationFromAlmanac()

	return almanac, nil
}

var stsLock sync.RWMutex
var stfLock sync.RWMutex
var ftwLock sync.RWMutex
var wtlLock sync.RWMutex
var lttLock sync.RWMutex
var tthLock sync.RWMutex
var htlLock sync.RWMutex

func (almanac *Almanac) updateAlmanacWithMapItem(currentLineType string, mapItem MapItem) {
	switch currentLineType {
	case SeedToSoil:
		//almanac.SeedToSoilRows = append(almanac.SeedToSoilRows, mapItem)
		for j := 0; j < mapItem.RangeLength; j++ {
			stsLock.Lock()
			almanac.SeedToSoilMap[mapItem.SourceRangeStart+j] = mapItem.DestinationRangeStart + j
			stsLock.Unlock()
		}
	case SoilToFertilizer:
		//almanac.SoilToFertilizerRows = append(almanac.SoilToFertilizerRows, mapItem)
		for j := 0; j < mapItem.RangeLength; j++ {
			stfLock.Lock()
			almanac.SoilToFertilizerMap[mapItem.SourceRangeStart+j] = mapItem.DestinationRangeStart + j
			stfLock.Unlock()
		}
	case FertilizerToWater:
		//almanac.FertilizerToWaterRows = append(almanac.FertilizerToWaterRows, mapItem)
		for j := 0; j < mapItem.RangeLength; j++ {
			ftwLock.Lock()
			almanac.FertilizerToWaterMap[mapItem.SourceRangeStart+j] = mapItem.DestinationRangeStart + j
			ftwLock.Unlock()
		}
	case WaterToLight:
		// almanac.WaterToLightRows = append(almanac.WaterToLightRows, mapItem)
		for j := 0; j < mapItem.RangeLength; j++ {
			wtlLock.Lock()
			almanac.WaterToLightMap[mapItem.SourceRangeStart+j] = mapItem.DestinationRangeStart + j
			wtlLock.Unlock()
		}
	case LightToTemperature:
		// almanac.LightToTemperatureRows = append(almanac.LightToTemperatureRows, mapItem)
		for j := 0; j < mapItem.RangeLength; j++ {
			lttLock.Lock()
			almanac.LightToTemperatureMap[mapItem.SourceRangeStart+j] = mapItem.DestinationRangeStart + j
			lttLock.Unlock()
		}
	case TemperatureToHumidity:
		// almanac.TemperatureToHumidityRows = append(almanac.TemperatureToHumidityRows, mapItem)
		for j := 0; j < mapItem.RangeLength; j++ {
			tthLock.Lock()
			almanac.TemperatureToHumidityMap[mapItem.SourceRangeStart+j] = mapItem.DestinationRangeStart + j
			tthLock.Unlock()
		}
	case HumidityToLocation:
		// almanac.HumidityToLocationRows = append(almanac.HumidityToLocationRows, mapItem)
		for j := 0; j < mapItem.RangeLength; j++ {
			htlLock.Lock()
			almanac.HumidityToLocationMap[mapItem.SourceRangeStart+j] = mapItem.DestinationRangeStart + j
			htlLock.Unlock()
		}
	default:
	}
}

//nolint:unused
func (almanac *Almanac) fillEmptyListIndicies() {
	almanac.SeedToSoilMap = fillEmptyIndicies(almanac.SeedToSoilMap, map[int]int{})
	almanac.SoilToFertilizerMap = fillEmptyIndicies(almanac.SoilToFertilizerMap, almanac.SeedToSoilMap)
	almanac.FertilizerToWaterMap = fillEmptyIndicies(almanac.FertilizerToWaterMap, almanac.SoilToFertilizerMap)
	almanac.WaterToLightMap = fillEmptyIndicies(almanac.WaterToLightMap, almanac.FertilizerToWaterMap)
	almanac.LightToTemperatureMap = fillEmptyIndicies(almanac.LightToTemperatureMap, almanac.WaterToLightMap)
	almanac.TemperatureToHumidityMap = fillEmptyIndicies(almanac.TemperatureToHumidityMap, almanac.LightToTemperatureMap)
	almanac.HumidityToLocationMap = fillEmptyIndicies(almanac.HumidityToLocationMap, almanac.TemperatureToHumidityMap)
}

//nolint:unused
func fillEmptyIndicies(input map[int]int, sourceMap map[int]int) map[int]int {
	var maxSourceValue int
	if len(sourceMap) > 0 {
		for _, source := range sourceMap {
			if maxSourceValue < source {
				maxSourceValue = source
			}
		}
	} else {
		// map containing source values
		keys := []int{}

		// get all source values
		for _, source := range input {
			keys = append(keys, source)
		}

		// sort so last index is highest
		slices.Sort(keys)

		maxSourceValue = keys[len(keys)-1]
	}

	// fill any non-mapped source values to equal itself e.g 13 - 13 if not mapped
	for i := 0; i < maxSourceValue; i++ {
		_, ok := input[i]
		if !ok {
			input[i] = i
		}
	}
	return input
}

func getDestinationFromMap(input map[int]int, source int) int {
	destination, ok := input[source]
	if !ok {
		return source
	}
	return destination
}

func (almanac *Almanac) updateLowestLocationFromAlmanac() {
	for _, seed := range almanac.Seeds {

		soil := getDestinationFromMap(almanac.SeedToSoilMap, seed)

		fertilizer := getDestinationFromMap(almanac.SoilToFertilizerMap, soil)

		water := getDestinationFromMap(almanac.FertilizerToWaterMap, fertilizer)

		light := getDestinationFromMap(almanac.WaterToLightMap, water)

		temperature := getDestinationFromMap(almanac.LightToTemperatureMap, light)

		humidity := getDestinationFromMap(almanac.TemperatureToHumidityMap, temperature)

		location := getDestinationFromMap(almanac.HumidityToLocationMap, humidity)

		fmt.Printf("seed %d, location %d\n", seed, location)
		if location < almanac.MinLocation {
			almanac.MinLocation = location
		}
	}
}

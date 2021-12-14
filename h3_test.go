package h3

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromGeo(t *testing.T) {

	req := require.New(t)

	file, err := os.Open("test_data.txt")
	if err != nil {
		t.FailNow()
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(arr[0], 64)
		if err != nil {
			t.FailNow()
		}
		long, err := strconv.ParseFloat(arr[1], 64)
		if err != nil {
			t.FailNow()
		}
		exp, err := strconv.ParseUint(arr[2], 10, 64)
		if err != nil {
			t.FailNow()
		}

		res := FromGeo(GeoCoord{
			Latitude:  lat,
			Longitude: long,
		}, 11)

		req.Equal(exp, uint64(res))
	}
}
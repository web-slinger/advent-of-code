package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSumOfCalibrationValues(t *testing.T) {

	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	sumOfValues, err := getSumOfCalibrationValues(lines)
	require.NoError(t, err)

	assert.Equal(t, 281, sumOfValues)
}

func TestGetCalibrationValueFromLine(t *testing.T) {
	tests := []struct {
		line                string
		expCalibrationValue int
	}{
		{
			line:                "eightsevenvqvzlqxkbm6rqhsgqpnine7twonex",
			expCalibrationValue: 81,
		},
		{
			line:                "zoneight234",
			expCalibrationValue: 14,
		},
		{
			line:                "hgn6mnpmdmcpzceighteightnxkvjjfrninejmpkrfzcgv3",
			expCalibrationValue: 63,
		},
		{
			line:                "qhsevenvpg4hzffldbrvpxxpthreeqpvvdndv",
			expCalibrationValue: 73,
		},
		{
			line:                "fourpcjhfjrxdhvzf2dkmszvtjx",
			expCalibrationValue: 42,
		},
		{
			line:                "ninegldjhplfthreetnqcbrllpvjtlthn9xkbqkfourthree",
			expCalibrationValue: 93,
		},
		{
			line:                "hpspsgtfxvxtmdsqcninelcjhfb2mhffpvxkdxdlvkqxnine",
			expCalibrationValue: 99,
		},
		{
			line:                "fivexrhxhtfivesevenone3d",
			expCalibrationValue: 53,
		},
		{
			line:                "doneightthtpmjlzhgpxdc18229twofive",
			expCalibrationValue: 15,
		},
		{
			line:                "sixmone3phrxxdninetwosix",
			expCalibrationValue: 66,
		},
		{
			line:                "shnnn3nqcgfbgpzzfrtchbseven5dk3",
			expCalibrationValue: 33,
		},
		{
			line:                "onethreetwofive5sevenfdmrmczlqs",
			expCalibrationValue: 17,
		},
		{
			line:                "fivetmbnqlchhtqmbcsssvxjzvlxdvznlbfive7",
			expCalibrationValue: 57,
		},
		{
			line:                "fq2qbmone",
			expCalibrationValue: 21,
		},
		{
			line:                "rczhlkcqpfkgcjmcggztjlqsgxmdxstwo88bdxmqjvlfl",
			expCalibrationValue: 28,
		},
		{
			line:                "grbhpnjrtvrbslnfgthree47vbpncxqfourfp",
			expCalibrationValue: 34,
		},
		{
			line:                "threektqgtgcrccbsnsqpcfxtb3vxtdfour2hgvdg",
			expCalibrationValue: 32,
		},
		{
			line:                "5443",
			expCalibrationValue: 53,
		},
		{
			line:                "eighthree",
			expCalibrationValue: 83,
		},
	}

	for _, tc := range tests {
		t.Run(tc.line, func(t *testing.T) {
			numbers := getDigitsFromLineIncludingStringNumbers(tc.line)
			require.Len(t, numbers, 2)

			ActualCalibrationValue, err := strconv.Atoi(numbers[0] + numbers[1])
			require.NoError(t, err)

			assert.Equal(t, tc.expCalibrationValue, ActualCalibrationValue)
		})
	}
}

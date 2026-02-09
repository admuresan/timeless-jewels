package main

import (
	"strconv"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

// assertValidResult checks that the result from Calculate is well-formed and deterministic.
func assertValidResult(t *testing.T, jewel data.JewelType, conqueror data.Conqueror, passive uint32, seed uint32) {
	t.Helper()
	result1 := calculator.Calculate(passive, seed, jewel, conqueror)
	result2 := calculator.Calculate(passive, seed, jewel, conqueror)
	testza.AssertEqual(t, result1, result2, "Calculate should be deterministic")

	hasSkill := result1.AlternatePassiveSkill != nil
	hasAdditions := len(result1.AlternatePassiveAdditionInformations) > 0
	testza.AssertTrue(t, hasSkill || hasAdditions)
}

func TestGloriousVanity(t *testing.T) {
	const seed = 2000
	passives := []uint32{2286, 411, 519, 1190, 88}
	conquerors := []data.Conqueror{data.Xibaqua, data.Zerphi, data.Ahuana, data.Doryani}

	for _, conq := range conquerors {
		t.Run(string(conq), func(t *testing.T) {
			for _, passive := range passives {
				t.Run(strconv.Itoa(int(passive)), func(t *testing.T) {
					assertValidResult(t, data.GloriousVanity, conq, passive, seed)
				})
			}
		})
	}
}

func TestLethalPride(t *testing.T) {
	const seed = 12000
	passives := []uint32{2286, 411, 519, 1190, 88}
	conquerors := []data.Conqueror{data.Kaom, data.Rakiata, data.Kiloava, data.Akoya}

	for _, conq := range conquerors {
		t.Run(string(conq), func(t *testing.T) {
			for _, passive := range passives {
				t.Run(strconv.Itoa(int(passive)), func(t *testing.T) {
					assertValidResult(t, data.LethalPride, conq, passive, seed)
				})
			}
		})
	}
}

func TestBrutalRestraint(t *testing.T) {
	const seed = 2000
	passives := []uint32{2286, 411, 519, 1190, 88}
	conquerors := []data.Conqueror{data.Deshret, data.Balbala, data.Asenath, data.Nasima}

	for _, conq := range conquerors {
		t.Run(string(conq), func(t *testing.T) {
			for _, passive := range passives {
				t.Run(strconv.Itoa(int(passive)), func(t *testing.T) {
					assertValidResult(t, data.BrutalRestraint, conq, passive, seed)
				})
			}
		})
	}
}

func TestMilitantFaith(t *testing.T) {
	const seed = 2000
	passives := []uint32{2286, 411, 519, 1190, 88}
	conquerors := []data.Conqueror{data.Venarius, data.Maxarius, data.Dominus, data.Avarius}

	for _, conq := range conquerors {
		t.Run(string(conq), func(t *testing.T) {
			for _, passive := range passives {
				t.Run(strconv.Itoa(int(passive)), func(t *testing.T) {
					assertValidResult(t, data.MilitantFaith, conq, passive, seed)
				})
			}
		})
	}
}

func TestElegantHubris(t *testing.T) {
	const seed = 2000
	passives := []uint32{2286, 411, 519, 1190, 88}
	conquerors := []data.Conqueror{data.Cadiro, data.Victario, data.Chitus, data.Caspiro}

	for _, conq := range conquerors {
		t.Run(string(conq), func(t *testing.T) {
			for _, passive := range passives {
				t.Run(strconv.Itoa(int(passive)), func(t *testing.T) {
					assertValidResult(t, data.ElegantHubris, conq, passive, seed)
				})
			}
		})
	}
}

func BenchmarkAll(b *testing.B) {
	applicable := data.GetApplicablePassives()

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		for jewelType := range data.TimelessJewelConquerors {
			var firstConqueror data.Conqueror
			for conqueror := range data.TimelessJewelConquerors[jewelType] {
				firstConqueror = conqueror
				break
			}

			seedMin := data.TimelessJewelSeedRanges[jewelType].Min
			seedMax := data.TimelessJewelSeedRanges[jewelType].Max

			if data.TimelessJewelSeedRanges[jewelType].Special {
				seedMin /= 20
				seedMax /= 20
			}

			for seed := seedMin; seed <= seedMax; seed++ {
				realSeed := seed
				if data.TimelessJewelSeedRanges[jewelType].Special {
					realSeed *= 20
				}

				for _, skill := range applicable {
					calculator.Calculate(skill.Index, realSeed, jewelType, firstConqueror)
				}
			}
		}
	}
}

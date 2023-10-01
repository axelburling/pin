package utils

type Pin struct {
	FullYear     string
	Century      string
	Year         string
	Month        string
	Day          string
	Place        string
	Gender       string
	ControlDigit string
}

var (
	countyNumbers = map[int]string{
		0:  "Stockholm",
		1:  "Stockholm",
		2:  "Stockholm",
		3:  "Stockholm",
		4:  "Stockholm",
		5:  "Stockholm",
		6:  "Stockholm",
		7:  "Stockholm",
		8:  "Stockholm",
		9:  "Stockholm",
		10: "Stockholm",
		11: "Stockholm",
		12: "Stockholm",
		13: "Stockholm",
		14: "Uppsala",
		15: "Uppsala",
		16: "Södermanland",
		17: "Södermanland",
		18: "Södermanland",
		19: "Östergötland",
		20: "Östergötland",
		21: "Östergötland",
		22: "Östergötland",
		23: "Östergötland",
		24: "Jönköping",
		25: "Jönköping",
		26: "Jönköping",
		27: "Kronoberg",
		28: "Kronoberg",
		29: "Kalmar",
		30: "Kalmar",
		31: "Kalmar",
		32: "Gotland",
		33: "Blekinge",
		34: "Blekinge",
		35: "Kristianstad",
		36: "Kristianstad",
		37: "Kristianstad",
		38: "Kristianstad",
		39: "Malmöhus",
		40: "Malmöhus",
		41: "Malmöhus",
		42: "Malmöhus",
		43: "Malmöhus",
		44: "Malmöhus",
		45: "Malmöhus",
		46: "Halland",
		47: "Halland",
		48: "Göteborg och Bohus",
		49: "Göteborg och Bohus",
		50: "Göteborg och Bohus",
		51: "Göteborg och Bohus",
		52: "Göteborg och Bohus",
		53: "Göteborg och Bohus",
		54: "Göteborg och Bohus",
		55: "Älvsborg",
		56: "Älvsborg",
		57: "Älvsborg",
		58: "Älvsborg",
		59: "Skaraborg",
		60: "Skaraborg",
		61: "Skaraborg",
		62: "Värmland",
		63: "Värmland",
		64: "Värmland",
		65: "Extra numbers",
		66: "Örebro",
		67: "Örebro",
		68: "Örebro",
		69: "Västmanland",
		70: "Västmanland",
		71: "Kopparberg",
		72: "Kopparberg",
		73: "Kopparberg",
		74: "Extra numbers",
		75: "Gävleborg",
		76: "Gävleborg",
		77: "Gävleborg",
		78: "Västernorrland",
		79: "Västernorrland",
		80: "Västernorrland",
		81: "Västernorrland",
		82: "Jämtland",
		83: "Jämtland",
		84: "Jämtland",
		85: "Västerbotten",
		86: "Västerbotten",
		87: "Västerbotten",
		88: "Västerbotten",
		89: "Norrbotten",
		90: "Norrbotten",
		91: "Norrbotten",
		92: "Norrbotten",
		93: "Extra numbers",
		94: "Extra numbers",
		95: "Extra numbers",
		96: "Extra numbers",
		97: "Extra numbers",
		98: "Extra numbers",
		99: "Extra numbers",
	}
)

func (p *Pin) GetPlaceOfBirth() string {
	place, ok := countyNumbers[StringToInt(p.Place)]

	if !ok {
		return ""
	}

	return place
}

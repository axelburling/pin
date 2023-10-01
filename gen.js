// Do not modify this file! (unless you know what you are doing)

// Data is from https://sv.wikipedia.org/wiki/Personnummer_i_Sverige#Personnumrets_uppbyggnad "Födelsenummer", Data might not be perfect but it's good enough for this project.

const states = [
    {
        name: 'Stockholm',
        min: 0,
        max: 13
    },
    {
        name: 'Uppsala',
        min: 14,
        max: 15
    },
     {
        name: 'Södermanland',
        min: 16,
        max: 18
    },
    {
        name: 'Östergötland',
        min: 19,
        max: 23
    },
    {
        name: 'Jönköping',
        min: 24,
        max: 26
    },
    {
        name: 'Kronoberg',
        min: 27,
        max: 28
    }, 
    {
        name: 'Kalmar',
        min: 29,
        max: 31
    },
    {
        name: 'Gotland',
        min: 32,
        max: 32
    },
    {
        name: 'Blekinge',
        min: 33,
        max: 34
    },
    {
        name: 'Kristianstad',
        min: 35,
        max: 38
    },
    {
        name: 'Malmöhus',
        min: 39,
        max: 45
    },
    {
        name: 'Halland',
        min: 46,
        max: 47
    },
    {
        name: 'Göteborg och Bohus',
        min: 48,
        max: 54
    },
    {
        name: 'Älvsborg',
        min: 55,
        max: 58
    }, 
    {
        name: 'Skaraborg',
        min: 59,
        max: 61
    },
    {
        name: 'Värmland',
        min: 62,
        max: 64
    }, 
    {
        name: 'Extra numbers',
        min: 65,
        max: 65
    },
    {
        name: 'Örebro',
        min: 66,
        max: 68
    },
    {
        name: 'Västmanland',
        min: 69,
        max: 70
    },
    {
        name: 'Kopparberg',
        min: 71,
        max: 73
    },
    {
        name: 'Extra numbers',
        min: 74,
        max: 74
    },
    {
        name: 'Gävleborg',
        min: 75,
        max: 77
    },
    {
        name: 'Västernorrland',
        min: 78,
        max: 81
    },
    {
        name: 'Jämtland',
        min: 82,
        max: 84
    },
    {
        name: 'Västerbotten',
        min: 85,
        max: 88
    },
    {
        name: 'Norrbotten',
        min: 89,
        max: 92
    },
    {
        name: 'Extra numbers',
        min: 93,
        max: 99
    }
]

function isBetween(min, max, val) {
    return val >= min && val <= max
}

for(let i = 0; i < 100; i++) {
    states.map(state => {
        const { min, max, name } = state
        if(isBetween(min, max, i)) {
            console.log(`${i}: "${name}",`)
        }
    })
}
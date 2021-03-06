package core

import (
	"fmt"
	"math/rand"
)

var (
	left = [...]string{
		"admiring",
		"adoring",
		"affectionate",
		"agitated",
		"amazing",
		"angry",
		"awesome",
		"beautiful",
		"blissful",
		"bold",
		"boring",
		"brave",
		"busy",
		"charming",
		"clever",
		"cool",
		"compassionate",
		"competent",
		"condescending",
		"confident",
		"cranky",
		"crazy",
		"dazzling",
		"determined",
		"distracted",
		"dreamy",
		"eager",
		"ecstatic",
		"elastic",
		"elated",
		"elegant",
		"eloquent",
		"epic",
		"exciting",
		"fervent",
		"festive",
		"flamboyant",
		"focused",
		"friendly",
		"frosty",
		"funny",
		"gallant",
		"gifted",
		"goofy",
		"gracious",
		"great",
		"happy",
		"hardcore",
		"heuristic",
		"hopeful",
		"hungry",
		"infallible",
		"inspiring",
		"interesting",
		"intelligent",
		"jolly",
		"jovial",
		"keen",
		"kind",
		"laughing",
		"loving",
		"lucid",
		"magical",
		"mystifying",
		"modest",
		"musing",
		"naughty",
		"nervous",
		"nice",
		"nifty",
		"nostalgic",
		"objective",
		"optimistic",
		"peaceful",
		"pedantic",
		"pensive",
		"practical",
		"priceless",
		"quirky",
		"quizzical",
		"recursing",
		"relaxed",
		"reverent",
		"romantic",
		"sad",
		"serene",
		"sharp",
		"silly",
		"sleepy",
		"stoic",
		"strange",
		"stupefied",
		"suspicious",
		"sweet",
		"tender",
		"thirsty",
		"trusting",
		"unruffled",
		"upbeat",
		"vibrant",
		"vigilant",
		"vigorous",
		"wizardly",
		"wonderful",
		"xenodochial",
		"youthful",
		"zealous",
		"zen",
	}

	// Docker, starting from 0.7.x, generates names from notable scientists and hackers.
	// Please, for any amazing man that you add to the list, consider adding an equally amazing woman to it, and vice versa.
	right = [...]string{
		"aardvark",
		"abyssinian",
		"affenpinscher",
		"akbash",
		"akita",
		"albatross",
		"alligator",
		"alpaca",
		"angelfish",
		"ant",
		"anteater",
		"antelope",
		"ape",
		"armadillo",
		"ass",
		"avocet",
		"axolotl",
		"baboon",
		"badger",
		"balinese",
		"bandicoot",
		"barb",
		"barnacle",
		"barracuda",
		"bat",
		"beagle",
		"bear",
		"beaver",
		"bee",
		"beetle",
		"binturong",
		"bird",
		"birman",
		"bison",
		"bloodhound",
		"boar",
		"bobcat",
		"bombay",
		"bongo",
		"bonobo",
		"booby",
		"budgerigar",
		"buffalo",
		"bulldog",
		"bullfrog",
		"burmese",
		"butterfly",
		"caiman",
		"camel",
		"capybara",
		"caracal",
		"caribou",
		"cassowary",
		"cat",
		"caterpillar",
		"catfish",
		"cattle",
		"centipede",
		"chameleon",
		"chamois",
		"cheetah",
		"chicken",
		"chihuahua",
		"chimpanzee",
		"chinchilla",
		"chinook",
		"chipmunk",
		"chough",
		"cichlid",
		"clam",
		"coati",
		"cobra",
		"cockroach",
		"cod",
		"collie",
		"coral",
		"cormorant",
		"cougar",
		"cow",
		"coyote",
		"crab",
		"crane",
		"crocodile",
		"crow",
		"curlew",
		"cuscus",
		"cuttlefish",
		"dachshund",
		"dalmatian",
		"deer",
		"dhole",
		"dingo",
		"dinosaur",
		"discus",
		"dodo",
		"dog",
		"dogfish",
		"dolphin",
		"donkey",
		"dormouse",
		"dotterel",
		"dove",
		"dragonfly",
		"drever",
		"duck",
		"dugong",
		"dunker",
		"dunlin",
		"eagle",
		"earwig",
		"echidna",
		"eel",
		"eland",
		"elephant",
		"elk",
		"emu",
		"falcon",
		"ferret",
		"finch",
		"fish",
		"flamingo",
		"flounder",
		"fly",
		"fossa",
		"fox",
		"frigatebird",
		"frog",
		"galago",
		"gar",
		"gaur",
		"gazelle",
		"gecko",
		"gerbil",
		"gharial",
		"gibbon",
		"giraffe",
		"gnat",
		"gnu",
		"goat",
		"goldfinch",
		"goldfish",
		"goose",
		"gopher",
		"gorilla",
		"goshawk",
		"grasshopper",
		"greyhound",
		"grouse",
		"guanaco",
		"gull",
		"guppy",
		"hamster",
		"hare",
		"harrier",
		"havanese",
		"hawk",
		"hedgehog",
		"heron",
		"herring",
		"himalayan",
		"hippopotamus",
		"hornet",
		"horse",
		"human",
		"hummingbird",
		"hyena",
		"ibis",
		"iguana",
		"impala",
		"indri",
		"insect",
		"jackal",
		"jaguar",
		"javanese",
		"jay",
		"jellyfish",
		"kakapo",
		"kangaroo",
		"kingfisher",
		"kiwi",
		"koala",
		"kouprey",
		"kudu",
		"labradoodle",
		"ladybird",
		"lapwing",
		"lark",
		"lemming",
		"lemur",
		"leopard",
		"liger",
		"lion",
		"lionfish",
		"lizard",
		"llama",
		"lobster",
		"locust",
		"loris",
		"louse",
		"lynx",
		"lyrebird",
		"macaw",
		"magpie",
		"mallard",
		"maltese",
		"manatee",
		"mandrill",
		"markhor",
		"marten",
		"mastiff",
		"mayfly",
		"meerkat",
		"millipede",
		"mink",
		"mole",
		"molly",
		"mongoose",
		"mongrel",
		"monkey",
		"moorhen",
		"moose",
		"mosquito",
		"moth",
		"mouse",
		"mule",
		"narwhal",
		"neanderthal",
		"newfoundland",
		"newt",
		"nightingale",
		"numbat",
		"ocelot",
		"octopus",
		"okapi",
		"olm",
		"opossum",
		"orangutan",
		"oryx",
		"ostrich",
		"otter",
		"owl",
		"ox",
		"oyster",
		"pademelon",
		"panther",
		"parrot",
		"partridge",
		"peacock",
		"peafowl",
		"pekingese",
		"pelican",
		"penguin",
		"persian",
		"pheasant",
		"pig",
		"pigeon",
		"pika",
		"pike",
		"piranha",
		"platypus",
		"pointer",
		"pony",
		"poodle",
		"porcupine",
		"porpoise",
		"possum",
		"prawn",
		"puffin",
		"pug",
		"puma",
		"quail",
		"quelea",
		"quetzal",
		"quokka",
		"quoll",
		"rabbit",
		"raccoon",
		"ragdoll",
		"rail",
		"ram",
		"rat",
		"rattlesnake",
		"raven",
		"reindeer",
		"rhinoceros",
		"robin",
		"rook",
		"rottweiler",
		"ruff",
		"salamander",
		"salmon",
		"sandpiper",
		"saola",
		"sardine",
		"scorpion",
		"seahorse",
		"seal",
		"serval",
		"shark",
		"sheep",
		"shrew",
		"shrimp",
		"siamese",
		"siberian",
		"skunk",
		"sloth",
		"snail",
		"snake",
		"snowshoe",
		"somali",
		"sparrow",
		"spider",
		"sponge",
		"squid",
		"squirrel",
		"starfish",
		"starling",
		"stingray",
		"stinkbug",
		"stoat",
		"stork",
		"swallow",
		"swan",
		"tang",
		"tapir",
		"tarsier",
		"termite",
		"tetra",
		"tiffany",
		"tiger",
		"toad",
		"tortoise",
		"toucan",
		"tropicbird",
		"trout",
		"tuatara",
		"turkey",
		"turtle",
		"uakari",
		"uguisu",
		"umbrellabird",
		"viper",
		"vulture",
		"wallaby",
		"walrus",
		"warthog",
		"wasp",
		"weasel",
		"whale",
		"whippet",
		"wildebeest",
		"wolf",
		"wolverine",
		"wombat",
		"woodcock",
		"woodlouse",
		"woodpecker",
		"worm",
		"wrasse",
		"wren",
		"yak",
		"zebra",
		"zebu",
		"zonkey",
		"zorse",
	}
)

func GetRandomName(retry int) string {
	name := fmt.Sprintf("%s_%s_%s", left[rand.Intn(len(left))], left[rand.Intn(len(left))], right[rand.Intn(len(right))])
	if retry > 0 {
		name = fmt.Sprintf("%s_%d", name, rand.Intn(10))
	}
	return name
}

package bun

import (
	"net/url"
	"strings"
)

var defaultCommand = "g"

// Commands is a map from the list of available commands to the function that will handle the redirect
var Commands = map[string]Command{

	"bad": Command{
		Name:         "Blockchain address",
		Key:          "bad",
		redirectFunc: simpleAppend("https://blockchain.info/address/"),
		Help:         "Finds the specific bitcoin address on blockchain.info",
		private:      false,
	},
	"bh": Command{
		Name:         "Behance",
		Key:          "bh",
		redirectFunc: simpleQuery("https://www.behance.net/search", "search"),
		Help:         "Searches Behance",
		private:      false,
	},
	"btx": Command{
		Name:         "Blockchain tx",
		Key:          "btx",
		redirectFunc: simpleAppend("https://blockchain.info/tx/"),
		Help:         "Finds the specific bitcoin transaction on blockchain.info",
		private:      false,
	},
	"d": Command{
		Name:         "Google drive",
		Key:          "d",
		redirectFunc: simpleQuery("https://drive.google.com/drive/search", "q"),
		Help:         "Searches Google Drive",
		private:      false,
	},
	"db": Command{
		Name:         "Dribble",
		Key:          "db",
		redirectFunc: simpleQuery("https://dribbble.com/search", "q"),
		Help:         "Searches Dribbble",
		private:      false,
	},
	"etx": Command{
		Name:         "Ethereum transaction",
		Key:          "etx",
		redirectFunc: simpleAppend("https://etherscan.io/tx/"),
		Help:         "Finds the ethereum transaction on Etherscan",
		private:      false,
	},
	"ead": Command{
		Name:         "Ethereum address",
		Key:          "ead",
		redirectFunc: simpleAppend("https://etherscan.io/address/"),
		Help:         "Finds the ethereum address on Etherscan",
		private:      false,
	},
	"fi": Command{
		Name:         "FlatIcon",
		Key:          "fi",
		redirectFunc: simpleQuery("https://www.flaticon.com/search", "word"),
		Help:         "FlatIcon search",
		private:      false,
	},
	"fp": Command{
		Name:         "FreePic",
		Key:          "fp",
		redirectFunc: simpleAppend("http://www.freepik.com/index.php?goto=2&searchform=1&k="),
		Help:         "FreePic search",
		private:      false,
	},
	"g": Command{
		Name:         "Google",
		Key:          "g",
		redirectFunc: simpleQuery("https://www.google.com/search", "q"),
		Help:         "Google search",
		private:      false,
	},
	"gd": Command{
		Name:         "Google Design",
		Key:          "gd",
		redirectFunc: simpleQuery("https://design.google/search/", "q"),
		Help:         "Google design search",
		private:      false,
	},
	"gm": Command{
		Name:         "Gmail",
		Key:          "gm",
		redirectFunc: simpleAppend("https://mail.google.com/mail/#search/"),
		Help:         "Searches your Gmail inbox",
		private:      false,
	},
	"h": Command{
		Name:         "help",
		Key:          "h",
		redirectFunc: simpleRedirect("/"),
		Help:         "Shows help (You're looking at it!)",
		private:      false,
	},
	"np": Command{
		Name:         "Noun Project",
		Key:          "np",
		redirectFunc: simpleQuery("https://thenounproject.com/search/", "q"),
		Help:         "Noun Project search",
		private:      false,
	},
	"pin": Command{
		Name:         "Pinterest",
		Key:          "pin",
		redirectFunc: simpleQuery("https://za.pinterest.com/search/pins/", "q"),
		Help:         "Pinterest search",
		private:      false,
	},
	"so": Command{
		Name:         "StackOverflow",
		Key:          "so",
		redirectFunc: simpleQuery("https://stackoverflow.com/search", "q"),
		Help:         "StackOverflow search",
		private:      false,
	},
	"sl": Command{
		Name:         "Shelflife",
		Key:          "sl",
		redirectFunc: simpleQuery("https://www.shelflife.co.za/search", "search"),
		Help:         "For your sneaker needs",
		private:      false,
	},
	"tk": Command{
		Name:         "Takealot",
		Key:          "tk",
		redirectFunc: simpleQuery("https://www.takealot.com/all", "qsearch"),
		Help:         "Takealot search",
		private:      false,
	},
	"tw": Command{
		Name:         "Twitter",
		Key:          "tw",
		redirectFunc: simpleQuery("https://www.twitter.com/search", "q"),
		Help:         "Twitter search",
		private:      false,
	},
	"us": Command{
		Name:         "Usplash",
		Key:          "us",
		redirectFunc: usplash,
		Help:         "Usplash search",
		private:      false,
	},
	"wk": Command{
		Name:         "Wikipedia",
		Key:          "wk",
		redirectFunc: simpleQuery("https://en.wikipedia.org/w/index.php", "search"),
		Help:         "Search Wikipedia",
		private:      false,
	},
	"yt": Command{
		Name:         "YouTube",
		Key:          "yt",
		redirectFunc: simpleQuery("https://www.youtube.com/results", "search_query"),
		Help:         "Search YouTube",
		private:      false,
	},
}

// simpleQuery handles search urls in the format
// baseURL?key=query
func simpleQuery(baseURL, key string) redirector {
	return func(query string) string {
		redirectURL, _ := url.Parse(baseURL)
		q := redirectURL.Query()
		q.Set(key, query)
		redirectURL.RawQuery = q.Encode()
		return redirectURL.String()
	}
}

// simpleRedirect handles situations where search is not
// possible, just always redirects to specific url
func simpleRedirect(baseURL string) redirector {
	return func(query string) string {
		return baseURL
	}
}

// simpleAppend handles search urls in the format
// baseURL/query
func simpleAppend(baseURL string) redirector {
	return func(query string) string {
		return baseURL + query
	}
}

// usplash uses '-'instead of '+' in their query URLS
func usplash(query string) string {
	wrongURL := simpleAppend("https://unsplash.com/search/")(query)
	return strings.Replace(wrongURL, "+", "-", -1)
}

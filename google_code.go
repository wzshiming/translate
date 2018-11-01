package translate

import (
	"fmt"
)

type GoogleCode uint32

const (
	GoogleAuto GoogleCode = iota //
	GoogleAF                     // Burry (Afrikaans),     布尔语(南非荷兰语), Burry (Afrikaans),
	GoogleAM                     // Amharic,               阿姆哈拉语,         አማርኛ,
	GoogleAR                     // Arabic,                阿拉伯语,           العربية,
	GoogleAZ                     // Azerbaijani,           阿塞拜疆语,         Azərbaycan,
	GoogleBE                     // Belarusian,            白俄罗斯语,         беларускі,
	GoogleBG                     // Bulgarian,             保加利亚语,         български,
	GoogleBN                     // Bengali,               孟加拉语,           বাংলা ভাষার,
	GoogleBS                     // Bosnian,               波斯尼亚语,         Bosanski,
	GoogleCA                     // Catalan,               加泰罗尼亚语,       Català,
	GoogleCEB                    // Cebu,                  宿务语,             Sugbo,
	GoogleCO                     // Corsican,              科西嘉语,           Corsa,
	GoogleCS                     // Czech,                 捷克语,             Česky,
	GoogleCY                     // Welsh,                 威尔士语,           Cymraeg,
	GoogleDA                     // Danish,                丹麦语,             dansk,
	GoogleDE                     // German,                德语,               Deutsche Sprache,
	GoogleEL                     // Greek,                 希腊语,             Ελληνικά,
	GoogleEN                     // English,               英语,               English,
	GoogleEO                     // Esperanto,             世界语,             Esperanto,
	GoogleES                     // Spanish,               西班牙语,           Español,
	GoogleET                     // Estonian,              爱沙尼亚语,         Eesti keel,
	GoogleEU                     // Basque,                巴斯克语,           Euskal,
	GoogleFA                     // Persian,               波斯语,             فارسی,
	GoogleFI                     // Finnish,               芬兰语,             suomalainen,
	GoogleFR                     // French,                法语,               Français,
	GoogleFY                     // Frisian,               弗里西语,           Frysk,
	GoogleGA                     // Irish,                 爱尔兰语,           Gaeilge,
	GoogleGD                     // Scottish Gaelic,       苏格兰盖尔语,       Gàidhlig na h-Alba,
	GoogleGL                     // Galician,              加利西亚语,         Galego,
	GoogleGU                     // Gujarati,              古吉拉特语,         ગુજરાતી,
	GoogleHA                     // Hausa,                 豪萨语,             Hausa,
	GoogleHAW                    // Hawaiian,              夏威夷语,           Hawaiian,
	GoogleHI                     // Hindi,                 印地语,             हिन्दी,
	GoogleHMN                    // Miao,                  苗语,               Miao,
	GoogleHR                     // Croatian,              克罗地亚语,         hrvatski,
	GoogleHT                     // Haitian Creole,        海地克里奥尔语,     Kreyòl Ayisyen,
	GoogleHU                     // Hungarian,             匈牙利语,           magyar,
	GoogleHY                     // Armenian,              亚美尼亚语,         Հայերեն,
	GoogleID                     // Indonesian,            印尼语,             Bahasa indonesia,
	GoogleIG                     // Ibo language,          伊博语,             Asụsụ Ibo,
	GoogleIS                     // Icelandic,             冰岛语,             Íslensku,
	GoogleIT                     // Italian,               意大利语,           lingua italiana,
	GoogleIW                     // Hebrew,                希伯来语,           עברית,
	GoogleJA                     // Japanese,              日语,               日本語,
	GoogleJW                     // Javanese,              印尼爪哇语,         Wong Jawa,
	GoogleKA                     // Georgian,              格鲁吉亚语,         ქართული,
	GoogleKK                     // Kazakh,                哈萨克语,           Қазақша,
	GoogleKM                     // Cambodian,             高棉语,             ភាសាខ្មែរ,
	GoogleKN                     // Kannada,               卡纳达语,           ಕನ್ನಡ,
	GoogleKO                     // Korean,                韩语,               한국어,
	GoogleKU                     // Kurdish,               库尔德语,           Kurdî,
	GoogleKY                     // Kyrgyz,                吉尔吉斯语,         Кыргыз тили,
	GoogleLA                     // Latin,                 拉丁语,             Latine,
	GoogleLB                     // Luxembourgish,         卢森堡语,           Lëtzebuergesch,
	GoogleLO                     // Lao,                   老挝语,             ລາວ,
	GoogleLT                     // Lithuanian,            立陶宛语,           Lietuviškai,
	GoogleLV                     // Latvian,               拉脱维亚语,         Latviešu,
	GoogleMG                     // Malagasy,              马尔加什语,         Malagasy,
	GoogleMI                     // Maori,                 毛利语,             Maori,
	GoogleMK                     // Macedonian,            马其顿语,           Македонски,
	GoogleML                     // Malayalam,             马拉雅拉姆语,       മലയാളം,
	GoogleMN                     // Mongolian,             蒙古语,             Монгол хэл,
	GoogleMR                     // Marathi,               马拉地语,           मराठी,
	GoogleMS                     // Malay,                 马来语,             Melayu,
	GoogleMT                     // Maltese,               马耳他语,           Malti,
	GoogleMY                     // Burmese,               缅甸语,             မြန်မာ,
	GoogleNE                     // Nepali,                尼泊尔语,           नेपाली,
	GoogleNL                     // Dutch,                 荷兰语,             Nederlands,
	GoogleNO                     // Norwegian,             挪威语,             norsk språk,
	GoogleNY                     // Chichewa,              齐切瓦语,           Chichewa,
	GooglePA                     // Punjabi,               旁遮普语,           ਪੰਜਾਬੀ,
	GooglePL                     // Polish,                波兰语,             Polski,
	GooglePS                     // Pashto,                普什图语,           پښتو,
	GooglePT                     // Portuguese,            葡萄牙语,           Português,
	GoogleRO                     // Romanian,              罗马尼亚语,         românesc,
	GoogleRU                     // Russian,               俄语,               Русский язык,
	GoogleSD                     // Sindhi,                信德语,             سنڌي,
	GoogleSI                     // Sinhalese,             僧伽罗语,           සිංහල,
	GoogleSK                     // Slovak,                斯洛伐克语,         slovenského jazyk,
	GoogleSL                     // Slovenian,             斯洛文尼亚语,       Slovenščina,
	GoogleSM                     // Samoan,                萨摩亚语,           Samoa,
	GoogleSN                     // Shinra,                修纳语,             Shinra,
	GoogleSO                     // Somali,                索马里语,           Somali,
	GoogleSQ                     // Albanian,              阿尔巴尼亚语,       shqiptar,
	GoogleSR                     // Serbian,               塞尔维亚语,         Српски,
	GoogleST                     // Sesotho,               塞索托语,           Sesotho,
	GoogleSU                     // Indonesian Sundanese,  印尼巽他语,         Sunda Indonesian,
	GoogleSV                     // Swedish,               瑞典语,             Svenska,
	GoogleSW                     // Swahili,               斯瓦希里语,         Kiswahili,
	GoogleTA                     // Tamil,                 泰米尔语,           தமிழ் மொழி,
	GoogleTE                     // Telugu,                泰卢固语,           తెలుగు,
	GoogleTG                     // Tajik,                 塔吉克语,           Тоҷикӣ,
	GoogleTH                     // Thai,                  泰语,               ไทย,
	GoogleTL                     // Filipino,              菲律宾语,           Filipino,
	GoogleTR                     // Turkish,               土耳其语,           Türk dili,
	GoogleUK                     // Ukrainian,             乌克兰语,           Українська,
	GoogleUR                     // Urdu,                  乌尔都语,           اردو,
	GoogleUZ                     // Uzbek,                 乌兹别克语,         O'zbek,
	GoogleVI                     // Vietnamese,            越南语,             Người việt nam,
	GoogleXH                     // South African Xhosa,   南非科萨语,         IsiXhosa saseMzantsi Afrika,
	GoogleYI                     // Yiddish,               意第绪语,           ייִדיש,
	GoogleYO                     // Yoruba,                约鲁巴语,           Yorùbá,
	GoogleZHCN                   // Chinese (Simplified),  中文(简体),         中文(简体),
	GoogleZHTW                   // Chinese (Traditional), 中文(繁体),         中文(繁體),
	GoogleZU                     // South African Zulu,    南非祖鲁语,         I-South African Zulu,
)

func (c GoogleCode) String() string {
	s, ok := GoogleCodeMap[c]
	if ok {
		return s[0]
	}
	return fmt.Sprintf("GoogleCode(%d)", uint32(c))
}

var GoogleCodeMap = map[GoogleCode][]string{
	GoogleAuto: {"auto", "auto  - auto"},
	GoogleAF:   {"af", "af    - Burry (Afrikaans),     布尔语(南非荷兰语), Burry (Afrikaans),"},
	GoogleAM:   {"am", "am    - Amharic,               阿姆哈拉语,         አማርኛ,"},
	GoogleAR:   {"ar", "ar    - Arabic,                阿拉伯语,           العربية,"},
	GoogleAZ:   {"az", "az    - Azerbaijani,           阿塞拜疆语,         Azərbaycan,"},
	GoogleBE:   {"be", "be    - Belarusian,            白俄罗斯语,         беларускі,"},
	GoogleBG:   {"bg", "bg    - Bulgarian,             保加利亚语,         български,"},
	GoogleBN:   {"bn", "bn    - Bengali,               孟加拉语,           বাংলা ভাষার,"},
	GoogleBS:   {"bs", "bs    - Bosnian,               波斯尼亚语,         Bosanski,"},
	GoogleCA:   {"ca", "ca    - Catalan,               加泰罗尼亚语,       Català,"},
	GoogleCEB:  {"ceb", "ceb   - Cebu,                  宿务语,             Sugbo,"},
	GoogleCO:   {"co", "co    - Corsican,              科西嘉语,           Corsa,"},
	GoogleCS:   {"cs", "cs    - Czech,                 捷克语,             Česky,"},
	GoogleCY:   {"cy", "cy    - Welsh,                 威尔士语,           Cymraeg,"},
	GoogleDA:   {"da", "da    - Danish,                丹麦语,             dansk,"},
	GoogleDE:   {"de", "de    - German,                德语,               Deutsche Sprache,"},
	GoogleEL:   {"el", "el    - Greek,                 希腊语,             Ελληνικά,"},
	GoogleEN:   {"en", "en    - English,               英语,               English,"},
	GoogleEO:   {"eo", "eo    - Esperanto,             世界语,             Esperanto,"},
	GoogleES:   {"es", "es    - Spanish,               西班牙语,           Español,"},
	GoogleET:   {"et", "et    - Estonian,              爱沙尼亚语,         Eesti keel,"},
	GoogleEU:   {"eu", "eu    - Basque,                巴斯克语,           Euskal,"},
	GoogleFA:   {"fa", "fa    - Persian,               波斯语,             فارسی,"},
	GoogleFI:   {"fi", "fi    - Finnish,               芬兰语,             suomalainen,"},
	GoogleFR:   {"fr", "fr    - French,                法语,               Français,"},
	GoogleFY:   {"fy", "fy    - Frisian,               弗里西语,           Frysk,"},
	GoogleGA:   {"ga", "ga    - Irish,                 爱尔兰语,           Gaeilge,"},
	GoogleGD:   {"gd", "gd    - Scottish Gaelic,       苏格兰盖尔语,       Gàidhlig na h-Alba,"},
	GoogleGL:   {"gl", "gl    - Galician,              加利西亚语,         Galego,"},
	GoogleGU:   {"gu", "gu    - Gujarati,              古吉拉特语,         ગુજરાતી,"},
	GoogleHA:   {"ha", "ha    - Hausa,                 豪萨语,             Hausa,"},
	GoogleHAW:  {"haw", "haw   - Hawaiian,              夏威夷语,           Hawaiian,"},
	GoogleHI:   {"hi", "hi    - Hindi,                 印地语,             हिन्दी,"},
	GoogleHMN:  {"hmn", "hmn   - Miao,                  苗语,               Miao,"},
	GoogleHR:   {"hr", "hr    - Croatian,              克罗地亚语,         hrvatski,"},
	GoogleHT:   {"ht", "ht    - Haitian Creole,        海地克里奥尔语,     Kreyòl Ayisyen,"},
	GoogleHU:   {"hu", "hu    - Hungarian,             匈牙利语,           magyar,"},
	GoogleHY:   {"hy", "hy    - Armenian,              亚美尼亚语,         Հայերեն,"},
	GoogleID:   {"id", "id    - Indonesian,            印尼语,             Bahasa indonesia,"},
	GoogleIG:   {"ig", "ig    - Ibo language,          伊博语,             Asụsụ Ibo,"},
	GoogleIS:   {"is", "is    - Icelandic,             冰岛语,             Íslensku,"},
	GoogleIT:   {"it", "it    - Italian,               意大利语,           lingua italiana,"},
	GoogleIW:   {"iw", "iw    - Hebrew,                希伯来语,           עברית,"},
	GoogleJA:   {"ja", "ja    - Japanese,              日语,               日本語,"},
	GoogleJW:   {"jw", "jw    - Javanese,              印尼爪哇语,         Wong Jawa,"},
	GoogleKA:   {"ka", "ka    - Georgian,              格鲁吉亚语,         ქართული,"},
	GoogleKK:   {"kk", "kk    - Kazakh,                哈萨克语,           Қазақша,"},
	GoogleKM:   {"km", "km    - Cambodian,             高棉语,             ភាសាខ្មែរ,"},
	GoogleKN:   {"kn", "kn    - Kannada,               卡纳达语,           ಕನ್ನಡ,"},
	GoogleKO:   {"ko", "ko    - Korean,                韩语,               한국어,"},
	GoogleKU:   {"ku", "ku    - Kurdish,               库尔德语,           Kurdî,"},
	GoogleKY:   {"ky", "ky    - Kyrgyz,                吉尔吉斯语,         Кыргыз тили,"},
	GoogleLA:   {"la", "la    - Latin,                 拉丁语,             Latine,"},
	GoogleLB:   {"lb", "lb    - Luxembourgish,         卢森堡语,           Lëtzebuergesch,"},
	GoogleLO:   {"lo", "lo    - Lao,                   老挝语,             ລາວ,"},
	GoogleLT:   {"lt", "lt    - Lithuanian,            立陶宛语,           Lietuviškai,"},
	GoogleLV:   {"lv", "lv    - Latvian,               拉脱维亚语,         Latviešu,"},
	GoogleMG:   {"mg", "mg    - Malagasy,              马尔加什语,         Malagasy,"},
	GoogleMI:   {"mi", "mi    - Maori,                 毛利语,             Maori,"},
	GoogleMK:   {"mk", "mk    - Macedonian,            马其顿语,           Македонски,"},
	GoogleML:   {"ml", "ml    - Malayalam,             马拉雅拉姆语,       മലയാളം,"},
	GoogleMN:   {"mn", "mn    - Mongolian,             蒙古语,             Монгол хэл,"},
	GoogleMR:   {"mr", "mr    - Marathi,               马拉地语,           मराठी,"},
	GoogleMS:   {"ms", "ms    - Malay,                 马来语,             Melayu,"},
	GoogleMT:   {"mt", "mt    - Maltese,               马耳他语,           Malti,"},
	GoogleMY:   {"my", "my    - Burmese,               缅甸语,             မြန်မာ,"},
	GoogleNE:   {"ne", "ne    - Nepali,                尼泊尔语,           नेपाली,"},
	GoogleNL:   {"nl", "nl    - Dutch,                 荷兰语,             Nederlands,"},
	GoogleNO:   {"no", "no    - Norwegian,             挪威语,             norsk språk,"},
	GoogleNY:   {"ny", "ny    - Chichewa,              齐切瓦语,           Chichewa,"},
	GooglePA:   {"pa", "pa    - Punjabi,               旁遮普语,           ਪੰਜਾਬੀ,"},
	GooglePL:   {"pl", "pl    - Polish,                波兰语,             Polski,"},
	GooglePS:   {"ps", "ps    - Pashto,                普什图语,           پښتو,"},
	GooglePT:   {"pt", "pt    - Portuguese,            葡萄牙语,           Português,"},
	GoogleRO:   {"ro", "ro    - Romanian,              罗马尼亚语,         românesc,"},
	GoogleRU:   {"ru", "ru    - Russian,               俄语,               Русский язык,"},
	GoogleSD:   {"sd", "sd    - Sindhi,                信德语,             سنڌي,"},
	GoogleSI:   {"si", "si    - Sinhalese,             僧伽罗语,           සිංහල,"},
	GoogleSK:   {"sk", "sk    - Slovak,                斯洛伐克语,         slovenského jazyk,"},
	GoogleSL:   {"sl", "sl    - Slovenian,             斯洛文尼亚语,       Slovenščina,"},
	GoogleSM:   {"sm", "sm    - Samoan,                萨摩亚语,           Samoa,"},
	GoogleSN:   {"sn", "sn    - Shinra,                修纳语,             Shinra,"},
	GoogleSO:   {"so", "so    - Somali,                索马里语,           Somali,"},
	GoogleSQ:   {"sq", "sq    - Albanian,              阿尔巴尼亚语,       shqiptar,"},
	GoogleSR:   {"sr", "sr    - Serbian,               塞尔维亚语,         Српски,"},
	GoogleST:   {"st", "st    - Sesotho,               塞索托语,           Sesotho,"},
	GoogleSU:   {"su", "su    - Indonesian Sundanese,  印尼巽他语,         Sunda Indonesian,"},
	GoogleSV:   {"sv", "sv    - Swedish,               瑞典语,             Svenska,"},
	GoogleSW:   {"sw", "sw    - Swahili,               斯瓦希里语,         Kiswahili,"},
	GoogleTA:   {"ta", "ta    - Tamil,                 泰米尔语,           தமிழ் மொழி,"},
	GoogleTE:   {"te", "te    - Telugu,                泰卢固语,           తెలుగు,"},
	GoogleTG:   {"tg", "tg    - Tajik,                 塔吉克语,           Тоҷикӣ,"},
	GoogleTH:   {"th", "th    - Thai,                  泰语,               ไทย,"},
	GoogleTL:   {"tl", "tl    - Filipino,              菲律宾语,           Filipino,"},
	GoogleTR:   {"tr", "tr    - Turkish,               土耳其语,           Türk dili,"},
	GoogleUK:   {"uk", "uk    - Ukrainian,             乌克兰语,           Українська,"},
	GoogleUR:   {"ur", "ur    - Urdu,                  乌尔都语,           اردو,"},
	GoogleUZ:   {"uz", "uz    - Uzbek,                 乌兹别克语,         O'zbek,"},
	GoogleVI:   {"vi", "vi    - Vietnamese,            越南语,             Người việt nam,"},
	GoogleXH:   {"xh", "xh    - South African Xhosa,   南非科萨语,         IsiXhosa saseMzantsi Afrika,"},
	GoogleYI:   {"yi", "yi    - Yiddish,               意第绪语,           ייִדיש,"},
	GoogleYO:   {"yo", "yo    - Yoruba,                约鲁巴语,           Yorùbá,"},
	GoogleZHCN: {"zh-CN", "zh-CN - Chinese (Simplified),  中文(简体),         中文(简体),"},
	GoogleZHTW: {"zh-TW", "zh-TW - Chinese (Traditional), 中文(繁体),         中文(繁體),"},
	GoogleZU:   {"zu", "zu    - South African Zulu,    南非祖鲁语,         I-South African Zulu,"},
}

"Hello world"
"a ë A"
"Buy milk: \u221A"
"Logotype for my band: \U00010299"
"A cool cat 😎"

// Different ways of embedding the same Unicode text:
"日本語"                                // UTF-8 input text
"\u65e5\u672c\u8a9e"                    // explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // explicit UTF-8 bytes

/*
    U+1F469 woman
  + U+1F3FD skin tone modifier
  + U+200D  zero width joiner
  + U+1F680
  = astronaut (woman) with skin tone modifier */
"👩🏽‍🚀" // UTF-8 input text
"\U0001F469\U0001F3FD\u200D\U0001F680" // Unicode code points
"\xf0\x9f\x91\xa9\xf0\x9f\x8f\xbd\xe2\x80\x8d\xf0\x9f\x9a\x80" // UTF-8 bytes

// Bytes vs code points
"a \xEB A" // the byte 0xEB
"a \u00EB A" // the code point U+00EB

// piece strings
"x ${y + "foo ${or} yo" + b ar} z"
"x \${y} z"

// multiline strings
"multi
 line"
message = "
           multi
           line
           "

// character literals
'c' 'B' 'a'
'\n' '\t' '\f'
'\0' '\x00' '\x0A'
'\u221A' '\U00010299'

// "\uD800"     // illegal: surrogate half
// "\U00110000" // illegal: invalid Unicode code point

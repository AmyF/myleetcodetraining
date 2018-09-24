func uniqueMorseRepresentations(words []string) int {
    if len(words) < 2 { return len(words)}
    dict := map[string]string{}
    for i := 0; i < len(words); i++ {
        word := words[i]
        dict[convert(word)] = ""
    }
    return len(dict)
}

func convert(word string) string {
    result := ""
    for i:=0;i<len(word);i++ {
        result = result + convertList[word[i]-'a']
    }
    return result
}

var convertList = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

/*
input: non empty string s, dictionary wordDict with a list of non empty words
output: boolean
 */


class Trie {
    constructor(value) {
        this.value = value;
        this.children = {};
    }

    addWord(word) {

        let i = 0;

        let node = this;

        while (i < word.length) {
            let letter = word[i];

            if (node.children[letter]) {
                node = node.children[letter];
                i++;
            } else {
                node.children[letter] = new Trie(letter);
                i++;
                node = node.children[letter];
            }
        }

        node.children['*'] = true;
    }
}


const wordBreak = (s, wordDict, node = new Trie('root'), i = 0) => {

    let result;

    if (i === s.length) {
        return true;
    }

    wordDict.forEach(word => {
        node.addWord(word);
    });

    let char = s[i];
    console.log(char);

    if (node.children[char]) {
        result = result || wordBreak(s, wordDict, node.children[char], ++i);
    }

    if (node.children['*']) {
        result = result || wordBreak(s, wordDict, new Trie('root'), i);
    }

    if (!node.children[char] && !node.children['*']) return false;

    return result;
}

let testString = 'catsandog';
let testDict = ['cats', 'dog', 'sand', 'and', 'cat'];

console.log(wordBreak(testString, testDict));
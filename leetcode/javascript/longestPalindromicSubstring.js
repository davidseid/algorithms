function longestPalindromicSubstring(s) {
    let longest = '';

    for (i = 0; i < s.length; i++) {
        let left = i;
        let right = i;

        while (left >= 0 && right < s.length && s[left] == s[right]) {
            left--;
            right++;
        }

        const palindrome = s.slice(left+1, right); 

        longest = palindrome.length > longest.length ? palindrome : longest; 

        if (i+1 < s.length && s[i] == s[i+1]) {
            left = i
            right = i+1

            while (left >= 0 && right < s.length && s[left] == s[right]) {
                left--;
                right++;
            }
    
            const palindrome = s.slice(left+1, right); 
    
            longest = palindrome.length > longest.length ? palindrome : longest; 
        }
    }

    return longest;
}

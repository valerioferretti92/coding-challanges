class Solution {

    private Map<Character, Integer> digitValues = new HashMap<>() {{
        put('I', 1);
        put('V', 5);
        put('X', 10);
        put('L', 50);
        put('C', 100);
        put('D', 500);
        put('M', 1000);
    }};

    public int romanToInt(String s) {
        int sum = 0;
        char c1 = '0', c2 = '0';
        for (int i = 0; i < s.length(); i++) {
            c1 = s.charAt(i);
            if (i+1 == s.length()) {
                return sum + valueOf(c1);
            }

            c2 = s.charAt(i+1);
            if (valueOf(c1) >= valueOf(c2)) {
                sum = sum + valueOf(c1);
                continue;
            }

            sum = sum + valueOfPair(c1, c2);
            i++;
        }
        return sum;
    }

    private int valueOf(char c) {
        return digitValues.get(c);      
    }

    private int valueOfPair(char c1, char c2) {
        return valueOf(c2) - valueOf(c1);
    }
}

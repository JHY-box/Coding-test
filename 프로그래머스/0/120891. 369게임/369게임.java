class Solution {
    public int solution(int order) {
        int answer = 0;
        
        String orderStr = String.valueOf(order);

        for (char c : orderStr.toCharArray()) {

            if (c=='3' || c == '6' || c=='9') {
                answer++;
            }
        }
        return answer;
    }
}
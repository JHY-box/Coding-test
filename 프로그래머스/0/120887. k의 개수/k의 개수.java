class Solution {
    public int solution(int i, int j, int k) {
        int count = 0; // 결과를 저장할 변수 초기화
        String kStr = String.valueOf(k); // k를 문자열로 변환
        // i부터 j까지 반복하며 k의 등장 횟수 계산
        for (int n = i; n <= j; n++) {
        String numStr = String.valueOf(n); // 현재 숫자를 문자열로 변환
        count += numStr.length() - numStr.replace(kStr, "").length(); // k의 개수를 누적
        }
        return count; // 최종 등장 횟수 반환
        }
}
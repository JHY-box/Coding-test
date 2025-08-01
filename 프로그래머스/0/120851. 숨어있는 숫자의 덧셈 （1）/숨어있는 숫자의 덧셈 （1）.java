class Solution {
    public int solution(String myString) {
 int answer = 0; // 숫자의 합을 저장할 변수
 // 문자열의 각 문자를 순회
 for (char m : myString.toCharArray()) {
 // 숫자인 경우
 if (Character.isDigit(m)) {
 // 정수로 변환하여 더함
 answer += Character.getNumericValue(m);
 }
 }
 return answer; // 최종 결과 반환
}
}
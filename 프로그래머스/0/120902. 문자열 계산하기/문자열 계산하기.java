class Solution {
    public int solution(String myString) {
 // 문자열을 공백 기준으로 분리하여 배열로 변환
 String[] myList = myString.split(" ");
 // 첫 번째 숫자를 정수로 변환하여 초기값으로 설정
 int answer = Integer.parseInt(myList[0]);
 // 1부터 2씩 증가하면서 연산자와 숫자를 순차적으로 처리
 for (int i = 1; i < myList.length - 1; i += 2) {
 // 연산자 추출
 String operator = myList[i];
 // 숫자 추출
 int number = Integer.parseInt(myList[i + 1]);
 if (operator.equals("+")) {
 // "+" 연산자이면 다음 숫자를 더함
 answer += number;
 } else {
 // "-" 연산자이면 다음 숫자를 뺌
 answer -= number;
      }
 }
 return answer; // 계산된 최종 결과 반환
 }
}
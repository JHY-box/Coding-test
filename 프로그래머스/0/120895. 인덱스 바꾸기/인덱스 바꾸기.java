class Solution {
    public String solution(String myString, int num1, int num2) {
 // 문자열을 문자 배열로 변환
 char[] charArray = myString.toCharArray();
 // num1과 num2의 위치에 있는 문자 교환
 char temp = charArray[num1];
 charArray[num1] = charArray[num2];
 charArray[num2] = temp;
 // 문자 배열을 다시 문자열로 변환하여 반환
 return new String(charArray);
 }
}
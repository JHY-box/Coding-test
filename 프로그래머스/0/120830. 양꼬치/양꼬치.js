/**function solution(n, k) {
    let answer = (12000 * n) + (2000 * k) - (2000*n/10);
    return answer;
}
*/

function solution(n, k) {
  let answer = 0;
  
  answer = (n * 12000) + (k *2000) - (parseInt(n/10) *2000)
  return answer;
}
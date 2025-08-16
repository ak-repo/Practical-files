var isPowerOfFour = function (n) {
  if (n === 1 || n == 4) return true;

  let run = 1;
  while (run <= n) {
    run *= 4;
    if (run === n) return true;
  }
  return false;
};

console.log(isPowerOfFour(14));

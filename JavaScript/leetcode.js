class Calculator {
  constructor(value) {
    this.value = value;
  }

  add(value) {
    this.value = this.value + value;
    return this;
  }

  subtract(value) {
    this.value = this.value - value;
    return this;
  }

  multiply(value) {
    this.value = this.value * value;
  }

  divide(value) {
    this.value = this.value / value;
    return this;
  }

  power(value) {
    this.value = this.value ** value;
    return this;
  }

  getResult() {
    return this.value;
  }
}

// Input:
// actions = ["Calculator", "add", "subtract", "getResult"],
// values = [10, 5, 7]
// Output: 8
// Explanation:
const x =new Calculator(2).power(2).getResult()
console.log(x);

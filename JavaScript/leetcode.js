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
    return this;
  }

  divide(value) {
    if (value === 0) {
      this.value = "Division by zero is not allowed";
    } else {
      this.value = this.value / value;
      return this;
    }
  }

  power(value) {
    this.value = this.value ** value;
    return this;
  }

  getResult() {
    return this.value;
  }
}

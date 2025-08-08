async function sleep(millis) {
  await new Promise((resolve) => setTimeout(() => resolve(millis), millis));
}

console.log(sleep(100));
